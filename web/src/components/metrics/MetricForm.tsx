import * as React from 'react'
import { DatePicker } from '@mantine/dates'
import { useMutation } from '@tanstack/react-query'
import { NumberInput, Stack, Group, Button } from '@mantine/core'
import * as dayjs from 'dayjs'
import { Controller, SubmitHandler, useForm } from 'react-hook-form'
import * as api from '../../api'
import { ContextModalProps } from '@mantine/modals'

export type MetricFormState = {
  timespan: Date
  value: number
}

export const AddEntryModal = ({ context, id, innerProps }: ContextModalProps<{ bucket: string }>) => {
  const { control, handleSubmit } = useForm<MetricFormState>({
    defaultValues: {
      timespan: dayjs().toDate()
    }
  })

  const addMetricEntry = useMutation(['metric', innerProps.bucket], api.metrics(innerProps.bucket).addEntry)

  const handleSaveClick: SubmitHandler<MetricFormState> = ({ timespan, value }) => {
    addMetricEntry.mutate({ timespan: dayjs(timespan).unix(), value })
    context.closeModal(id)
  }

  return (
    <form onSubmit={handleSubmit(handleSaveClick)}>
      <Stack>
        <Controller
          name="timespan"
          control={control}
          rules={{ required: true }}
          render={({ field }) => (
            <DatePicker label="Date" withAsterisk {...field} />
          )} />
        <Controller
          name="value"
          control={control}
          rules={{ required: true }}
          render={({ field }) => (
            <NumberInput label="Value" precision={4} step={0.01} withAsterisk {...field} />
          )} />
      </Stack>
      <Group mt="xl" grow>
        <Button variant="outline" onClick={() => context.closeModal(id, true)}>Cancel</Button>
        <Button type="submit">Save</Button>
      </Group>
    </form>
  )
}
