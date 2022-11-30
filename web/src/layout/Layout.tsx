import * as React from 'react'
import { Outlet } from 'react-router-dom'
import { AppShell, Container } from '@mantine/core'
import { Navbar } from './Navbar'

export const Layout: React.FC = () => (
  <AppShell
    padding="md"
    header={<Navbar />}>
    <Container size="xl">
      <Outlet />
    </Container>
  </AppShell>
)