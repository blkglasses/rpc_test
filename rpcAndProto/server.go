package main

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"rpc_test/rpcAndProto/message"
	"time"
)

type OrderService struct {
}

func (os *OrderService) GetOrderInfo(req message.OrderRequest, resp *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201987300001", OrderName: "衣服", OrderStatus: "已付款"},
		"281987310801": message.OrderInfo{OrderId: "201987310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	current := time.Now().Unix()
	if req.Timestamp > current {
		*resp = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[req.OrderId]
		if result.OrderId != "" {
			*resp = orderMap[req.OrderId]
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	orderservice := new(OrderService)
	err := rpc.Register(orderservice)
	if err != nil {
		panic(err.Error())
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
