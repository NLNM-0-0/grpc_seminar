package client

import (
	"context"
	"google.golang.org/grpc"
	productv1 "order/gen/seminar/product/v1"
)

type ProductClient interface {
	GetProductByIds(ctx context.Context, ids []string) ([]*productv1.Product, error)
}

type productClient struct {
	conn          *grpc.ClientConn
	serviceClient productv1.ProductServiceClient
}

func NewProductClient(url string) (*productClient, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := productv1.NewProductServiceClient(conn)
	return &productClient{conn, c}, nil
}

func (c *productClient) Close() {
	c.conn.Close()
}

func (c *productClient) GetProductByIds(ctx context.Context, ids []string) ([]*productv1.Product, error) {
	getProductByIdsResponse, err := c.serviceClient.GetProductByIds(
		ctx,
		&productv1.GetProductsByIdsRequest{
			Ids: ids,
		},
	)
	if err != nil {
		return nil, err
	}

	return getProductByIdsResponse.Products, nil
}
