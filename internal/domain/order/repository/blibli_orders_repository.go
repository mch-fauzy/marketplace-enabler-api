package repository

import (
	"context"
	"fmt"

	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	selectBlibliDownload = `
	SELECT 
		order_no,
		order_status,
		awb_no,
		name_product,
		sku_code,
		product_price,
		total_quantity
	FROM 
		blibli_orders
	`
	whereBlibliBrandAndStore = `
	WHERE brand = ? AND store = ?
	`
	limitBlibliRow = `
	ORDER BY order_date DESC LIMIT ?
	`
	optionalBlibliDateRange = `
	AND order_date BETWEEN ? AND ?
	`
	selectBlibliUniqueBrand = `
	SELECT DISTINCT
		brand
	FROM 
		blibli_orders
	`
	selectBlibliUniqueStore = `
	SELECT DISTINCT
		store
	FROM
		blibli_orders
	`
)

type BlibliOrdersRepository interface {
	DownloadBlibliOrders(downloadCtx context.Context, downloadFilter model.DownloadOrdersByMarketFilter) (model.OrdersDownloadList, error)
	GetBlibliBrands(brandCtx context.Context) (model.OrdersBrandList, error)
	GetBlibliStores(storeCtx context.Context) (model.OrdersStoreList, error)
}

func (r *OrderRepositoryMySQL) DownloadBlibliOrders(downloadCtx context.Context, downloadFilter model.DownloadOrdersByMarketFilter) (model.OrdersDownloadList, error) {
	query := fmt.Sprintf(selectBlibliDownload)

	var args []interface{}
	query += whereBlibliBrandAndStore
	args = append(args, downloadFilter.Brand, downloadFilter.Store)

	// (Optional Query)
	if !downloadFilter.DateFrom.IsZero() && !downloadFilter.DateTo.IsZero() {
		query += optionalBlibliDateRange
		args = append(args, downloadFilter.DateFrom, downloadFilter.DateTo)
	}

	query += limitBlibliRow
	args = append(args, r.CFG.DB.RowLimit)

	var blibliDownloadList model.BlibliDownloadList
	err := r.DB.Read.Select(&blibliDownloadList, query, args...)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[DownloadBlibliOrders] Failed to get blibli orders")
		err = failure.InternalError(err)
		return model.OrdersDownloadList{}, err
	}

	return model.BlibliNewOrdersDownloadList(blibliDownloadList), nil
}

func (r *OrderRepositoryMySQL) GetBlibliBrands(brandCtx context.Context) (model.OrdersBrandList, error) {
	query := fmt.Sprintf(selectBlibliUniqueBrand)

	var blibliBrandList model.BlibliBrandList
	err := r.DB.Read.Select(&blibliBrandList, query)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[GetBlibliBrands] Failed to get blibli brands")
		err = failure.InternalError(err)
		return model.OrdersBrandList{}, err
	}

	return model.BlibliNewOrdersBrandList(blibliBrandList), nil
}

func (r *OrderRepositoryMySQL) GetBlibliStores(storeCtx context.Context) (model.OrdersStoreList, error) {
	query := fmt.Sprintf(selectBlibliUniqueStore)

	var blibliStoreList model.BlibliStoreList
	err := r.DB.Read.Select(&blibliStoreList, query)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[GetBlibliStores] Failed to get blibli stores")
		err = failure.InternalError(err)
		return model.OrdersStoreList{}, nil
	}
	return model.BlibliNewOrdersStoreList(blibliStoreList), nil
}
