import * as React from 'react'
import { createRoot } from 'react-dom/client'
import * as dayjs from 'dayjs'
import * as localizedFormat from 'dayjs/plugin/localizedFormat'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { App } from './components/App'

dayjs.extend(localizedFormat)

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchInterval: 5000
    }
  }
})

document.addEventListener('DOMContentLoaded', () => {
  const root = createRoot(document.getElementById('root') as HTMLElement)
  root.render(
    <QueryClientProvider client={queryClient}>
      <App />
    </QueryClientProvider>)
})