<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="common-layout" id="main">
    <el-container>
      <el-header>
        <el-menu
          mode="horizontal"
          :ellipsis="false"
          router
          >
          <el-menu-item class="el-menu-item-no-style">
            <el-image src="/image/svg/matrilx.svg" style="width: 56px;" />
            <h2>MatrilxArena</h2>
          </el-menu-item>
          <div class="flex-grow" />
          <el-menu-item index="/">
            <el-text class="mx-1" size="large">Index</el-text>
          </el-menu-item>
          <el-menu-item index="/challenge">
            <el-text class="mx-1" size="large">Challenge</el-text>
          </el-menu-item>
          <el-menu-item class="el-menu-item-no-style">
            <el-button
              @click="toggleTheme"
              :icon="icon"
              circle
              >
            </el-button>
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <RouterView />
      </el-main>
      <el-footer>
        <el-text class="mx-1" size="small">Powered By CoolCoolTomato</el-text>
        <el-text class="mx-1" size="small">|</el-text>
        <el-text class="mx-1" size="small">© 2024 Matrilx. All rights reserved.</el-text>
      </el-footer>
    </el-container>
  </div>
</template>
<script>
import {Moon, Sunny} from '@element-plus/icons-vue'

export default {
  name: 'APP',
  data() {
    return {
      theme: 'light',
      icon: Sunny,
    };
  },
  methods: {
    toggleTheme(event) {
      document.documentElement.style.setProperty('--x', event.clientX + 'px')
      document.documentElement.style.setProperty('--y', event.clientY + 'px')
      this.theme = this.theme === 'light' ? 'dark' : 'light';
      this.icon = this.theme === 'light' ? Sunny : Moon;
      localStorage.setItem('theme', this.theme);
      if (document.startViewTransition) {
        document.startViewTransition(() => {
          setTimeout(() => {
            document.documentElement.className = this.theme
          }, 100)
        });
      } else {
        document.documentElement.className = this.theme
      }
    },
  },
  mounted() {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
      this.theme = savedTheme;
      this.icon = this.theme === 'light' ? Sunny : Moon;
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
::v-deep .el-menu-item-no-style > .el-button > [class^="el-icon"] {
  margin-right: 0;
}
.el-container{
  width: 100%;
  height: 100%;
}
.flex-grow {
  flex-grow: 1;
}
.el-footer{
  display: flex;
  justify-content: center;
  align-items: center;
}
.el-footer .el-text{
  margin: 0 5px;
}
</style>