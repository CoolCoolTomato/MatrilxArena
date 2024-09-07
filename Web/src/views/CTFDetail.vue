<template>
  <el-container>
    <el-aside :width="isMenuOpen ? '190px' : '64px'">
      <el-scrollbar>
        <el-menu :collapse="!isMenuOpen" style="border: none;" router>
          <el-menu-item @click="toggleSidebar">
            <el-icon>
              <Menu />
            </el-icon>
            <template #title>{{ $t('CTFDetail.Menu') }}</template>
          </el-menu-item>
          <el-sub-menu :index="'/ctf/' + ctf.ID">
            <template #title>
              <el-icon>
                <Category/>
              </el-icon>
              <el-text style="color: var(--el-text-color-primary);">{{ $t('CTFDetail.Category') }}</el-text>
            </template>
            <el-menu-item :index="'/ctf/' + ctf.ID">
              <template #title>{{ $t('CTFDetail.All') }}</template>
            </el-menu-item>
            <el-menu-item
              v-for="category in categoryList"
              :key="category.ID"
              :index="'/ctf/' + ctf.ID + '/' + category.Name"
            >
              <template #title>{{ category.Name }}</template>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item @click="userCTFChallengeVisible = true">
            <el-icon>
              <Finish/>
            </el-icon>
            <template #title>{{ $t('CTFDetail.Solved') }}</template>
          </el-menu-item>
          <el-menu-item @click="userCTFContainerListVisible = true">
            <el-icon>
              <Container/>
            </el-icon>
            <template #title>{{ $t('CTFDetail.Containers') }}</template>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-affix :offset="100">
          <div style="display: flex; margin: 0 20px 20px 20px; align-items: center">
            <el-breadcrumb separator="/" style="width: 400px; font-size: 20px;">
              <el-breadcrumb-item to="/ctf">{{ $t('CTFDetail.CTF') }}</el-breadcrumb-item>
              <el-breadcrumb-item :to="'/ctf/' + ctf.ID">{{ ctf.Name }}</el-breadcrumb-item>
              <el-breadcrumb-item
                v-if="this.$route.params.category !== ''"
                :to="'/ctf/' + ctf.ID + '/' + this.$route.params.category"
                >
                {{ this.$route.params.category }}
              </el-breadcrumb-item>
            </el-breadcrumb>
            <el-input v-model="ctfChallengeQueryTitle" :placeholder="$t('CTFDetail.FindChallenges')" size="large"/>
            <el-button @click="GetCTFChallengeList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('CTFDetail.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
          </div>
        </el-affix>
        <el-row>
          <el-col v-for="ctfChallenge in ctfChallengeList" :span="12">
            <div class="ctfChallenge" @click="OpenCTFChallengeDetail(ctfChallenge)">
              <h2 style="color: var(--el-text-color-primary)">{{ ctfChallenge.Title }}<el-text v-if="checkChallengeSolved(ctfChallenge.ID)">{{ $t('CTFDetail.Solved') }}</el-text></h2>
              <el-text truncated>{{ ctfChallenge.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
      <el-dialog
        v-model="ctfChallengeDetailVisible"
        :title="$t('CTFDetail.ChallengeDetail')"
        width="600"
        @close="ClearCTFChallengeDetail"
        >
        <h2 style="color: var(--el-text-color-primary)">{{ ctfChallengeDetail.Title }} <el-text v-if="checkChallengeSolved(ctfChallengeDetail.ID)">({{ $t('CTFDetail.Solved') }})</el-text></h2>
        <el-text>{{ ctfChallengeDetail.Description }}</el-text>
        <div style="margin-top: 5px;">
          <el-text>{{ $t('CTFDetail.Category') }}: </el-text>
          <el-text>{{ ctfChallengeDetail.CategoryName === "" ? $t('CTFDetail.Uncategorized') : ctfChallengeDetail.CategoryName }}</el-text>
        </div>
        <div v-if="ctfChallengeDetail.Attachment.ID !== 0" style="margin-top: 5px;">
          <el-text>{{ $t('CTFDetail.Attachment') }}: </el-text>
          <el-link
            @click="DownloadAttachment(ctfChallengeDetail.Attachment)"
            style="vertical-align: unset"
            :underline="false"
          >
            {{ ctfChallengeDetail.Attachment.FileName }}
          </el-link>
        </div>
        <div
          v-if="checkContainerInUse(ctfChallengeDetail.ID)"
          v-for="portMap in userCTFContainer.PortMaps"
          style="margin-top: 15px">
          <el-text>
            {{ portMap.PortProtocol }} -> {{ portMap.Link }}
          </el-text>
        </div>
        <el-progress
          v-if="checkContainerInUse(ctfChallengeDetail.ID)"
          :percentage="(userCTFContainer.RemainingTime - Date.now()) / 36000"
          style="margin-top: 15px"
          >
          <el-countdown
            :value="userCTFContainer.RemainingTime"
            :finish="this.GetCTFContainerListByUser"
            value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
          />
        </el-progress>
        <div class="operations">
          <el-button
            @click="ResetuserCTFChallenge(ctfChallengeDetail.ID)"
            >
            {{ $t('CTFDetail.Reset') }}
          </el-button>
          <el-button
            @click="CreateContainerByUser(createCTFContainerByUserData)"
            v-if="!checkContainerInUse(ctfChallengeDetail.ID) && ctfChallengeDetail.ImageID !== 0"
            >
            {{ $t('CTFDetail.Create') }}
          </el-button>
          <el-button
            @click="DestroyContainerByUser(destroyCTFContainerByUserData)"
            v-if="checkContainerInUse(ctfChallengeDetail.ID)"
            >
            {{ $t('CTFDetail.Destroy') }}
          </el-button>
          <el-button
            @click="DelayContainerByUser(delayCTFContainerByUserData)"
            v-if="checkContainerInUse(ctfChallengeDetail.ID)"
            >
            {{ $t('CTFDetail.Delay') }}
          </el-button>
        </div>
        <div style="display:flex; margin-top: 15px" v-if="checkContainerInUse(ctfChallengeDetail.ID) || ctfChallengeDetail.ImageID === 0">
          <el-input v-model="checkFlagData.Flag" :placeholder="$t('CTFDetail.InputFlag')"/>
          <el-button @click="CheckCTFChallengeFlag">{{ $t('CTFDetail.Submit') }}</el-button>
        </div>
      </el-dialog>
      <el-dialog
        v-model="userCTFContainerListVisible"
        :title="$t('CTFDetail.ContainerList')"
        width="600"
        >
        <el-text v-if="userCTFContainerList.length === 0" size="large">{{ $t('CTFDetail.NoContainerIsRunning') }}</el-text>
        <el-card
          v-for="userCTFContainer in userCTFContainerList"
          style="margin: 20px 0 20px 0;"
          >
          <h2 style="color: var(--el-text-color-primary)">
          {{ getContainerChallenge(userCTFContainer.CTFChallengeID).CTF.Name }}
          /
          {{ getContainerChallenge(userCTFContainer.CTFChallengeID) === undefined ? "" : getContainerChallenge(userCTFContainer.CTFChallengeID).Title }}
          </h2>
          <div
            v-for="portMap in userCTFContainer.PortMaps"
            style="margin-top: 15px">
            <el-text>
              {{ portMap.PortProtocol }} -> {{ portMap.Link }}
            </el-text>
          </div>
          <el-progress
            :percentage="(userCTFContainer.RemainingTime - Date.now()) / 36000"
            style="margin-top: 15px"
            >
            <el-countdown
              :value="userCTFContainer.RemainingTime"
              :finish="this.GetCTFContainerListByUser"
              value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
            />
          </el-progress>
          <div class="operations">
            <el-button
              @click="DestroyContainerByUser(userCTFContainer)"
              >
              {{ $t('CTFDetail.Destroy') }}
            </el-button>
            <el-button
              @click="DelayContainerByUser(userCTFContainer)"
              >
              {{ $t('CTFDetail.Delay') }}
            </el-button>
          </div>
        </el-card>
      </el-dialog>
      <el-dialog
        v-model="userCTFChallengeVisible"
        :title="$t('CTFDetail.SolvedChallenges')"
        width="70%"
      >
        <el-scrollbar style="height: 60vh;">
          <el-row>
            <el-col v-for="ctfChallenge in userCTFChallengeList" :span="12">
              <div class="ctfChallenge" @click="OpenCTFChallengeDetail(ctfChallenge)">
                <h2 style="color: var(--el-text-color-primary)">{{ ctfChallenge.Title }}<el-text v-if="checkChallengeSolved(ctfChallenge.ID)">{{ $t('CTFDetail.Solved') }}</el-text></h2>
                <el-text truncated>{{ ctfChallenge.Description }}</el-text>
              </div>
            </el-col>
          </el-row>
        </el-scrollbar>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import ctfApi from "@/api/ctf.js"
import categoryApi from "@/api/category.js"
import ctfChallengeApi from "@/api/ctfChallenge.js"
import userCTFContainerApi from "@/api/userCTFContainer.js"
import userCTFChallengeApi from "@/api/userCTFChallenge.js"
import { ElMessage } from "element-plus"
import attachmentApi from "@/api/attachment.js"

import Menu from "@/components/icons/Menu.vue"
import Search from "@/components/icons/Search.vue"
import Category from "@/components/icons/Category.vue"
import Finish from "@/components/icons/Finish.vue"
import Container from "@/components/icons/Container.vue"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: { Menu, Search, Category, Finish, Container },
  data() {
    return {
      isMenuOpen: false,
      ctf: {
        "ID": 0,
      },
      userCTFContainerList: [],
      userCTFContainerListVisible: false,
      userCTFContainer: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "CTFChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      },
      categoryList: [],
      totalCTFChallengeList: [],
      ctfChallengeQueryCategoryID: 0,
      ctfChallengeQueryTitle: "",
      ctfChallengeList: [],
      ctfChallengeDetailVisible: false,
      ctfChallengeDetail: {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryName": "",
        "ImageID": 0,
        "Attachment": {}
      },
      createCTFContainerByUserData: {
        "ID": 0,
      },
      destroyCTFContainerByUserData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      },
      delayCTFContainerByUserData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      },
      userCTFChallengeVisible: false,
      userCTFChallengeList: [],
      resetUserCTFChallengeData: {
        "CTFChallengeID": 0
      },
      checkFlagData: {
        "CTFChallengeID": 0,
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "Flag": "",
      },
      intervalId: null
    }
  },
  methods: {
    toggleSidebar() {
      this.isMenuOpen = !this.isMenuOpen;
    },
    GetCTF() {
      ctfApi.GetCTF(this.ctf).then(res => {
        if (res.code === 0) {
          this.ctf = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    async GetCategoryList() {
      return categoryApi.GetCategoryList().then(res => {
        if (res.code === 0) {
          this.categoryList = res.data
          this.categoryList.sort((a, b) => a.Order - b.Order)
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    async GetCTFContainerListByUser() {
      return userCTFContainerApi.GetCTFContainerListByUser().then(res => {
        if (res.code === 0) {
          this.userCTFContainerList = res.data == null ? [] : res.data
          this.userCTFContainerList = this.userCTFContainerList.map(container => {
            const newContainer = {...container}
            newContainer.RemainingTime = Math.trunc(newContainer.RemainingTime / 1000000 + Date.now())
            newContainer.CTFChallengeID = newContainer.ChallengeID
            return newContainer
          })
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetCTFChallengeList() {
      ctfChallengeApi.GetCTFChallengeList({}).then(res => {
        if (res.code === 0) {
          this.totalCTFChallengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
      ctfChallengeApi.GetCTFChallengeList({
        "CTFID": this.ctf.ID,
        "CategoryID": this.ctfChallengeQueryCategoryID,
        "Title": this.ctfChallengeQueryTitle
      }).then(res => {
        if (res.code === 0) {
          this.ctfChallengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCTFChallengeDetail(ctfChallenge) {
      this.ctfChallengeDetailVisible = true
      this.ctfChallengeDetail = {
        "ID": ctfChallenge.ID,
        "Title": ctfChallenge.Title,
        "Description": ctfChallenge.Description,
        "CategoryName": ctfChallenge.Category.Name,
        "ImageID": ctfChallenge.ImageID,
        "Attachment": ctfChallenge.Attachment
      }
      this.createCTFContainerByUserData = {
        "ID": ctfChallenge.ID,
      }
      this.checkFlagData = {
        "CTFChallengeID": ctfChallenge.ID,
      }
      if (this.checkContainerInUse(ctfChallenge.ID)) {
        this.getContainerInUse(ctfChallenge.ID)
        this.destroyCTFContainerByUserData = {
          "DockerNodeID": this.userCTFContainer.DockerNodeID,
          "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
        }
        this.delayCTFContainerByUserData = {
          "DockerNodeID": this.userCTFContainer.DockerNodeID,
          "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
        }
        this.checkFlagData = {
          "CTFChallengeID": this.ctfChallengeDetail.ID,
          "DockerNodeID": this.userCTFContainer.DockerNodeID,
          "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
        }
      }
      this.resetUserCTFChallengeData = {
        "CTFChallengeID": ctfChallenge.ID
      }
    },
    ClearCTFChallengeDetail() {
      this.ctfChallengeDetail = {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryName": "",
        "ImageID": 0,
        "Attachment": {}
      }
      this.createCTFContainerByUserData = {
        "ID": 0,
      }
      this.destroyCTFContainerByUserData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      }
      this.delayCTFContainerByUserData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      }
      this.userCTFContainer = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "CTFChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      }
      this.resetUserCTFChallengeData = {
        "CTFChallengeID": 0
      }
      this.checkFlagData = {
        "CTFChallengeID": 0,
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "Flag": "",
      }
    },
    CreateContainerByUser(createCTFContainerByUserData) {
      userCTFContainerApi.CreateCTFContainerByUser(createCTFContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetCTFContainerListByUser()
        if (this.checkContainerInUse(this.ctfChallengeDetail.ID)) {
          this.getContainerInUse(this.ctfChallengeDetail.ID)
          this.destroyCTFContainerByUserData = {
            "DockerNodeID": this.userCTFContainer.DockerNodeID,
            "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
          }
          this.delayCTFContainerByUserData = {
            "DockerNodeID": this.userCTFContainer.DockerNodeID,
            "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
          }
          this.checkFlagData = {
            "CTFChallengeID": this.ctfChallengeDetail.ID,
            "DockerNodeID": this.userCTFContainer.DockerNodeID,
            "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
          }
        }
      }).catch(error => {
        console.log(error)
      })
    },
    DestroyContainerByUser(destroyCTFContainerByUserData) {
      userCTFContainerApi.DestroyCTFContainerByUser(destroyCTFContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetCTFContainerListByUser()
        this.destroyCTFContainerByUserData = {
          "DockerNodeID": 0,
          "DockerNodeContainerID": ""
        }
        this.delayCTFContainerByUserData = {
          "DockerNodeID": 0,
          "DockerNodeContainerID": ""
        }
        this.userCTFContainer = {
          "DockerNodeContainerID": "",
          "DockerNodeID": 0,
          "CTFChallengeID": 0,
          "RemainingTime": 0,
          "PortMaps": []
        }
        this.checkFlagData = {
          "CTFChallengeID": 0,
          "DockerNodeID": 0,
          "DockerNodeContainerID": "",
          "Flag": "",
        }
      }).catch(error => {
        console.log(error)
      })
    },
    DelayContainerByUser(delayCTFContainerByUserData) {
      userCTFContainerApi.DelayCTFContainerByUser(delayCTFContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetCTFContainerListByUser()
        if (this.checkContainerInUse(this.ctfChallengeDetail.ID)) {
          this.getContainerInUse(this.ctfChallengeDetail.ID)
          this.destroyCTFContainerByUserData = {
            "DockerNodeID": this.userCTFContainer.DockerNodeID,
            "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
          }
          this.delayCTFContainerByUserData = {
            "DockerNodeID": this.userCTFContainer.DockerNodeID,
            "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
          }
          this.checkFlagData = {
            "CTFChallengeID": this.ctfChallengeDetail.ID,
            "DockerNodeID": this.userCTFContainer.DockerNodeID,
            "DockerNodeContainerID": this.userCTFContainer.DockerNodeContainerID
          }
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetuserCTFChallengeList() {
      userCTFChallengeApi.GetUserCTFChallengeList({
        "ID": this.ctf.ID
      }).then(res => {
        if (res.code === 0) {
          this.userCTFChallengeList = res.data == null ? [] : res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ResetuserCTFChallenge() {
      userCTFChallengeApi.ResetUserCTFChallenge(this.resetUserCTFChallengeData).then(res => {
        if (res.code === 0) {
          this.GetuserCTFChallengeList()
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
    CheckCTFChallengeFlag() {
      userCTFChallengeApi.CheckCTFChallengeFlag(this.checkFlagData).then(res => {
        if (res.code === 0) {
          this.GetuserCTFChallengeList()
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
    DownloadAttachment(attachment) {
      try {
        new URL(attachment.FilePath)
        window.open(attachment.FilePath, '_blank')
      } catch (err) {
        attachmentApi.DownloadAttachment(attachment.ID).then(ok => {
          if (!ok) {
            ElMessage.error(this.t('CTFDetail.DownloadAttachmentFail'))
          }
        }).catch(error => {
          console.log(error)
        })
      }
    },
    checkContainerInUse(ctfChallengeID) {
      return this.userCTFContainerList.some(container => container.CTFChallengeID === ctfChallengeID)
    },
    getContainerInUse(ctfChallengeID) {
      this.userCTFContainer = this.userCTFContainerList.find(container => container.CTFChallengeID === ctfChallengeID)
    },
    getContainerChallenge(ctfChallengeID) {
      return this.totalCTFChallengeList.find(ctfChallenge => ctfChallenge.ID === ctfChallengeID)
    },
    checkChallengeSolved(ctfChallengeID) {
      return this.userCTFChallengeList.some(userCTFChallenge => userCTFChallenge.ID === ctfChallengeID)
    },
    getCategory(categoryName) {
      return this.categoryList.find(category => category.Name === categoryName)
    }
  },
  async mounted() {
    this.ctf.ID = Number(this.$route.params.id)
    this.GetCTF()
    await this.GetCTFContainerListByUser()
    await this.GetCategoryList()
    const category = this.getCategory(this.$route.params.category)
    this.ctfChallengeQueryCategoryID = category === undefined ? 0 : category.ID
    this.GetCTFChallengeList()
    this.GetuserCTFChallengeList()
  },
  async beforeRouteUpdate(to, from, next) {
    this.ctf.ID = Number(to.params.id)
    this.GetCTF()
    await this.GetCTFContainerListByUser()
    await this.GetCategoryList()
    const category = this.getCategory(to.params.category)
    this.ctfChallengeQueryCategoryID = category === undefined ? 0 : category.ID
    this.ctfChallengeQueryTitle = ""
    this.GetCTFChallengeList()
    this.GetuserCTFChallengeList()
    next()
  },
}
</script>

<style scoped>
.el-aside {
  overflow-x: hidden;
  transition: width var(--el-transition-duration);
  border-right: 1px solid var(--el-border-color);
}

.ctfChallenge {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
  cursor: pointer;
}
.ctfChallenge h2{
  margin: 20px;
}
.ctfChallenge .el-text{
  margin: 20px;
  width: calc(100% - 40px);
}
.operations{
  margin-top: 15px;
 }
.el-pagination{
  justify-content: center;
}
</style>
