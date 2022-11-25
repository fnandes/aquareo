export type Config = {
  name: string
  temperatureController: Partial<TemperatureCtrlConfig>
  customMetrics: CustomMetric[]
}

export type TemperatureCtrlConfig = {
  enabled: boolean
  deviceId: string
  tickInterval: number
}

export type CustomMetric = {
  id: string
  displayName: string
  metricUnit: string
}

export type MetricEntry = {
  timespan: number
  value: number
}

export type ControllableGear = {
  id: string
  name: string
  lastUpdate: number
  isOn: boolean
}