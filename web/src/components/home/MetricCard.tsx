import * as React from 'react'
import { useQuery } from '@tanstack/react-query'
import { CartesianGrid, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as colors from 'tailwindcss/colors'
import * as moment from 'moment'
import * as api from '../../api'
import { Card } from '../card'

export type MetricCardProps = {
  bucket: string
  title: string
  metricUnit: string
}

export const MetricCard: React.FC<MetricCardProps> = ({ bucket, title, metricUnit }) => {
  const tickFormatter = (val: number) => moment.unix(val).format('D/M hh:mm')
  const { data = [] } = useQuery(['metric', bucket], () => api.fetchMetricData(bucket))

  return (
    <Card title={title}>
      <ResponsiveContainer height={150} width="99%">
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="timespan" tickFormatter={tickFormatter} stroke={colors.purple[500]} fontSize="small" />
          <YAxis stroke={colors.purple[500]} />
          <Tooltip />
          <Line type="monotone" connectNulls dataKey="value" stroke={colors.purple[500]} strokeWidth={3} unit={metricUnit} />
        </LineChart>
      </ResponsiveContainer>
    </Card>
  )
}