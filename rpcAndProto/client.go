package main

import (
	"fmt"
	"net/rpc"
	"rpc_test/rpcAndProto/message"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8001")
	if err != nil {
		panic(err.Error())
	}
	request := message.OrderRequest{OrderId: "201907300001", Timestamp: time.Now().Unix()}
	resp := new(message.OrderInfo)
	err = client.Call("OrderService.GetOrderInfo", request, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp)
}
