import { Card, Table, Title, ActionIcon, Group, Button } from '@mantine/core'
import { IconTrash } from '@tabler/icons'
import { useParams } from 'react-router-dom'
import * as React from 'react'
import { useQuery, useMutation } from '@tanstack/react-query'
import * as api from '../../api'
import * as dayjs from 'dayjs'
import { useConfig } from '../../hooks/useConfig'
import { openContextModal } from '@mantine/modals'

export const MeasurementsList: React.FC = () => {
  const config = useConfig()

  const { testId = '' } = useParams<{ testId: string }>()
  const bucket = `cm_${testId}`
  const customMetric = React.useMemo(() => config.customMetrics.find(m => m.id === testId), [testId])

  const { data = [] } = useQuery(['metric', bucket], api.metrics(bucket).fetchAll)

  const deleteMetricEntry = useMutation(['metric', bucket], api.metrics(bucket).deleteEntry)

  const handleDeleteClick = (timespan: number) => {
    if (confirm('Are you sure?')) {
      deleteMetricEntry.mutate(timespan)
    }
  }

  return customMetric ? (
    <div>
      <Group position="apart" mb="sm">
        <Title order={2} mb="lg">{customMetric.displayName}</Title>
        <Button
          variant="filled"
          onClick={() => openContextModal({
            modal: 'addMetricEntry',
            title: `Log ${customMetric.displayName}`,
            innerProps: { bucket }
          })}>Log Entry</Button>
      </Group>
      <Card radius="sm" shadow="xs" withBorder>
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
  ) : <span>metric not found</span>
}
