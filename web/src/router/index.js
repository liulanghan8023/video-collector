import { createRouter, createWebHashHistory } from 'vue-router'
import Product from '../pages/product/index.vue'
import Video from '../pages/video/index.vue'

const routes = [
  {
    path: '/products',
    name: 'Product',
    component: Product
  },
  {
    path: '/videos',
    name: 'Video',
    component: Video
  },
  {
    path: '/',
    redirect: '/products'
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
