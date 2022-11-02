package config

import (
	"flag"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type ConfigService struct {
	ListenAddress   string
	SerialDevice    string
	SerialBaud      int
	SerialSize      byte
	SerialStopBits  serial.StopBits
	BackOffInterval time.Duration
	LogLevel        log.Level
}

var instance *ConfigService = nil
var lock = &sync.Mutex{}

// Returns the singleton instance
func GetConfigService() *ConfigService {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = newConfigService()
	}
	return instance
}

// Construct a new parser
func newConfigService() *ConfigService {
	listenAddress := flag.String("listen", defaultListenAddress, "Listen address and port")
	serialDevice := flag.String("device", defaultSerialDevice, "Serial device")
	serialBaud := flag.Int("baud", defaultSerialBaud, "Baud rate")
	sLogLevel := flag.String("loglevel", defaultLogLevel, "Log Level")
	flag.Parse()

	logLevel, err := log.ParseLevel(*sLogLevel)
	if err != nil {
		log.Error(fmt.Sprintf("ConfigService: %v, setting default (%s)", err, defaultLogLevel))
		logLevel, err = log.ParseLevel(defaultLogLevel)
		if err != nil {
			log.Fatal(fmt.Sprintf("ConfigService: Issue setting default Log level: %v", err))
		}
	}

	service := &ConfigService{
		ListenAddress:   *listenAddress,
		SerialDevice:    *serialDevice,
		SerialBaud:      *serialBaud,
		SerialSize:      defaultSerialSize,
		SerialStopBits:  defaultSerialStopBits,
		BackOffInterval: defaultBackOffInterval * time.Second,
		LogLevel:        logLevel,
	}
	service.info()

	return service
}

// Some informational logging
func (service *ConfigService) info() {
	log.Infof("ConfigService: configured value for ListenAddress: `%s`", service.ListenAddress)
	log.Infof("ConfigService: configured value for SerialDevice: `%s`", service.SerialDevice)
	log.Infof("ConfigService: configured value for SerialBaud: `%d`", service.SerialBaud)
	log.Infof("ConfigService: configured value for SerialSize: `%d`", service.SerialSize)
	log.Infof("ConfigService: configured value for SerialStopBits: `%d`", service.SerialStopBits)
	log.Infof("ConfigService: configured value for BackOffInterval: `%s`", service.BackOffInterval.String())
	log.Infof("ConfigService: configured value for LogLevel: `%s`", service.LogLevel.String())
}
