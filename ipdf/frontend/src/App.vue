<script lang="ts" setup>
import { RouterLink, RouterView, useRoute } from 'vue-router'
import { FileStack, Layers2, Minimize2, ImageDown, Images } from '@lucide/vue'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import { ScrollArea } from '@/components/ui/scroll-area'
import { cn } from '@/lib/utils'

type NavItem = {
  title: string
  to?: string
  soon?: boolean
  icon?: any
}

const route = useRoute()

const ready: NavItem[] = [
  { title: 'PDF 合并', to: '/tools/merge', icon: Layers2 },
  { title: 'PDF 分割', to: '/tools/split', icon: FileStack },
  { title: 'PDF 压缩', to: '/tools/compress', icon: Minimize2 },
  { title: 'PDF 转图片', to: '/tools/pdf-to-image', icon: ImageDown },
  { title: '图片转 PDF', to: '/tools/image-to-pdf', icon: Images },
]

const soon: NavItem[] = [
  { title: '旋转 / 重排', soon: true },
  { title: '加水印 / 页码', soon: true },
  { title: '加密 / 解密', soon: true },
  { title: 'Office 转换', soon: true },
]
</script>

<template>
  <div class="flex min-h-dvh bg-background text-left">
    <aside class="flex w-60 shrink-0 flex-col border-r bg-sidebar text-sidebar-foreground">
      <div class="px-4 py-5">
        <div class="text-xl font-bold tracking-tight text-sidebar-primary">iPDF</div>
        <p class="mt-1 text-xs text-muted-foreground">本地 PDF 工具箱</p>
      </div>
      <Separator />
      <ScrollArea class="flex-1 px-2 py-3">
        <div class="mb-2 px-2 text-[11px] font-semibold uppercase tracking-wide text-muted-foreground">
          可用
        </div>
        <nav class="space-y-1">
          <RouterLink
            v-for="item in ready"
            :key="item.title"
            :to="item.to!"
            :class="cn(
              'flex items-center gap-2 rounded-md px-2.5 py-2 text-sm font-medium transition-colors',
              route.path === item.to
                ? 'bg-sidebar-accent text-sidebar-accent-foreground'
                : 'text-sidebar-foreground/80 hover:bg-sidebar-accent/70 hover:text-sidebar-accent-foreground',
            )"
          >
            <component :is="item.icon" class="size-4 shrink-0 opacity-80" />
            <span>{{ item.title }}</span>
          </RouterLink>
        </nav>

        <div class="mb-2 mt-6 px-2 text-[11px] font-semibold uppercase tracking-wide text-muted-foreground">
          即将推出
        </div>
        <div class="space-y-1">
          <div
            v-for="item in soon"
            :key="item.title"
            class="flex items-center justify-between rounded-md px-2.5 py-2 text-sm text-muted-foreground/70"
          >
            <span>{{ item.title }}</span>
            <Badge variant="outline" class="text-[10px]">Soon</Badge>
          </div>
        </div>
      </ScrollArea>
      <div class="border-t px-4 py-3 text-[11px] text-muted-foreground">
        文件仅在本地处理，不上传云端
      </div>
    </aside>

    <main class="min-w-0 flex-1 overflow-auto p-6 md:p-8">
      <div class="mx-auto max-w-3xl">
        <RouterView />
      </div>
    </main>
  </div>
</template>
