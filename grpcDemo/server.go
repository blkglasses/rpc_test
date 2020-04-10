package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rpc_test/grpcDemo/message"
	"time"
)

type OrderServiceImpl struct {
}

func (osi *OrderServiceImpl) GetOrderInfo(c context.Context, req *message.OrderRequest) (resp *message.OrderInfo, err error) {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201987300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201987310001": message.OrderInfo{OrderId: "201987310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	resp = new(message.OrderInfo)
	current := time.Now().Unix()
	fmt.Println(current)
	if req.Timestamp > current {
		*resp = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		fmt.Println(req)
		result := orderMap[req.OrderId]
		fmt.Println(orderMap[req.OrderId])
		if result.OrderId != "" {
			fmt.Println(result)
			*resp = orderMap[req.OrderId]
			return
		} else {
			return nil, errors.New("server error")
		}
	}
	return
}
func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
