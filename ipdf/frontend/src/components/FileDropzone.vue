<script lang="ts" setup>
defineProps<{
  acceptLabel?: string
  multiple?: boolean
  files: string[]
}>()

const emit = defineEmits<{
  (e: 'pick'): void
  (e: 'remove', index: number): void
  (e: 'move', index: number, dir: -1 | 1): void
  (e: 'clear'): void
}>()
</script>

<template>
  <div class="dropzone">
    <div class="actions">
      <button type="button" class="btn primary" @click="emit('pick')">
        {{ acceptLabel || '选择文件' }}
      </button>
      <button
        v-if="files.length"
        type="button"
        class="btn ghost"
        @click="emit('clear')"
      >
        清空
      </button>
    </div>
    <ul v-if="files.length" class="file-list">
      <li v-for="(f, i) in files" :key="f + i">
        <span class="name" :title="f">{{ f }}</span>
        <span v-if="multiple && files.length > 1" class="reorder">
          <button type="button" class="icon" :disabled="i === 0" @click="emit('move', i, -1)">↑</button>
          <button type="button" class="icon" :disabled="i === files.length - 1" @click="emit('move', i, 1)">↓</button>
        </span>
        <button type="button" class="icon danger" @click="emit('remove', i)">×</button>
      </li>
    </ul>
    <p v-else class="hint">尚未选择文件。点击上方按钮从本机选择（处理全程本地完成）。</p>
  </div>
</template>

<style scoped>
.dropzone {
  border: 1px dashed var(--border);
  border-radius: var(--radius);
  background: var(--surface);
  padding: 1rem;
}
.actions {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
}
.file-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}
.file-list li {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.45rem 0.6rem;
  background: var(--bg);
  border-radius: 6px;
  border: 1px solid var(--border);
}
.name {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 0.875rem;
}
.reorder {
  display: inline-flex;
  gap: 0.15rem;
}
.hint {
  margin: 0;
  color: var(--muted);
  font-size: 0.875rem;
}
.btn {
  border-radius: 6px;
  border: 1px solid transparent;
  padding: 0.45rem 0.9rem;
  font-weight: 600;
}
.btn.primary {
  background: var(--accent);
  color: #fff;
}
.btn.primary:hover {
  background: var(--accent-hover);
}
.btn.ghost {
  background: transparent;
  border-color: var(--border);
  color: var(--text);
}
.icon {
  border: none;
  background: transparent;
  color: var(--muted);
  font-size: 1rem;
  line-height: 1;
  padding: 0.2rem 0.35rem;
}
.icon:disabled {
  opacity: 0.35;
}
.icon.danger {
  color: var(--danger);
}
</style>
