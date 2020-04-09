package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8001")
	if err != nil {
		panic(err.Error())
	}
	var req float64 = 3
	var resp float64
	err = client.Call("MathUtil.Cal", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp)
}
