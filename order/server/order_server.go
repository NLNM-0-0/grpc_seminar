package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"order/client"
	orderv1 "order/gen/seminar/order/v1"
	productv1 "order/gen/seminar/product/v1"
	"order/model"
	"order/service"
	"order/utils"
	"sync"
)

type orderServer struct {
	orderService  service.OrderService
	userClient    client.UserClient
	productClient client.ProductClient
}

func ListenGRPC(service service.OrderService, userUrl, productUrl string, port int) error {
	userClient, err := client.NewUserClient(userUrl)
	if err != nil {
		return err
	}

	productClient, err := client.NewProductClient(productUrl)
	if err != nil {
		userClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		userClient.Close()
		productClient.Close()
		return err
	}

	serv := grpc.NewServer()
	orderv1.RegisterOrderServiceServer(serv, &orderServer{
		service,
		userClient,
		productClient,
	})

	reflection.Register(serv)

	return serv.Serve(lis)
}

func (server *orderServer) PostOrder(
	ctx context.Context,
	request *orderv1.PostOrderRequest,
) (*orderv1.PostOrderResponse, error) {
	ids := getProductIdsFromRequest(request)
	stream, err := server.productClient.GetProductByIds(ctx, ids)
	if err != nil {
		log.Println(err)
		return nil, utils.NewGrpcErrorWithMetadata(
			codes.Aborted,
			"Failed to receive stream request",
			"STREAM_RECEIVE_ERROR",
			err,
			nil,
		)
	}

	var orderProducts []model.OrderProduct
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, utils.NewGrpcErrorWithMetadata(
				codes.Internal,
				"Failed to get product by ids",
				"GET_PRODUCT_BY_IDS_ERROR",
				err,
				nil,
			)
		}

		grpcProduct := res.GetProduct()

		orderProduct := model.OrderProduct{
			ProductId: grpcProduct.Id,
			Quantity:  0,
			Price:     grpcProduct.Price,
			Name:      grpcProduct.Name,
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	for _, requestProduct := range request.Products {
		for i := range orderProducts {
			if requestProduct.ProductId == orderProducts[i].ProductId {
				orderProducts[i].Quantity += requestProduct.Quantity
				break
			}
		}
	}

	if _, err = server.userClient.GetUser(ctx, request.UserId); err != nil {
		log.Println(err)
		return nil, utils.NewGrpcErrorWithMetadata(
			codes.Internal,
			fmt.Sprintf("Failed to get user with ID %s", request.UserId),
			"GET_USER_ERROR",
			err,
			map[string]string{
				"user_id": request.UserId,
			},
		)
	}

	order, err := server.orderService.CreateOrder(ctx, request.UserId, orderProducts)
	if err != nil {
		log.Println(err)
		return nil, utils.NewGrpcErrorWithMetadata(
			codes.Internal,
			"Failed to create order",
			"CREATE_ORDER_ERROR",
			err,
			nil,
		)
	}

	orderProto := &orderv1.Order{
		Id:         order.Id,
		UserId:     order.UserId,
		TotalPrice: order.TotalPrice,
		Products:   []*orderv1.Order_OrderProduct{},
	}
	orderProto.CreatedAt, _ = order.CreatedAt.MarshalBinary()
	for _, p := range order.Products {
		orderProto.Products = append(orderProto.Products, &orderv1.Order_OrderProduct{
			ProductId: p.ProductId,
			Name:      p.Name,
			Price:     p.Price,
			Quantity:  p.Quantity,
		})
	}
	return &orderv1.PostOrderResponse{
		Order: orderProto,
	}, nil
}

func getProductIdsFromRequest(request *orderv1.PostOrderRequest) []string {
	var productIds []string
	for _, product := range request.Products {
		productIds = append(productIds, product.ProductId)
	}
	return productIds
}

func (server *orderServer) RateOrder(ctx context.Context, request *orderv1.RateOrderRequest) (*orderv1.RateOrderResponse, error) {
	stream, err := server.productClient.RateProductByIds(ctx)
	if err != nil {
		log.Println(err)
		return nil, utils.NewGrpcErrorWithMetadata(
			codes.Unknown,
			"Failed to call RateProductByIds",
			"RATE_PRODUCT_ERROR",
			err,
			nil,
		)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, rating := range request.Products {
			if err := stream.Send(rating); err != nil {
				log.Println(err)
				select {
				case errCh <- utils.NewGrpcErrorWithMetadata(
					codes.Aborted,
					fmt.Sprintf("Failed to send stream request"),
					"STREAM_SEND_ERROR",
					err,
					nil,
				):
				default:
				}
				return
			}
		}

		if err := stream.CloseSend(); err != nil {
			log.Println(err)
			select {
			case errCh <- utils.NewGrpcErrorWithMetadata(
				codes.Aborted,
				fmt.Sprintf("Failed to close stream request"),
				"STREAM_CLOSE_ERROR",
				err,
				nil,
			):
			default:
			}
		}
	}()

	var responses []*productv1.RateProductByIdsResponse

	for {
		select {
		case err := <-errCh:
			return nil, err
		default:
		}

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, utils.NewGrpcErrorWithMetadata(
				codes.Internal,
				"Failed to receive stream response",
				"STREAM_RECEIVE_ERROR",
				err,
				nil,
			)
		}
		responses = append(responses, res)
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return nil, err
	default:
	}

	return &orderv1.RateOrderResponse{
		Ratings: responses,
	}, nil
}
