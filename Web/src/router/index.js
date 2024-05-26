import { createRouter, createWebHistory } from 'vue-router'
import Index from "@/views/Index.vue";
import Challenge from "@/views/Challenge.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: "/", component: Index},
    {path: "/challenge", component: Challenge},
  ]
})

export default router
