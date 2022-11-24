import { useIsFetching, useQuery } from '@tanstack/react-query'
import * as React from 'react'
import * as api from '../api'
import { ConfigProvider } from '../hooks/useConfig'
import { Config } from '../types'
import { Home } from './home'

export const App: React.FC = () => {
  const isFetching = useIsFetching()
  const { data: config } = useQuery(['config'], api.fetchConfig)

  return (
    <ConfigProvider config={config || {} as Config}>
      <div className='font-sans bg-slate-200 flex flex-col w-screen h-screen overflow-hidden'>
        <nav className='p-4 mb-4 flex'>
          <h1 className='text-2xl leading-none font-semibold'>Aquareo</h1>
          <div>
            {isFetching ? 'loading data ...' : null}
          </div>
        </nav>
        <main className='flex-1 overflow-auto'>
          {config ? (
            <div className='m-4'>
              <Home />
            </div>

          ) : 'Loading ...'}
        </main>
      </div>
    </ConfigProvider>
  )
}