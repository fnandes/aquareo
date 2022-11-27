
import { Stack, Button, Grid, Paper } from '@mui/material'
import * as React from 'react'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from './MetricCard'

export type HomeProps = unknown

export const Home: React.FC<HomeProps> = () => {
  const config = useConfig()

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Paper elevation={1} sx={{ p: 2 }}>
          <Stack spacing={2} direction="row">
            <Button variant="outlined" size="small">Temperature</Button>
            {config.customMetrics?.map(metric => (
              <Button key={metric.id} variant="outlined" size="small">{metric.displayName}</Button>
            ))}
          </Stack>
        </Paper>
      </Grid>
      {config?.temperatureController?.enabled ? (
        <Grid item xs={12}>
          <MetricCard bucket="temperature" title="Temperature" metricUnit="C" />
        </Grid>
      ) : null}
      {config.customMetrics.length && config.customMetrics.map(metric => (
        <Grid key={metric.id} item xs={6}>
          <MetricCard
            bucket={`cm_${metric.id}`}
            title={metric.displayName}
            metricUnit={metric.metricUnit} />
        </Grid>
      ))}
    </Grid>
  )
}