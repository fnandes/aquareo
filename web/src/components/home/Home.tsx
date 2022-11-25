import * as React from 'react'
import { ChartsWidget } from '../charts-widget'

export type HomeProps = {}

export const Home: React.FC<HomeProps> = () => {
  return (
    <div className='mx-4'>
      <div className='mb-4'>
        <ChartsWidget />
      </div>
    </div>
  )
}