package business

type OrderService struct {
	// inject repo layer here
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
