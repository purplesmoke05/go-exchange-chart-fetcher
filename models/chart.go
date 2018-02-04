package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Chart struct {
	gorm.Model

	Last       float64
	Open       float64
	High       float64
	Low        float64
	Volume     float64
	Duration   int        `gorm:"index"`
	Datetime   time.Time  `gorm:"index"`
	Pair       string     `gorm:"index"`
	ExchangeID ExchangeID `gorm:"index"`
}
