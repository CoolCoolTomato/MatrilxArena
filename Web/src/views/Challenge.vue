<template>
  <el-container>
    <el-aside :width="isMenuOpen ? '190px' : '64px'">
      <el-scrollbar>
        <el-menu :collapse="!isMenuOpen" style="border: none;" router>
          <el-menu-item @click="toggleSidebar">
            <el-icon>
              <Menu />
            </el-icon>
            <template #title>{{ $t('Challenge.Menu') }}</template>
          </el-menu-item>
          <el-sub-menu index="challenge">
            <template #title>
              <el-icon>
                <Category/>
              </el-icon>
              <el-text style="color: var(--el-text-color-primary);">{{ $t('Challenge.Category') }}</el-text>
            </template>
            <el-menu-item index="/challenge">
              <template #title>{{ $t('Challenge.All') }}</template>
            </el-menu-item>
            <el-menu-item
              v-for="category in categoryList"
              :key="category.ID"
              :index="'/challenge/' + category.Name"
            >
              <template #title>{{ category.Name }}</template>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item @click="userChallengeVisible = true">
            <el-icon>
              <Finish/>
            </el-icon>
            <template #title>{{ $t('Challenge.Solved') }}</template>
          </el-menu-item>
          <el-menu-item @click="userContainerListVisible = true">
            <el-icon>
              <Container/>
            </el-icon>
            <template #title>{{ $t('Challenge.Containers') }}</template>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-affix :offset="100">
          <div style="display: flex; margin: 0 20px 20px 20px;">
            <el-input v-model="challengeQueryTitle" :placeholder="$t('Challenge.FindChallenges')" size="large"/>
            <el-button @click="GetChallengeList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('Challenge.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
          </div>
        </el-affix>
        <el-row>
          <el-col v-for="challenge in challengeList" :span="12">
            <div class="challenge" @click="OpenChallengeDetail(challenge)">
              <h2 style="color: var(--el-text-color-primary)">{{ challenge.Title }}<el-text v-if="checkChallengeSolved(challenge.ID)">{{ $t('Challenge.Solved') }}</el-text></h2>
              <el-text truncated>{{ challenge.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
      <el-dialog
        v-model="challengeDetailVisible"
        :title="$t('Challenge.ChallengeDetail')"
        width="600"
        @close="ClearChallengeDetail"
        >
        <h2 style="color: var(--el-text-color-primary)">{{ challengeDetail.Title }} <el-text v-if="checkChallengeSolved(challengeDetail.ID)">({{ $t('Challenge.Solved') }})</el-text></h2>
        <el-text>{{ challengeDetail.Description }}</el-text>
        <div style="margin-top: 5px;">
          <el-text>{{ $t('Challenge.Category') }}: </el-text>
          <el-text>{{ challengeDetail.CategoryName === "" ? $t('Challenge.Uncategorized') : challengeDetail.CategoryName }}</el-text>
        </div>
        <div v-if="challengeDetail.Attachment.ID !== 0" style="margin-top: 5px;">
          <el-text>{{ $t('Challenge.Attachment') }}: </el-text>
          <el-link
            @click="DownloadAttachment(challengeDetail.Attachment)"
            style="vertical-align: unset"
            :underline="false"
          >
            {{ challengeDetail.Attachment.FileName }}
          </el-link>
        </div>
        <div
          v-if="checkContainerInUse(challengeDetail.ID)"
          v-for="portMap in userContainer.PortMaps"
          style="margin-top: 15px">
          <el-text>
            {{ portMap.PortProtocol }} -> {{ portMap.Link }}
          </el-text>
        </div>
        <el-progress
          v-if="checkContainerInUse(challengeDetail.ID)"
          :percentage="(userContainer.RemainingTime - Date.now()) / 36000"
          style="margin-top: 15px"
          >
          <el-countdown
            :value="userContainer.RemainingTime"
            :finish="this.GetContainerListByUser"
            value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
          />
        </el-progress>
        <div class="operations">
          <el-button
            @click="ResetUserChallenge(challengeDetail.ID)"
            >
            {{ $t('Challenge.Reset') }}
          </el-button>
          <el-button
            @click="CreateContainerByUser(createContainerByUserData)"
            v-if="!checkContainerInUse(challengeDetail.ID) && challengeDetail.ImageID !== 0"
            >
            {{ $t('Challenge.Create') }}
          </el-button>
          <el-button
            @click="DestroyContainerByUser(destroyContainerByUserData)"
            v-if="checkContainerInUse(challengeDetail.ID)"
            >
            {{ $t('Challenge.Destroy') }}
          </el-button>
          <el-button
            @click="DelayContainerByUser(delayContainerByUserData)"
            v-if="checkContainerInUse(challengeDetail.ID)"
            >
            {{ $t('Challenge.Delay') }}
          </el-button>
        </div>
        <div style="display:flex; margin-top: 15px" v-if="checkContainerInUse(challengeDetail.ID) || challengeDetail.ImageID === 0">
          <el-input v-model="checkFlagData.Flag" :placeholder="$t('Challenge.InputFlag')"/>
          <el-button @click="CheckChallengeFlag">{{ $t('Challenge.Submit') }}</el-button>
        </div>
      </el-dialog>
      <el-dialog
        v-model="userContainerListVisible"
        :title="$t('Challenge.ContainerList')"
        width="600"
        >
        <el-text v-if="userContainerList.length === 0" size="large">{{ $t('Challenge.NoContainerIsRunning') }}</el-text>
        <el-card
          v-for="userContainer in userContainerList"
          style="margin: 20px 0 20px 0;"
          >
          <h2 style="color: var(--el-text-color-primary)">{{ getContainerChallenge(userContainer.ChallengeID) === undefined ? "" : getContainerChallenge(userContainer.ChallengeID).Title }}</h2>
          <div
            v-for="portMap in userContainer.PortMaps"
            style="margin-top: 15px">
            <el-text>
              {{ portMap.PortProtocol }} -> {{ portMap.Link }}
            </el-text>
          </div>
          <el-progress
            :percentage="(userContainer.RemainingTime - Date.now()) / 36000"
            style="margin-top: 15px"
            >
            <el-countdown
              :value="userContainer.RemainingTime"
              :finish="this.GetContainerListByUser"
              value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
            />
          </el-progress>
          <div class="operations">
            <el-button
              @click="DestroyContainerByUser(userContainer)"
              >
              {{ $t('Challenge.Destroy') }}
            </el-button>
            <el-button
              @click="DelayContainerByUser(userContainer)"
              >
              {{ $t('Challenge.Delay') }}
            </el-button>
          </div>
        </el-card>
      </el-dialog>
      <el-dialog
        v-model="userChallengeVisible"
        :title="$t('Challenge.SolvedChallenges')"
        width="70%"
      >
        <el-scrollbar style="height: 60vh;">
          <el-row>
            <el-col v-for="challenge in userChallengeList" :span="12">
              <div class="challenge" @click="OpenChallengeDetail(challenge)">
                <h2 style="color: var(--el-text-color-primary)">{{ challenge.Title }}<el-text v-if="checkChallengeSolved(challenge.ID)">{{ $t('Challenge.Solved') }}</el-text></h2>
                <el-text truncated>{{ challenge.Description }}</el-text>
              </div>
            </el-col>
          </el-row>
        </el-scrollbar>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import categoryApi from "@/api/category.js"
import challengeApi from "@/api/challenge.js"
import userContainerApi from "@/api/userContainer.js"
import userChallengeApi from "@/api/userChallenge.js"
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
      userContainerList: [],
      userContainerListVisible: false,
      userContainer: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "ChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      },
      categoryList: [],
      totalChallengeList: [],
      challengeQueryCategoryID: 0,
      challengeQueryTitle: "",
      challengeList: [],
      challengeDetailVisible: false,
      challengeDetail: {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryName": "",
        "ImageID": 0,
        "Attachment": {}
      },
      createContainerByUserData: {
        "ID": 0,
      },
      destroyContainerByUserData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      },
      delayContainerByUserData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      },
      userChallengeVisible: false,
      userChallengeList: [],
      resetUserChallengeData: {
        "ChallengeID": 0
      },
      checkFlagData: {
        "ChallengeID": 0,
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
    async GetContainerListByUser() {
      return userContainerApi.GetContainerListByUser().then(res => {
        if (res.code === 0) {
          this.userContainerList = res.data == null ? [] : res.data
          this.userContainerList = this.userContainerList.map(container => {
            const newContainer = {...container}
            newContainer.RemainingTime = Math.trunc(newContainer.RemainingTime / 1000000 + Date.now())
            return newContainer
          })
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetChallengeList() {
      challengeApi.GetChallengeList({}).then(res => {
        if (res.code === 0) {
          this.totalChallengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
      challengeApi.GetChallengeList({
        "CategoryID": this.challengeQueryCategoryID,
        "Title": this.challengeQueryTitle
      }).then(res => {
        if (res.code === 0) {
          this.challengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenChallengeDetail(challenge) {
      this.challengeDetailVisible = true
      this.challengeDetail = {
        "ID": challenge.ID,
        "Title": challenge.Title,
        "Description": challenge.Description,
        "CategoryName": challenge.Category.Name,
        "ImageID": challenge.ImageID,
        "Attachment": challenge.Attachment
      }
      this.createContainerByUserData = {
        "ID": challenge.ID,
      }
      this.checkFlagData = {
        "ChallengeID": challenge.ID,
      }
      if (this.checkContainerInUse(challenge.ID)) {
        this.getContainerInUse(challenge.ID)
        this.destroyContainerByUserData = {
          "DockerNodeID": this.userContainer.DockerNodeID,
          "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
        }
        this.delayContainerByUserData = {
          "DockerNodeID": this.userContainer.DockerNodeID,
          "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
        }
        this.checkFlagData = {
          "ChallengeID": this.challengeDetail.ID,
          "DockerNodeID": this.userContainer.DockerNodeID,
          "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
        }
      }
      this.resetUserChallengeData = {
        "ChallengeID": challenge.ID
      }
    },
    ClearChallengeDetail() {
      this.challengeDetail = {
        "ID": 0,
        "Title": "",
        "Description": "",
        "ImageID": 0,
        "Attachment": {}
      }
      this.createContainerByUserData = {
        "ID": 0,
      }
      this.destroyContainerByUserData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      }
      this.delayContainerByUserData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      }
      this.userContainer = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "ChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      }
      this.resetUserChallengeData = {
        "ChallengeID": 0
      }
      this.checkFlagData = {
        "ChallengeID": 0,
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "Flag": "",
      }
    },
    CreateContainerByUser(createContainerByUserData) {
      userContainerApi.CreateContainerByUser(createContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetContainerListByUser()
        if (this.checkContainerInUse(this.challengeDetail.ID)) {
          this.getContainerInUse(this.challengeDetail.ID)
          this.destroyContainerByUserData = {
            "DockerNodeID": this.userContainer.DockerNodeID,
            "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
          }
          this.delayContainerByUserData = {
            "DockerNodeID": this.userContainer.DockerNodeID,
            "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
          }
          this.checkFlagData = {
            "ChallengeID": this.challengeDetail.ID,
            "DockerNodeID": this.userContainer.DockerNodeID,
            "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
          }
        }
      }).catch(error => {
        console.log(error)
      })
    },
    DestroyContainerByUser(destroyContainerByUserData) {
      userContainerApi.DestroyContainerByUser(destroyContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetContainerListByUser()
        this.destroyContainerByUserData = {
          "DockerNodeID": 0,
          "DockerNodeContainerID": ""
        }
        this.delayContainerByUserData = {
          "DockerNodeID": 0,
          "DockerNodeContainerID": ""
        }
        this.userContainer = {
          "DockerNodeContainerID": "",
          "DockerNodeID": 0,
          "ChallengeID": 0,
          "RemainingTime": 0,
          "PortMaps": []
        }
        this.checkFlagData = {
          "ChallengeID": 0,
          "DockerNodeID": 0,
          "DockerNodeContainerID": "",
          "Flag": "",
        }
      }).catch(error => {
        console.log(error)
      })
    },
    DelayContainerByUser(delayContainerByUserData) {
      userContainerApi.DelayContainerByUser(delayContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetContainerListByUser()
        if (this.checkContainerInUse(this.challengeDetail.ID)) {
          this.getContainerInUse(this.challengeDetail.ID)
          this.destroyContainerByUserData = {
            "DockerNodeID": this.userContainer.DockerNodeID,
            "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
          }
          this.delayContainerByUserData = {
            "DockerNodeID": this.userContainer.DockerNodeID,
            "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
          }
          this.checkFlagData = {
            "ChallengeID": this.challengeDetail.ID,
            "DockerNodeID": this.userContainer.DockerNodeID,
            "DockerNodeContainerID": this.userContainer.DockerNodeContainerID
          }
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetUserChallengeList() {
      userChallengeApi.GetUserChallengeList().then(res => {
        if (res.code === 0) {
          this.userChallengeList = res.data == null ? [] : res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ResetUserChallenge() {
      userChallengeApi.ResetUserChallenge(this.resetUserChallengeData).then(res => {
        if (res.code === 0) {
          this.GetUserChallengeList()
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
    CheckChallengeFlag() {
      userChallengeApi.CheckChallengeFlag(this.checkFlagData).then(res => {
        if (res.code === 0) {
          this.GetUserChallengeList()
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
        attachmentApi.DownloadAttachment(attachment).then(ok => {
          if (!ok) {
            ElMessage.error(this.t('Challenge.DownloadAttachmentFail'))
          }
        }).catch(error => {
          console.log(error)
        })
      }
    },
    checkContainerInUse(challengeID) {
      return this.userContainerList.some(container => container.ChallengeID === challengeID)
    },
    getContainerInUse(challengeID) {
      this.userContainer = this.userContainerList.find(container => container.ChallengeID === challengeID)
    },
    getContainerChallenge(challengeID) {
      return this.totalChallengeList.find(challenge => challenge.ID === challengeID)
    },
    checkChallengeSolved(challengeID) {
      return this.userChallengeList.some(userChallenge => userChallenge.ID === challengeID)
    },
  },
  async mounted() {
    await this.GetContainerListByUser()
    await this.GetCategoryList()
    const category = this.categoryList.find(category => category.Name === this.$route.params.category)
    this.challengeQueryCategoryID = category === undefined ? 0 : category.ID
    this.GetChallengeList()
    this.GetUserChallengeList()
  },
  async beforeRouteUpdate(to, from, next) {
    await this.GetContainerListByUser()
    await this.GetCategoryList()
    const category = this.categoryList.find(category => category.Name === to.params.category)
    this.challengeQueryCategoryID = category === undefined ? 0 : category.ID
    this.challengeQueryTitle = ""
    this.GetChallengeList()
    this.GetUserChallengeList()
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

.challenge {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
  cursor: pointer;
}
.challenge h2{
  margin: 20px;
}
.challenge .el-text{
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
