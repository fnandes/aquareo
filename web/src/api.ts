import { Metric, MetricRecord } from './types'

const baseUrl = 'http://raspberrypi.local:8082'

export const listMetrics = async (): Promise<Metric[]> => {
  return await fetch(`${baseUrl}/metrics`).then(res => res.json())
}

export const fetchMetricData = async (metricId: string): Promise<MetricRecord[]> => {
  return await fetch(`${baseUrl}/metrics/${metricId}`).then(res => res.json())
}