package aquareo

type Config struct {
	Name                  string                      `json:"name"`
	TemperatureController TemperatureControllerConfig `json:"temperatureController"`
	CustomMetrics         []CustomMetric              `json:"customMetrics"`
}

type TemperatureControllerConfig struct {
	Enabled          bool   `json:"enabled"`
	DeviceId         string `json:"deviceId"`
	TickInterval     int32  `json:"tickInterval"`
	SnapshotInterval int32  `json:"snapshotInterval"`
}

type CustomMetric struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	MetricUnit  string `json:"metricUnit"`
}
