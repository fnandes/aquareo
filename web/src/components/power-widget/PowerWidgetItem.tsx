import * as moment from 'moment'
import * as React from 'react'
import { ControllableGear } from 'web/src/types'

export type PowerWidgetItemProps = {
  gear: ControllableGear
  onToggle: () => void
}

export const PowerWidgetItem: React.FC<PowerWidgetItemProps> = ({ gear, onToggle }) => (
  <div key={gear.id} className='px-4 py-2 leading-1 border-b flex items-center border-black/10'>
    <div className='flex-1'>
      <p className='font-semibold text-lg'>{gear.name}</p>
      <p className='text-sm italic text-slate-500'>updated {moment(gear.lastUpdate).fromNow()}</p>
    </div>
    <div>
      <label className="relative flex justify-between items-center group text-xl">
        <input
          type="checkbox"
          checked={gear.isOn}
          onChange={onToggle}
          className="absolute left-1/2 -translate-x-1/2 w-full h-full peer appearance-none rounded-md" />
        <span className="w-16 h-10 flex items-center flex-shrink-0 ml-4 p-1 bg-gray-300 rounded-full duration-300 ease-in-out peer-checked:bg-green-500 after:w-8 after:h-8 after:bg-white after:rounded-full after:shadow-md after:duration-300 peer-checked:after:translate-x-6"></span>
      </label>
    </div>
  </div>
)