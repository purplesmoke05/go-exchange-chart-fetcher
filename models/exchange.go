package models

//go:generate enumer -type=ExchangeID
type ExchangeID int

const (
	Poloniex ExchangeID = iota + 1
	Hitbtc
	Bitflyer

	UnknownExchange
)
