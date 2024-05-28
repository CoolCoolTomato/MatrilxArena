<template>
  <el-button
    @click="toggleTheme"
    :icon="icon"
    circle
    >
    </el-button>
</template>
<script>
import {Moon, Sunny} from '@element-plus/icons-vue'

export default {
  name: 'ToggleThemeButton',
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

</style>