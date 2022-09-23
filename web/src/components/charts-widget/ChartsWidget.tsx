import * as React from 'react'
import * as api from '../../api'
import { Metric } from '../../types'
import { MetricCard } from './MetricCard'

export type ChartsWidgetProps = {}

export const ChartsWidget: React.FC<ChartsWidgetProps> = () => {
  const [metrics, setMetrics] = React.useState<Metric[]>([])

  React.useEffect(() => {
    api.listMetrics()
      .then(setMetrics)
      .catch(console.error)
  }, [])

  if (!metrics) {
    return null
  }

  return (
    <div>
      {metrics.map(metric => (
        <div key={metric.id} className='basis-1/2 mb-4'>
          <MetricCard key={metric.id} id={metric.id} name={metric.name} />
        </div>
      ))}
    </div>
  )
}