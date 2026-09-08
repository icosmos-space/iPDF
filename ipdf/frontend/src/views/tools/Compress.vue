<script lang="ts" setup>
import { ref } from 'vue'
import FileDropzone from '../../components/FileDropzone.vue'
import ProgressPanel from '../../components/ProgressPanel.vue'
import ResultActions from '../../components/ResultActions.vue'
import { CompressPDF, SelectFile, SelectSaveFile } from '../../../wailsjs/go/main/App'

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
  <section class="tool">
    <router-link class="back" to="/">← 返回工具列表</router-link>
    <h1>PDF 压缩</h1>
    <p class="desc">通过优化 PDF 内部结构减小体积。效果因文件而异。</p>
    <FileDropzone
      accept-label="选择 PDF"
      :files="files"
      @pick="pick"
      @remove="() => (files = [])"
      @clear="files = []"
    />
    <div class="options">
      <label class="field">
        <span>压缩等级</span>
        <select v-model="quality">
          <option value="low">较小（优先体积）</option>
          <option value="medium">均衡</option>
          <option value="high">较高质量</option>
        </select>
      </label>
    </div>
    <button class="run" type="button" :disabled="loading" @click="run">压缩并保存</button>
    <ProgressPanel :loading="loading" :error="error" :message="message" />
    <ResultActions :paths="outputs" />
  </section>
</template>

