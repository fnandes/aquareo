import * as React from 'react'
import { Config } from '../types'

const context = React.createContext<Config>({
  name: '',
  customMetrics: [],
  temperatureController: {
    enabled: false
  }
})

export type ConfigProviderProps = {
  config: Config
  children: React.ReactElement
}
export const ConfigProvider: React.FC<ConfigProviderProps> = ({ children, config }) => (
  <context.Provider value={config}>
    {children}
  </context.Provider>
)

export const useConfig = () => React.useContext(context)