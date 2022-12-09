import * as React from 'react'
import { Outlet } from 'react-router-dom'
import { AppShell, Container, Footer } from '@mantine/core'
import { AppHeader } from './AppHeader'
import { AppNav } from './AppNav'

export const Layout: React.FC = () => (
  <AppShell
    padding="md"
    header={<AppHeader />}
    navbar={<AppNav />}
    footer={<Footer height={30} px="sm">Aquareo</Footer>}>
    <Container size="xl">
      <Outlet />
    </Container>
  </AppShell>
)