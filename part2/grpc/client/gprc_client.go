package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"tohrux.com/grpc/service"
)

func main() {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务端出错")
	}
	defer conn.Close()

	proClient := service.NewProdServiceClient(conn)

	request := &service.ProductRequest{
		ProdId: 123,
	}

	stockResponse, err := proClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("出错", err)
	}
	fmt.Println(stockResponse.ProdStock)
}
