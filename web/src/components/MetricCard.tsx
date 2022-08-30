import * as React from 'react'
import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as colors from 'tailwindcss/colors'
import * as moment from 'moment'
import { Metric, MetricRecord } from '../types'
import * as api from '../api'


const EVERY_5_SECONDS = 5000

export type MetricCardProps = Metric & {}
export const MetricCard: React.FC<MetricCardProps> = ({ id, name }) => {
  const [data, setData] = React.useState<MetricRecord[]>([])

  React.useEffect(() => {
    api.fetchMetricData(id).then(setData)
    const interval = setInterval(() => {
      api.fetchMetricData(id).then(setData)
    }, EVERY_5_SECONDS)
    return () => clearInterval(interval)
  }, [])

  return (
    <div className='flex bg-white shadow-md rounded-md overflow-hidden'>
      <div className='flex-1'>
        <ResponsiveContainer height={300} width="99%">
          <LineChart data={data.sort((a, b) => a.Timespan - b.Timespan)}>
            <XAxis dataKey="Timespan" tickFormatter={val => moment.unix(val).format('hh:mm')}
              stroke={colors.purple[500]} />
            <YAxis domain={['auto', 'auto']} stroke={colors.purple[500]} strokeWidth={0} padding={{}} />
            <Tooltip />
            <Line type="monotone" dataKey="Value" stroke={colors.purple[500]} strokeWidth={3} unit="C" />
          </LineChart>
        </ResponsiveContainer>
      </div>
      <div className='flex-none basis-1/6 bg-gradient-to-tr from-indigo-500 to-indigo-900 text-white flex flex-col items-center justify-center'>
        <div className='mx-4'>
          <span className='text-6xl font-semibold'>43.3</span>
        </div>
        <div className='mx-4'>
          <h4>Temperature</h4>
        </div>
      </div>
    </div>
  )
}