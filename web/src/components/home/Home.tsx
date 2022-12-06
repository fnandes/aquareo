import * as React from 'react'
import { Grid, Title } from '@mantine/core'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from '../metrics'

export const Home: React.FC = () => {
  const config = useConfig()

  return (
    <>
      <Title my="sm" order={3}>Sensors</Title>
      <Grid gutter="sm">
        {config?.temperatureController?.enabled ? (
          <Grid.Col xs={6}>
            <MetricCard bucket="temperature" title="Temperature" metricUnit="C" />
          </Grid.Col>
        ) : null}
      </Grid>
    </>
  )
}