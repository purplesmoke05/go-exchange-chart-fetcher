[![Build Status](https://secure.travis-ci.org/airking05/go-exchange-chart-fetcher.png?branch=master)](http://travis-ci.org/airking05/go-exchange-chart-fetcher)
[![Coverage Status](https://coveralls.io/repos/airking05/go-exchange-chart-fetcher/badge.svg?branch=master)](https://coveralls.io/r/airking05/go-exchange-chart-fetcher?branch=master)
[![GoDoc](https://godoc.org/github.com/airking05/go-exchange-chart-fetcher?status.svg)](https://godoc.org/github.com/airking05/go-exchange-chart-fetcher)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/airking05/go-exchange-chart-fetcher/blob/master/LICENSE)

Go-exchange-chart-fetcher is a small server for "Fetching CryptoCurrency Data", implemented by golang.
It's pretty much stable server.

## How to Install :

```bash
go get github.com/airking05/go-exchange-chart-fetcher
```

## What is this :

- this is a tool for fetching chart data.
- the data can be used for machine learning or so...

|        	| Poloniex       	| Hitbtc          	| Bitflyer     	| Binance 	| Zaif 	|
|--------	|----------------	|-----------------	|--------------	|---------	|------	|
| ID     	| 1              	| 2               	| 3            	|         	|      	|
| Status 	| Done(99 paris) 	| Done(421 paris) 	| Done(1 pair) 	|         	|      	|

![fetched_data](https://i.imgur.com/Qt4gUf9.png)

## How to Use :

```Go
package main
import(
	"github.com/jinzhu/gorm"
	"github.com/airking05/go-exchange-chart-fetcher/models"
	"github.com/airking05/go-exchange-chart-fetcher/server"
)

// declare exchanges you want to fetch
var ExchangeIDs = []models.ExchangeID{
	models.Bitflyer,
	models.Poloniex,
	models.Hitbtc,
}

func main() {
	dbConf := "mysql:mysql@tcp(localhost:3306)/chart_fetcher?charset=utf8&parseTime=True&loc=UTC"
    // setup DB
    db, err := gorm.Open("mysql", dbConf)
    if err != nil {
        panic("failed to connect db")
    }
    db.AutoMigrate(&models.Chart{})
    
    // then start server!
    server := server.NewServer(ExchangeIDs, db)
    server.Run()
}
```

### running server

```bash
2018-02-04T18:50:40.918+0900	INFO	go-exchange-chart-fetcher/main.go:51	starting chart_server...
2018-02-04T18:50:40.918+0900	INFO	go-exchange-chart-fetcher/server.go:31	checking currecy pairs updates
2018-02-04T18:50:40.918+0900	INFO	go-exchange-chart-fetcher/server.go:192	starting chart writer...
2018-02-04T18:50:41.136+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Bitflyer] BTC/JPY
2018-02-04T18:50:41.136+0900	INFO	go-exchange-chart-fetcher/server.go:31	checking currecy pairs updates
2018-02-04T18:50:41.781+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Poloniex] CLAM/BTC
2018-02-04T18:50:41.781+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Poloniex] BTCD/XMR
2018-02-04T18:50:41.781+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Poloniex] BTCD/BTC
2018-02-04T18:50:41.781+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Poloniex] ETH/USDT
2018-02-04T18:50:41.781+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Poloniex] ETH/BTC
....
2018-02-04T18:59:08.583+0900	INFO	go-exchange-chart-fetcher/server.go:31	checking currecy pairs updates
2018-02-04T18:59:09.105+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Hitbtc] DOGE/USD
2018-02-04T18:59:09.105+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Hitbtc] DOGE/ETH
2018-02-04T18:59:09.105+0900	INFO	go-exchange-chart-fetcher/server.go:69	starting pair watcher [Hitbtc] COV/BTC
....
```


