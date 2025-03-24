package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	userv1 "order/gen/seminar/user/v1"
)

type UserClient interface {
	GetUser(ctx context.Context, userId string) (*userv1.User, error)
}

type userClient struct {
	conn          *grpc.ClientConn
	serviceClient userv1.UserServiceClient
}

func NewUserClient(url string) (*userClient, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := userv1.NewUserServiceClient(conn)
	return &userClient{conn, c}, nil
}

func (client *userClient) Close() {
	client.conn.Close()
}

func (client *userClient) GetUser(ctx context.Context, userId string) (*userv1.User, error) {
	request, err := client.serviceClient.GetUser(
		ctx,
		&userv1.GetUserRequest{Id: userId},
	)
	if err != nil {
		return nil, err
	}
	return &userv1.User{
		Id:   request.User.Id,
		Name: request.User.Name,
	}, nil
}
