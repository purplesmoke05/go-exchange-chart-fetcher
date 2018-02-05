package server

import "github.com/jinzhu/gorm"
import (
	"testing"
	"github.com/airking05/go-exchange-chart-fetcher/models"
)

func TestNewServer(t *testing.T) {

	var ExchangeIDs = []models.ExchangeID{
		models.Bitflyer,
		models.Poloniex,
		models.Hitbtc,
	}

	_ = NewServer(ExchangeIDs, &gorm.DB{})
}
