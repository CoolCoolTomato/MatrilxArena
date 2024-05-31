import { createRouter, createWebHistory } from 'vue-router'
import Base from "@/components/Base.vue";
import Index from "@/views/Index.vue";
import Challenge from "@/views/Challenge.vue";

import BaseAdmin from "@/components/BaseAdmin.vue";
import AdminIndex from "@/views/admin/AdminIndex.vue"
import AdminNode from "@/views/admin/AdminNode.vue";
import AdminNodeDetail from "@/views/admin/AdminNodeDetail.vue"
import AdminImage from "@/views/admin/AdminImage.vue";
import AdminChallenge from "@/views/admin/AdminChallenge.vue";
import AdminUser from "@/views/admin/AdminUser.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: "/", component: Base, children: [
      {path: "", component: Index},
      {path: "challenge", component: Challenge},
    ]},
    {path: "/admin", component: BaseAdmin, children: [
      {path: "", component: AdminIndex},
      {path: "node", component: AdminNode},
      {path: "node/:id(\\d+)", component: AdminNodeDetail},
      {path: "image", component: AdminImage},
      {path: "challenge", component: AdminChallenge},
      {path: "user", component: AdminUser},
    ]},
  ]
})

export default router