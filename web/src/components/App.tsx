import * as React from 'react'
import { MantineProvider, MantineThemeOverride } from '@mantine/core'
import { useQuery } from '@tanstack/react-query'
import * as api from '../api'
import { ConfigProvider } from '../hooks/useConfig'
import { Config } from '../types'
import { Home } from './home'
import { createHashRouter, RouterProvider } from 'react-router-dom'
import { ModalsProvider } from '@mantine/modals'
import { AddEntryModal, MeasurementsList } from './measurements'
import { Layout } from '../layout'
import { MeasurementsOverview } from './measurements'

const theme: MantineThemeOverride = {
  primaryColor: 'indigo',
  primaryShade: 9
}

const router = createHashRouter([{
  path: '/',
  element: <Layout />,
  children: [
    { index: true, element: <Home /> },
    { path: 'measurements', element: <MeasurementsOverview /> },
    { path: 'measurements/:testId', element: <MeasurementsList /> }
  ]
}])

export const App: React.FC = () => {
  const { data: config } = useQuery(['config'], api.fetchConfig)

  return (
    <MantineProvider withNormalizeCSS withGlobalStyles theme={theme}>
      <ModalsProvider
        labels={{ confirm: 'Save', cancel: 'Cancel' }}
        modals={{ addMetricEntry: AddEntryModal }}>
        <ConfigProvider config={config || {} as Config}>
          <>
            {config ? <RouterProvider router={router} /> : null}
          </>
        </ConfigProvider>
      </ModalsProvider>
    </MantineProvider>
  )
}