<template>
  <el-container>
    <el-header>
      <div style="display: flex; align-items: center;">
        <h2 style="color: var(--el-text-color-primary);">{{ t('AdminGroup.GroupManager') }}</h2>
        <div style="flex-grow: 1;"></div>
        <div style="margin-right: 50px;">
          <el-button
            style="margin: 10px;"
            @click="createGroupFormVisible = true"
            type="primary"
          >
            {{ t('AdminGroup.Add') }}
          </el-button>
        </div>
      </div>
    </el-header>
    <el-main>
      <el-table
        :data="groupList"
        table-layout="fixed"
        height="100%"
      >
        <el-table-column prop="Name" :label="t('AdminGroup.Name')"/>
        <el-table-column prop="Description" :label="t('AdminGroup.Description')"/>
        <el-table-column :label="t('AdminGroup.GroupChallenges')">
          <template #default="scope">
            {{ scope.row.GroupChallenges.length }}
          </template>
        </el-table-column>
        <el-table-column :label="t('AdminGroup.Users')">
          <template #default="scope">
            {{ scope.row.Users.length }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" :label="t('AdminGroup.Operations')" width="300">
          <template #default=scope>
            <el-button
              @click="OpenGroupDetail(scope.row)"
            >
              {{ $t('AdminGroup.Detail') }}
            </el-button>
            <el-button
              @click="OpenUpdateGroupForm(scope.row)"
            >
              {{ t('AdminGroup.Update') }}
            </el-button>
            <el-button
              @click="OpenDeleteGroupForm(scope.row)"
            >
              {{ t('AdminGroup.Delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createGroupFormVisible"
        :title="t('AdminGroup.CreateGroup')"
        width="500"
        @close="ClearCreateGroupForm"
      >
        <el-form :model="createGroupData">
          <el-form-item :label="t('AdminGroup.Name')" :label-width="labelWidth">
            <el-input v-model="createGroupData.Name"/>
          </el-form-item>
          <el-form-item :label="t('AdminGroup.Description')" :label-width="labelWidth">
            <el-input v-model="createGroupData.Description" type="textarea" autosize/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateGroup"
            type="primary"
          >
            {{ t('AdminGroup.Submit') }}
          </el-button>
          <el-button
            @click="createGroupFormVisible = false"
          >
            {{ t('AdminGroup.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateGroupFormVisible"
        :title="t('AdminGroup.UpdateGroup')"
        width="500"
        @close="ClearUpdateGroupForm"
      >
        <el-form :model=updateGroupData>
          <el-form-item :label="t('AdminGroup.Name')" :label-width="labelWidth">
            <el-input v-model="updateGroupData.Name"/>
          </el-form-item>
          <el-form-item :label="t('AdminGroup.Description')" :label-width="labelWidth">
            <el-input v-model="updateGroupData.Description" type="textarea" autosize/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateGroup"
            type="primary"
          >
            {{ t('AdminGroup.Submit') }}
          </el-button>
          <el-button
            @click="updateGroupFormVisible = false"
          >
            {{ t('AdminGroup.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteGroupFormVisible"
        :title="t('AdminGroup.DeleteGroup')"
        width="500"
        @close="ClearDeleteGroupForm"
      >
        <el-text>{{ t('AdminGroup.AreYouConfirmToDeleteTheGroup') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteGroup"
            type="primary"
          >
            {{ t('AdminGroup.Confirm') }}
          </el-button>
          <el-button
            @click="deleteGroupFormVisible = false"
          >
            {{ t('AdminGroup.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import groupApi from "@/api/group.js"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      labelWidth: 100,
      groupList: [],
      createGroupFormVisible: false,
      createGroupData: {
        "Name": "",
        "Description": "",
      },
      updateGroupFormVisible: false,
      updateGroupData: {
        "ID": 0,
        "Name": "",
        "Description": "",
      },
      deleteGroupFormVisible: false,
      deleteGroupData: {
        "ID": 0,
      }
    }
  },
  methods: {
    GetGroupList() {
      groupApi.GetGroupList().then(res => {
        if (res.code === 0) {
          this.groupList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenGroupDetail(row) {
      this.$router.push(`/admin/group/${row.ID}`);
    },
    CreateGroup() {
      groupApi.CreateGroup(this.createGroupData).then(res => {
        if (res.code === 0) {
          this.createGroupFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetGroupList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateGroupForm() {
      this.createGroupData = {
        "Name": "",
        "Description": "",
      }
    },
    UpdateGroup() {
      groupApi.UpdateGroup(this.updateGroupData).then(res => {
        if (res.code === 0) {
          this.updateGroupFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetGroupList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateGroupForm(row) {
      this.updateGroupData = {
        "ID": row.ID,
        "Name": row.Name,
        "Description": row.Description,
      }
      this.updateGroupFormVisible = true
    },
    ClearUpdateGroupForm() {
      this.updateGroupData = {
        "ID": 0,
        "Name": "",
        "Description": "",
      }
    },
    DeleteGroup() {
      groupApi.DeleteGroup(this.deleteGroupData).then(res => {
        if (res.code === 0) {
          this.deleteGroupFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetGroupList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteGroupForm(row) {
      this.deleteGroupData = {
        "ID": row.ID,
      }
      this.deleteGroupFormVisible = true
    },
    ClearDeleteGroupForm() {
      this.deleteGroupData = {
        "ID": 0,
      }
    },
  },
  mounted() {
    this.GetGroupList()
  }
}
</script>
