import Container from '@mui/material/Container'
import Toolbar from '@mui/material/Toolbar'
import { useQuery } from '@tanstack/react-query'
import * as React from 'react'
import * as api from '../api'
import { ConfigProvider } from '../hooks/useConfig'
import { Config } from '../types'
import { Home } from './home'
import { NavBar } from './navbar'

export const App: React.FC = () => {
  const { data: config } = useQuery(['config'], api.fetchConfig)

  return (
    <ConfigProvider config={config || {} as Config}>
      <div>
        <NavBar />
        <Toolbar />
        <Container>

          <main className='flex-1 overflow-auto'>
            <div className='container mx-auto'>
              {config ? (
                <div className='m-4'>
                  <Home />
                </div>

              ) : 'Loading ...'}
            </div>
          </main>
        </Container>
      </div>
    </ConfigProvider>
  )
}