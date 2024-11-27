package grpc

import (
	"fmt"
	"github.com/afanasjev/my-grpc-go-server/internal/port"
	"github.com/afanasjev/proto-repo/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	helloservice port.HelloServicePort
	port         int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloservice port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloservice: helloservice,
		port:         grpcPort,
	}
}

func (s *GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))

	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", s.port, err)
	}

	log.Printf("server is listening on port %v", s.port)

	grpcServer := grpc.NewServer()
	s.server = grpcServer
	hello.RegisterHelloServiceServer(grpcServer, s)

	if err := s.server.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %v: %v", s.port, err)
	}
}

func (s *GrpcAdapter) Stop() {
	s.server.GracefulStop()
}
