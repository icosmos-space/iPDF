<script lang="ts" setup>
import { ref, computed } from 'vue'
import FileDropzone from '@/components/FileDropzone.vue'
import PageRangeInput from '@/components/PageRangeInput.vue'
import ProgressPanel from '@/components/ProgressPanel.vue'
import ResultActions from '@/components/ResultActions.vue'
import { SplitPDF, SelectFile, SelectDirectory } from '../../../wailsjs/go/main/App'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const files = ref<string[]>([])
const mode = ref<'each' | 'range' | 'everyN'>('each')
const ranges = ref('')
const everyN = ref(2)
const loading = ref(false)
const error = ref('')
const message = ref('')
const outputs = ref<string[]>([])

const rangeArg = computed(() => {
  if (mode.value === 'everyN') return String(everyN.value)
  if (mode.value === 'range') return ranges.value
  return ''
})

async function pick() {
  const selected = await SelectFile('pdf')
  if (selected) files.value = [selected]
}

function onPaths(paths: string[]) {
  if (paths[0]) files.value = [paths[0]]
}

async function run() {
  error.value = ''
  message.value = ''
  outputs.value = []
  if (!files.value.length) {
    error.value = '请先选择一个 PDF'
    return
  }
  if (mode.value === 'range' && !ranges.value.trim()) {
    error.value = '请填写页码范围'
    return
  }
  const outDir = await SelectDirectory()
  if (!outDir) return
  loading.value = true
  try {
    const res = await SplitPDF(files.value[0], mode.value, rangeArg.value, outDir)
    message.value = res.message
    outputs.value = res.outputPaths || []
  } catch (e: any) {
    error.value = e?.message || String(e)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>PDF 分割</CardTitle>
      <CardDescription>按页拆分、按范围提取，或按每 N 页生成多个文件。</CardDescription>
    </CardHeader>
    <CardContent class="space-y-4">
      <FileDropzone
        accept-label="选择 PDF"
        :files="files"
        :accept-exts="['.pdf']"
        @pick="pick"
        @paths="onPaths"
        @remove="() => (files = [])"
        @clear="files = []"
      />
      <div class="flex flex-wrap gap-4">
        <div class="grid gap-2">
          <Label>模式</Label>
          <select
            v-model="mode"
            class="border-input bg-background h-9 rounded-md border px-3 text-sm shadow-xs"
          >
            <option value="each">每页一个文件</option>
            <option value="range">按页码范围提取</option>
            <option value="everyN">每 N 页一个文件</option>
          </select>
        </div>
        <PageRangeInput v-if="mode === 'range'" v-model="ranges" />
        <div v-if="mode === 'everyN'" class="grid gap-2">
          <Label>每 N 页</Label>
          <Input v-model.number="everyN" type="number" min="1" class="w-32" />
        </div>
      </div>
      <Button type="button" :disabled="loading" @click="run">分割并保存到目录</Button>
      <ProgressPanel :loading="loading" :error="error" :message="message" />
      <ResultActions :paths="outputs" />
    </CardContent>
  </Card>
</template>
