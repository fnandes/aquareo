import * as React from 'react'
import { useConfig } from '../../hooks/useConfig'
import { Button } from '../button'
import { Card, CardBody } from '../card'
import { MetricCard } from './MetricCard'

export type HomeProps = unknown

export const Home: React.FC<HomeProps> = () => {
  const config = useConfig()

  return (
    <div>
      <div className="mb-4">
        <Card title="Add manual entry">
          <CardBody>
            <Button label="Temperature" />
            {config.customMetrics?.map(metric => (
              <Button key={metric.id} label={metric.displayName} />
            ))}
          </CardBody>
        </Card>
      </div>
      {config?.temperatureController?.enabled ? (
        <div className='mb-4'>
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