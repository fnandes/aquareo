import * as React from 'react'
import { Outlet } from 'react-router-dom'
import { AppShell, Container, Footer } from '@mantine/core'
import { Navbar } from './Navbar'

export const Layout: React.FC = () => (
  <AppShell
    padding="md"
    header={<Navbar />}
    footer={<Footer>Aquareo</Footer>}>
    <Container size="xl">
      <Outlet />
    </Container>
  </AppShell>
)