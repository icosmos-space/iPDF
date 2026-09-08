<script lang="ts" setup>
import { ref } from 'vue'
import FileDropzone from '../../components/FileDropzone.vue'
import ProgressPanel from '../../components/ProgressPanel.vue'
import ResultActions from '../../components/ResultActions.vue'
import { MergePDFs, SelectFiles, SelectSaveFile } from '../../../wailsjs/go/main/App'

const files = ref<string[]>([])
const loading = ref(false)
const error = ref('')
const message = ref('')
const outputs = ref<string[]>([])

async function pick() {
  const selected = await SelectFiles('pdf')
  if (selected?.length) {
    files.value = [...files.value, ...selected]
  }
}

function remove(i: number) {
  files.value.splice(i, 1)
}

function move(i: number, dir: -1 | 1) {
  const j = i + dir
  if (j < 0 || j >= files.value.length) return
  const tmp = files.value[i]
  files.value[i] = files.value[j]
  files.value[j] = tmp
}

async function run() {
  error.value = ''
  message.value = ''
  outputs.value = []
  if (files.value.length < 2) {
    error.value = '请至少选择 2 个 PDF'
    return
  }
  const out = await SelectSaveFile('merged.pdf')
  if (!out) return
  loading.value = true
  try {
    const res = await MergePDFs(files.value, out)
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
  <section class="tool">
    <h1>PDF 合并</h1>
    <p class="desc">按列表顺序合并多个 PDF。可用上下箭头调整顺序。</p>
    <FileDropzone
      accept-label="添加 PDF"
      :multiple="true"
      :files="files"
      @pick="pick"
      @remove="remove"
      @move="move"
      @clear="files = []"
    />
    <button class="run" type="button" :disabled="loading" @click="run">合并并保存</button>
    <ProgressPanel :loading="loading" :error="error" :message="message" />
    <ResultActions :paths="outputs" />
  </section>
</template>

