package server

import (
	"context"

	"github.com/FadyGamilH/ordermangment/services/common/protos/orders"
	"github.com/FadyGamilH/ordermangment/services/orders/service"
	"google.golang.org/grpc"
)

// the impl of RPCs
type OrderGrpcServer struct {
	// inject business (service layer)
	orderService *service.OrderService

	// compose this provided interface from protobuf generated code
	// so now we have to implement the RPCs
	orders.UnimplementedOrderServiceServer
}

// this is the grpc service server
// receive the grpc server
// receive the services (business) to be injected
func NewGrpcServer(grpc *grpc.Server, orderService service.OrderService) {
	orderGrpcServer := &OrderGrpcServer{
		orderService: &orderService,
	}

	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, orderGrpcServer)
}

func (ogs *OrderGrpcServer) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := orders.Order{
		OrderId:    1,
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Qty:        req.Qty,
	}

	// call service layer
	err := ogs.orderService.CreateOrder(ctx, &order)
	if err != nil {
		return nil, err
	}
	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}
