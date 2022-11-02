package metrics

const (
	ObisConsumedTariff1 = "1-0:1.8.1"
	ObisConsumedTariff2 = "1-0:1.8.2"
	ObisInjectedTariff1 = "1-0:2.8.1"
	ObisInjectedTariff2 = "1-0:2.8.2"
	ObisActiveTariff    = "0-0:96.14.0"

	ObisCurrentConsumption = "1-0:1.7.0"
	ObisCurrentInjection   = "1-0:2.7.0"

	ObisCurrentConsumptionL1 = "1-0:21.7.0"
	ObisCurrentConsumptionL2 = "1-0:41.7.0"
	ObisCurrentConsumptionL3 = "1-0:61.7.0"
	ObisCurrentInjectionL1   = "1-0:22.7.0"
	ObisCurrentInjectionL2   = "1-0:42.7.0"
	ObisCurrentInjectionL3   = "1-0:62.7.0"

	ObisVoltageL1 = "1-0:32.7.0"
	ObisVoltageL2 = "1-0:52.7.0"
	ObisVoltageL3 = "1-0:72.7.0"
	ObisCurrentL1 = "1-0:31.7.0"
	ObisCurrentL2 = "1-0:51.7.0"
	ObisCurrentL3 = "1-0:71.7.0"

	ObisElectricitySwitchState = "0-0:96.3.10"
	ObisGasSwitchState         = "0-1:24.4.0"

	ObisGasConsumption = "0-1:24.2.3"
)

type Metric struct {
	Id    string
	Value float64
	Unit  string
}
