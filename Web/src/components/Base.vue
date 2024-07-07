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
            <el-image src="/image/svg/matrilx.svg" style="width: 40px; margin-bottom: 16px; margin-right: 5px;" />
            <h2 style="color: var(--el-text-color-primary);">{{ $t('Base.MatrilxArena') }}</h2>
          </el-menu-item>
          <div class="flex-grow" />
          <el-menu-item index="/challenge">
            <el-text size="large" tag="b">{{ $t('Base.Challenge') }}</el-text>
          </el-menu-item>
          <el-menu-item index="/group">
            <el-text size="large" tag="b">{{ $t('Base.Group') }}</el-text>
          </el-menu-item>
          <el-menu-item index="/awd">
            <el-text size="large" tag="b">{{ $t('Base.AWD') }}</el-text>
          </el-menu-item>
          <el-menu-item index="/profile">
            <el-text size="large" tag="b">{{ $t('Base.Profile') }}</el-text>
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <RouterView />
      </el-main>
      <el-footer>
        <Copyright />
      </el-footer>
    </el-container>
  </div>
</template>
<script>
import { RouterView } from 'vue-router'
import Copyright from "@/components/Copyright.vue"
import {useI18n} from "vue-i18n"

export default {
  components: {
    RouterView,
    Copyright
  },
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      theme: "light",
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
#main{
  width: 100%;
  height: 100%;
  position: absolute;
}
.el-menu-item-no-style{
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
.el-container{
  width: 100%;
  height: 100%;
}
.el-footer{
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
