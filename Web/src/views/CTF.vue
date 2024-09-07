<template>
  <el-container>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-affix :offset="100">
          <div style="display: flex; margin: 0 20px 20px 20px;" v-if="!showUserCTFList">
            <el-input v-model="visibleUserCTFQueryName" :placeholder="$t('CTF.FindCTFs')" size="large"/>
            <el-button @click="GetVisibleUserCTFList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('CTF.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
            <el-switch
              style="margin-left: 10px; white-space: nowrap;"
              v-model="showUserCTFList"
              @change="SwitchChange"
              :active-text="$t('CTF.ShowAddedCTFs')"
              size="large"
            />
          </div>
          <div style="display: flex; margin: 0 20px 20px 20px;" v-else>
            <el-input v-model="userCTFQueryName" :placeholder="$t('CTF.FindCTFs')" size="large"/>
            <el-button @click="GetUserCTFList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('CTF.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
            <el-switch
              style="margin-left: 10px; white-space: nowrap;"
              v-model="showUserCTFList"
              @change="SwitchChange"
              :active-text="$t('CTF.ShowAddedCTFs')"
              size="large"
            />
          </div>
        </el-affix>
        <el-row>
          <el-col v-for="ctf in visibleUserCTFList" :span="8" v-if="!showUserCTFList">
            <div class="ctf" @click="CTFClick(ctf)">
              <h2 style="color: var(--el-text-color-primary)">{{ ctf.Name }}</h2>
              <el-text truncated>{{ ctf.Description }}</el-text>
            </div>
          </el-col>
          <el-col v-for="ctf in userCTFList" :span="8" v-if="showUserCTFList">
            <div class="ctf" @click="CTFClick(ctf)">
              <h2 style="color: var(--el-text-color-primary)">{{ ctf.Name }}</h2>
              <el-text truncated>{{ ctf.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
      <el-dialog
        v-model="userCTFFormVisible"
        :title="$t('CTF.CTFDetail')"
        width="600"
        @close="ClearUserCTFForm"
        >
        <h2 style="color: var(--el-text-color-primary)">{{ userCTFData.Name }}</h2>
        <el-text truncated>{{ userCTFData.Description }}</el-text>
        <template #footer>
          <el-button
            v-if="userInCTF"
            @click="OpenCTFDetail"
            type="primary"
          >
            {{ $t('CTF.Open') }}
          </el-button>
          <el-button
            v-if="!userInCTF"
            @click="AddUserCTF"
            type="primary"
          >
            {{ $t('CTF.Add') }}
          </el-button>
          <el-button
            v-if="userInCTF"
            @click="RemoveUserCTF"
          >
            {{ $t('CTF.Remove') }}
          </el-button>
          <el-button
            @click="userCTFFormVisible = false"
          >
            {{ $t('CTF.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import userCTFApi from "@/api/userCTF.js"
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
      userCTFQueryName: "",
      userCTFList: [],
      visibleUserCTFQueryName: "",
      visibleUserCTFList: [],
      userCTFFormVisible: false,
      userCTFData: {
        "ID": 0,
        "Name": "",
        "Description": ""
      },
      userInCTF: false,
      showUserCTFList: false
    }
  },
  methods: {
    GetUserCTFList() {
      userCTFApi.GetUserCTFList({
        "Name": this.userCTFQueryName
      }).then(res => {
        if (res.code === 0) {
          this.userCTFList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetVisibleUserCTFList() {
      userCTFApi.GetVisibleUserCTFList({
        "Name": this.visibleUserCTFQueryName
      }).then(res => {
        if (res.code === 0) {
          this.visibleUserCTFList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CTFClick(ctf) {
      this.userCTFData = ctf
      this.userInCTF = this.checkUserInCTF(this.userCTFData.ID)
      this.userCTFFormVisible = true
    },
    OpenCTFDetail() {
      this.$router.push(`/ctf/${this.userCTFData.ID}`)
    },
    AddUserCTF() {
      userCTFApi.AddUserCTF(this.userCTFData).then(res => {
        if (res.code === 0) {
          this.userCTFFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetUserCTFList()
          this.GetVisibleUserCTFList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    RemoveUserCTF() {
      userCTFApi.RemoveUserCTF(this.userCTFData).then(res => {
        if (res.code === 0) {
          this.userCTFFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetUserCTFList()
          this.GetVisibleUserCTFList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearUserCTFForm() {
      this.userCTFData = {
        "ID": 0,
        "Name": "",
        "Description": ""
      }
      this.userInCTF = false
    },
    SwitchChange() {
      this.userCTFQueryName = ""
      this.visibleUserCTFQueryName = ""
      this.GetUserCTFList()
      this.GetVisibleUserCTFList()
    },
    checkUserInCTF(id) {
      return this.userCTFList.some(item => item.ID === id)
    }
  },
  mounted() {
    this.GetUserCTFList()
    this.GetVisibleUserCTFList()
  }
}
</script>
<style scoped>
.ctf {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
  cursor: pointer;
}
.ctf h2{
  margin: 20px;
}
.ctf .el-text{
  margin: 20px;
  width: calc(100% - 40px);
}
</style>
