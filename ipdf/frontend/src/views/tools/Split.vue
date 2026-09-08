<script lang="ts" setup>
import { ref, computed } from 'vue'
import FileDropzone from '../../components/FileDropzone.vue'
import PageRangeInput from '../../components/PageRangeInput.vue'
import ProgressPanel from '../../components/ProgressPanel.vue'
import ResultActions from '../../components/ResultActions.vue'
import { SplitPDF, SelectFile, SelectDirectory } from '../../../wailsjs/go/main/App'

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
  <section class="tool">
    <router-link class="back" to="/">← 返回工具列表</router-link>
    <h1>PDF 分割</h1>
    <p class="desc">按页拆分、按范围提取，或按每 N 页生成多个文件。</p>
    <FileDropzone
      accept-label="选择 PDF"
      :files="files"
      @pick="pick"
      @remove="() => (files = [])"
      @clear="files = []"
    />
    <div class="options">
      <label class="field">
        <span>模式</span>
        <select v-model="mode">
          <option value="each">每页一个文件</option>
          <option value="range">按页码范围提取</option>
          <option value="everyN">每 N 页一个文件</option>
        </select>
      </label>
      <PageRangeInput v-if="mode === 'range'" v-model="ranges" />
      <label v-if="mode === 'everyN'" class="field">
        <span>每 N 页</span>
        <input v-model.number="everyN" type="number" min="1" />
      </label>
    </div>
    <button class="run" type="button" :disabled="loading" @click="run">分割并保存到目录</button>
    <ProgressPanel :loading="loading" :error="error" :message="message" />
    <ResultActions :paths="outputs" />
  </section>
</template>

