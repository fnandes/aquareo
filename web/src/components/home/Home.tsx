import * as React from 'react'
import { Button, Paper, Grid, Group } from '@mantine/core'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from './MetricCard'
import { MetricFormModal } from '../metricFormModal/MetricFormModal'

export type HomeProps = unknown

export const Home: React.FC<HomeProps> = () => {
  const config = useConfig()
  const [metricModalOpen, setMetricModalOpen] = React.useState(false)
  const [logMetricId, setLogMetricId] = React.useState<string>('')

  const handleMetricBtnClick = (metricId: string) => {
    setMetricModalOpen(true)
    setLogMetricId(metricId)
  }

  const handleMetricModalClose = () => {
    setMetricModalOpen(false)
    setLogMetricId('')
  }

  const handleMetricModalSave = () => {
    setMetricModalOpen(false)
    setLogMetricId('')
    alert('test')
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
                  variant="outline"
                  size="sm"
                  onClick={() => handleMetricBtnClick(metric.id)}>{metric.displayName}</Button>
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
      <MetricFormModal
        isOpen={metricModalOpen}
        onClose={handleMetricModalClose}
        metricId={logMetricId} />
    </>
  )
}