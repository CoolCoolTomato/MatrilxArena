<template>
  <div id="login">
    <el-container>
      <el-header>
        <el-menu
          mode="horizontal"
          :ellipsis="false"
          style="border: none"
        >
          <el-menu-item class="el-menu-item-no-style">
            <el-image src="/image/svg/matrilx.svg" style="width: 40px; margin-bottom: 16px; margin-right: 5px;" />
            <h2 style="color: var(--el-text-color-primary);">MatrilxArena</h2>
          </el-menu-item>
          <div class="flex-grow" />
          <el-menu-item class="el-menu-item-no-style">
            <ToggleThemeButton />
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <div id="login_form">
          <el-form
            v-model="userForm"
            label-position="left"
            label-width="auto"
            style="width: 320px"
          >
            <el-form-item>
              <h2 style="width: 100%; text-align: center; margin: 0; color: var(--el-text-color-regular); font-size: 25px">MtrilxArena</h2>
            </el-form-item>
            <el-form-item label="Username">
              <el-input v-model="userForm.Username"/>
            </el-form-item>
            <el-form-item label="Password">
              <el-input v-model="userForm.Password"/>
            </el-form-item>
            <el-form-item style="display:flex;">
              <el-button @click="Login" style="width: 100%">Login</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-main>
      <el-footer>
        <Copyright />
      </el-footer>
    </el-container>
  </div>
</template>
<script>
import authApi from "@/api/auth.js";
import {ElMessage} from "element-plus";
import ToggleThemeButton from "@/components/ToggleThemeButton.vue";
import Copyright from "@/components/Copyright.vue";

export default {
  components: {Copyright, ToggleThemeButton},
  data() {
    return {
      userForm: {
        "Username": "",
        "Password": "",
        "Email": "",
      }
    }
  },
  methods: {
    Login() {
      authApi.Login(this.userForm).then(res => {
        if (res.code === 0) {
          localStorage.setItem("token", res.data);
          this.$router.push("/");
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
  }
}
</script>
<style scoped>
#login{
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
::v-deep(.el-menu-item-no-style > .el-button > [class^="el-icon"]) {
  margin-right: 0;
}
.el-container{
  width: 100%;
  height: 100%;
}
#login_form{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.el-footer{
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
