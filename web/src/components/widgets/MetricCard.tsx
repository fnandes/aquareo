import * as React from 'react'
import { useMantineTheme, Card, Group, Text } from '@mantine/core'
import { useQuery } from '@tanstack/react-query'
import { CartesianGrid, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as moment from 'dayjs'
import * as api from '../../api'

export type MetricCardProps = {
  bucket: string
  title: string
  metricUnit: string
}

export const MetricCard: React.FC<MetricCardProps> = ({ bucket, title, metricUnit }) => {
  const { colors } = useMantineTheme()
  const tickFormatter = (val: number) => moment.unix(val).format('L LT')
  const { data = [] } = useQuery(['metric', bucket], api.metrics(bucket).fetchAll)

  return (
    <Card shadow="xs" p="sm">
      <Group position="apart" mb="sm">
        <Text fw={500}>{title}</Text>
      </Group>
      <Card.Section>
        <ResponsiveContainer height={150} width="99%">
          <LineChart data={data}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="timespan" tickFormatter={tickFormatter} stroke={colors.violet[5]} fontSize="small" />
            <YAxis stroke={colors.violet[5]} />
            <Tooltip />
            <Line type="monotone" connectNulls dataKey="value" stroke={colors.grape[5]} strokeWidth={3} unit={metricUnit} />
          </LineChart>
        </ResponsiveContainer>
      </Card.Section>
    </Card>
  )
}