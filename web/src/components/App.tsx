import * as React from 'react'
import { MantineProvider, AppShell, Container, MantineThemeOverride } from '@mantine/core'
import { useQuery } from '@tanstack/react-query'
import * as api from '../api'
import { ConfigProvider } from '../hooks/useConfig'
import { Config } from '../types'
import { Home } from './home'
import { NavBar } from './navbar'
import { createHashRouter, RouterProvider } from 'react-router-dom'
import { ModalsProvider } from '@mantine/modals'
import { AddEntryModal, MetricEntries } from './metrics'

const theme: MantineThemeOverride = {
  primaryColor: 'indigo',
  primaryShade: 9
}

const router = createHashRouter([{
  path: '/',
  element: <Home />
}, {
  path: '/metrics/:bucket',
  element: <MetricEntries />
}])

export const App: React.FC = () => {
  const { data: config } = useQuery(['config'], api.fetchConfig)

  return (
    <MantineProvider withNormalizeCSS withGlobalStyles theme={theme}>
      <ModalsProvider
        labels={{ confirm: 'Save', cancel: 'Cancel' }}
        modals={{ addMetricEntry: AddEntryModal }}>
        <ConfigProvider config={config || {} as Config}>
          <AppShell
            padding="md"
            header={<NavBar />}>
            <Container size="xl">
              {config ? (
                <RouterProvider router={router} />
              ) : 'Loading ...'}
            </Container>
          </AppShell>
        </ConfigProvider>
      </ModalsProvider>
    </MantineProvider>
  )
}