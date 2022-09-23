import * as React from 'react'
import { Metric } from '../../types'
import * as api from '../../api'
import { useQuery } from '@tanstack/react-query'
import { MetricCardLineChart } from './MetricCardLineChart'

export type MetricCardProps = Metric & {}

export const MetricCard: React.FC<MetricCardProps> = ({ id, name }) => {
  const { data } = useQuery(['metric', id], () => api.fetchMetricData(id))

  return data ? (
    <MetricCardLineChart
      data={data}
      domain={[20, 30]}
      unit='C'
      title={name} />
  ) : null
}