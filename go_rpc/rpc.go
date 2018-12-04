package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	proto "rpc/proto"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	go_grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"google.golang.org/grpc"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:3003", "go.micro.srv.greeter address")
)

func cli() {
	time.Sleep(2 * time.Second)
	service := go_grpc.NewService()
	service.Init()

	example := proto.NewExampleService("example", service.Client())
	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	for i := 0; i < 100; i++ {
		rsp, err := example.Call(ctx, &proto.CallRequest{Name: "xiaozhongting"})

		if err != nil {
			log.Fatal("xxxxxxxxxxx", err)
		}

		fmt.Printf("rsp messge: %+v", rsp)
	}
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

	opts := []micro.Option{
		func(opt *micro.Options) {
			opt.Server.Init(
				server.Address("127.0.0.1:3003"),
				server.Name("example"),
			)
		},
	}

	service := go_grpc.NewService(
		opts...,
	)

	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterGWExampleHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func run_gateway() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Errorf("xxxxxxxxxxxx: ", err)
	}
}

func main() {
	go cli()
	go run_gateway()
	run_server()
}
