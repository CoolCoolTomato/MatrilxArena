<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">{{ $t('AdminUser.UserManager') }}</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createUserFormVisible = true">{{ $t('AdminUser.Add') }}</el-button>
      <el-table
        :data="userList"
        table-layout="fixed"
      >
        <el-table-column prop="Username" :label="$t('AdminUser.Username')"/>
        <el-table-column prop="Email" :label="$t('AdminUser.Email')"/>
        <el-table-column :label="$t('AdminUser.Role')">
          <template #default="scope">
            {{ formatRole(scope.row.Role) }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" :label="$t('AdminUser.Operations')" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateUserForm(scope.row)"
              >
              {{ $t('AdminUser.Update') }}
            </el-button>
            <el-button
              @click="OpenDeleteUserForm(scope.row)"
              >
              {{ $t('AdminUser.Delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createUserFormVisible"
        :title="$t('AdminUser.CreateUser')"
        width="500"
        @close="ClearCreateUserForm"
        >
        <el-form :model=createUserData>
          <el-form-item :label="$t('AdminUser.Username')" :label-width="labelWidth">
            <el-input v-model="createUserData.Username"/>
          </el-form-item>
          <el-form-item :label="$t('AdminUser.Email')" :label-width="labelWidth">
            <el-input v-model="createUserData.Email"/>
          </el-form-item>
          <el-form-item :label="$t('AdminUser.Role')" :label-width="labelWidth">
            <el-select
              v-model="createUserData.Role"
              :placeholder="$t('AdminUser.Select')"
            >
              <el-option v-for="role in roles" :label="role.label" :key="role.value" :value="role.value"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminUser.Password')" :label-width="labelWidth">
            <el-input v-model="createUserData.Password"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createUserFormVisible = false">{{ $t('AdminUser.Cancel') }}</el-button>
          <el-button @click="CreateUser">{{ $t('AdminUser.Submit') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateUserFormVisible"
        :title="$t('AdminUser.UpdateUser')"
        width="500"
        @close="ClearUpdateUserForm"
        >
        <el-form :model=updateUserData>
          <el-form-item :label="$t('AdminUser.Username')" :label-width="labelWidth">
            <el-input v-model="updateUserData.Username"/>
          </el-form-item>
          <el-form-item :label="$t('AdminUser.Email')" :label-width="labelWidth">
            <el-input v-model="updateUserData.Email"/>
          </el-form-item>
          <el-form-item :label="$t('AdminUser.Role')" :label-width="labelWidth">
            <el-select
              v-model="updateUserData.Role"
              :placeholder="$t('AdminUser.Select')"
            >
              <el-option v-for="role in roles" :label="role.label" :key="role.value" :value="role.value"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminUser.Password')" :label-width="labelWidth">
            <el-input v-model="updateUserData.Password"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateUserFormVisible = false">{{ $t('AdminUser.Cancel') }}</el-button>
          <el-button @click="UpdateUser">{{ $t('AdminUser.Submit') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteUserFormVisible"
        :title="$t('AdminUser.DeleteUser')"
        width="500"
        @close="ClearDeleteUserForm"
        >
        <el-text>{{ $t('AdminUser.AreYouConfirmToDeleteTheUser') }}</el-text>
        <template #footer>
          <el-button @click="deleteUserFormVisible = false">{{ $t('AdminUser.Cancel') }}</el-button>
          <el-button @click="DeleteUser">{{ $t('AdminUser.Confirm') }}</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import userApi from "@/api/user.js"
import { ElMessage } from 'element-plus'
import {useI18n} from "vue-i18n";

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      labelWidth: 100,
      userList: [],
      createUserFormVisible: false,
      createUserData: {
        "Username": "",
        "Password": "",
        "Email": "",
        "Role": null,
      },
      updateUserFormVisible: false,
      updateUserData: {
        "ID": 0,
        "Username": "",
        "Password": "",
        "Email": "",
        "Role": null,
      },
      deleteUserFormVisible: false,
      deleteUserData: {
        "ID": 0,
      },
      roles: [
        {
          "value": 1,
          "label": this.t('AdminUser.Admin')
        },
        {
          "value": 2,
          "label": this.t('AdminUser.User')
        }
      ]
    }
  },
  methods: {
    GetUserList() {
      userApi.GetUserList().then(res => {
        if (res.code === 0) {
          this.userList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateUser() {
      userApi.CreateUser(this.createUserData).then(res => {
        if (res.code === 0) {
          this.createUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetUserList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateUserForm() {
      this.createUserData = {
        "Username": "",
        "Password": "",
        "Email": "",
        "Role": null,
      }
    },
    UpdateUser() {
      userApi.UpdateUser(this.updateUserData).then(res => {
        if (res.code === 0) {
          this.updateUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetUserList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateUserForm(row) {
      this.updateUserData = {
        "ID": row.ID,
        "Username": row.Username,
        "Password": row.Password,
        "Email": row.Email,
        "Role": row.Role,
      }
      this.updateUserFormVisible = true
    },
    ClearUpdateUserForm() {
      this.updateUserData = {
        "ID": 0,
        "Username": "",
        "Password": "",
        "Email": "",
        "Role": null,
      }
    },
    DeleteUser() {
      userApi.DeleteUser(this.deleteUserData).then(res => {
        if (res.code === 0) {
          this.deleteUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetUserList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteUserForm(row) {
      this.deleteUserData = {
        "ID": row.ID,
      }
      this.deleteUserFormVisible = true
    },
    ClearDeleteUserForm() {
      this.deleteUserData = {
        "ID": 0,
      }
    },
    formatRole(role) {
      if (role === 1) {
        return this.t('AdminUser.Admin')
      } else if (role === 2) {
        return this.t('AdminUser.User')
      } else {
        return this.t('AdminUser.Null')
      }
    },
  },
  mounted() {
    this.GetUserList()
  },
}
</script>
