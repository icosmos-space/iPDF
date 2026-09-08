import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './views/Home.vue'
import Merge from './views/tools/Merge.vue'
import Split from './views/tools/Split.vue'
import Compress from './views/tools/Compress.vue'
import PdfToImage from './views/tools/PdfToImage.vue'
import ImageToPdf from './views/tools/ImageToPdf.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', name: 'home', component: Home },
    { path: '/tools/merge', name: 'merge', component: Merge },
    { path: '/tools/split', name: 'split', component: Split },
    { path: '/tools/compress', name: 'compress', component: Compress },
    { path: '/tools/pdf-to-image', name: 'pdf-to-image', component: PdfToImage },
    { path: '/tools/image-to-pdf', name: 'image-to-pdf', component: ImageToPdf },
  ],
})

export default router
