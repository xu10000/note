// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	proto/api.proto

It has these top-level messages:
	CallRequest
	CallResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "example"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Example.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) error {
	type example interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type Example struct {
		example
	}
	h := &exampleHandler{hdlr}
	return s.Handle(s.NewHandler(&Example{h}, opts...))
}

type exampleHandler struct {
	ExampleHandler
}

func (h *exampleHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.ExampleHandler.Call(ctx, in, out)
}
