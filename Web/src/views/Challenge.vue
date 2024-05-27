<template>
  <el-container>
    <el-aside :width="isMenuOpen ? '150px' : '64px'">
      <el-menu :collapse="!isMenuOpen" style="border: none;">
        <el-menu-item index="1" @click="toggleSidebar">
          <el-icon>
            <Menu />
          </el-icon>
          <span slot="title">Menu</span>
        </el-menu-item>
        <el-menu-item index="2">
          <el-icon>
            <Aim />
          </el-icon>
          <span slot="title">Web</span>
        </el-menu-item>
        <el-menu-item index="3">
          <el-icon>
            <Aim />
          </el-icon>
          <span slot="title">Misc</span>
        </el-menu-item>
        <el-menu-item index="4">
          <el-icon>
            <Aim />
          </el-icon>
          <span slot="title">Pwn</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-row>
          <el-col v-for="i in 10" :span=12>
            <div class="challenge">
              <h2>Web{{i}}</h2>
              <el-text>这是一个Web{{i}}的题目，你能找到它的flag吗？</el-text>
            </div>
          </el-col>
        </el-row>
        <el-pagination
          layout="prev, pager, next"
          :pager-count="11"
          :total="200"
          />
      </el-scrollbar>
    </el-main>
  </el-container>
</template>

<script>
import {Aim, Menu} from '@element-plus/icons-vue'

export default {
  components: {Aim, Menu},
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
.el-aside {
  overflow-x: hidden;
  transition: width var(--el-transition-duration);
  border-right: 1px solid var(--el-border-color);
}

.challenge {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
}
.challenge h2{
  margin: 20px;
}
.challenge .el-text{
  margin: 20px;
}
.el-pagination{
  justify-content: center;
}
</style>
