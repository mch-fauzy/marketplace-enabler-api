package model

import "time"

type DownloadOrdersByMarketFilter struct {
	Brand    string
	Store    string
	DateFrom time.Time
	DateTo   time.Time
}
