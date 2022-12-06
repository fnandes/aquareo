import { Group, Button, Grid, Paper } from '@mantine/core'
import { openContextModal } from '@mantine/modals'
import * as React from 'react'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from './MetricCard'

export const Metrics: React.FC = () => {
  const config = useConfig()

  const openAddEntryModal = (metricId: string, bucket: string) => {
    const metric = config.customMetrics.find(m => m.id === metricId)
    if (metric) {
      openContextModal({
        modal: 'addMetricEntry',
        title: `Log ${metric.displayName}`,
        innerProps: { bucket }
      })
    }
  }

  return (
    <>
      <Grid gutter="md">
        <Grid.Col xs={12}>
          <Paper p="xs" shadow="md">
            <Group spacing="xs">
              {config.customMetrics?.map(metric => (
                <Button
                  key={metric.id}
                  onClick={() => openAddEntryModal(metric.id, `cm_${metric.id}`)}>{metric.displayName}</Button>
              ))}
            </Group>
          </Paper>
        </Grid.Col>
        {config.customMetrics.length && config.customMetrics.map(metric => (
          <Grid.Col key={metric.id} sm={6}>
            <MetricCard
              bucket={`cm_${metric.id}`}
              title={metric.displayName}
              metricUnit={metric.metricUnit} />
          </Grid.Col>
        ))}
      </Grid>
    </>
  )
}