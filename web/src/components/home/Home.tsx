import * as React from 'react'
import { Button, Paper, Grid, Group } from '@mantine/core'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from '../metrics'
import { openContextModal } from '@mantine/modals'

export type HomeProps = unknown

export const Home: React.FC<HomeProps> = () => {
  const config = useConfig()

  const openAddMetricModal = (metricId: string, bucket: string) => {
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
      <Grid gutter="sm">
        <Grid.Col xs={12}>
          <Paper shadow="sm" p="xs">
            <Group spacing="xs">
              {config.customMetrics?.map(metric => (
                <Button
                  key={metric.id}
                  onClick={() => openAddMetricModal(metric.id, `cm_${metric.id}`)}>{metric.displayName}</Button>
              ))}
            </Group>
          </Paper>
        </Grid.Col>
        {config?.temperatureController?.enabled ? (
          <Grid.Col xs={12}>
            <MetricCard bucket="temperature" title="Temperature" metricUnit="C" />
          </Grid.Col>
        ) : null}
        {config.customMetrics.length && config.customMetrics.map(metric => (
          <Grid.Col key={metric.id} xs={6}>
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