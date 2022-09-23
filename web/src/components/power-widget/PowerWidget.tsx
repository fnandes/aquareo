import { useQuery } from '@tanstack/react-query'
import * as React from 'react'
import * as api from '../../api'
import { ControllableGear } from '../../types'
import { PowerWidgetItem } from './PowerWidgetItem'

export type PowerWidgetProps = {}

export const PowerWidget: React.FC<PowerWidgetProps> = () => {
  const { data: gears } = useQuery(['gears'], api.fetchGears)

  const handleGearToggle = (gear: ControllableGear) => {
    console.log(gear)
  }

  return gears ? (
    <div className='border border-black/25 bg-white'>
      {gears.length ? gears.map(gear => (
        <PowerWidgetItem key={gear.id} gear={gear} onToggle={() => handleGearToggle(gear)} />
      )) : <div>not</div>}
    </div>
  ) : null
}