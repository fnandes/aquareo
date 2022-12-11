import * as React from 'react'
import { Navbar, NavLink } from '@mantine/core'
import { IconHome, IconTestPipe } from '@tabler/icons'
import { NavLink as RouterLink } from 'react-router-dom'
import { useConfig } from '../hooks/useConfig'
import { getMetricName } from '../utils'


export const AppNav: React.FC = () => {
  const { customMetrics } = useConfig()

  return (
    <Navbar width={{ base: 300 }} p="xs">
      <Navbar.Section grow mt="xs">
        <NavLink icon={<IconHome size={20} />} label="Dashboard" component={RouterLink} to="/" />
        <NavLink icon={<IconTestPipe size={20} />} label="Test Measurements" childrenOffset={38}>
          <NavLink label="Overview" component={RouterLink} to="/measurements" />
          {customMetrics.map(metric => (
            <NavLink
              key={metric}
              label={getMetricName(metric)}
              component={RouterLink}
              to={`/measurements/${metric}`} />
          ))}
        </NavLink>
      </Navbar.Section>
    </Navbar>
  )
}