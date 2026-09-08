<script lang="ts" setup>
defineProps<{
  loading?: boolean
  error?: string
  message?: string
}>()
</script>

<template>
  <div class="panel">
    <div v-if="loading" class="loading">正在处理，请稍候…</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else-if="message" class="ok">{{ message }}</div>
  </div>
</template>

<style scoped>
.panel:empty {
  display: none;
}
.loading, .error, .ok {
  margin-top: 1rem;
  padding: 0.75rem 1rem;
  border-radius: var(--radius);
  font-size: 0.9rem;
}
.loading {
  background: #eef4fb;
  color: var(--accent);
}
.error {
  background: #fdecea;
  color: var(--danger);
}
.ok {
  background: #e8f5ee;
  color: var(--ok);
}
</style>
