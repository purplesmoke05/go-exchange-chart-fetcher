Go-exchange-chart-fetcher is a small server for "Fetching CryptoCurrency Data", implemented by golang.
It's pretty much stable server.

## How to Install :

```bash
git clone github.com/airking05/go-exchange-chart-fetcher
```

## What is this :

- this is a tool for fetching chart data.
- the data can be used for machine learning or so...

| Poloniex       	| Hitbtc          	| Bitflyer     	| Binance 	| Zaif 	|
|----------------	|-----------------	|--------------	|---------	|------	|
| Done(99 paris) 	| Done(421 paris) 	| Done(1 pair) 	|         	|      	|

![fetched_data](https://imgur.com/a/RujeJ)

## How to Use :

### make config.yml

```bash
cd go-exchange-chart-fetcher
make glide
cp config_sample.yml config.yml
nano config.yml
// edit mysql connector setting
```

### config.yml

write your DB connection setting
```
debug: false
test: false
db_connection: mysql:mysql@tcp(localhost:3306)/chart_fetcher?charset=utf8&parseTime=True&loc=UTC
```

### run server

```bash
go run *.go config.yml

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


