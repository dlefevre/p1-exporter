package metrics

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

type MetricStore interface {
	Clear()
	Commit()
	SetMetric(Metric)
	GetMetric(string) (Metric, error)
	Current() map[string]Metric
}

type MetricStoreImpl struct {
	collector map[string]Metric
	current   map[string]Metric

	// only used to lock access to current set of metrics
	lock *sync.RWMutex
}

// Construct a new MetricsStore
func NewMetricStore() MetricStore {
	return &MetricStoreImpl{
		collector: make(map[string]Metric),
		current:   make(map[string]Metric),
		lock:      &sync.RWMutex{},
	}
}

// Clear the collector
func (store *MetricStoreImpl) Clear() {
}

// Commit the collected metrics to the current set
func (store *MetricStoreImpl) Commit() {
	store.lock.Lock()
	defer store.lock.Unlock()
	store.current = store.collector
	log.Debugf("MetricStore: committed %d metrics", len(store.current))
}

// Set a single metric
func (store *MetricStoreImpl) SetMetric(metric Metric) {
	store.collector[metric.Id] = metric
	log.Tracef("MetricStore: metric set: %+v", metric)
}

// Get a single metric (not guaranteed to yield a consistent set over multiple calls)
func (store *MetricStoreImpl) GetMetric(id string) (Metric, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	if metric, found := store.current[id]; found {
		return metric, nil
	} else {
		return Metric{}, fmt.Errorf("no metric with id `%s` found", id)
	}
}

// Get the current metrics map (so we won't copy...
// and thus trust our caller to leave the content untouched)
func (store *MetricStoreImpl) Current() map[string]Metric {
	store.lock.RLock()
	defer store.lock.RUnlock()
	return store.current
}
