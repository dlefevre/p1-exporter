package metrics

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var multiplier = map[string]float64{
	"MW":  1.e+6,
	"kW":  1.e+3,
	"W":   1.,
	"MWh": 1.e+6,
	"kWh": 1.e+3,
	"Wh":  1.,
	"kV":  1.e+3,
	"V":   1.,
	"kA":  1.e+3,
	"A":   1.,
	"m3":  1.,
	"":    1.,
}

type MetricBuilder struct {
	re_raw *regexp.Regexp
	parts  []string
}

// Get new metric builder
func NewMetricBuilder() MetricBuilder {
	return MetricBuilder{
		re_raw: regexp.MustCompile(`(.+?)\((.*?)\)(?:\((.*?)\))?`),
	}
}

// Build metrics
func (builder *MetricBuilder) CreateMetric(raw string) (Metric, MetricBuilderError) {
	parts := builder.re_raw.FindStringSubmatch(raw)

	if parts == nil || len(parts) < 3 {
		return Metric{}, NewMetricBuilderError(fmt.Sprintf("could not parse raw string: `%s`", raw), false)
	}
	builder.parts = parts
	id := parts[1]

	switch id {
	case ObisConsumedTariff1, ObisConsumedTariff2, ObisInjectedTariff1, ObisInjectedTariff2:
		return builder.parseMetric("kWh")
	case ObisCurrentConsumption, ObisCurrentConsumptionL1, ObisCurrentConsumptionL2,
		ObisCurrentConsumptionL3, ObisCurrentInjection, ObisCurrentInjectionL1,
		ObisCurrentInjectionL2, ObisCurrentInjectionL3:
		return builder.parseMetric("W")
	case ObisVoltageL1, ObisVoltageL2, ObisVoltageL3:
		return builder.parseMetric("V")
	case ObisCurrentL1, ObisCurrentL2, ObisCurrentL3:
		return builder.parseMetric("A")
	case ObisElectricitySwitchState, ObisGasSwitchState:
		return builder.parseMetric("")
	case ObisGasConsumption:
		if len(parts) < 4 {
			return Metric{}, NewMetricBuilderError(fmt.Sprintf("gas consumption line has an invalid format: %s", raw), false)
		}
		parts[2] = parts[3]
		return builder.parseMetric("m3")
	default:
		return Metric{}, NewMetricBuilderError(fmt.Sprintf("unsupported metric: %s", id), true)
	}
}

// Parse
func (builder *MetricBuilder) parseMetric(unit string) (Metric, MetricBuilderError) {
	tmp := strings.Split(builder.parts[2], "*")
	if len(tmp) < 1 {
		return Metric{}, NewMetricBuilderError(fmt.Sprintf("could not parse metric: %s", builder.parts[2]), false)
	}

	value, err := strconv.ParseFloat(tmp[0], 64)
	if err != nil {
		return Metric{}, NewMetricBuilderError(fmt.Sprintf("could not parse metric's value: %s", tmp[0]), false)
	}

	if len(tmp) == 2 && tmp[1] != unit {
		if _, found := multiplier[tmp[1]]; !found {
			return Metric{}, NewMetricBuilderError(fmt.Sprintf("invalid unit for source metric: %s", tmp[1]), false)
		}
		if _, found := multiplier[unit]; !found {
			return Metric{}, NewMetricBuilderError(fmt.Sprintf("invalid unit for target metric: %s", unit), false)
		}

		from := multiplier[tmp[1]]
		to := multiplier[unit]
		value = value * from / to
	}

	return Metric{
		Id:    builder.parts[1],
		Value: value,
		Unit:  unit,
	}, nil
}
