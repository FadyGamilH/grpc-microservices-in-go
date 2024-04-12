package service

import (
	"context"

	"github.com/FadyGamilH/ordermangment/services/common/protos/orders"
)

type OrderService struct {
	// inject repo layer here
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (os *OrderService) CreateOrder(ctx context.Context, req *orders.Order) error {
	return nil
}
