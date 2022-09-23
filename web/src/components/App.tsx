import { useIsFetching } from '@tanstack/react-query'
import * as React from 'react'
import { Home } from './home'

export const App: React.FC = () => {
  const isFetching = useIsFetching()

  return (
    <div className='font-sans bg-slate-200 flex flex-col w-screen h-screen overflow-hidden'>
      <nav className='p-4 mb-4 flex'>
        <h1 className='text-2xl leading-none font-semibold'>Aquareo</h1>
        <div>
          {isFetching ? 'loading data ...' : null}
        </div>
      </nav>
      <main className='flex-1 overflow-auto'>
        <div className='m-4'>
          <Home />
        </div>
      </main>
    </div>
  )
}