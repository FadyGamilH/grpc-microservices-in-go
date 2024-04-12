package main

import (
	"log"

	"github.com/FadyGamilH/ordermangment/services/orders/transport"
)

func main() {
	grpcServer := transport.NewGrpcServer(":50051")
	err := grpcServer.Run()
	log.Fatal(err)
}
