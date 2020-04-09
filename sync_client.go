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
	var req float64 = 5
	var resp float64
	call := client.Go("MathUtil.Cal", req, &resp, nil)
	replayDone := <-call.Done
	fmt.Println(replayDone)
	fmt.Println(resp)
}
