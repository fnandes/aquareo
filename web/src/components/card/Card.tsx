import * as React from 'react'

export type CardProps = {
  title: string
  children: React.ReactElement
}
export const Card: React.FC<CardProps> = ({ title, children }) => (
  <div className='bg-white rounded shadow-sm'>
    <div className='py-2 px-4 font-semibold text-sm text-slate-500'>
      {title}
    </div>
    <div>
      {children}
    </div>
  </div>
)