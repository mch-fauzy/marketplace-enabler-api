package order

import (
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/order/model/dto"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

// Download orders by marketplace with optional date range filter
// @Summary Download Orders By Marketplace
// @Description Get a list of orders field to be downloaded
// @Tags orders
// @Produce json
// @Param market path string true "Market"
// @Param brand query string true "Brand"
// @Param store query string true "Store"
// @Param dateFrom query string false "From (timestamp)" Format(datetime)
// @Param dateTo query string false "To (timestamp)" Format(datetime)
// @Success 200 {object} response.Base(data=dto.OrdersDownloadResponseList)
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/orders/{market}/download [get]
func (h *OrderHandler) DownloadOrdersByMarket(w http.ResponseWriter, r *http.Request) {
	marketplace := chi.URLParam(r, shared.MarketPathField)

	brandName := r.URL.Query().Get(shared.BrandNameQueryField)
	storeName := r.URL.Query().Get(shared.StoreNameQueryField)
	dateFromStr := r.URL.Query().Get(shared.StartDateQueryField)
	dateToStr := r.URL.Query().Get(shared.EndDateQueryField)
	downloadFilter := dto.NewDownloadOrdersByMarketFilterRequests(marketplace, brandName, storeName, dateFromStr, dateToStr)

	err := downloadFilter.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	downloadCtx := r.Context()
	download, err := h.OrderService.DownloadOrdersByMarket(downloadCtx, downloadFilter)
	if err != nil {
		log.Warn().
			Err(err).
			Msg("[DownloadOrdersByMarket] Failed to download orders")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, download)
}

// Get list of unique brands by marketplace
// @Summary Get Brands By Marketplace
// @Description Get a list of unique brands by marketplace
// @Tags orders
// @Produce json
// @Param market path string true "Market"
// @Success 200 {object} response.Base(data=dto.OrdersBrandResponseList)
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/orders/{market}/brand [get]
func (h *OrderHandler) GetBrandsByMarket(w http.ResponseWriter, r *http.Request) {
	marketplace := chi.URLParam(r, shared.MarketPathField)
	brandFilter := dto.NewGetBrandsByMarketFilterRequests(marketplace)

	err := brandFilter.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	brandCtx := r.Context()
	brand, err := h.OrderService.GetBrandsByMarket(brandCtx, brandFilter)
	if err != nil {
		log.Warn().
			Err(err).
			Msg("[GetBrandsByMarket] Failed to get brands")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, brand)
}

// Get list of unique stores by marketplace
// @Summary Get Stores By Marketplace
// @Description Get a list of unique stores by marketplace
// @Tags orders
// @Produce json
// @Param market path string true "Market"
// @Success 200 {object} response.Base(data=dto.OrdersStoreResponseList)
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/orders/{market}/store [get]
func (h *OrderHandler) GetStoresByMarket(w http.ResponseWriter, r *http.Request) {
	marketplace := chi.URLParam(r, shared.MarketPathField)
	storeFilter := dto.NewGetStoresByMarketFilterRequests(marketplace)

	err := storeFilter.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	storeCtx := r.Context()
	store, err := h.OrderService.GetStoresByMarket(storeCtx, storeFilter)
	if err != nil {
		log.Warn().
			Err(err).
			Msg("[GetStoresByMarket] Failed to get brands")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, store)
}
