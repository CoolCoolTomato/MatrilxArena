<template>
  <el-select v-model="lang" @change="changeLang">
    <el-option value="en" label="English"/>
    <el-option value="zh" label="中文"/>
  </el-select>
</template>

<script>

export default {
  data() {
    return {
      lang: 'en'
    }
  },
  methods: {
    changeLang(value) {
      this.lang = value
      this.$i18n.locale = value
      localStorage.setItem('lang', value)
    },
  },
  created() {
    const savedLanguage = localStorage.getItem('lang') || 'en'
    this.lang = savedLanguage
    this.$i18n.locale = savedLanguage
  }
}
</script>
