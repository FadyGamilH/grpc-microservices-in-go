package transport

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGrpcServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (grpcSrv *gRPCServer)  Run() error {
	// define a tcp connection to send it to the serve() api provided by the grpc pkg
	listener, err := net.Listen("tcp", grpcSrv.addr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err.Error())
	}
	// use the grpc pkg to get a new server instance
	server := grpc.NewServer()

	// register our grpc service (the impl of RPCs / the business logic)

	// call Serve() and sned on tcp connection listener
	return server.Serve(listener)
}
