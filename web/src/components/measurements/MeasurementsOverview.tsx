import { Grid } from '@mantine/core'
import * as React from 'react'
import { getMetricName } from '../../utils'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from '../widgets'

export const MeasurementsOverview: React.FC = () => {
  const config = useConfig()

  return (
    <>
      <Grid gutter="md">
        {config.customMetrics.length && config.customMetrics.map(metric => (
          <Grid.Col key={metric} sm={6}>
            <MetricCard
              bucket={`cm_${metric}`}
              title={getMetricName(metric)} />
          </Grid.Col>
        ))}
      </Grid>
    </>
  )
}