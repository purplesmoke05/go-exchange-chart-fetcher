package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/airking05/go-exchange-chart-fetcher/config"
	"github.com/airking05/go-exchange-chart-fetcher/models"
	"github.com/airking05/go-exchange-chart-fetcher/logger"
	"github.com/airking05/go-exchange-chart-fetcher/api"
)

func help_and_exit() {
	fmt.Fprintf(os.Stderr, "%s config.yml\n", os.Args[0])
	os.Exit(1)
}

var ExchangeIDs = []models.ExchangeID{
	models.Bitflyer,
	models.Poloniex,
	models.Hitbtc,
}
func main() {
	if len(os.Args) != 2 {
		help_and_exit()
	}
	confPath := os.Args[1]

	conf := config.ReadConfig(confPath)
	var eapis []api.ExchangeApi
	for _, exchangeId := range ExchangeIDs {
		ea, err := api.NewExchangeAPI(exchangeId)
		if err != nil {
			logger.Get().Info(err)
			continue
		}
		eapis = append(eapis, ea)

	}

	db, err := gorm.Open("mysql", conf.DBConnection)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect db"))
	}
	db.AutoMigrate(&models.Chart{})

	server := NewServer(eapis, db)
	logger.Get().Info("starting chart_server...")
	server.Run()
}
