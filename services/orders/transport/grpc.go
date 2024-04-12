package transport

import (
	"log"
	"net"

	"github.com/FadyGamilH/ordermangment/services/orders/server"
	"github.com/FadyGamilH/ordermangment/services/orders/service"
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

func (grpcSrv *gRPCServer) Run() error {
	// define a tcp connection to send it to the serve() api provided by the grpc pkg
	listener, err := net.Listen("tcp", grpcSrv.addr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err.Error())
	}
	// use the grpc pkg to get a new server instance
	grpcServer := grpc.NewServer()

	// register our grpc service (the business logic)
	// define dependeices
	orderService := service.NewOrderService()
	server.NewGrpcServer(grpcServer, *orderService)

	// log for running
	log.Println("grpc server is running on : ", grpcSrv.addr)

	// call Serve() and sned on tcp connection listener
	return grpcServer.Serve(listener)
}
