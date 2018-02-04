package main

import (
	"github.com/airking05/go-exchange-chart-fetcher/api"
)

type watcherMap map[string]map[string]*PairWatcher

func (m watcherMap) keys() []api.CurrencyPair {
	keys := make([]api.CurrencyPair, 0, len(m)*3)
	for trading, sm := range m {
		for settlement := range sm {
			keys = append(keys, api.CurrencyPair{Trading: trading, Settlement: settlement})
		}
	}
	return keys
}

func (m watcherMap) get(pair *api.CurrencyPair) (*PairWatcher, bool) {
	if sm, ok := m[pair.Trading]; !ok {
		return nil, false
	} else if w, ok := sm[pair.Settlement]; !ok {
		return nil, false
	} else {
		return w, true
	}
}

func (m watcherMap) put(pair *api.CurrencyPair, watcher *PairWatcher) {
	sm, ok := m[pair.Trading]
	if !ok {
		sm = make(map[string]*PairWatcher)
		m[pair.Trading] = sm
	}

	sm[pair.Settlement] = watcher
}
