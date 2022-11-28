import { Config, MetricEntry } from './types'

const baseUrl = 'http://raspberrypi.local:8082/api'

export const fetchConfig = async (): Promise<Config> =>
  await fetch(`${baseUrl}/config`).then(res => res.json())

export const fetchMetricData = async (bucket: string): Promise<MetricEntry[]> =>
  await fetch(`${baseUrl}/metrics/${bucket}`).then(res => res.json())

export const addMetricEntry = async (bucket: string, entry: MetricEntry) =>
  await fetch(`${baseUrl}/metrics/${bucket}`, {
    headers: {
      'Accept': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify(entry)
  })