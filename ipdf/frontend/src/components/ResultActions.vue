<script lang="ts" setup>
import { FolderOpen } from '@lucide/vue'
import { OpenInExplorer } from '../../wailsjs/go/main/App'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

const props = defineProps<{
  paths: string[]
}>()

async function openFirst() {
  if (!props.paths.length) return
  await OpenInExplorer(props.paths[0])
}

async function openDir() {
  if (!props.paths.length) return
  const p = props.paths[0]
  const dir = p.replace(/[\\/][^\\/]+$/, '') || p
  await OpenInExplorer(dir)
}
</script>

<template>
  <Card v-if="paths.length" class="mt-4">
    <CardHeader class="pb-3">
      <CardTitle class="text-base">输出</CardTitle>
    </CardHeader>
    <CardContent class="space-y-3">
      <ul class="space-y-1 text-sm text-muted-foreground break-all">
        <li v-for="p in paths" :key="p">{{ p }}</li>
      </ul>
      <div class="flex flex-wrap gap-2">
        <Button type="button" @click="openFirst">
          <FolderOpen class="size-4" />
          在资源管理器中显示
        </Button>
        <Button v-if="paths.length > 1" type="button" variant="outline" @click="openDir">
          打开输出目录
        </Button>
      </div>
    </CardContent>
  </Card>
</template>
