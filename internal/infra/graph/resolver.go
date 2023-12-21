package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/GoExpertCurso/Desafio-CLA/internal/usecase"
)

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrdersUseCase usecase.GetOrdersUseCase
}
