import * as React from 'react'

export type CardProps = React.PropsWithChildren & {
  title: string
}
export const Card: React.FC<CardProps> = ({ title, children }) => (
  <div className='bg-white rounded shadow'>
    <div className='py-2 px-4 font-semibold text-sm text-slate-500'>
      {title}
    </div>
    {children}
  </div>
)

export const CardBody: React.FC<React.PropsWithChildren> = ({ children }) => (
  <div className="px-4 py-2">{children}</div>
)