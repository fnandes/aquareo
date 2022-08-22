import * as React from 'react'
import {Metric} from '../types'
import * as api from '../api'
import {MetricCard} from './MetricCard'
import Grid from '@mui/material/Grid'

export type DashboardProps = {}

export const Dashboard: React.FC<DashboardProps> = () => {
  const [metrics, setMetrics] = React.useState<Metric[]>([])

  React.useEffect(() => {
    api.listMetrics().then(setMetrics)
  }, [])

  return metrics ? (
    <Grid container spacing={2}>
      {metrics.map(metric => (
        <Grid key={metric.id} item md={6} sm={4} xs={12}>
          <MetricCard key={metric.id} id={metric.id} name={metric.name}/>
        </Grid>
      ))}
    </Grid>
  ) : <span>Nothing to see here ... =/</span>
}