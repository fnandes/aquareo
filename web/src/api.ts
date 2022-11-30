import { Config, MetricEntry } from './types'

const baseUrl = 'http://raspberrypi.local:8082/api'

export const fetchConfig = async (): Promise<Config> =>
  await fetch(`${baseUrl}/config`).then(res => res.json())

export const metrics = (bucket: string) => ({
  fetchAll: async (): Promise<MetricEntry[]> => await fetch(`${baseUrl}/metrics/${bucket}`).then(res => res.json()),

  addEntry: async (entry: MetricEntry) => await fetch(`${baseUrl}/metrics/${bucket}`, {
    headers: { 'Accept': 'application/json' },
    method: 'POST',
    body: JSON.stringify(entry)
  }),

  deleteEntry: async (timestamp: number) => await fetch(`${baseUrl}/metrics/${bucket}/${timestamp}`, {
    headers: { 'Accept': 'application/json' },
    method: 'DELETE'
  })
})