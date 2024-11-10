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
            <ThemeButton />
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-row style="margin: 20px">
          <el-col :span="12">
            <div id="login_content">
              <div style="display: flex; flex-direction: column; justify-items: center; align-items: center;">
                <el-text style="font-size: 60px; margin: 60px 0 10px 50px; color: var(--el-text-color-primary);" tag="b">Advanced Cybersecurity Learning Platform</el-text>
                <el-text style="font-size: 20px; margin: 10px 0 10px 50px; color: var(--el-text-color-primary);">MatrilxArena is an advanced cybersecurity learning platform, it uses multiple nodes to manage docker servers and provides flexible challenge, image, container management.</el-text>
              </div>
              <div style="margin: 20px 0 0 30px;">
                <el-row>
                  <el-col :span="12">
                    <el-card style="margin: 25px; border-radius: 10px; color: var(--el-text-color-primary);">
                      <p><b>Multi-node docker server management</b></p>
                    </el-card>
                  </el-col>
                  <el-col :span="12">
                    <el-card style="margin: 25px; border-radius: 10px; color: var(--el-text-color-primary);">
                      <p><b>Support for custom docker repositories</b></p>
                    </el-card>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :span="12">
                    <el-card style="margin: 25px; border-radius: 10px; color: var(--el-text-color-primary);">
                      <p><b>Flexible topic management</b></p>
                    </el-card>
                  </el-col>
                  <el-col :span="12">
                    <el-card style="margin: 25px; border-radius: 10px; color: var(--el-text-color-primary);">
                      <p><b>Cross-platform software installation</b></p>
                    </el-card>
                  </el-col>
                </el-row>
              </div>
            </div>
          </el-col>
          <el-col :span="12">
            <div id="login_form">
              <el-form
                v-model="userForm"
                label-position="left"
                label-width="auto"
                style="width: 500px"
              >
                <el-form-item>
                  <h2 style="width: 100%; text-align: center; margin: 10px; color: var(--el-text-color-regular); font-size: 30px">MtrilxArena</h2>
                </el-form-item>
                <el-form-item label="Username" size="large">
                  <el-input v-model="userForm.Username" size="large"/>
                </el-form-item>
                <el-form-item label="Password" size="large">
                  <el-input v-model="userForm.Password" type="password" size="large"/>
                </el-form-item>
                <el-form-item style="display:flex;">
                  <el-button @click="Login" style="width: 100%" size="large">Login</el-button>
                </el-form-item>
              </el-form>
            </div>
          </el-col>
        </el-row>
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
import ThemeButton from "@/components/ThemeButton.vue";
import Copyright from "@/components/Copyright.vue";

export default {
  components: {Copyright, ThemeButton},
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
.el-main{
  display: flex;
}
#login_content{
  height: 100%;
}
#login_form{
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
