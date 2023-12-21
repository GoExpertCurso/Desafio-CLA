package usecase

import (
	"github.com/GoExpertCurso/Desafio-CLA/internal/entity"
)

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}
func (g *GetOrdersUseCase) GetOrder(id string) (OrderOutputDTO, error) {
	order, err := g.OrderRepository.GetOrder(id)
	if err != nil {
		return OrderOutputDTO{}, err
	}

	return OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil 
}

func (g *GetOrdersUseCase) GetOrders() ([]OrderOutputDTO, error) {
	orders, err := g.OrderRepository.GetOrders()
	if err != nil {
		return nil, err
	}
	dto := make([]OrderOutputDTO, len(orders))
	for i, order := range orders {
		item := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dto[i] = item
	}
	return dto, nil
}