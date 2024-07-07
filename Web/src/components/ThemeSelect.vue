<template>
  <el-select v-model="theme" @change="toggleTheme">
    <el-option value="light" :label="$t('ThemeSelect.Light')"/>
    <el-option value="dark" :label="$t('ThemeSelect.Dark')"/>
  </el-select>
</template>
<script>
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      theme: 'light',
    }
  },
  methods: {
    toggleTheme(value) {
      document.documentElement.style.setProperty('--x',  '100%')
      document.documentElement.style.setProperty('--y',  '0px')
      this.theme = value
      localStorage.setItem('theme', this.theme)
      if (document.startViewTransition) {
      document.startViewTransition(() => {
        setTimeout(() => {
          document.documentElement.className = this.theme
        }, 100)
      })
      } else {
        document.documentElement.className = this.theme
      }
    },
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
