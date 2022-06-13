import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Map',
    component: () => import('@/views/Map.vue')
  },
  {
    path: '/data',
    name: 'Data',
    component: () => import('@/views/Data.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
