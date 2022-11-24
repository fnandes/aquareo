import { demoData } from './tests/dummy-data'
import { Config, ControllableGear, MetricEntry } from './types'

const baseUrl = 'http://raspberrypi.local:8082/api'

export const fetchConfig = async (): Promise<Config> =>
  await fetch(`${baseUrl}/config`).then(res => res.json())

export const fetchMetricData = async (bucket: string): Promise<MetricEntry[]> =>
  await fetch(`${baseUrl}/metrics/${bucket}`).then(res => res.json())

export const fetchGears = async (): Promise<ControllableGear[]> => {
  return Promise.resolve(demoData.gears)
}