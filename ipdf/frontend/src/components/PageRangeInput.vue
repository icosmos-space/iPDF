<script lang="ts" setup>
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

defineProps<{
  modelValue: string
  placeholder?: string
}>()
const emit = defineEmits<{ (e: 'update:modelValue', v: string): void }>()
</script>

<template>
  <div class="grid gap-2">
    <Label>页码范围</Label>
    <Input
      :model-value="modelValue"
      :placeholder="placeholder || '例如 1-3,5'"
      @update:model-value="emit('update:modelValue', String($event))"
    />
  </div>
</template>
