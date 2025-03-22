package server

import (
	"context"
	"fmt"
	"log"
	"net"
	productv1 "product/gen/seminar/product/v1"
	"product/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		return nil, err
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
		return nil, err
	}
	return &productv1.GetProductResponse{
		Product: &productv1.Product{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		},
	}, nil
}

func (server *productServer) GetProductByIds(ctx context.Context, request *productv1.GetProductsByIdsRequest) (*productv1.GetProductsByIdsResponse, error) {
	products, err := server.productService.GetProductByIds(ctx, request.Ids)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var productResponses []*productv1.Product
	for _, product := range products {
		productResponses = append(productResponses, &productv1.Product{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return &productv1.GetProductsByIdsResponse{
		Products: productResponses,
	}, nil
}
