import { Card, Table, Title, ActionIcon } from '@mantine/core'
import { IconTrash } from '@tabler/icons'
import { useParams } from 'react-router-dom'
import * as React from 'react'
import { useQuery, useMutation } from '@tanstack/react-query'
import * as api from '../../api'
import * as dayjs from 'dayjs'

export const MetricEntries: React.FC = () => {
  const { bucket } = useParams<{ bucket: string }>()
  const { data = [] } = useQuery(['metric', bucket], api.metrics(bucket).fetchAll)

  const deleteMetricEntry = useMutation(['metric', bucket], api.metrics(bucket).deleteEntry)

  const handleDeleteClick = (timespan: number) => {
    if (confirm('Are you sure?')) {
      deleteMetricEntry.mutate(timespan)
    }
  }

  return (
    <div>
      <Title order={1} size="h2" mb="lg">Phosphate</Title>
      <Card shadow="sm" withBorder>
        <Card.Section>
          <Table>
            <thead>
              <tr>
                <th>Timespan</th>
                <th>Value</th>
                <th></th>
              </tr>
            </thead>
            {data.length ? (
              <tbody>
                {data.map(entry => (
                  <tr key={entry.timespan}>
                    <td>{dayjs.unix(entry.timespan).format('L LT')}</td>
                    <td>{entry.value}</td>
                    <td>
                      <ActionIcon variant="subtle" color="red" onClick={() => handleDeleteClick(entry.timespan)}>
                        <IconTrash />
                      </ActionIcon>
                    </td>
                  </tr>
                ))}
              </tbody>
            ) : (
              <tbody>
                <tr>
                  <td colSpan={2}>no records found</td>
                </tr>
              </tbody>
            )}
          </Table>
        </Card.Section>
      </Card>
    </div>
  )
}
