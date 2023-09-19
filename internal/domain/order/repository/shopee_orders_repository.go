package repository

import (
	"context"
	"fmt"

	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	selectShopeeDownload = `
	SELECT 
		order_no,
		order_status,
		receipt_no,
		name_product,
		variation,
		initial_price,
		discounted_price,
		product_discount_seller,
		total_payment,
		quantity
	FROM 
		shopee_orders
	`
	whereShopeeBrandAndStore = `
	WHERE brand = ? AND store = ?
	`
	limitShopeeRow = `
	ORDER BY order_created_date DESC LIMIT ?
	`
	optionalShopeeDateRange = `
	AND order_created_date BETWEEN ? AND ?
	`
	selectShopeeUniqueBrand = `
	SELECT DISTINCT
		brand
	FROM 
		shopee_orders
	`
)

type ShopeeOrdersRepository interface {
	DownloadShopeeOrders(downloadCtx context.Context, downloadFilter model.DownloadOrdersByMarketFilter) (model.OrdersDownloadList, error)
	GetShopeeBrands(brandCtx context.Context) (model.OrdersBrandList, error)
}

func (r *OrderRepositoryMySQL) DownloadShopeeOrders(downloadCtx context.Context, downloadFilter model.DownloadOrdersByMarketFilter) (model.OrdersDownloadList, error) {
	query := fmt.Sprintf(selectShopeeDownload)

	var args []interface{}
	query += whereShopeeBrandAndStore
	args = append(args, downloadFilter.Brand, downloadFilter.Store)

	// (Optional Query) Check if datetime is not equal to 0001-01-01 00:00:00 +0000 UTC or ""
	if !downloadFilter.DateFrom.IsZero() && !downloadFilter.DateTo.IsZero() {
		query += optionalShopeeDateRange
		args = append(args, downloadFilter.DateFrom, downloadFilter.DateTo)
	}

	query += limitShopeeRow
	args = append(args, r.CFG.DB.RowLimit)

	var shopeeDownloadList model.ShopeeDownloadList
	err := r.DB.Read.Select(&shopeeDownloadList, query, args...)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[DownloadShopeeOrders] Failed to get shopee orders")
		err = failure.InternalError(err)
		return model.OrdersDownloadList{}, err
	}

	return model.ShopeeNewOrdersDownloadList(shopeeDownloadList), nil
}

func (r *OrderRepositoryMySQL) GetShopeeBrands(brandCtx context.Context) (model.OrdersBrandList, error) {
	query := fmt.Sprintf(selectShopeeUniqueBrand)

	var shopeeBrandList model.ShopeeBrandList
	err := r.DB.Read.Select(&shopeeBrandList, query)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[GetShopeeBrands] Failed to get shopee brands")
		err = failure.InternalError(err)
		return model.OrdersBrandList{}, err
	}

	return model.ShopeeNewOrdersBrandList(shopeeBrandList), nil
}
