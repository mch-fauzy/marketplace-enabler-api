//go:build wireinject
// +build wireinject

package main

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/infras"
	orderRepository "github.com/evermos/boilerplate-go/internal/domain/order/repository"
	orderService "github.com/evermos/boilerplate-go/internal/domain/order/service"
	orderHandler "github.com/evermos/boilerplate-go/internal/handlers/order"
	"github.com/evermos/boilerplate-go/transport/http"
	"github.com/evermos/boilerplate-go/transport/http/router"
	"github.com/google/wire"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvideMySQLConn,
)

// Wiring for domain.
var domainOrder = wire.NewSet(
	// Service interface and implementation
	orderService.ProvideOrderServiceImpl,
	wire.Bind(new(orderService.OrderService), new(*orderService.OrderServiceImpl)),
	// Repository interface and implementation
	orderRepository.ProvideOrderRepositoryMySQL,
	wire.Bind(new(orderRepository.OrderRepository), new(*orderRepository.OrderRepositoryMySQL)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainOrder,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "OrderHandler"),
	orderHandler.ProvideOrderHandler,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
