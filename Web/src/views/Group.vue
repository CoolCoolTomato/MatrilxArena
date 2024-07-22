<template>
  <el-container>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-affix :offset="100">
          <div style="display: flex; margin: 0 20px 20px 20px;" v-if="!showUserGroupList">
            <el-input v-model="visibleUserGroupQueryName" :placeholder="$t('Group.FindGroups')" size="large"/>
            <el-button @click="GetVisibleUserGroupList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('Group.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
            <el-switch
              style="margin-left: 10px; white-space: nowrap;"
              v-model="showUserGroupList"
              @change="SwitchChange"
              :active-text="$t('Group.ShowAddedGroups')"
              size="large"
            />
          </div>
          <div style="display: flex; margin: 0 20px 20px 20px;" v-else>
            <el-input v-model="userGroupQueryName" :placeholder="$t('Group.FindGroups')" size="large"/>
            <el-button @click="GetUserGroupList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('Group.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
            <el-switch
              style="margin-left: 10px; white-space: nowrap;"
              v-model="showUserGroupList"
              @change="SwitchChange"
              :active-text="$t('Group.ShowAddedGroups')"
              size="large"
            />
          </div>
        </el-affix>
        <el-row>
          <el-col v-for="group in visibleUserGroupList" :span="8" v-if="!showUserGroupList">
            <div class="group" @click="GroupClick(group)">
              <h2 style="color: var(--el-text-color-primary)">{{ group.Name }}</h2>
              <el-text truncated>{{ group.Description }}</el-text>
            </div>
          </el-col>
          <el-col v-for="group in userGroupList" :span="8" v-if="showUserGroupList">
            <div class="group" @click="GroupClick(group)">
              <h2 style="color: var(--el-text-color-primary)">{{ group.Name }}</h2>
              <el-text truncated>{{ group.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
      <el-dialog
        v-model="userGroupFormVisible"
        :title="$t('Group.GroupDetail')"
        width="600"
        @close="ClearUserGroupForm"
        >
        <h2 style="color: var(--el-text-color-primary)">{{ userGroupData.Name }}</h2>
        <el-text truncated>{{ userGroupData.Description }}</el-text>
        <template #footer>
          <el-button
            v-if="userInGroup"
            @click="OpenGroupDetail"
            type="primary"
          >
            {{ $t('Group.Open') }}
          </el-button>
          <el-button
            v-if="!userInGroup"
            @click="AddUserGroup"
            type="primary"
          >
            {{ $t('Group.Add') }}
          </el-button>
          <el-button
            v-if="userInGroup"
            @click="RemoveUserGroup"
          >
            {{ $t('Group.Remove') }}
          </el-button>
          <el-button
            @click="userGroupFormVisible = false"
          >
            {{ $t('Group.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import userGroupApi from "@/api/userGroup.js"
import { ElMessage } from "element-plus"
import {useI18n} from "vue-i18n"
import Search from "@/components/icons/Search.vue"

export default {
  components: { Search },
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      userGroupQueryName: "",
      userGroupList: [],
      visibleUserGroupQueryName: "",
      visibleUserGroupList: [],
      userGroupFormVisible: false,
      userGroupData: {
        "ID": 0,
        "Name": "",
        "Description": ""
      },
      userInGroup: false,
      showUserGroupList: false
    }
  },
  methods: {
    GetUserGroupList() {
      userGroupApi.GetUserGroupList({
        "Name": this.userGroupQueryName
      }).then(res => {
        if (res.code === 0) {
          this.userGroupList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetVisibleUserGroupList() {
      userGroupApi.GetVisibleUserGroupList({
        "Name": this.visibleUserGroupQueryName
      }).then(res => {
        if (res.code === 0) {
          this.visibleUserGroupList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GroupClick(group) {
      this.userGroupData = group
      this.userInGroup = this.checkUserInGroup(this.userGroupData.ID)
      this.userGroupFormVisible = true
    },
    OpenGroupDetail() {
      this.$router.push(`/group/${this.userGroupData.ID}`)
    },
    AddUserGroup() {
      userGroupApi.AddUserGroup(this.userGroupData).then(res => {
        if (res.code === 0) {
          this.userGroupFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetUserGroupList()
          this.GetVisibleUserGroupList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    RemoveUserGroup() {
      userGroupApi.RemoveUserGroup(this.userGroupData).then(res => {
        if (res.code === 0) {
          this.userGroupFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetUserGroupList()
          this.GetVisibleUserGroupList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearUserGroupForm() {
      this.userGroupData = {
        "ID": 0,
        "Name": "",
        "Description": ""
      }
      this.userInGroup = false
    },
    SwitchChange() {
      this.userGroupQueryName = ""
      this.visibleUserGroupQueryName = ""
      this.GetUserGroupList()
      this.GetVisibleUserGroupList()
    },
    checkUserInGroup(id) {
      return this.userGroupList.some(item => item.ID === id)
    }
  },
  mounted() {
    this.GetUserGroupList()
    this.GetVisibleUserGroupList()
  }
}
</script>
<style scoped>
.group {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
  cursor: pointer;
}
.group h2{
  margin: 20px;
}
.group .el-text{
  margin: 20px;
  width: calc(100% - 40px);
}
</style>
