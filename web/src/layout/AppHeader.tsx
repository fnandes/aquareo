import * as React from 'react'
import { Header, Text } from '@mantine/core'

export const AppHeader: React.FC = () => (
  <Header height={50} p="sm" bg="indigo.9">
    <Text c="white" fw={500} size="lg">Aquareo</Text>
  </Header>
)