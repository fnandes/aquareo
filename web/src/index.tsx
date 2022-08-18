import * as React from 'react'
import { createRoot } from 'react-dom/client'
import { App } from './components/App'

import './theme.css'

document.addEventListener('DOMContentLoaded', () => {
  const root = createRoot(document.getElementById('root') as HTMLElement)
  root.render(<App />)
})