import { createRouter, createWebHistory } from 'vue-router'
import Base from "@/components/Base.vue";
import Login from "@/views/Login.vue";
import Index from "@/views/Index.vue";
import Challenge from "@/views/Challenge.vue";

import BaseAdmin from "@/components/BaseAdmin.vue";
import AdminIndex from "@/views/admin/AdminIndex.vue"
import AdminNode from "@/views/admin/AdminNode.vue";
import AdminNodeDetail from "@/views/admin/AdminNodeDetail.vue"
import AdminImage from "@/views/admin/AdminImage.vue";
import AdminChallenge from "@/views/admin/AdminChallenge.vue";
import AdminUser from "@/views/admin/AdminUser.vue";
import AdminAttachment from "@/views/admin/AdminAttachment.vue";
import AdminChallengeClass from "@/views/admin/AdminChallengeClass.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/login", component: Login },
    {path: "/", component: Base, children: [
      {path: "", component: Index},
      {path: "challenge", component: Challenge},
    ]},
    {path: "/admin", component: BaseAdmin, children: [
      {path: "", component: AdminIndex},
      {path: "node", component: AdminNode},
      {path: "node/:id(\\d+)", component: AdminNodeDetail},
      {path: "image", component: AdminImage},
      {path: "challengeClass", component: AdminChallengeClass},
      {path: "challenge", component: AdminChallenge},
      {path: "attachment", component: AdminAttachment},
      {path: "user", component: AdminUser},
    ]},
  ]
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");
  if (!token && to.path !== "/login") {
    next("/login");
  } else {
    next();
  }
});


export default router
