import { Card, Table, Title } from '@mantine/core'
import * as React from 'react'

export const MetricEntries: React.FC = () => (
  <div>
    <Title order={1} size="h2" mb="lg">Phosphate</Title>
    <Card shadow="sm" withBorder>
      <Card.Section>
        <Table>
          <thead>
            <tr>
              <th>Timespan</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>25.11.2022</td>
              <td>3.1</td>
            </tr>
          </tbody>
        </Table>
      </Card.Section>
    </Card>
  </div>
)