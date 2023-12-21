package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetOrders() ([]Order, error)
	GetOrder(id string) (Order, error)
	// GetTotal() (int, error)
}
