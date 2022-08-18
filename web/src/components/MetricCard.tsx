import * as React from 'react'
import * as colors from 'tailwindcss/colors'
import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
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
    <div className='basis-1/2'>
      <div className='bg-gray-800 shadow-lg rounded-lg m-4 hover:shadow-sm'>

        <div className='px-4 pt-2 pb-4 font-semibold'>{name}</div>
        <div className='relative overflow-hidden'>
          <ResponsiveContainer height={300} width="99%">
            <LineChart data={data} margin={{}}>
              <XAxis dataKey="Timespan" tickFormatter={val => moment(val).format('HH:mm')} stroke={colors.indigo[500]} />
              <YAxis domain={['auto', 'auto']} stroke={colors.indigo[400]} strokeWidth={0} padding={{}} />
              <Tooltip />
              <Line type="monotone" dataKey="Value" stroke={colors.purple[500]} strokeWidth={2} unit="C" />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </div>
    </div>
  )
}