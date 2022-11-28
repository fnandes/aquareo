import * as React from 'react'
import { DatePicker } from '@mantine/dates'
import { useMutation } from '@tanstack/react-query'
import { Modal, NumberInput, Stack, Group, Button } from '@mantine/core'
import * as dayjs from 'dayjs'
import { Controller, SubmitHandler, useForm } from 'react-hook-form'
import { useConfig } from '../../hooks/useConfig'
import * as api from '../../api'

export type MetricFormState = {
  timespan: Date
  value: number
}

export type MetricEntryFormModalProps = {
  metricId: string
  isOpen: boolean
  onClose: () => any
}
export const MetricEntryFormModal: React.FC<MetricEntryFormModalProps> = ({ metricId, isOpen, onClose }) => {
  const config = useConfig()
  const metric = config.customMetrics.find(m => m.id === metricId)

  const { control, handleSubmit, reset } = useForm<MetricFormState>({
    defaultValues: {
      timespan: dayjs().toDate(),
      value: 0
    }
  })

  const addMetricEntry = useMutation((d: MetricFormState) => api.addMetricEntry(`cm_${metricId}`, {
    timespan: dayjs(d.timespan).unix(),
    value: d.value
  }))

  const handleSaveClick: SubmitHandler<MetricFormState> = values => {
    addMetricEntry.mutate(values)
    reset()
    onClose()
  }

  return (
    <Modal
      opened={!!metric && isOpen}
      onClose={onClose}
      title={metric ? `Log ${metric?.displayName}` : ''}>
      <form onSubmit={handleSubmit(handleSaveClick)}>
        <Stack>
          <Controller
            name="timespan"
            control={control}
            rules={{ required: true }}
            render={({ field }) => (
              <DatePicker
                label="Date"
                withAsterisk
                {...field} />
            )} />
          <Controller
            name="value"
            control={control}
            rules={{ required: true }}
            render={({ field }) => (
              <NumberInput
                label="Value"
                defaultValue={0}
                precision={4}
                step={0.01}
                withAsterisk
                {...field} />
            )} />
        </Stack>
        <Group mt="xl" grow>
          <Button variant="outline" onClick={onClose}>Cancel</Button>
          <Button type="submit">Save</Button>
        </Group>
      </form>
    </Modal>
  )
}