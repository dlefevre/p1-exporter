package collector

import (
	"net/http"
	"sync"
	"time"

	"github.com/dlefevre/p1-exporter/config"
	"github.com/dlefevre/p1-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type P1Collector struct {
	descriptors map[string]*prometheus.Desc
	types       map[string]prometheus.ValueType
	store       metrics.MetricStore
}

var instance *P1Collector = nil
var lock = &sync.Mutex{}

// Returns the singleton instance
func GetP1Collector() *P1Collector {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = newP1Collector()
	}
	return instance
}

// Constructor
func newP1Collector() *P1Collector {
	return &P1Collector{
		descriptors: createPrometheusDescriptors(),
		types:       createPrometheusTypes(),
		store:       nil,
	}
}

// Send describers to prometheus channel
func (collector *P1Collector) Describe(ch chan<- *prometheus.Desc) {
	log.Debug("P1Collector: Describing metrics.")
	current := collector.store.Current()

	for obisId, descriptor := range collector.descriptors {
		// Only push descriptors for metrics that are published on the P1 serial interface
		if _, found := current[obisId]; found {
			ch <- descriptor
		}
	}
}

// Collect metrics and send to prometheus channel
func (collector *P1Collector) Collect(ch chan<- prometheus.Metric) {
	log.Debug("P1Collector: Collecting metrics.")

	current := collector.store.Current()
	ts := time.Now()

	for obisId, descriptor := range collector.descriptors {
		if metric, found := current[obisId]; found {
			m, err := prometheus.NewConstMetric(
				descriptor,
				collector.types[obisId],
				metric.Value,
			)
			if err != nil {
				log.Errorf("P1Collector: Could not create ConstMetric for %s (%v)", obisId, err)
				continue
			}
			m = prometheus.NewMetricWithTimestamp(ts, m)
			ch <- m
		}
	}
}

// Listen for incomming connections
func (collector *P1Collector) Run() {
	config := config.GetConfigService()

	collector.validate()
	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	log.Infof("P1Collector: Listening on %s", config.ListenAddress)

	http.Handle("/metrics", handler)
	err := http.ListenAndServe(config.ListenAddress, nil)
	log.Fatalf("P1Collector: Listener borked: %v", err)
}

// Check consistency of service setup
func (collector *P1Collector) validate() {
	if collector.store == nil {
		log.Fatal("P1Collector: No MetricStore configured in P1Collector")
	}
}

// Set the metricstore
func (collector *P1Collector) SetMetricStore(store metrics.MetricStore) {
	collector.store = store
}
