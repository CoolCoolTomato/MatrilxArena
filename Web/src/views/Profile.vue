<template>
  <el-container>
    <el-main>
      <div style="display: flex; flex-direction: column; align-items: center">
        <div style="width: 600px">
          <h2 style="color: var(--el-text-color-primary)">{{ t('Profile.UserInformation') }}</h2>
        </div>
        <div style="width: 600px">
          <h2 style="color: var(--el-text-color-primary)">{{ t('Profile.Setting') }}</h2>
          <div style="display: flex; margin: 20px">
            <el-text size="large" style="width: 100px;">{{ t('Profile.Language') }}</el-text>
            <LanguageSelect />
          </div>
          <div style="display: flex; margin: 20px">
            <el-text size="large" style="width: 100px;">{{ t('Profile.Theme') }}</el-text>
            <ThemeSelect />
          </div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>
<script>
import LanguageSelect from "@/components/LanguageSelect.vue"
import ThemeSelect from "@/components/ThemeSelect.vue"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: { LanguageSelect, ThemeSelect }
}
</script>
