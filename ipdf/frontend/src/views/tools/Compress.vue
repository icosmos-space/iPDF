<script lang="ts" setup>
import { ref } from 'vue'
import FileDropzone from '@/components/FileDropzone.vue'
import ProgressPanel from '@/components/ProgressPanel.vue'
import ResultActions from '@/components/ResultActions.vue'
import { CompressPDF, SelectFile, SelectSaveFile } from '../../../wailsjs/go/main/App'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'

const files = ref<string[]>([])
const quality = ref('medium')
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
  const out = await SelectSaveFile('compressed.pdf')
  if (!out) return
  loading.value = true
  try {
    const res = await CompressPDF(files.value[0], quality.value, out)
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
      <CardTitle>PDF 压缩</CardTitle>
      <CardDescription>通过优化 PDF 内部结构减小体积，效果因文件而异。</CardDescription>
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
      <div class="grid max-w-xs gap-2">
        <Label>压缩等级</Label>
        <select
          v-model="quality"
          class="border-input bg-background h-9 rounded-md border px-3 text-sm shadow-xs"
        >
          <option value="low">较小（优先体积）</option>
          <option value="medium">均衡</option>
          <option value="high">较高质量</option>
        </select>
      </div>
      <Button type="button" :disabled="loading" @click="run">压缩并保存</Button>
      <ProgressPanel :loading="loading" :error="error" :message="message" />
      <ResultActions :paths="outputs" />
    </CardContent>
  </Card>
</template>
