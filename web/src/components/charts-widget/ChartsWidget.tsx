import * as React from 'react'
import { useConfig } from '../../hooks/useConfig'
import { MetricCard } from './MetricCard'

export type ChartsWidgetProps = {}

export const ChartsWidget: React.FC<ChartsWidgetProps> = () => {
  const config = useConfig()

  return (
    <div>
      {config?.temperatureController?.enabled ? (
        <div className='basis-1/2 mb-4'>
          <MetricCard bucket="temperature" title="Temperature" metricUnit="C" />
        </div>
      ) : null}
    </div>
  )
}