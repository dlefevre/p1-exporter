package collector

import (
	"github.com/dlefevre/p1-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// Helper functions to construct the promethus collector for the P1 interface
// This keeps the clutter out of P1Collector.go

// Get a map of metric descriptors
func createPrometheusDescriptors() map[string]*prometheus.Desc {
	return map[string]*prometheus.Desc{
		metrics.ObisConsumedTariff1: prometheus.NewDesc(
			"p1_electricity_consumed_tariff1",
			"Total amount of consumed electricity (tariff 1) in kWh",
			nil, nil),
		metrics.ObisConsumedTariff2: prometheus.NewDesc(
			"p1_electricity_consumed_tariff2",
			"Total amount of consumed electricity (tariff 2) in kWh",
			nil, nil),
		metrics.ObisInjectedTariff1: prometheus.NewDesc(
			"p1_electricity_injected_tariff1",
			"Total amount of injected electricity (tariff 1) in kWh",
			nil, nil),
		metrics.ObisInjectedTariff2: prometheus.NewDesc(
			"p1_electricity_injected_tariff2",
			"Total amount of injected electricity (tariff 2) in kWh",
			nil, nil),
		metrics.ObisCurrentConsumption: prometheus.NewDesc(
			"p1_electricity_current_consumption",
			"Actual power consumption for all lines in Watt",
			nil, nil),
		metrics.ObisCurrentInjection: prometheus.NewDesc(
			"p1_electricity_current_injection",
			"Actual power injection for all lines in Watt",
			nil, nil),
		metrics.ObisCurrentConsumptionL1: prometheus.NewDesc(
			"p1_electricity_current_consumption_l1",
			"Actual power consumption for Line 1 in Watt",
			nil, nil),
		metrics.ObisCurrentConsumptionL2: prometheus.NewDesc(
			"p1_electricity_current_consumption_l2",
			"Actual power consumption for Line 2 in Watt",
			nil, nil),
		metrics.ObisCurrentConsumptionL3: prometheus.NewDesc(
			"p1_electricity_current_consumption_l3",
			"Actual power consumption for Line 3 in Watt",
			nil, nil),
		metrics.ObisCurrentInjectionL1: prometheus.NewDesc(
			"p1_electricity_current_injection_l1",
			"Actual power injection for Line 1 in Watt",
			nil, nil),
		metrics.ObisCurrentInjectionL2: prometheus.NewDesc(
			"p1_electricity_current_injection_l2",
			"Actual power injection for Line 2 in Watt",
			nil, nil),
		metrics.ObisCurrentInjectionL3: prometheus.NewDesc(
			"p1_electricity_current_injection_l3",
			"Actual power injection for Line 3 in Watt",
			nil, nil),
		metrics.ObisVoltageL1: prometheus.NewDesc(
			"p1_electricity_voltage_l1",
			"Actual voltage on Line 1 in Volt",
			nil, nil),
		metrics.ObisVoltageL2: prometheus.NewDesc(
			"p1_electricity_voltage_l2",
			"Actual voltage on Line 2 in Volt",
			nil, nil),
		metrics.ObisVoltageL3: prometheus.NewDesc(
			"p1_electricity_voltage_l3",
			"Actual voltage on Line 3 in Volt",
			nil, nil),
		metrics.ObisCurrentL1: prometheus.NewDesc(
			"p1_electricity_current_l1",
			"Actual current on Line 1 in Amps",
			nil, nil),
		metrics.ObisCurrentL2: prometheus.NewDesc(
			"p1_electricity_current_l2",
			"Actual current on Line 2 in Amps",
			nil, nil),
		metrics.ObisCurrentL3: prometheus.NewDesc(
			"p1_electricity_current_l3",
			"Actual current on Line 3 in Amps",
			nil, nil),
		metrics.ObisElectricitySwitchState: prometheus.NewDesc(
			"p1_electricity_switch",
			"State of the electricity switch on/off",
			nil, nil),
		metrics.ObisGasSwitchState: prometheus.NewDesc(
			"p1_gas_switch",
			"State of the gas valve on/off",
			nil, nil),
		metrics.ObisGasConsumption: prometheus.NewDesc(
			"p1_gas_consumption",
			"Total volume of consumed gas in cubic meters",
			nil, nil),
	}
}

// Get a map of matched types
func createPrometheusTypes() map[string]prometheus.ValueType {
	return map[string]prometheus.ValueType{
		metrics.ObisConsumedTariff1:        prometheus.CounterValue,
		metrics.ObisConsumedTariff2:        prometheus.CounterValue,
		metrics.ObisInjectedTariff1:        prometheus.CounterValue,
		metrics.ObisInjectedTariff2:        prometheus.CounterValue,
		metrics.ObisCurrentConsumption:     prometheus.GaugeValue,
		metrics.ObisCurrentInjection:       prometheus.GaugeValue,
		metrics.ObisCurrentConsumptionL1:   prometheus.GaugeValue,
		metrics.ObisCurrentConsumptionL2:   prometheus.GaugeValue,
		metrics.ObisCurrentConsumptionL3:   prometheus.GaugeValue,
		metrics.ObisCurrentInjectionL1:     prometheus.GaugeValue,
		metrics.ObisCurrentInjectionL2:     prometheus.GaugeValue,
		metrics.ObisCurrentInjectionL3:     prometheus.GaugeValue,
		metrics.ObisVoltageL1:              prometheus.GaugeValue,
		metrics.ObisVoltageL2:              prometheus.GaugeValue,
		metrics.ObisVoltageL3:              prometheus.GaugeValue,
		metrics.ObisCurrentL1:              prometheus.GaugeValue,
		metrics.ObisCurrentL2:              prometheus.GaugeValue,
		metrics.ObisCurrentL3:              prometheus.GaugeValue,
		metrics.ObisElectricitySwitchState: prometheus.UntypedValue,
		metrics.ObisGasSwitchState:         prometheus.UntypedValue,
		metrics.ObisGasConsumption:         prometheus.CounterValue,
	}
}
