<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">User Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createUserFormVisible = true">Add</el-button>
      <el-table
        :data="userList"
        table-layout="fixed"
      >
        <el-table-column prop="Username" label="Username"/>
        <el-table-column prop="Email" label="Email"/>
        <el-table-column label="Role">
          <template #default="scope">
            {{ formatRole(scope.row.Role) }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateUserForm(scope.row)"
              >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteUserForm(scope.row)"
              >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createUserFormVisible"
        title="Create User"
        width="500"
        @close="ClearCreateUserForm"
        >
        <el-form :model=createUserData>
          <el-form-item label="Username">
            <el-input v-model="createUserData.Username"/>
          </el-form-item>
          <el-form-item label="Email">
            <el-input v-model="createUserData.Email"/>
          </el-form-item>
          <el-form-item label="Role">
            <el-select v-model="createUserData.Role">
              <el-option v-for="role in roles" :label="role.label" :key="role.value" :value="role.value"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Password">
            <el-input v-model="createUserData.Password"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createUserFormVisible = false">Cancel</el-button>
          <el-button @click="CreateUser">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateUserFormVisible"
        title="Update Node"
        width="500"
        @close="ClearUpdateUserForm"
        >
        <el-form :model=updateUserData>
          <el-form-item label="Username">
            <el-input v-model="updateUserData.Username"/>
          </el-form-item>
          <el-form-item label="Email">
            <el-input v-model="updateUserData.Email"/>
          </el-form-item>
          <el-form-item label="Role">
            <el-select v-model="updateUserData.Role">
              <el-option v-for="role in roles" :label="role.label" :key="role.value" :value="role.value"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Password">
            <el-input v-model="updateUserData.Password"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateUserFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateUser">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteUserFormVisible"
        title="Delete User"
        width="500"
        @close="ClearDeleteUserForm"
        >
        <el-text>Are you confirm to delete the user?</el-text>
        <template #footer>
          <el-button @click="deleteUserFormVisible = false">Cancel</el-button>
          <el-button @click="DeleteUser">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import userApi from "@/api/user.js"
import { ElMessage } from 'element-plus'

export default {
  data() {
    return {
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
          "label": "admin"
        },
        {
          "value": 2,
          "label": "user"
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
        return "admin"
      } else if (role === 2) {
        return "user"
      } else {
        return "select"
      }
    },
  },
  mounted() {
    this.GetUserList()
  },
}
</script>
