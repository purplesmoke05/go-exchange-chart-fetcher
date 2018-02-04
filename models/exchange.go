package models

//go:generate enumer -type=ExchangeID
type ExchangeID int

const (
	Bitflyer ExchangeID = iota + 1
	Poloniex
	Hitbtc

	UnknownExchange
)
