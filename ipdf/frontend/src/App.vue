<script lang="ts" setup>
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="app-shell">
    <header class="topbar">
      <router-link class="brand" to="/">iPDF</router-link>
      <span class="tagline">本地 PDF 工具箱</span>
    </header>
    <main class="main">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100dvh;
  display: flex;
  flex-direction: column;
  text-align: left;
}
.topbar {
  display: flex;
  align-items: baseline;
  gap: 0.75rem;
  padding: 0.9rem 1.5rem;
  border-bottom: 1px solid var(--border);
  background: var(--surface);
  position: sticky;
  top: 0;
  z-index: 100;
}
.brand {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--accent);
  text-decoration: none;
  letter-spacing: 0.02em;
}
.tagline {
  color: var(--muted);
  font-size: 0.875rem;
}
.main {
  flex: 1;
  padding: 1.5rem;
  max-width: 1100px;
  width: 100%;
  margin: 0 auto;
  box-sizing: border-box;
}
</style>
