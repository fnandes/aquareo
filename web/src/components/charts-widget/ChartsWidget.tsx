import * as React from 'react'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from './MetricCard'

export type ChartsWidgetProps = {}

export const ChartsWidget: React.FC<ChartsWidgetProps> = () => {
  const config = useConfig()

  return (
    <div>
      {config?.temperatureController?.enabled ? (
        <div className='mb-8'>
          <MetricCard bucket="temperature" title="Temperature" metricUnit="C" />
        </div>
      ) : null}
      {config.customMetrics.length ? (
        <div className="flex flex-wrap -m-2">
          {config.customMetrics.map(metric => (
            <div key={metric.id} className="basis-1/2 p-2">
              <MetricCard bucket={`cm_${metric.id}`} title={metric.displayName} metricUnit={metric.metricUnit} />
            </div>
          ))}
        </div>
      ) : null}
    </div>
  )
}