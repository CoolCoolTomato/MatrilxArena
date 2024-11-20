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
              <el-progress
                :percentage="ctf.Progress"
              >
                <el-countdown
                  :value="ctf.RemainingTime"
                  value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
                />
              </el-progress>
              <div style="display:flex;">
                <el-text style="margin: 0 20px 0 20px;">{{ formatDateTime(ctf.StartTime) }}</el-text>
                <div style="flex: 1"></div>
                <el-text style="margin: 0 20px 0 20px;">{{ formatDateTime(ctf.EndTime) }}</el-text>
              </div>
            </div>
          </el-col>
          <el-col v-for="ctf in userCTFList" :span="8" v-if="showUserCTFList">
            <div class="ctf" @click="CTFClick(ctf)">
              <h2 style="color: var(--el-text-color-primary)">{{ ctf.Name }}</h2>
              <el-text truncated>{{ ctf.Description }}</el-text>
              <el-progress
                :percentage="ctf.Progress"
              >
                <el-countdown
                  :value="ctf.RemainingTime"
                  value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
                />
              </el-progress>
              <div style="display:flex;">
                <el-text style="margin: 0 20px 0 20px;">{{ formatDateTime(ctf.StartTime) }}</el-text>
                <div style="flex: 1"></div>
                <el-text style="margin: 0 20px 0 20px;">{{ formatDateTime(ctf.EndTime) }}</el-text>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
      <el-dialog
        v-model="ctfDetailVisible"
        :title="$t('CTF.CTFDetail')"
        width="600"
        @close="ClearUserCTFForm"
        >
        <h2 style="color: var(--el-text-color-primary)">{{ ctfDetail.Name }}</h2>
        <el-text truncated>{{ ctfDetail.Description }}</el-text>
        <el-progress
          :percentage="ctfDetail.Progress"
        >
          <el-countdown
            :value="ctfDetail.RemainingTime"
            value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
          />
        </el-progress>
        <div style="display:flex;">
          <el-text>{{ formatDateTime(ctfDetail.StartTime) }}</el-text>
          <div style="flex: 1"></div>
          <el-text>{{ formatDateTime(ctfDetail.EndTime) }}</el-text>
        </div>
        <template #footer>
          <el-button
            v-if="userInCTF && ctfDetail.Active === 3"
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
            @click="ctfDetailVisible = false"
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
      ctfDetailVisible: false,
      ctfDetail: {
        "ID": 0,
        "Name": "",
        "Description": "",
        "StartTime": "",
        "EndTime": "",
        "Active": 1,
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
          this.userCTFList = this.userCTFList.map(ctf => {
            const newCTF = {...ctf}
            newCTF.Progress = this.getProgress(ctf)
            if (newCTF.Progress === 0) {
              newCTF.Active = 1
            } else if (newCTF.Progress === 1) {
              newCTF.Active = 2
            } else {
              newCTF.Active = 3
            }
            newCTF.RemainingTime = this.toTimestamp(ctf.EndTime) * 1000
            return newCTF
          })
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
          this.visibleUserCTFList = this.visibleUserCTFList.map(ctf => {
            const newCTF = {...ctf}
            newCTF.Progress = this.getProgress(ctf)
            if (newCTF.Progress === 0) {
              newCTF.Active = 1
            } else if (newCTF.Progress === 1) {
              newCTF.Active = 2
            } else {
              newCTF.Active = 3
            }
            newCTF.RemainingTime = this.toTimestamp(ctf.EndTime) * 1000
            return newCTF
          })
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CTFClick(ctf) {
      this.ctfDetail = ctf
      this.userInCTF = this.checkUserInCTF(this.ctfDetail.ID)
      this.ctfDetailVisible = true
    },
    OpenCTFDetail() {
      this.$router.push(`/ctf/${this.ctfDetail.ID}`)
    },
    AddUserCTF() {
      userCTFApi.AddUserCTF(this.ctfDetail).then(res => {
        if (res.code === 0) {
          this.ctfDetailVisible = false
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
      userCTFApi.RemoveUserCTF(this.ctfDetail).then(res => {
        if (res.code === 0) {
          this.ctfDetailVisible = false
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
      this.ctfDetail = {
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
    },
    getProgress(ctf) {
      const remain = this.toTimestamp(ctf.EndTime) - Math.floor(Date.now() / 1000)
      const total = this.toTimestamp(ctf.EndTime) - this.toTimestamp(ctf.StartTime)
      if (remain <= 0) {
        return 0
      }
      if (remain >= total) {
        return 100
      }
      return Math.floor(remain / total * 100)
    },
    toTimestamp(dateStr) {
      const date = new Date(dateStr)
      return Math.floor(date.getTime() / 1000)
    },
    formatDateTime(dateTime) {
      dateTime = new Date(dateTime)
      const year = dateTime.getFullYear()
      const month = String(dateTime.getMonth() + 1).padStart(2, '0')
      const day = String(dateTime.getDate()).padStart(2, '0')
      const hours = String(dateTime.getHours()).padStart(2, '0')
      const minutes = String(dateTime.getMinutes()).padStart(2, '0')
      return `${year}-${month}-${day} ${hours}:${minutes}`
    },
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
  margin: 10px 20px 10px 20px;
}
.ctf > .el-text{
  margin: 10px 20px 10px 20px;
  width: calc(100% - 40px);
}
.ctf .el-progress {
  margin: 10px 20px 10px 20px;
}
</style>
