<script lang="ts" setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { ArrowDown, ArrowUp, FileUp, Trash2, X } from '@lucide/vue'
import { OnFileDrop, OnFileDropOff } from '../../wailsjs/runtime/runtime'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { cn } from '@/lib/utils'

const props = withDefaults(
  defineProps<{
    acceptLabel?: string
    multiple?: boolean
    files: string[]
    acceptExts?: string[]
    hint?: string
  }>(),
  {
    acceptLabel: '选择文件',
    multiple: false,
    acceptExts: () => [],
    hint: '拖入文件到此处，或点击下方按钮选择（全程本地处理）',
  },
)

const emit = defineEmits<{
  (e: 'pick'): void
  (e: 'paths', paths: string[]): void
  (e: 'remove', index: number): void
  (e: 'move', index: number, dir: -1 | 1): void
  (e: 'clear'): void
}>()

const dragging = ref(false)
const dropError = ref('')
const rootEl = ref<HTMLElement | null>(null)

const acceptText = computed(() => {
  if (!props.acceptExts.length) return '任意文件'
  return props.acceptExts.map((e) => e.toUpperCase().replace('.', '')).join(' / ')
})

function matchExt(path: string) {
  if (!props.acceptExts.length) return true
  const lower = path.toLowerCase()
  return props.acceptExts.some((ext) => lower.endsWith(ext.toLowerCase()))
}

function ingestPaths(paths: string[]) {
  dropError.value = ''
  const filtered = paths.filter(Boolean).filter(matchExt)
  if (!filtered.length) {
    dropError.value = `请拖入 ${acceptText.value} 文件`
    return
  }
  emit('paths', filtered)
}

function onNativeDrop(_x: number, _y: number, paths: string[]) {
  if (!paths?.length) return
  // Only accept drops that land on this dropzone when using CSS drop target.
  ingestPaths(paths)
}

function onDragOver(e: DragEvent) {
  e.preventDefault()
  dragging.value = true
}

function onDragLeave(e: DragEvent) {
  if (e.currentTarget === rootEl.value) dragging.value = false
}

function onHtmlDrop(e: DragEvent) {
  e.preventDefault()
  dragging.value = false
  const list = e.dataTransfer?.files
  if (!list?.length) return
  const paths: string[] = []
  for (let i = 0; i < list.length; i++) {
    const f = list.item(i) as File & { path?: string }
    if (f?.path) paths.push(f.path)
  }
  if (paths.length) {
    ingestPaths(props.multiple ? paths : paths.slice(0, 1))
  } else {
    dropError.value = '未能读取文件路径，请改用「选择文件」或从资源管理器拖入'
  }
}

onMounted(() => {
  OnFileDrop(onNativeDrop, true)
})

onUnmounted(() => {
  OnFileDropOff()
})
</script>

<template>
  <div
    ref="rootEl"
    class="drop-target rounded-xl border border-dashed p-4 transition-colors"
    :class="cn(
      dragging
        ? 'border-primary bg-primary/5'
        : 'border-border bg-card hover:border-primary/40',
    )"
    style="--wails-drop-target: drop"
    @dragover="onDragOver"
    @dragleave="onDragLeave"
    @drop="onHtmlDrop"
  >
    <div class="flex flex-col items-center gap-2 py-4 text-center">
      <div class="flex size-11 items-center justify-center rounded-full bg-primary/10 text-primary">
        <FileUp class="size-5" />
      </div>
      <div class="text-sm font-medium">{{ dragging ? '松开以添加文件' : '拖拽文件到此处' }}</div>
      <p class="max-w-md text-xs text-muted-foreground">{{ hint }}</p>
      <Badge variant="secondary" class="mt-1">{{ acceptText }}</Badge>
    </div>

    <div class="mt-2 flex flex-wrap gap-2">
      <Button type="button" @click="emit('pick')">{{ acceptLabel }}</Button>
      <Button
        v-if="files.length"
        type="button"
        variant="outline"
        @click="emit('clear')"
      >
        <Trash2 class="size-4" />
        清空
      </Button>
    </div>

    <p v-if="dropError" class="mt-3 text-sm text-destructive">{{ dropError }}</p>

    <ul v-if="files.length" class="mt-4 space-y-2">
      <li
        v-for="(f, i) in files"
        :key="f + i"
        class="flex items-center gap-2 rounded-lg border bg-background px-3 py-2"
      >
        <span class="min-w-0 flex-1 truncate text-sm" :title="f">{{ f }}</span>
        <div v-if="multiple && files.length > 1" class="flex items-center gap-1">
          <Button
            type="button"
            variant="ghost"
            size="icon-sm"
            :disabled="i === 0"
            @click="emit('move', i, -1)"
          >
            <ArrowUp class="size-4" />
          </Button>
          <Button
            type="button"
            variant="ghost"
            size="icon-sm"
            :disabled="i === files.length - 1"
            @click="emit('move', i, 1)"
          >
            <ArrowDown class="size-4" />
          </Button>
        </div>
        <Button
          type="button"
          variant="ghost"
          size="icon-sm"
          class="text-destructive hover:text-destructive"
          @click="emit('remove', i)"
        >
          <X class="size-4" />
        </Button>
      </li>
    </ul>
  </div>
</template>
