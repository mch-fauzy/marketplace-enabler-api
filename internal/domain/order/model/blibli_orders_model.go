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

// type BlibliOrders struct {
// 	ID              int32       `db:"id"`
// 	Brand           null.String `db:"brand"`
// 	Store           null.String `db:"store"`
// 	OrderNo         null.String `db:"order_no"`
// 	OrderStatus     null.String `db:"order_status"`
// 	SKUCode         null.String `db:"sku_code"`
// 	TotalQuantity   null.Int    `db:"total_quantity"`
// 	ProductPrice    null.Int    `db:"product_price"`
// 	OrderDate       null.Time   `db:"order_date"`
// 	NameProduct     null.String `db:"name_product"`
// 	OrderItemNo     null.String `db:"order_item_no"`
// 	BuyerName       null.String `db:"buyer_name"`
// 	PacketNo        null.String `db:"packet_no"`
// 	AWBNo           null.String `db:"awb_no"`
// 	PickupPointCode null.String `db:"pickup_point_code"`
// 	CreatedAt       null.Time   `db:"created_at"`
// 	UpdatedAt       null.Time   `db:"updated_at"`
// 	CreatedBy       null.String `db:"created_by"`
// 	UpdatedBy       null.String `db:"updated_by"`
// }

// func (b BlibliOrders) GetFields() []interface{} {
// 	return []interface{}{
// 		b.ID,
// 		b.Brand,
// 		b.Store,
// 		b.OrderNo,
// 		b.OrderStatus,
// 		b.SKUCode,
// 		b.TotalQuantity,
// 		b.ProductPrice,
// 		b.OrderDate,
// 		b.NameProduct,
// 		b.OrderItemNo,
// 		b.BuyerName,
// 		b.PacketNo,
// 		b.AWBNo,
// 		b.PickupPointCode,
// 		b.CreatedAt,
// 		b.UpdatedAt,
// 		b.CreatedBy,
// 		b.UpdatedBy,
// 	}
// }

// func (b BlibliBrands) GetFields() []interface{} {
// 	return []interface{}{
// 		b.Label,
// 		b.Value,
// 	}
// }

// type BlibliStores struct {
// 	Label null.String `db:"store"`
// 	Value null.String `db:"store"`
// }

// func (b BlibliStores) GetFields() []interface{} {
// 	return []interface{}{
// 		b.Label,
// 		b.Value,
// 	}
// }
