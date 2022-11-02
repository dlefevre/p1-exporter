package parser

import (
	"bufio"
	"strings"
	"sync"
	"time"

	"github.com/dlefevre/p1-exporter/config"
	"github.com/dlefevre/p1-exporter/metrics"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type P1ParserService struct {
	serialConf *serial.Config
	port       *serial.Port
	reader     *bufio.Reader
	store      metrics.MetricStore
}

var instance *P1ParserService = nil
var lock = &sync.Mutex{}

// Returns the singleton instance
func GetP1ParserService() *P1ParserService {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = newP1ParserService()
	}
	return instance
}

// Construct a new parser
func newP1ParserService() *P1ParserService {
	config := config.GetConfigService()

	serialConf := &serial.Config{
		Name:     config.SerialDevice,
		Baud:     config.SerialBaud,
		Size:     config.SerialSize,
		StopBits: config.SerialStopBits,
	}

	return &P1ParserService{
		serialConf: serialConf,
		port:       nil,
		reader:     nil,
		store:      nil,
	}
}

// Open serial port
func (service *P1ParserService) openPort() {
	for {
		port, err := serial.OpenPort(service.serialConf)
		if err != nil {
			log.Errorf("P1ParserService: Could not open port: %s", err)
			time.Sleep(config.GetConfigService().BackOffInterval)
		}
		service.port = port
		service.reader = bufio.NewReader(service.port)
		break
	}
}

// Reopen port
func (service *P1ParserService) reopenPort() {
	service.port.Close()
	service.openPort()
}

// Active thread
func (service *P1ParserService) Run() {
	config := config.GetConfigService()

	service.openPort()
	defer service.port.Close()

	service.validate()

	var fullDatagram bool = false
	builder := metrics.NewMetricBuilder()
	log.Infof("P1ParserService: Parsing serial data on %s", config.SerialDevice)

	for {
		str, err := service.reader.ReadString('\n')
		if err != nil {
			log.Errorf("P1ParserService: Read error (%s), reopening port", err)
			service.reopenPort()
		}

		str = strings.Trim(str, " \t\r\n\x00")
		if str == "" {
			continue
		}
		if strings.HasPrefix(str, "!") {
			if fullDatagram {
				service.store.Commit()
			}
			continue
		}
		if strings.HasPrefix(str, "/") {
			fullDatagram = true
			service.store.Clear()
			continue
		}

		metric, res := builder.CreateMetric(str)
		if res != nil {
			if !res.Ignore() {
				log.Warn(res)
			}
		} else {
			service.store.SetMetric(metric)
		}
	}
}

// Check consistency of service setup
func (service *P1ParserService) validate() {
	if service.store == nil {
		log.Fatal("P1ParserService: No MetricStore configured in P1ParserService")
	}
}

// Set the metricstore
func (service *P1ParserService) SetMetricStore(store metrics.MetricStore) {
	service.store = store
}
