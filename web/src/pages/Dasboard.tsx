import * as React from 'react'
import { Metric } from '../types'
import * as api from '../api'
import { MetricCard } from '../components/MetricCard'

export type DashboardProps = {}

export const Dashboard: React.FC<DashboardProps> = () => {
  const [metrics, setMetrics] = React.useState<Metric[]>([])

  React.useEffect(() => {
    api.listMetrics().then(setMetrics)
  }, [])

  return metrics ? (
    <div className='flex flex-wrap'>
      {metrics.map(metric => (
        <div key={metric.id} className='basis-1/2 mb-6'>
          <div className='mr-4'>
            <MetricCard key={metric.id} id={metric.id} name={metric.name} />
          </div>
        </div>
      ))}
    </div>
  ) : <span>Nothing to see here ... =/</span>
}