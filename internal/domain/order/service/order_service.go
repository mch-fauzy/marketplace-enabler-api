package service

import (
	"context"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/order/model"
	"github.com/evermos/boilerplate-go/internal/domain/order/model/dto"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/rs/zerolog/log"
)

type OrderExtensionService interface {
	DownloadOrdersByMarket(downloadCtx context.Context, downloadFilter dto.DownloadOrdersByMarketFilterRequest) (dto.OrdersDownloadResponseList, error)
	GetBrandsByMarket(brandCtx context.Context, brandFilter dto.BrandOrdersByMarketFilterRequest) (dto.OrdersBrandResponseList, error)
}

func (s *OrderServiceImpl) DownloadOrdersByMarket(downloadCtx context.Context, downloadFilter dto.DownloadOrdersByMarketFilterRequest) (dto.OrdersDownloadResponseList, error) {
	var downloadResults model.OrdersDownloadList
	var err error

	downloadQueryFilter := dto.ConvertDownloadOrdersByMarketFilterRequests(downloadFilter)
	switch downloadFilter.Marketplace {
	case shared.SHOPEE:
		downloadResults, err = s.OrderRepository.DownloadShopeeOrders(downloadCtx, downloadQueryFilter)
	case shared.BLIBLI:
		downloadResults, err = s.OrderRepository.DownloadBlibliOrders(downloadCtx, downloadQueryFilter)
	default:
		return dto.OrdersDownloadResponseList{}, nil
	}

	if err != nil {
		if failure.GetCode(err) == http.StatusInternalServerError {
			log.Error().
				Err(err).
				Msg("[DownloadOrdersByMarket] Internal server error occurred")
			err = failure.InternalError(err)
			return dto.OrdersDownloadResponseList{}, err
		}
		log.Warn().
			Msg("[DownloadOrdersByMarket] Error occurred")
		return dto.OrdersDownloadResponseList{}, err
	}

	return dto.ConvertDownloadResponses(downloadResults), nil
}

func (s *OrderServiceImpl) GetBrandsByMarket(brandCtx context.Context, brandFilter dto.BrandOrdersByMarketFilterRequest) (dto.OrdersBrandResponseList, error) {
	var brandResults model.OrdersBrandList
	var err error

	switch brandFilter.Marketplace {
	case shared.SHOPEE:
		brandResults, err = s.OrderRepository.GetShopeeBrands(brandCtx)
	case shared.BLIBLI:
		brandResults, err = s.OrderRepository.GetBlibliBrands(brandCtx)
	default:
		return dto.OrdersBrandResponseList{}, nil
	}

	if err != nil {
		if failure.GetCode(err) == http.StatusInternalServerError {
			log.Error().
				Err(err).
				Msg("[GetBrandsByMarket] Internal server error occurred")
			err = failure.InternalError(err)
			return dto.OrdersBrandResponseList{}, err
		}
		log.Warn().
			Msg("[GetBrandsByMarket] Error occurred")
		return dto.OrdersBrandResponseList{}, err
	}

	return dto.ConvertBrandResponses(brandResults), nil
}
