package dto

import (
	"fmt"
	"time"

	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
)

type DownloadOrdersByMarketFilterRequest struct {
	Marketplace string
	Brand       string
	Store       string
	DateFromStr string
	DateToStr   string
	DateFrom    time.Time
	DateTo      time.Time
}

func (f *DownloadOrdersByMarketFilterRequest) Validate() error {
	var err error
	validator := shared.GetValidator()

	//Parse and validate dates in DownloadOrdersByMarketFilterRequest struct, "" will be defaulted into 0001-01-01 00:00:00 +0000 UTC
	if f.DateFromStr != "" && f.DateToStr != "" {
		f.DateFrom, err = time.Parse(shared.DefaultDateFormat, f.DateFromStr)
		if err != nil {
			return failure.BadRequestFromString(fmt.Sprintf("Invalid dateFrom parameter %s", err))
		}

		f.DateTo, err = time.Parse(shared.DefaultDateFormat, f.DateToStr)
		if err != nil {
			return failure.BadRequestFromString(fmt.Sprintf("Invalid dateTo parameter %s", err))
		}

		if f.DateTo.Before(f.DateFrom) {
			return failure.BadRequestFromString("DateTo must be after DateFrom")
		}
	}
	return validator.Struct(f)
}

// Map field of DownloadOrdersByMarketFilterRequest based on the input (for init filter request variable)
func NewDownloadOrdersByMarketFilterRequests(marketplace, brand, store, dateFromStr, dateToStr string) DownloadOrdersByMarketFilterRequest {
	return DownloadOrdersByMarketFilterRequest{
		Marketplace: marketplace,
		Brand:       brand,
		Store:       store,
		DateFromStr: dateFromStr,
		DateToStr:   dateToStr,
	}
}

// Convert dto of Filter Request into model of Filter to interact with repo/db
func ConvertDownloadOrdersByMarketFilterRequests(downloadOrdersByMarketFilterRequest DownloadOrdersByMarketFilterRequest) model.DownloadOrdersByMarketFilter {
	return model.DownloadOrdersByMarketFilter{
		Brand:    downloadOrdersByMarketFilterRequest.Brand,
		Store:    downloadOrdersByMarketFilterRequest.Store,
		DateFrom: downloadOrdersByMarketFilterRequest.DateFrom,
		DateTo:   downloadOrdersByMarketFilterRequest.DateTo,
	}
}

type BrandOrdersByMarketFilterRequest struct {
	Marketplace string
}

func (f *BrandOrdersByMarketFilterRequest) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(f)
}

func NewBrandOrdersByMarketFilterRequests(marketplace string) BrandOrdersByMarketFilterRequest {
	return BrandOrdersByMarketFilterRequest{
		Marketplace: marketplace,
	}
}
