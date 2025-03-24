package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	productv1 "product/gen/seminar/product/v1"
	"product/service"
	"product/utils"
)

type productServer struct {
	productService service.ProductService
}

func ListenGRPC(service service.ProductService, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	productv1.RegisterProductServiceServer(serv, &productServer{
		productService: service,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (server *productServer) PostProduct(ctx context.Context, request *productv1.PostProductRequest) (*productv1.PostProductResponse, error) {
	product, err := server.productService.PostProduct(ctx, request.Name, request.Price)
	if err != nil {
		log.Println(err)
		var appErr *utils.AppError
		if errors.As(err, &appErr) {
			if appErr.Key == utils.DB_DUPLICATE_ERROR_CODE {
				return nil, utils.NewGrpcErrorWithMetadata(
					codes.AlreadyExists,
					"Product name already exists",
					utils.DB_DUPLICATE_ERROR_CODE,
					err,
					nil,
				)
			}
		}

		return nil, utils.NewGrpcErrorWithMetadata(
			codes.Internal,
			"Failed to create product",
			"INTERNAL_ERROR",
			err,
			nil,
		)
	}
	return &productv1.PostProductResponse{Product: &productv1.Product{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}}, nil
}

func (server *productServer) GetProduct(ctx context.Context, request *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	product, err := server.productService.GetProduct(ctx, request.Id)
	if err != nil {
		log.Println(err)
		return nil, server.handleGetProductDBError(err, request.Id)
	}
	return &productv1.GetProductResponse{
		Product: &productv1.Product{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		},
	}, nil
}

func (server *productServer) GetProductByIds(request *productv1.GetProductsByIdsRequest, stream grpc.ServerStreamingServer[productv1.GetProductsByIdsResponse]) error {
	for _, id := range request.GetIds() {
		product, err := server.productService.GetProduct(stream.Context(), id)
		if err != nil {
			log.Println(err)
			return server.handleGetProductDBError(err, id)
		}

		res := &productv1.GetProductsByIdsResponse{
			Product: &productv1.Product{
				Id:    product.Id,
				Name:  product.Name,
				Price: product.Price,
			},
		}

		if err := stream.Send(res); err != nil {
			return utils.NewGrpcErrorWithMetadata(
				codes.Aborted,
				"Failed to send stream request",
				"STREAM_SEND_ERROR",
				err,
				nil,
			)
		}
	}

	return nil
}

func (server *productServer) RateProductByIds(stream grpc.BidiStreamingServer[productv1.RateProductByIdsRequest, productv1.RateProductByIdsResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println(err)
			return utils.NewGrpcErrorWithMetadata(
				codes.Aborted,
				"Failed to receive stream request",
				"STREAM_RECEIVE_ERROR",
				err,
				map[string]string{},
			)
		}

		productId := req.GetId()
		score := req.GetScore()

		if _, err = server.productService.GetProduct(stream.Context(), productId); err != nil {
			return server.handleGetProductDBError(err, productId)
		}

		rating, err := server.productService.AddRating(stream.Context(), productId, score)
		if err != nil {
			log.Println(err)
			return utils.NewGrpcErrorWithMetadata(
				codes.Internal,
				fmt.Sprintf("Error adding rating product ID %s", productId),
				"ADD_RATING_ERROR",
				err,
				map[string]string{
					"product_id": productId,
				},
			)
		}

		res := &productv1.RateProductByIdsResponse{
			Id:           productId,
			RatedCount:   rating.RatingCount,
			AverageScore: float64(rating.RatingSum) / float64(rating.RatingCount),
		}

		if err := stream.Send(res); err != nil {
			log.Println(err)
			return utils.NewGrpcErrorWithMetadata(
				codes.Aborted,
				fmt.Sprintf("Error rating product ID %s", productId),
				"STREAM_SEND_ERROR",
				err,
				nil,
			)
		}
	}
}

func (server *productServer) handleGetProductDBError(err error, id string) error {
	var appErr *utils.AppError
	if errors.As(err, &appErr) {
		if appErr.Key == utils.DB_RECORD_NOT_FOUND_CODE {
			return utils.NewGrpcErrorWithMetadata(
				codes.NotFound,
				fmt.Sprintf("Product with ID %s not found", id),
				utils.DB_RECORD_NOT_FOUND_CODE,
				err,
				map[string]string{
					"product_id": id,
				},
			)
		}
	}

	return utils.NewGrpcErrorWithMetadata(
		codes.Internal,
		"Failed to get product",
		"PRODUCT_DATABASE_ERROR",
		err,
		nil,
	)
}
