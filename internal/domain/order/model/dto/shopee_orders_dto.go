package dto

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/guregu/null"
)

type ShopeeDownloadResponse struct {
	OrderNo               null.String `json:"orderNo"`
	OrderStatus           null.String `json:"orderStatus"`
	ReceiptNo             null.String `json:"receiptNo"`
	NameProduct           null.String `json:"nameProduct"`
	Variation             null.String `json:"variation"`
	InitialPrice          null.Int    `json:"initialPrice"`
	DiscountedPrice       null.Int    `json:"discountedPrice"`
	ProductDiscountSeller null.Int    `json:"productDiscountSeller"`
	TotalPayment          null.Int    `json:"totalPayment"`
	Quantity              null.Int    `json:"quantity"`
}

func mapShopeeDownloadResponses(shopeeDownload model.ShopeeDownload) ShopeeDownloadResponse {
	return ShopeeDownloadResponse{
		OrderNo:               shopeeDownload.OrderNo,
		OrderStatus:           shopeeDownload.OrderStatus,
		ReceiptNo:             shopeeDownload.ReceiptNo,
		NameProduct:           shopeeDownload.NameProduct,
		Variation:             shopeeDownload.Variation,
		InitialPrice:          shopeeDownload.InitialPrice,
		DiscountedPrice:       shopeeDownload.DiscountedPrice,
		ProductDiscountSeller: shopeeDownload.ProductDiscountSeller,
		TotalPayment:          shopeeDownload.TotalPayment,
		Quantity:              shopeeDownload.Quantity,
	}
}

type ShopeeDownloadResponseList []*ShopeeDownloadResponse

func convertShopeeDownloadResponses(ordersDownloadList model.OrdersDownloadList) ShopeeDownloadResponseList {
	var shopeeDownloadResponseList ShopeeDownloadResponseList = ShopeeDownloadResponseList{}

	if len(ordersDownloadList.Shopee) > 0 {
		for _, shopeeDownload := range ordersDownloadList.Shopee {
			shopeeDownloadResponse := mapShopeeDownloadResponses(*shopeeDownload)
			shopeeDownloadResponseList = append(shopeeDownloadResponseList, &shopeeDownloadResponse)
		}
	}

	return shopeeDownloadResponseList
}

type ShopeeBrandResponse struct {
	Brand null.String `json:"brand"`
}

// Shopee brand dto mapper
func mapShopeeBrandResponses(shopeeBrand model.ShopeeBrand) ShopeeBrandResponse {
	return ShopeeBrandResponse{
		Brand: shopeeBrand.Brand,
	}
}

type ShopeeBrandResponseList []*ShopeeBrandResponse

// Shopee brand dto converter
func convertShopeeBrandResponses(ordersBrandList model.OrdersBrandList) ShopeeBrandResponseList {
	var shopeeBrandResponseList ShopeeBrandResponseList = ShopeeBrandResponseList{}

	if len(ordersBrandList.Shopee) > 0 {
		for _, shopeeBrand := range ordersBrandList.Shopee {
			shopeeBrandResponse := mapShopeeBrandResponses(*shopeeBrand)
			shopeeBrandResponseList = append(shopeeBrandResponseList, &shopeeBrandResponse)
		}
	}
	return shopeeBrandResponseList
}
