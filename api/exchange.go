package api

import (
	"time"

	"github.com/pkg/errors"
	"github.com/airking05/go-exchange-chart-fetcher/models"
)

func NewExchangeAPI(id models.ExchangeID) (ExchangeApi, error) {
	switch id {
	case models.Poloniex:
		return NewPoloniexApiUsingConfigFunc(func(c *PoloniexApiConfig) {
			c.RateCacheDuration = 3 * time.Second
			c.ExchangeId = models.Poloniex
		})
	case models.Hitbtc:
		return NewHitbtcApiUsingConfigFunc(func(c *HitbtcApiConfig) {
			c.RateCacheDuration = 6 * time.Second
			c.ExchangeId = models.Hitbtc
		})

	case models.Bitflyer:
		return NewBitflyerApiUsingConfigFunc(func(c *BitflyerApiConfig) {
			c.RateCacheDuration = 3 * time.Second
			c.ExchangeId = models.Bitflyer
		})
	}
	return nil, errors.Errorf("there is no exchange api for %s", id)
}
