package service

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/repository"
)

type OrderService interface {
	OrderExtensionService
}

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func ProvideOrderServiceImpl(orderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
	}
}
