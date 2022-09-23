import * as React from 'react'
import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as colors from 'tailwindcss/colors'
import * as moment from 'moment'
import { MetricRecord } from '../../types'

export type MetricCardLineChartProps = {
  title: string
  unit: string
  data: MetricRecord[]
  domain: [number, number]
}

export const MetricCardLineChart: React.FC<MetricCardLineChartProps> = ({ title, unit, data, domain }) => {
  const tickFormatter = (val: number) => moment.unix(val).format('hh:mm')

  return (
    <div className='flex bg-white shadow-md shadow-black/5 rounded-md overflow-hidden'>
      <div className='flex-1'>
        <div className='p-2 font-semibold text-slate-500'>
          <h3>{title}</h3>
        </div>
        <ResponsiveContainer height={150} width="99%">
          <LineChart data={data.sort((a, b) => a.Timespan - b.Timespan)}>
            <XAxis dataKey="Timespan" tickFormatter={tickFormatter} stroke={colors.purple[500]} />
            <YAxis domain={domain} stroke={colors.purple[500]} strokeWidth={0} padding={{}} />
            <Tooltip />
            <Line type="monotone" dataKey="Value" stroke={colors.purple[500]} strokeWidth={3} unit={unit} />
          </LineChart>
        </ResponsiveContainer>
      </div>
      <div className='flex-none basis-1/6 bg-gradient-to-tr from-indigo-500 to-indigo-900 text-white flex flex-col items-center justify-center'>
        <div className='mx-4'>
          <span className='text-4xl font-semibold'>43.3</span>
        </div>
      </div>
    </div>
  )
}