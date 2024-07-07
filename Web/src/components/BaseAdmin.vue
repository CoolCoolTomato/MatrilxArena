<template>
  <div class="common-layout" id="main">
    <el-container>
      <el-header>
        <el-menu
          mode="horizontal"
          :ellipsis="false"
          router
        >
          <el-menu-item index="/">
            <el-image src="/image/svg/matrilx.svg" style="width: 40px; margin-bottom: 16px; margin-right: 5px;"/>
            <h2 style="color: var(--el-text-color-primary);">{{ $t('BaseAdmin.MatrilxArena') }}</h2>
          </el-menu-item>
          <div class="flex-grow"/>
        </el-menu>
      </el-header>
      <el-main>
        <el-container>
          <el-aside :width="isMenuOpen ? '190px' : '64px'">
            <el-menu
              :collapse="!isMenuOpen"
              style="border: none;"
              router
            >
              <el-menu-item @click="toggleSidebar">
                <el-icon>
                  <Menu/>
                </el-icon>
                <template #title>{{ $t('BaseAdmin.Menu') }}</template>
              </el-menu-item>
              <el-menu-item index="/admin">
                <el-icon>
                  <Home/>
                </el-icon>
                <template #title>{{ $t('BaseAdmin.Index') }}</template>
              </el-menu-item>
              <el-sub-menu index="docker">
                <template #title>
                  <el-icon>
                    <Docker/>
                  </el-icon>
                  <el-text style="color: var(--el-text-color-primary);">{{ $t('BaseAdmin.Docker') }}</el-text>
                </template>
                <el-menu-item index="/admin/node">
                  <template #title>{{ $t('BaseAdmin.Node') }}</template>
                </el-menu-item>
                <el-menu-item index="/admin/image">
                  <template #title>{{ $t('BaseAdmin.Image') }}</template>
                </el-menu-item>
              </el-sub-menu>
              <el-sub-menu index="challenge">
                <template #title>
                  <el-icon>
                    <Challenge/>
                  </el-icon>
                  <el-text style="color: var(--el-text-color-primary);">{{ $t('BaseAdmin.Challenge') }}</el-text>
                </template>
                <el-menu-item index="/admin/category">
                  <template #title>{{ $t('BaseAdmin.Category') }}</template>
                </el-menu-item>
                <el-menu-item index="/admin/challenge">
                  <template #title>{{ $t('BaseAdmin.Challenge') }}</template>
                </el-menu-item>
              </el-sub-menu>
              <el-menu-item index="/admin/attachment">
                <el-icon>
                  <Attachment/>
                </el-icon>
                <template #title>{{ $t('BaseAdmin.Attachment') }}</template>
              </el-menu-item>
              <el-menu-item index="/admin/user">
                <el-icon>
                  <User/>
                </el-icon>
                <template #title>{{ $t('BaseAdmin.User') }}</template>
              </el-menu-item>
            </el-menu>
          </el-aside>
          <el-main>
            <RouterView/>
          </el-main>
        </el-container>
      </el-main>
      <el-footer>
        <Copyright/>
      </el-footer>
     </el-container>
  </div>
</template>
<script>
import { RouterView } from 'vue-router'
import Copyright from "@/components/Copyright.vue"
import Menu from "@/components/icon/Menu.vue"
import Home from "@/components/icon/Home.vue"
import Docker from "@/components/icon/Docker.vue"
import Challenge from "@/components/icon/Challenge.vue"
import Attachment from "@/components/icon/Attachment.vue"
import User from "@/components/icon/User.vue"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: {
    RouterView,
    Copyright,
    Menu,
    Home,
    Docker,
    Challenge,
    Attachment,
    User
  },
  data() {
    return {
      theme: "light",
      isMenuOpen: false,
    }
  },
  methods: {
    toggleSidebar() {
      this.isMenuOpen = !this.isMenuOpen
    }
  },
  mounted() {
    const savedTheme = localStorage.getItem('theme')
    if (savedTheme) {
      this.theme = savedTheme
    }
    document.documentElement.className = this.theme
  },
}
</script>
<style scoped>
#main {
  width: 100%;
  height: 100%;
  position: absolute;
}
.el-menu-item-no-style {
  background-color: transparent !important;
  color: inherit !important;
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  cursor: default;
}
.flex-grow {
  flex-grow: 1;
}
::v-deep(.el-menu-item-no-style .el-button [class^="el-icon"]) {
  margin-right: 0;
}
.el-container {
  width: 100%;
  height: 100%;
}
.el-footer {
  display: flex;
  justify-content: center;
  align-items: center;
}
.el-aside {
  overflow-x: hidden;
  transition: width var(--el-transition-duration);
  border-right: 1px solid var(--el-border-color);
}
</style>
