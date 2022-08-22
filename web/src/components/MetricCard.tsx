import * as React from 'react'
import {Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from 'recharts'
import * as moment from 'moment'
import {Metric, MetricRecord} from '../types'
import * as api from '../api'
import Card from '@mui/material/Card'
import {useTheme} from '@mui/material/styles'
import {CardHeader} from '@mui/material'


const EVERY_5_SECONDS = 5000

export type MetricCardProps = Metric & {}
export const MetricCard: React.FC<MetricCardProps> = ({id, name}) => {
  const theme = useTheme()
  const [data, setData] = React.useState<MetricRecord[]>([])

  React.useEffect(() => {
    api.fetchMetricData(id).then(setData)
    const interval = setInterval(() => {
      api.fetchMetricData(id).then(setData)
    }, EVERY_5_SECONDS)
    return () => clearInterval(interval)
  }, [])

  return (
    <Card>
      <CardHeader title={name}/>
      <ResponsiveContainer height={300} width="99%">
        <LineChart data={data.sort((a, b) => a.Timespan - b.Timespan)}>
          <XAxis dataKey="Timespan" tickFormatter={val => moment.unix(val).format('hh:mm')}
                 stroke={theme.palette.secondary.light}/>
          <YAxis domain={['auto', 'auto']} stroke={theme.palette.secondary.light} strokeWidth={0} padding={{}}/>
          <Tooltip/>
          <Line type="monotone" dataKey="Value" stroke={theme.palette.primary.main} strokeWidth={3} unit="C"/>
        </LineChart>
      </ResponsiveContainer>
    </Card>
  )
}