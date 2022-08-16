package aquareo

const (
	SensorTemp1    = "temperature_1"
	SensorTemp2    = "temperature_2"
	SensorPh       = "ph"
	SensorSalinity = "salinity"
	DSL8B20        = "dsl8b20"
)

type Config struct {
	Name string `json:"name"`

	Sensors []SensorConfig `json:"sensors"`
}

type SensorConfig struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	SerialNr string `json:"serial_nr"`
}
