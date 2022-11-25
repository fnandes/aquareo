import * as React from 'react'

export type ButtonProps = {
  label: string
}
export const Button: React.FC<ButtonProps> = ({ label }) => (
  <button className="bg-blue-400 hover:bg-blue-600 rounded-lg ml-2 py-1 px-3">
    {label}
  </button>
)