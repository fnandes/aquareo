import * as React from 'react'

export type ButtonProps = {
  label: string
}
export const Button: React.FC<ButtonProps> = ({ label }) => (
  <button className="py-2 px-4 mr-2 mb-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300">
    {label}
  </button>
)