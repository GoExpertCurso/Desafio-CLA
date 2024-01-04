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

func (g *GetOrdersUseCase) ListOrders() ([]OrderOutputDTO, error) {
	orders, err := g.OrderRepository.ListOrders()
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
