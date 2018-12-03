package main

import (
	"context"
	"fmt"
	"log"
	proto "rpc/proto"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
)

func cli() {
	time.Sleep(2 * time.Second)
	example := proto.NewExampleService("example", nil)

	rsp, err := example.Call(context.TODO(), &proto.CallRequest{Name: "xiaozhongting"})

	if err != nil {
		log.Fatal("xxxxxxxxxxx", err)
	}

	fmt.Printf("rsp messge: %+v", rsp)
}

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Received Example Call request")

	if len(req.Name) == 0 {
		return errors.BadRequest("go micro api example xzt", "no content xzt")
	}

	rsp.Message = "go your request" + req.Name

	return nil
}

func run_server() {
	service := micro.NewService(
		micro.Name("example"),
	)

	server.Init(
		server.Address("127.0.0.1:3003"),
	)
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go cli()
	run_server()
}
