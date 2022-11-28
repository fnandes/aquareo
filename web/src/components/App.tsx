import * as React from 'react'
import { MantineProvider, AppShell, Container } from '@mantine/core'
import { useQuery } from '@tanstack/react-query'
import * as api from '../api'
import { ConfigProvider } from '../hooks/useConfig'
import { Config } from '../types'
import { Home } from './home'
import { NavBar } from './navbar'

export const App: React.FC = () => {
  const { data: config } = useQuery(['config'], api.fetchConfig)

  return (
    <MantineProvider
      withNormalizeCSS
      withGlobalStyles
      theme={{
        primaryColor: 'indigo',
        primaryShade: 9
      }}>
      <ConfigProvider config={config || {} as Config}>
        <AppShell
          padding="md"
          header={<NavBar />}>
          <Container size="xl">
            {config ? (
              <Home />
            ) : 'Loading ...'}
          </Container>
        </AppShell>
      </ConfigProvider>
    </MantineProvider>
  )
}