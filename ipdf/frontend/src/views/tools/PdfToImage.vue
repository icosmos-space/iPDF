<script lang="ts" setup>
import { ref } from 'vue'
import FileDropzone from '@/components/FileDropzone.vue'
import ProgressPanel from '@/components/ProgressPanel.vue'
import ResultActions from '@/components/ResultActions.vue'
import { PDFToImages, SelectFile, SelectDirectory } from '../../../wailsjs/go/main/App'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const files = ref<string[]>([])
const format = ref('jpg')
const dpi = ref(144)
const loading = ref(false)
const error = ref('')
const message = ref('')
const outputs = ref<string[]>([])

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
  const outDir = await SelectDirectory()
  if (!outDir) return
  loading.value = true
  try {
    const res = await PDFToImages(files.value[0], format.value, dpi.value, outDir)
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
      <CardTitle>PDF 转图片</CardTitle>
      <CardDescription>将每一页渲染为 JPG / PNG（go-fitz / MuPDF）。</CardDescription>
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
          <Label>格式</Label>
          <select
            v-model="format"
            class="border-input bg-background h-9 rounded-md border px-3 text-sm shadow-xs"
          >
            <option value="jpg">JPG</option>
            <option value="png">PNG</option>
          </select>
        </div>
        <div class="grid gap-2">
          <Label>DPI</Label>
          <Input v-model.number="dpi" type="number" min="72" max="300" step="12" class="w-32" />
        </div>
      </div>
      <Button type="button" :disabled="loading" @click="run">导出到目录</Button>
      <ProgressPanel :loading="loading" :error="error" :message="message" />
      <ResultActions :paths="outputs" />
    </CardContent>
  </Card>
</template>
