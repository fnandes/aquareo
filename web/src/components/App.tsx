import * as React from 'react'
import { Dashboard } from '../pages/Dasboard'

export const App: React.FC = () => (
  <div className='font-sans bg-slate-200 flex flex-col w-screen h-screen overflow-hidden'>
    <nav className='bg-purple-800 text-white p-4 shadow-lg'>
      <h1 className='text-2xl font-semibold'>Aquareo</h1>
    </nav>
    <main className='flex-1 overflow-auto'>
      <div className='m-4'>
        <Dashboard />
      </div>
    </main>
  </div>
)