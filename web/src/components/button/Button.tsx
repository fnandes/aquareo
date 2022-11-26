import * as React from 'react'

export type ButtonProps = {
  label: string
}
export const Button: React.FC<ButtonProps> = ({ label }) => (
  <button className="bg-indigo-600 hover:bg-indigo-500 text-slate-200 rounded-lg ml-2 py-1 px-3 font-semibold">
    {label}
  </button>
)