package main

import (
	"math"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/airking05/go-exchange-chart-fetcher/api"
	"github.com/airking05/go-exchange-chart-fetcher/logger"
	"github.com/airking05/go-exchange-chart-fetcher/models"
)

func NewServer(eapis []api.ExchangeApi, db *gorm.DB) *Server {
	return &Server{
		eapis:    eapis,
		watchers: make(watcherMap),
		writer:   NewChartWriter(db),
	}
}

type Server struct {
	eapis    []api.ExchangeApi
	watchers watcherMap
	writer   *ChartWriter
}

func (s *Server) Run() {
	s.writer.Start()
	for {
		for _, eapi := range s.eapis {
			logger.Get().Info("checking currecy pairs updates")
			pairs, err := eapi.CurrencyPairs()
			if err != nil {
				logger.Get().Errorf("failed to check currency pair: %s", err)
				time.Sleep(1 * time.Minute)
			}
			for _, pair := range pairs {
				if _, ok := s.watchers.get(pair); !ok {
					watcher := NewPairWatcher(pair, eapi, s.writer)
					s.watchers.put(pair, watcher)
					watcher.Start(eapi.GetExchangeId())
				}
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

/*
 * interval: rate check interval
 * duration: rate update duration
 */
func NewPairWatcher(pair *api.CurrencyPair, eapi api.ExchangeApi, writer *ChartWriter) *PairWatcher {
	return &PairWatcher{
		pair:   pair,
		eapi:   eapi,
		writer: writer,
	}
}

type PairWatcher struct {
	pair   *api.CurrencyPair
	eapi   api.ExchangeApi
	writer *ChartWriter
}

func (w *PairWatcher) Start(exchangeId models.ExchangeID) {
	logger.Get().Infof("starting pair watcher [%s] %s/%s",
		exchangeId.String(),
		(w.pair.Trading),
		(w.pair.Settlement))
	w1m := NewDurationPairWatcher(w.pair, w.eapi, w.writer, 3*time.Second, 1*time.Minute, exchangeId)
	w1h := NewDurationPairWatcher(w.pair, w.eapi, w.writer, 3*time.Second, 1*time.Hour, exchangeId)
	w1d := NewDurationPairWatcher(w.pair, w.eapi, w.writer, 3*time.Second, 24*time.Hour, exchangeId)
	w1m.Start()
	w1h.Start()
	w1d.Start()
}

func NewDurationPairWatcher(pair *api.CurrencyPair, eapi api.ExchangeApi, writer *ChartWriter,
	interval time.Duration, duration time.Duration, exchangeId models.ExchangeID) *DurationPairWatcher {
	return &DurationPairWatcher{
		pair:       pair,
		eapi:       eapi,
		writer:     writer,
		interval:   interval,
		duration:   duration,
		exchangeId: exchangeId,
	}
}

type DurationPairWatcher struct {
	pair       *api.CurrencyPair
	eapi       api.ExchangeApi
	writer     *ChartWriter
	exchangeId models.ExchangeID
	interval   time.Duration
	duration   time.Duration
}

func (w *DurationPairWatcher) Start() {
	go func() {
		intervalTick := time.Tick(w.interval)
		durationTick := time.Tick(w.duration)
		if w.interval >= w.duration {
			logger.Get().Fatalf("assertion error: interval(%d) >= duration(%d)",
				w.interval, w.duration)
		}

		open := float64(0)
		high := float64(0)
		low := float64(0)
		lastVolume := float64(0)

		rate, err := w.eapi.Rate(w.pair.Trading, w.pair.Settlement)
		if err != nil {
			return
		}
		volume, err := w.eapi.Volume(w.pair.Trading, w.pair.Settlement)
		if err != nil {
			return
		}
		open = rate
		high = rate
		low = rate
		lastVolume = volume

		for {
			select {
			case <-intervalTick:
				rate, err := w.eapi.Rate(w.pair.Trading, w.pair.Settlement)
				if err != nil {
					logger.Get().Errorf("rate not found: %s", err)
					continue
				}
				high = math.Max(high, rate)
				low = math.Min(low, rate)
			case <-durationTick:
				rate, err := w.eapi.Rate(w.pair.Trading, w.pair.Settlement)
				if err != nil {
					logger.Get().Errorf("rate not found: %s", err)
					continue
				}
				volume, err := w.eapi.Volume(w.pair.Trading, w.pair.Settlement)
				if err != nil {
					logger.Get().Errorf("volume not found: %s", err)
					continue
				}
				high = math.Max(high, rate)
				low = math.Min(low, rate)
				chart := models.Chart{
					Last:       rate,
					Open:       open,
					High:       high,
					Low:        low,
					Volume:     lastVolume - volume,
					Pair:       w.pair.Trading + "_" + w.pair.Settlement,
					Duration:   int(w.duration.Seconds()),
					Datetime:   time.Now(),
					ExchangeID: w.exchangeId,
				}
				w.writer.Appender() <- chart
				open = rate
				high = rate
				low = rate
				lastVolume = volume
			}

		}
	}()
}

type ChartWriter struct {
	ch chan models.Chart
	db *gorm.DB
}

func NewChartWriter(db *gorm.DB) *ChartWriter {
	return &ChartWriter{
		db: db,
		ch: make(chan models.Chart, 1024),
	}
}

func (w *ChartWriter) Appender() chan<- models.Chart {
	return w.ch
}

func (w *ChartWriter) Start() {
	go func() {
		logger.Get().Info("starting chart writer...")
		for {
			chart := <-w.ch
			if w.db.Create(&chart); w.db.Error != nil {
				logger.Get().Errorf("failed to create chart data: %s", w.db.Error)
			}
		}
	}()
}
