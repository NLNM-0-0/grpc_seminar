package main

import (
	"context"
	"fmt"
	orderv1 "gateway/gen/seminar/order/v1"
	productv1 "gateway/gen/seminar/product/v1"
	userv1 "gateway/gen/seminar/user/v1"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

type appConfig struct {
	Env  string
	Port int

	ProductServiceEndpoint string
	UserServiceEndpoint    string
	OrderServiceEndpoint   string
}

func run(cfg *appConfig) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := productv1.RegisterProductServiceHandlerFromEndpoint(ctx, mux, cfg.ProductServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, cfg.UserServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = orderv1.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, cfg.OrderServiceEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	if err := run(cfg); err != nil {
		grpclog.Fatal(err)
	}
}

func loadConfig() (*appConfig, error) {
	portString := os.Getenv("GO_PORT")
	portNumber, err := strconv.Atoi(portString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return &appConfig{
		Env:                    os.Getenv("GO_ENV"),
		Port:                   portNumber,
		ProductServiceEndpoint: os.Getenv("PRODUCT_SERVICE_ENDPOINT"),
		UserServiceEndpoint:    os.Getenv("USER_SERVICE_ENDPOINT"),
		OrderServiceEndpoint:   os.Getenv("ORDER_SERVICE_ENDPOINT"),
	}, nil
}
