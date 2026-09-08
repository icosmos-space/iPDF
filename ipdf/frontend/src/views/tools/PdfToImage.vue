<script lang="ts" setup>
import { ref } from 'vue'
import FileDropzone from '../../components/FileDropzone.vue'
import ProgressPanel from '../../components/ProgressPanel.vue'
import ResultActions from '../../components/ResultActions.vue'
import { PDFToImages, SelectFile, SelectDirectory } from '../../../wailsjs/go/main/App'

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
  <section class="tool">
    <router-link class="back" to="/">← 返回工具列表</router-link>
    <h1>PDF 转图片</h1>
    <p class="desc">
      使用 pdfcpu 提取 PDF 中的嵌入图片，并导出为 JPG / PNG。
      （无法将纯文字页栅格化；扫描件或图片合成的 PDF 效果最好。）
    </p>
    <FileDropzone
      accept-label="选择 PDF"
      :files="files"
      @pick="pick"
      @remove="() => (files = [])"
      @clear="files = []"
    />
    <div class="options">
      <label class="field">
        <span>格式</span>
        <select v-model="format">
          <option value="jpg">JPG</option>
          <option value="png">PNG</option>
        </select>
      </label>
      <label class="field">
        <span>DPI（预留）</span>
        <input v-model.number="dpi" type="number" min="72" max="300" step="12" />
      </label>
    </div>
    <button class="run" type="button" :disabled="loading" @click="run">导出到目录</button>
    <ProgressPanel :loading="loading" :error="error" :message="message" />
    <ResultActions :paths="outputs" />
  </section>
</template>
