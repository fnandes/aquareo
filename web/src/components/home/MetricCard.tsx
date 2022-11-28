import * as React from 'react'
import { useMantineTheme, Paper, Text } from '@mantine/core'
import { useQuery } from '@tanstack/react-query'
import { CartesianGrid, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as moment from 'moment'
import * as api from '../../api'

export type MetricCardProps = {
  bucket: string
  title: string
  metricUnit: string
}

export const MetricCard: React.FC<MetricCardProps> = ({ bucket, title, metricUnit }) => {
  const { colors } = useMantineTheme()
  const tickFormatter = (val: number) => moment.unix(val).format('D/M hh:mm')
  const { data = [] } = useQuery(['metric', bucket], () => api.fetchMetricData(bucket))

  return (
    <Paper shadow="xs" p={0}>
      <Text my="sm" px="md" weight={500}>{title}</Text>
      <ResponsiveContainer height={150} width="99%">
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="timespan" tickFormatter={tickFormatter} stroke={colors.violet[5]} fontSize="small" />
          <YAxis stroke={colors.violet[5]} />
          <Tooltip />
          <Line type="monotone" connectNulls dataKey="value" stroke={colors.grape[5]} strokeWidth={3} unit={metricUnit} />
        </LineChart>
      </ResponsiveContainer>
    </Paper>
  )
}