package dto

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/model"
)

type OrdersDownloadResponseList struct {
	Blibli BlibliDownloadResponseList `json:"blibli"`
	Shopee ShopeeDownloadResponseList `json:"shopee"`
}

// Public download dto converter for every marketplace
func ConvertDownloadResponses(ordersDownloadList model.OrdersDownloadList) OrdersDownloadResponseList {
	return OrdersDownloadResponseList{
		Blibli: convertBlibliDownloadResponses(ordersDownloadList),
		Shopee: convertShopeeDownloadResponses(ordersDownloadList),
	}
}

type OrdersBrandResponseList struct {
	Blibli BlibliBrandResponseList `json:"blibli"`
	Shopee ShopeeBrandResponseList `json:"shopee"`
}

// Public brand dto converter for every marketplace
func ConvertBrandResponses(ordersBrandList model.OrdersBrandList) OrdersBrandResponseList {
	return OrdersBrandResponseList{
		Blibli: convertBlibliBrandResponses(ordersBrandList),
		Shopee: convertShopeeBrandResponses(ordersBrandList),
	}
}

type OrdersStoreResponseList struct {
	Blibli BlibliStoreResponseList `json:"blibli"`
}

func ConvertStoreResponses(ordersStoreList model.OrdersStoreList) OrdersStoreResponseList {
	return OrdersStoreResponseList{
		Blibli: convertBlibliStoreResponses(ordersStoreList),
	}
}
