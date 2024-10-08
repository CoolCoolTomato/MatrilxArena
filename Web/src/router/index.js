import { createRouter, createWebHistory } from 'vue-router'
import Base from "@/components/Base.vue"
import Login from "@/views/Login.vue"
import Index from "@/views/Index.vue"
import Challenge from "@/views/Challenge.vue"
import CTF from "@/views/CTF.vue"
import CTFDetail from "@/views/CTFDetail.vue"
import Group from "@/views/Group.vue"
import GroupDetail from "@/views/GroupDetail.vue"
import BaseAdmin from "@/components/BaseAdmin.vue"
import AdminIndex from "@/views/admin/AdminIndex.vue"
import AdminNode from "@/views/admin/AdminNode.vue"
import AdminNodeDetail from "@/views/admin/AdminNodeDetail.vue"
import AdminImage from "@/views/admin/AdminImage.vue"
import AdminChallenge from "@/views/admin/AdminChallenge.vue"
import AdminGroup from "@/views/admin/AdminGroup.vue"
import AdminGroupDetail from "@/views/admin/AdminGroupDetail.vue"
import AdminUser from "@/views/admin/AdminUser.vue"
import AdminAttachment from "@/views/admin/AdminAttachment.vue"
import AdminCategory from "@/views/admin/AdminCategory.vue"
import AdminCTF from "@/views/admin/AdminCTF.vue"
import AdminCTFDetail from "@/views/admin/AdminCTFDetail.vue"
import authApi from "@/api/auth.js"
import Profile from "@/views/Profile.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/login", component: Login },
    { path: "/", component: Base, children: [
      { path: "", component: Index },
      { path: "challenge/:category?", component: Challenge },
      { path: "group", component: Group },
      { path: "group/:id(\\d+)/:category?", component: GroupDetail },
      { path: "ctf", component: CTF },
      { path: "ctf/:id(\\d+)/:category?", component: CTFDetail },
      { path: "profile", component: Profile },
    ]},
    { path: "/admin", component: BaseAdmin, children: [
      { path: "", component: AdminIndex },
      { path: "node", component: AdminNode },
      { path: "node/:id(\\d+)", component: AdminNodeDetail },
      { path: "image", component: AdminImage},
      { path: "category", component: AdminCategory },
      { path: "challenge", component: AdminChallenge },
      { path: "ctf", component: AdminCTF },
      { path: "ctf/:id(\\d+)", component: AdminCTFDetail },
      { path: "group", component: AdminGroup },
      { path: "group/:id(\\d+)", component: AdminGroupDetail },
      { path: "attachment", component: AdminAttachment },
      { path: "user", component: AdminUser },
    ]},
  ]
})

router.beforeEach((to, from, next) => {
  if (to.path === "/login" || to.path === "/") {
    next()
    return
  }
  authApi.GetUserByAuth().then(res => {
    if (res.code === 0) {
      const role = res.data.Role
      if (to.path.startsWith('/admin')) {
        if (role !== 1) {
          next("/login")
        } else {
          next()
        }
      } else {
        next()
      }
    } else {
      next("/login")
    }
  }).catch(error => {
    console.log(error)
    next("/login")
  })
})

export default router
