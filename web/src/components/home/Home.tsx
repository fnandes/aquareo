import * as React from 'react'
import { useConfig } from '../../hooks/useConfig'
import { Button } from '../button'
import { Card } from '../card'
import { ChartsWidget } from '../charts-widget'

export type HomeProps = unknown

export const Home: React.FC<HomeProps> = () => {
  const config = useConfig()

  return (
    <div className="flex">
      <div className="basis-1/3">
        <Card title="Manual log">
          <div className="p-4">
            {config.customMetrics.length && config.customMetrics.map(metric => (
              <Button key={metric.id} label={metric.displayName} />
            ))}
          </div>
        </Card>
      </div>
      <div className="ml-4 basis-2/3">
        <ChartsWidget />
      </div>
    </div>
  )
}