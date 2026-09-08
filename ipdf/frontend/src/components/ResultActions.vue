<script lang="ts" setup>
import { OpenInExplorer } from '../../wailsjs/go/main/App'

const props = defineProps<{
  paths: string[]
}>()

async function openFirst() {
  if (!props.paths.length) return
  const p = props.paths[0]
  await OpenInExplorer(p)
}

async function openDir() {
  if (!props.paths.length) return
  const p = props.paths[0]
  const dir = p.replace(/[\\/][^\\/]+$/, '') || p
  await OpenInExplorer(dir)
}
</script>

<template>
  <div v-if="paths.length" class="results">
    <h3>输出</h3>
    <ul>
      <li v-for="p in paths" :key="p">{{ p }}</li>
    </ul>
    <div class="actions">
      <button type="button" class="btn" @click="openFirst">在资源管理器中显示</button>
      <button v-if="paths.length > 1" type="button" class="btn ghost" @click="openDir">打开输出目录</button>
    </div>
  </div>
</template>

<style scoped>
.results {
  margin-top: 1.25rem;
  padding: 1rem;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
}
h3 {
  margin: 0 0 0.5rem;
  font-size: 1rem;
}
ul {
  margin: 0 0 0.75rem;
  padding-left: 1.1rem;
  font-size: 0.85rem;
  color: var(--muted);
  word-break: break-all;
}
.actions {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}
.btn {
  border-radius: 6px;
  border: 1px solid transparent;
  padding: 0.4rem 0.85rem;
  background: var(--accent);
  color: #fff;
  font-weight: 600;
}
.btn.ghost {
  background: #fff;
  border-color: var(--border);
  color: var(--text);
}
</style>
