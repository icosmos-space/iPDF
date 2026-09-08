<script lang="ts" setup>
import { AlertCircle, CheckCircle2, Loader2 } from '@lucide/vue'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'

defineProps<{
  loading?: boolean
  error?: string
  message?: string
}>()
</script>

<template>
  <div v-if="loading || error || message" class="mt-4">
    <Alert v-if="loading">
      <Loader2 class="size-4 animate-spin" />
      <AlertTitle>处理中</AlertTitle>
      <AlertDescription>正在处理，请稍候…</AlertDescription>
    </Alert>
    <Alert v-else-if="error" variant="destructive">
      <AlertCircle class="size-4" />
      <AlertTitle>出错了</AlertTitle>
      <AlertDescription>{{ error }}</AlertDescription>
    </Alert>
    <Alert v-else-if="message" class="border-emerald-200 bg-emerald-50 text-emerald-900">
      <CheckCircle2 class="size-4 text-emerald-600" />
      <AlertTitle>完成</AlertTitle>
      <AlertDescription>{{ message }}</AlertDescription>
    </Alert>
  </div>
</template>
