export type Config = {
  name: string
  temperatureController: Partial<TemperatureCtrlConfig>
  customMetrics: string[]
}

export type TemperatureCtrlConfig = {
  enabled: boolean
  deviceId: string
  tickInterval: number
}

export type MetricEntry = {
  timespan: number
  value: number
}
