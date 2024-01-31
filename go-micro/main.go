package main

import (
	"context"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Address(":8080"),
	)
	service.Init()
	service.Run()
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Helloworld struct{}

func (h *Helloworld) Greeting(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}
