<template>
  <el-container>
    <el-aside :width="isMenuOpen ? '150px' : '64px'">
      <el-menu :collapse="!isMenuOpen" style="border: none;">
        <el-menu-item index="1" @click="toggleSidebar">
          <el-icon>
            <Menu />
          </el-icon>
            <template #title>Menu</template>
        </el-menu-item>
        <el-menu-item index="2">
          <el-icon>
            <Aim />
          </el-icon>
          <template #title>Web</template>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-row>
          <el-col v-for="challenge in challengeList" :span=12>
            <div class="challenge" @click="OpenChallengeDetail(challenge)">
              <h2>{{ challenge.Title }}<el-text v-if="checkChallengeSolved(challenge.ID)">Solved</el-text></h2>
              <el-text>{{ challenge.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
        <el-pagination
          layout="prev, pager, next"
          :pager-count="11"
          :total="200"
          />
      </el-scrollbar>
      <el-dialog
        v-model="challengeDetailVisible"
        title="Challenge Detail"
        width="700"
        @close="ClearChallengeDetail"
        >
        <h2>{{ challengeDetail.Title }} <el-text v-if="checkChallengeSolved(challengeDetail.ID)">(Solved)</el-text></h2>
        <el-text>{{ challengeDetail.Description }}</el-text>
        <div style="margin-top: 5px;">
          <el-text v-if="challengeDetail.Attachment.ID !== 0">attachment: </el-text>
          <el-link
            @click="DownloadAttachment(challengeDetail.Attachment.ID)"
            style="vertical-align: unset"
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
            Reset
          </el-button>
          <el-button
            @click="CreateContainerByUser"
            v-if="!checkContainerInUse(challengeDetail.ID) && challengeDetail.ImageID !== 0"
            >
            Create
          </el-button>
          <el-button
            @click="DestroyContainerByUser"
            v-if="checkContainerInUse(challengeDetail.ID)"
            >
            Destroy
          </el-button>
          <el-button
            @click="DelayContainerByUser"
            v-if="checkContainerInUse(challengeDetail.ID)"
            >
            Delay
          </el-button>
        </div>
        <div style="display:flex; margin-top: 15px" v-if="checkContainerInUse(challengeDetail.ID) || challengeDetail.ImageID === 0">
          <el-input v-model="checkFlagData.Flag"/>
          <el-button @click="CheckFlag">Submit</el-button>
        </div>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import {Aim, Menu} from '@element-plus/icons-vue'
import challengeApi from "@/api/challenge.js"
import userContainerApi from "@/api/userContainer.js";
import userChallengeApi from "@/api/userChallenge.js";
import { ElMessage } from "element-plus";
import attachmentApi from "@/api/attachment.js";

export default {
  components: {Aim, Menu},
  data() {
    return {
      isMenuOpen: false,
      userContainerList: [],
      userContainer: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "ChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      },
      challengeList: [],
      challengeDetailVisible: false,
      challengeDetail: {
        "ID": 0,
        "Title": "",
        "Description": "",
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
      challengeApi.GetChallengeList().then(res => {
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
    CreateContainerByUser() {
      userContainerApi.CreateContainerByUser(this.createContainerByUserData).then(async res => {
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
    DestroyContainerByUser() {
      userContainerApi.DestroyContainerByUser(this.destroyContainerByUserData).then(async res => {
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
    DelayContainerByUser() {
      userContainerApi.DelayContainerByUser(this.delayContainerByUserData).then(async res => {
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
    CheckFlag() {
      userChallengeApi.CheckFlag(this.checkFlagData).then(res => {
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
    DownloadAttachment(id) {
      attachmentApi.DownloadAttachment(id).then(ok => {
        if (!ok) {
          ElMessage.error("Download attachment fail")
        }
      }).catch(error => {
        console.log(error)
      })
    },
    checkContainerInUse(challengeID) {
      return this.userContainerList.some(container => container.ChallengeID === challengeID)
    },
    getContainerInUse(challengeID) {
      this.userContainer = this.userContainerList.find(container => container.ChallengeID === challengeID)
    },
    checkChallengeSolved(challengeID) {
      return this.userChallengeList.some(userChallenge => userChallenge.ID === challengeID)
    },
  },
  mounted() {
    this.GetContainerListByUser()
    this.GetChallengeList()
    this.GetUserChallengeList()
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
}
.operations{
  margin-top: 15px;
 }
.el-pagination{
  justify-content: center;
}
</style>
