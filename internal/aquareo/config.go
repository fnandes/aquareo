package aquareo

const (
	SensorPh       = "ph"
	SensorSalinity = "salinity"
	SensorSysTemp  = "sys_temp"
	SensorSysMem   = "sys_mem"
	DSL8B20        = "dsl8b20"
)

type Config struct {
	Name          string                  `json:"name"`
	Sensors       map[string]SensorConfig `json:"sensors"`
	CustomMetrics []CustomMetric          `json:"customMetrics"`
}

type SensorConfig struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	MetricUnit  string `json:"metricUnit"`
	Type        string `json:"type"`
}

type CustomMetric struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	MetricUnit  string `json:"metricUnit"`
}
