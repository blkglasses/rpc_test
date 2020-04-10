package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"rpc_test/grpcDemo/message"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	req := message.OrderRequest{OrderId: "201987310001", Timestamp: time.Now().Unix()}
	client := message.NewOrderServiceClient(conn)
	resp, err := client.GetOrderInfo(context.Background(), &req)
	if err != nil {
		panic(err.Error())
	}
	if resp != nil {
		fmt.Println(resp)
	}
}
