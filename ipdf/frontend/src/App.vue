<script lang="ts" setup>
import { RouterLink, RouterView } from 'vue-router'

type NavItem = {
  title: string
  to?: string
  soon?: boolean
}

const ready: NavItem[] = [
  { title: 'PDF 合并', to: '/tools/merge' },
  { title: 'PDF 分割', to: '/tools/split' },
  { title: 'PDF 压缩', to: '/tools/compress' },
  { title: 'PDF 转图片', to: '/tools/pdf-to-image' },
  { title: '图片转 PDF', to: '/tools/image-to-pdf' },
]

const soon: NavItem[] = [
  { title: '旋转 / 重排', soon: true },
  { title: '加水印 / 页码', soon: true },
  { title: '加密 / 解密', soon: true },
  { title: 'Office 转换', soon: true },
]
</script>

<template>
  <div class="app-shell">
    <aside class="sidebar">
      <div class="brand-block">
        <div class="brand">iPDF</div>
        <div class="tagline">本地 PDF 工具箱</div>
      </div>

      <nav class="nav">
        <div class="nav-group">可用</div>
        <RouterLink
          v-for="item in ready"
          :key="item.title"
          class="nav-item"
          active-class="active"
          :to="item.to!"
        >
          {{ item.title }}
        </RouterLink>

        <div class="nav-group soon-group">即将推出</div>
        <div
          v-for="item in soon"
          :key="item.title"
          class="nav-item disabled"
        >
          {{ item.title }}
        </div>
      </nav>

      <div class="sidebar-foot">文件仅在本地处理</div>
    </aside>

    <main class="main">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.app-shell {
  display: flex;
  min-height: 100dvh;
  text-align: left;
  background: var(--bg);
}

.sidebar {
  width: 220px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  background: var(--surface);
  border-right: 1px solid var(--border);
  padding: 1.1rem 0.75rem;
}

.brand-block {
  padding: 0.35rem 0.65rem 1rem;
}

.brand {
  font-size: 1.35rem;
  font-weight: 700;
  color: var(--accent);
  letter-spacing: 0.02em;
  line-height: 1.2;
}

.tagline {
  margin-top: 0.2rem;
  color: var(--muted);
  font-size: 0.75rem;
}

.nav {
  display: flex;
  flex-direction: column;
  gap: 0.15rem;
  flex: 1;
  overflow-y: auto;
}

.nav-group {
  margin: 0.65rem 0.65rem 0.35rem;
  font-size: 0.72rem;
  font-weight: 600;
  color: var(--muted);
  letter-spacing: 0.04em;
}

.soon-group {
  margin-top: 1.25rem;
}

.nav-item {
  display: block;
  padding: 0.55rem 0.75rem;
  border-radius: 6px;
  color: var(--text);
  text-decoration: none;
  font-size: 0.92rem;
  font-weight: 500;
  transition: background 0.15s, color 0.15s;
}

.nav-item:hover {
  background: var(--bg);
  color: var(--accent);
}

.nav-item.active {
  background: #e8f1fb;
  color: var(--accent);
  font-weight: 600;
}

.nav-item.disabled {
  opacity: 0.45;
  pointer-events: none;
  font-weight: 400;
}

.sidebar-foot {
  margin-top: auto;
  padding: 0.85rem 0.65rem 0.25rem;
  font-size: 0.72rem;
  color: var(--muted);
  border-top: 1px solid var(--border);
}

.main {
  flex: 1;
  min-width: 0;
  padding: 1.5rem 1.75rem 2rem;
  overflow: auto;
}

@media (max-width: 720px) {
  .app-shell {
    flex-direction: column;
  }
  .sidebar {
    width: 100%;
    border-right: none;
    border-bottom: 1px solid var(--border);
    padding-bottom: 0.5rem;
  }
  .nav {
    flex-direction: row;
    flex-wrap: wrap;
    gap: 0.25rem;
  }
  .nav-group,
  .sidebar-foot {
    width: 100%;
  }
  .soon-group,
  .nav-item.disabled,
  .sidebar-foot {
    display: none;
  }
}
</style>
