import * as React from 'react'
import { Modal, TextInput, NumberInput, Stack, Group, Button } from '@mantine/core'
import { useConfig } from '../../hooks/useConfig'

export type MetricFormModalProps = {
  metricId: string
  isOpen: boolean
  onClose: () => any
}
export const MetricFormModal: React.FC<MetricFormModalProps> = ({ metricId, isOpen, onClose }) => {
  const config = useConfig()
  const metric = config.customMetrics.find(m => m.id === metricId)

  const handleSaveClick = () => {
    onClose()
  }

  return (
    <Modal
      opened={!!metric && isOpen}
      onClose={onClose}
      title={metric ? `Log ${metric?.displayName}` : ''}>
      <Stack>
        <TextInput
          label="Date"
          withAsterisk
        />
        <NumberInput
          label="Value"
          precision={4}
          step={0.01}
          withAsterisk
        />
      </Stack>
      <Group mt="xl" grow>
        <Button variant="outline" onClick={onClose}>
          Cancel
        </Button>
        <Button onClick={handleSaveClick}>
          Save
        </Button>
      </Group>
    </Modal>
  )
}