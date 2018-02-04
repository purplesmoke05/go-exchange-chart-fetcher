package api

import (
	"github.com/airking05/go-exchange-chart-fetcher/models"
)

type ExchangeApi interface {
	GetExchangeId() models.ExchangeID
	fetchRate() error
	Volume(trading string, settlement string) (float64, error)
	CurrencyPairs() ([]*CurrencyPair, error)
	Rate(trading string, settlement string) (float64, error)
}
