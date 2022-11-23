package aquareo

type Config struct {
	Name                  string                      `json:"name"`
	TemperatureController TemperatureControllerConfig `json:"temperatureController"`
	CustomMetrics         []CustomMetric              `json:"customMetrics"`
}

type TemperatureControllerConfig struct {
	DeviceId     string `json:"deviceId"`
	TickInterval uint   `json:"tickInterval"`
}

type CustomMetric struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	MetricUnit  string `json:"metricUnit"`
}
