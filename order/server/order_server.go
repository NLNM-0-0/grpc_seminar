package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"order/client"
	orderv1 "order/gen/seminar/order/v1"
	"order/model"
	"order/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	_, err := server.userClient.GetUser(ctx, request.UserId)
	if err != nil {
		log.Println("Error getting user: ", err)
		return nil, errors.New("user not found")
	}

	productIDs := []string{}
	for _, p := range request.Products {
		productIDs = append(productIDs, p.ProductId)
	}
	orderProducts, err := server.productClient.GetProductByIds(ctx, productIDs)
	if err != nil {
		log.Println("Error getting products: ", err)
		return nil, errors.New("products not found")
	}

	products := []model.OrderProduct{}
	for _, p := range orderProducts {
		product := model.OrderProduct{
			Id:       p.Id,
			Quantity: 0,
			Price:    p.Price,
			Name:     p.Name,
		}
		for _, rp := range request.Products {
			if rp.ProductId == p.Id {
				product.Quantity += rp.Quantity
				break
			}
		}

		if product.Quantity != 0 {
			products = append(products, product)
		}
	}

	order, err := server.orderService.CreateOrder(ctx, request.UserId, products)
	if err != nil {
		log.Println("Error posting order: ", err)
		return nil, errors.New("could not post order")
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

func (server *orderServer) GetOrdersForUser(
	ctx context.Context,
	request *orderv1.GetOrdersForUserRequest,
) (*orderv1.GetOrdersForUserResponse, error) {
	userOrders, err := server.orderService.GetOrdersForUser(ctx, request.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	productIDMap := map[string]bool{}
	for _, o := range userOrders {
		for _, p := range o.Products {
			productIDMap[p.Id] = true
		}
	}
	var productIDs []string
	for id := range productIDMap {
		productIDs = append(productIDs, id)
	}
	products, err := server.productClient.GetProductByIds(ctx, productIDs)
	if err != nil {
		log.Println("Error getting user products: ", err)
		return nil, err
	}

	var orders []*orderv1.Order
	for _, o := range userOrders {
		order := &orderv1.Order{
			UserId:     o.UserId,
			Id:         o.Id,
			TotalPrice: o.TotalPrice,
			Products:   []*orderv1.Order_OrderProduct{},
		}
		order.CreatedAt, _ = o.CreatedAt.MarshalBinary()

		for _, product := range o.Products {
			for _, p := range products {
				if p.Id == product.ProductId {
					product.Name = p.Name
					product.Price = p.Price
					break
				}
			}

			order.Products = append(order.Products, &orderv1.Order_OrderProduct{
				ProductId: product.Id,
				Name:      product.Name,
				Price:     product.Price,
				Quantity:  product.Quantity,
			})
		}

		orders = append(orders, order)
	}
	return &orderv1.GetOrdersForUserResponse{Orders: orders}, nil
}
