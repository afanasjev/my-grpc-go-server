package main

import (
	"github.com/afanasjev/my-grpc-go-server/internal/adapter/grpc"
	"github.com/afanasjev/my-grpc-go-server/internal/application"
)

func main() {
	hs := &application.HelloService{}
	grpcAdapter := grpc.NewGrpcAdapter(hs, 9099)
	grpcAdapter.Run()
}
