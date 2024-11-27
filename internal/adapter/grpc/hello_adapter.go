package grpc

import (
	"context"
	"github.com/afanasjev/proto-repo/protogen/go/hello"
)

//import (
// "github.com/afanasjev/proto-repo/protogen/go/hello"
//)

func (a GrpcAdapter) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	greet := a.helloservice.GenerateHello(request.Name)

	return &hello.HelloResponse{Greet: greet}, nil
}
