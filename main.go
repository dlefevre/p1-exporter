package main

import (
	"sync"

	"github.com/dlefevre/p1-exporter/collector"
	"github.com/dlefevre/p1-exporter/config"
	"github.com/dlefevre/p1-exporter/metrics"
	"github.com/dlefevre/p1-exporter/parser"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.GetConfigService()

	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(config.LogLevel)

	parser := parser.GetP1ParserService()
	collector := collector.GetP1Collector()

	store := metrics.NewMetricStore()
	parser.SetMetricStore(store)
	collector.SetMetricStore(store)

	var wg sync.WaitGroup
	wg.Add(2)

	go parser.Run()
	go collector.Run()

	wg.Wait()
}
