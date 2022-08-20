import * as React from 'react'
import { Dashboard } from './Dasboard'

export const App: React.FC = () => (
  <div className="bg-gray-900 text-white font-sans h-screen w-screen overflow-auto relative">
    <nav className="bg-purple-800 p-4 mb-4 shadow-lg">
      test
    </nav>
    <main className="mx-4">
      <Dashboard />
    </main>
  </div>
)