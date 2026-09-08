<script lang="ts" setup>
type Tool = {
  title: string
  desc: string
  to?: string
  soon?: boolean
}

const ready: Tool[] = [
  { title: 'PDF 合并', desc: '将多个 PDF 按顺序合并为一个文件', to: '/tools/merge' },
  { title: 'PDF 分割', desc: '按页、按范围或按每 N 页拆分', to: '/tools/split' },
  { title: 'PDF 压缩', desc: '优化体积，便于分享与存储', to: '/tools/compress' },
  { title: 'PDF 转图片', desc: '提取 PDF 内嵌图片为 JPG / PNG', to: '/tools/pdf-to-image' },
  { title: '图片转 PDF', desc: '将 JPG / PNG / WEBP 拼成 PDF', to: '/tools/image-to-pdf' },
]

const soon: Tool[] = [
  { title: '旋转 / 重排', desc: '调整页面方向与顺序', soon: true },
  { title: '加水印 / 页码', desc: '批注与页码工具', soon: true },
  { title: '加密 / 解密', desc: '密码保护与解除', soon: true },
  { title: 'Office 转换', desc: 'Word / Excel / PPT ↔ PDF', soon: true },
]
</script>

<template>
  <section>
    <h1>工具</h1>
    <p class="lead">全部在本地处理，文件不会上传到云端。</p>

    <h2>可用</h2>
    <div class="grid">
      <router-link
        v-for="t in ready"
        :key="t.title"
        class="card"
        :to="t.to!"
      >
        <strong>{{ t.title }}</strong>
        <span>{{ t.desc }}</span>
      </router-link>
    </div>

    <h2>即将推出</h2>
    <div class="grid">
      <div v-for="t in soon" :key="t.title" class="card disabled">
        <strong>{{ t.title }}</strong>
        <span>{{ t.desc }}</span>
        <em>即将推出</em>
      </div>
    </div>
  </section>
</template>

<style scoped>
h1 {
  margin: 0 0 0.35rem;
  font-size: 1.75rem;
}
.lead {
  margin: 0 0 1.5rem;
  color: var(--muted);
}
h2 {
  margin: 1.5rem 0 0.75rem;
  font-size: 1rem;
  color: var(--muted);
  font-weight: 600;
  text-transform: none;
}
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 0.85rem;
}
.card {
  display: flex;
  flex-direction: column;
  gap: 0.35rem;
  padding: 1rem;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  box-shadow: var(--shadow);
  text-decoration: none;
  color: inherit;
  transition: border-color 0.15s, transform 0.15s;
}
.card:hover {
  border-color: var(--accent);
  transform: translateY(-1px);
}
.card strong {
  font-size: 1.05rem;
}
.card span {
  color: var(--muted);
  font-size: 0.875rem;
}
.card.disabled {
  opacity: 0.55;
  pointer-events: none;
  position: relative;
}
.card em {
  font-style: normal;
  font-size: 0.75rem;
  color: var(--muted);
  margin-top: 0.25rem;
}
</style>
