<template>
  <el-container>
    <el-main>
      <div style="display: flex; flex-direction: column; align-items: center">
        <div style="width: 600px">
          <h2 style="color: var(--el-text-color-primary)">{{ t('Profile.UserInformation') }}</h2>
          <div style="display: flex; margin: 20px">
            <el-text size="large" style="width: 100px; text-align: right; margin-right: 20px">{{ t('Profile.Username') }}</el-text>
            <el-input v-model="user.Username"/>
          </div>
          <div style="display: flex; margin: 20px">
            <el-text size="large" style="width: 100px; text-align: right; margin-right: 20px">{{ t('Profile.Email') }}</el-text>
            <el-input v-model="user.Email"/>
          </div>
          <div style="display: flex; margin: 20px">
            <div style="width: 100px; margin: 15px 20px 0 0; text-align: right;">
              <el-text size="large" style="width: 100px;">{{ t('Profile.Password') }}</el-text>
            </div>
            <div style="width: 100%;">
              <el-input
                style="margin: 10px 0 10px 0;"
                :placeholder="t('Profile.InputNewPassword')"
                v-model="resetPasswordData.password1"
                type="password"
                />
              <el-input
                style="margin: 10px 0 10px 0;"
                :placeholder="t('Profile.InputNewPasswordAgain')"
                v-model="resetPasswordData.password2"
                type="password"
                />
              <div style="width: 100%; display: flex; margin: 10px 0 0 0;">
                <div style="flex-grow: 1;"></div>
                <el-button
                  @click="ResetPassword"
                  type="primary"
                  >
                  {{ t('Profile.ResetPassword') }}
                </el-button>
              </div>
            </div>
          </div>
        </div>
        <div style="width: 600px">
          <h2 style="color: var(--el-text-color-primary)">{{ t('Profile.Settings') }}</h2>
          <div style="display: flex; margin: 20px">
            <el-text size="large" style="width: 100px; text-align: right; margin-right: 20px">{{ t('Profile.Language') }}</el-text>
            <LanguageSelect />
          </div>
          <div style="display: flex; margin: 20px">
            <el-text size="large" style="width: 100px; text-align: right; margin-right: 20px">{{ t('Profile.Theme') }}</el-text>
            <ThemeSelect />
          </div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>
<script>
import authApi from "@/api/auth.js"
import LanguageSelect from "@/components/LanguageSelect.vue"
import ThemeSelect from "@/components/ThemeSelect.vue"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: { LanguageSelect, ThemeSelect },
  data() {
    return {
      "user": {},
      "resetPasswordData": {
        "password1": "",
        "password2": ""
      }
    }
  },
  methods: {
    GetUserByAuth() {
      authApi.GetUserByAuth().then(res => {
        if (res.code === 0) {
          this.user = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ResetPassword() {
      if (this.resetPasswordData.password1 === "" || this.resetPasswordData.password2 === "") {
        ElMessage.error(this.t('Profile.EmptyPassword'))
        return
      }
      if (this.resetPasswordData.password1 !== this.resetPasswordData.password2) {
        ElMessage.error(this.t('Profile.InconsistencyPassword'))
        return
      }
      authApi.ResetPassword({
        "Password": this.resetPasswordData.password1
      }).then(res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
  },
  mounted() {
    this.GetUserByAuth()
  }
}
</script>
