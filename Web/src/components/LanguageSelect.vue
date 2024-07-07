<template>
  <el-select v-model="currentLang" @change="changeLang">
    <el-option value="en" label="English"/>
    <el-option value="zh" label="中文"/>
  </el-select>
</template>

<script>

export default {
  data() {
    return {
      currentLang: 'en'
    }
  },
  methods: {
    changeLang(value) {
      this.currentLang = value
      this.$i18n.locale = value
      localStorage.setItem('lang', value)
    },
  },
  created() {
    const localLanguage = localStorage.getItem('lang') || 'en'
    this.currentLang = localLanguage
    this.$i18n.locale = localLanguage
  }
}
</script>
