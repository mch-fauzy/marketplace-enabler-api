package model

import "github.com/guregu/null"

type ShopeeDownload struct {
	OrderNo               null.String `db:"order_no"`
	OrderStatus           null.String `db:"order_status"`
	ReceiptNo             null.String `db:"receipt_no"`
	NameProduct           null.String `db:"name_product"`
	Variation             null.String `db:"variation"`
	InitialPrice          null.Int    `db:"initial_price"`
	DiscountedPrice       null.Int    `db:"discounted_price"`
	ProductDiscountSeller null.Int    `db:"product_discount_seller"`
	TotalPayment          null.Int    `db:"total_payment"`
	Quantity              null.Int    `db:"quantity"`
}

type ShopeeDownloadList []*ShopeeDownload

func ShopeeNewOrdersDownloadList(shopeeDownloadList ShopeeDownloadList) OrdersDownloadList {
	return OrdersDownloadList{
		Shopee: shopeeDownloadList,
	}
}
