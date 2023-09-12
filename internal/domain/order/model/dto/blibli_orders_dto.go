package dto

import (
	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/guregu/null"
)

type BlibliDownloadResponse struct {
	OrderNo       null.String `json:"orderNo"`
	OrderStatus   null.String `json:"orderStatus"`
	AWBNo         null.String `json:"awbNo"`
	NameProduct   null.String `json:"nameProduct"`
	SKUCode       null.String `json:"skuCode"`
	ProductPrice  null.Int    `json:"productPrice"`
	TotalQuantity null.Int    `json:"totalQuantity"`
}

// Map dto Response of Blibli Download based on model Blibli Download
func mapBlibliDownloadResponses(blibliDownload model.BlibliDownload) BlibliDownloadResponse {
	return BlibliDownloadResponse{
		OrderNo:       blibliDownload.OrderNo,
		OrderStatus:   blibliDownload.OrderStatus,
		AWBNo:         blibliDownload.AWBNo,
		NameProduct:   blibliDownload.NameProduct,
		SKUCode:       blibliDownload.SKUCode,
		ProductPrice:  blibliDownload.ProductPrice,
		TotalQuantity: blibliDownload.TotalQuantity,
	}
}

// Response List of Blibli Orders Download
type BlibliDownloadResponseList []*BlibliDownloadResponse

// convert model List of Orders Download into dto Response List of Orders for specific marketplace (in this case Blibli)
func convertBlibliDownloadResponses(ordersDownloadList model.OrdersDownloadList) BlibliDownloadResponseList {
	var blibliDownloadResponseList BlibliDownloadResponseList = BlibliDownloadResponseList{}

	if len(ordersDownloadList.Blibli) > 0 {
		for _, blibliDownload := range ordersDownloadList.Blibli {
			blibliDownloadResponse := mapBlibliDownloadResponses(*blibliDownload)
			blibliDownloadResponseList = append(blibliDownloadResponseList, &blibliDownloadResponse)
		}
	}

	return blibliDownloadResponseList
}

type BlibliBrandResponse struct {
	Brand null.String `json:"brand"`
}

// Blibli brand dto mapper
func mapBlibliBrandResponses(blibliBrand model.BlibliBrand) BlibliBrandResponse {
	return BlibliBrandResponse{
		Brand: blibliBrand.Brand,
	}
}

type BlibliBrandResponseList []*BlibliBrandResponse

// Blibli brand dto converter
func convertBlibliBrandResponses(ordersBrandList model.OrdersBrandList) BlibliBrandResponseList {
	var blibliBrandResponseList BlibliBrandResponseList = BlibliBrandResponseList{}

	if len(ordersBrandList.Blibli) > 0 {
		for _, blibliBrand := range ordersBrandList.Blibli {
			blibliBrandResponse := mapBlibliBrandResponses(*blibliBrand)
			blibliBrandResponseList = append(blibliBrandResponseList, &blibliBrandResponse)
		}
	}

	return blibliBrandResponseList
}
