import * as React from 'react'
import * as api from '../../api'
import { useQuery } from '@tanstack/react-query'
import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as colors from 'tailwindcss/colors'
import * as moment from 'moment'

export type MetricCardProps = {
  bucket: string
  title: string
  metricUnit: string
}

export const MetricCard: React.FC<MetricCardProps> = ({ bucket, title, metricUnit }) => {
  const tickFormatter = (val: number) => moment.unix(val).format('hh:mm')
  const { data } = useQuery(['metric', bucket], () => api.fetchMetricData(bucket))

  return data ? (
    <div className='bg-white rounded shadow-md'>
      <div className='p-2 font-semibold text-slate-500'>
        <h3>{title}</h3>
      </div>
      <ResponsiveContainer height={150} width="99%">
        <LineChart data={data.sort((a, b) => a.Timespan - b.Timespan)}>
          <XAxis dataKey="Timespan" tickFormatter={tickFormatter} stroke={colors.purple[500]} />
          <YAxis stroke={colors.purple[500]} />
          <Tooltip />
          <Line type="monotone" dataKey="Value" stroke={colors.purple[500]} strokeWidth={3} unit={metricUnit} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  ) : null
}