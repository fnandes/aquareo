import * as React from 'react'
import { PowerWidget } from '../power-widget/PowerWidget'
import { ChartsWidget } from '../charts-widget'

export type HomeProps = {}

export const Home: React.FC<HomeProps> = () => {
  return (
    <div className='flex mb-8'>
      <div className='flex-1 px-4'>
        <PowerWidget />
      </div>
      <div className='flex-1 px-4'>
        <ChartsWidget />
      </div>
    </div>
  )
}