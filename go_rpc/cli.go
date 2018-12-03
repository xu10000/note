package main

import (
	"context"
	"log"
	proto "rpc/proto"

	"github.com/micro/go-micro/errors"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Received Example Call request")

	if len(req.Name) == 0 {
		return errors.BadRequest("go micro api example xzt", "no content xzt")
	}

	rsp.Message = "go your request" + req.Name

	return nil
}

func main() {

	cli()
}
