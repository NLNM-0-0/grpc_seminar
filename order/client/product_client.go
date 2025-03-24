package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	productv1 "order/gen/seminar/product/v1"
)

type ProductClient interface {
	GetProductByIds(
		ctx context.Context,
		ids []string,
	) (productv1.ProductService_GetProductByIdsClient, error)
	RateProductByIds(ctx context.Context) (productv1.ProductService_RateProductByIdsClient, error)
}

type productClient struct {
	conn          *grpc.ClientConn
	serviceClient productv1.ProductServiceClient
}

func NewProductClient(url string) (*productClient, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := productv1.NewProductServiceClient(conn)
	return &productClient{conn, c}, nil
}

func (c *productClient) Close() {
	c.conn.Close()
}

func (c *productClient) GetProductByIds(
	ctx context.Context,
	ids []string,
) (productv1.ProductService_GetProductByIdsClient, error) {
	request := &productv1.GetProductsByIdsRequest{Ids: ids}
	stream, err := c.serviceClient.GetProductByIds(ctx, request)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *productClient) RateProductByIds(ctx context.Context) (productv1.ProductService_RateProductByIdsClient, error) {
	stream, err := c.serviceClient.RateProductByIds(ctx)
	if err != nil {
		return nil, err
	}
	return stream, nil
}
