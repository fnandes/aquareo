export type Metric = {
  id: string
  name: string
}

export type MetricRecord = {
  Timespan: number
  Value: number
}

export type ControllableGear = {
  id: string
  name: string
  lastUpdate: number
  isOn: boolean
}