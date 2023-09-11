package order

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/service"
	"github.com/go-chi/chi"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func ProvideOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{
		OrderService: orderService,
	}
}

// Router sets up the router for this domain.
func (h *OrderHandler) Router(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Get("/orders/{market}/download", h.DownloadOrdersByMarket)
		r.Get("/orders/{market}/brand", h.GetBrandsByMarket)
	})
}
