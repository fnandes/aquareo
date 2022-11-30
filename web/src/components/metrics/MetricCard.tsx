import * as React from 'react'
import { useMantineTheme, Text, Card, Group, Button } from '@mantine/core'
import { IconList } from '@tabler/icons'
import { useQuery } from '@tanstack/react-query'
import { CartesianGrid, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import * as moment from 'dayjs'
import * as api from '../../api'
import { useHref } from 'react-router-dom'

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
    <Card shadow="xs" p="xs">
      <Group position="apart" mb="sm">
        <Text weight={500}>{title}</Text>
        <div>
          <Button component="a" leftIcon={<IconList />} href={useHref(`metrics/${bucket}`)} variant="light" size="xs">view list</Button>
        </div>
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
    </Card >
  )
}