<script setup>
import {RouterLink, RouterView} from 'vue-router'
import ToggleThemeButton from "@/components/ToggleThemeButton.vue";
import Copyright from "@/components/Copyright.vue";
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
            <el-image src="/image/svg/matrilx.svg"
                  style="width: 40px; margin-bottom: 16px; margin-right: 5px;"/>
            <h2 style="color: var(--el-text-color-primary);">MatrilxArena</h2>
          </el-menu-item>
          <div class="flex-grow"/>
          <el-menu-item class="el-menu-item-no-style">
            <ToggleThemeButton/>
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-container>
          <el-aside :width="isMenuOpen ? '150px' : '64px'">
            <el-menu
              :collapse="!isMenuOpen"
              style="border: none;"
              router
            >
              <el-menu-item @click="toggleSidebar">
                <el-icon>
                  <Menu/>
                </el-icon>
                <template #title>Menu</template>
              </el-menu-item>
              <el-menu-item index="/admin">
                <el-icon>
                  <HomeFilled/>
                </el-icon>
                <template #title>Index</template>
              </el-menu-item>
              <el-menu-item index="/admin/node">
                <el-icon>
                  <Briefcase/>
                </el-icon>
                <template #title>Node</template>
              </el-menu-item>
              <el-menu-item index="/admin/image">
                <el-icon>
                  <Shop/>
                </el-icon>
                <template #title>Image</template>
              </el-menu-item>
              <el-menu-item index="/admin/challenge">
                <el-icon>
                  <List/>
                </el-icon>
                <template #title>Challenge</template>
              </el-menu-item>
              <el-menu-item index="/admin/attachment">
                <el-icon>
                  <UploadFilled/>
                </el-icon>
                <template #title>Attachment</template>
              </el-menu-item>
              <el-menu-item index="/admin/user">
                <el-icon>
                  <UserFilled/>
                </el-icon>
                <template #title>User</template>
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
import {Menu, Briefcase, Shop, HomeFilled, List, UploadFilled, UserFilled} from '@element-plus/icons-vue'

export default {
  components: {Menu, Briefcase, Shop, HomeFilled, List, UploadFilled, UserFilled},
  data() {
    return {
      isMenuOpen: false,
    };
  },
  methods: {
    toggleSidebar() {
      this.isMenuOpen = !this.isMenuOpen;
    }
  }
};
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
::v-deep .el-menu-item-no-style > .el-button > [class^="el-icon"] {
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
