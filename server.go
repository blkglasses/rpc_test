package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

func (mu *MathUtil) Cal(req float64, resp *float64) error {
	*resp = math.Pi * req * req
	return nil
}
func main() {
	mathUtil := new(MathUtil)
	err := rpc.Register(mathUtil)
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
