package model

import (
	"github.com/guregu/null"
)

type BlibliDownload struct {
	OrderNo       null.String `db:"order_no"`
	OrderStatus   null.String `db:"order_status"`
	AWBNo         null.String `db:"awb_no"`
	NameProduct   null.String `db:"name_product"`
	SKUCode       null.String `db:"sku_code"`
	ProductPrice  null.Int    `db:"product_price"`
	TotalQuantity null.Int    `db:"total_quantity"`
}

type BlibliDownloadList []*BlibliDownload

// Map List of Blibli Download into List of Orders Download
func BlibliNewOrdersDownloadList(blibliDownloadList BlibliDownloadList) OrdersDownloadList {
	return OrdersDownloadList{
		Blibli: blibliDownloadList,
	}
}

type BlibliBrand struct {
	Brand null.String `db:"brand"`
}

type BlibliBrandList []*BlibliBrand

func BlibliNewOrdersBrandList(blibliBrandList BlibliBrandList) OrdersBrandList {
	return OrdersBrandList{
		Blibli: blibliBrandList,
	}
}

type BlibliStore struct {
	Store null.String `db:"store"`
}

type BlibliStoreList []*BlibliStore

func BlibliNewOrdersStoreList(blibliStoreList BlibliStoreList) OrdersStoreList {
	return OrdersStoreList{
		Blibli: blibliStoreList,
	}
}
