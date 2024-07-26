<template>
  <el-container>
    <el-aside :width="isMenuOpen ? '190px' : '64px'">
      <el-scrollbar>
        <el-menu :collapse="!isMenuOpen" style="border: none;" router>
          <el-menu-item @click="toggleSidebar">
            <el-icon>
              <Menu />
            </el-icon>
            <template #title>{{ $t('GroupDetail.Menu') }}</template>
          </el-menu-item>
          <el-sub-menu :index="'/group/' + group.ID">
            <template #title>
              <el-icon>
                <Category/>
              </el-icon>
              <el-text style="color: var(--el-text-color-primary);">{{ $t('GroupDetail.Category') }}</el-text>
            </template>
            <el-menu-item :index="'/group/' + group.ID">
              <template #title>{{ $t('GroupDetail.All') }}</template>
            </el-menu-item>
            <el-menu-item
              v-for="category in categoryList"
              :key="category.ID"
              :index="'/group/' + group.ID + '/' + category.Name"
            >
              <template #title>{{ category.Name }}</template>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item @click="userGroupChallengeVisible = true">
            <el-icon>
              <Finish/>
            </el-icon>
            <template #title>{{ $t('GroupDetail.Solved') }}</template>
          </el-menu-item>
          <el-menu-item @click="userGroupContainerListVisible = true">
            <el-icon>
              <Container/>
            </el-icon>
            <template #title>{{ $t('GroupDetail.Containers') }}</template>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-affix :offset="100">
          <div style="display: flex; margin: 0 20px 20px 20px; align-items: center">
            <el-breadcrumb separator="/" style="width: 400px; font-size: 20px;">
              <el-breadcrumb-item to="/group">{{ $t('GroupDetail.Group') }}</el-breadcrumb-item>
              <el-breadcrumb-item :to="'/group/' + group.ID">{{ group.Name }}</el-breadcrumb-item>
              <el-breadcrumb-item
                v-if="this.$route.params.category !== ''"
                :to="'/group/' + group.ID + '/' + this.$route.params.category"
                >
                {{ this.$route.params.category }}
              </el-breadcrumb-item>
            </el-breadcrumb>
            <el-input v-model="groupChallengeQueryTitle" :placeholder="$t('GroupDetail.FindChallenges')" size="large"/>
            <el-button @click="GetGroupChallengeList" type="primary" style="margin-left: 15px;" size="large">
              {{ $t('GroupDetail.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
          </div>
        </el-affix>
        <el-row>
          <el-col v-for="groupChallenge in groupChallengeList" :span="12">
            <div class="groupChallenge" @click="OpenGroupChallengeDetail(groupChallenge)">
              <h2 style="color: var(--el-text-color-primary)">{{ groupChallenge.Title }}<el-text v-if="checkChallengeSolved(groupChallenge.ID)">{{ $t('GroupDetail.Solved') }}</el-text></h2>
              <el-text truncated>{{ groupChallenge.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
      <el-dialog
        v-model="groupChallengeDetailVisible"
        :title="$t('GroupDetail.ChallengeDetail')"
        width="600"
        @close="ClearGroupChallengeDetail"
        >
        <h2 style="color: var(--el-text-color-primary)">{{ groupChallengeDetail.Title }} <el-text v-if="checkChallengeSolved(groupChallengeDetail.ID)">({{ $t('GroupDetail.Solved') }})</el-text></h2>
        <el-text>{{ groupChallengeDetail.Description }}</el-text>
        <div style="margin-top: 5px;">
          <el-text>{{ $t('GroupDetail.Category') }}: </el-text>
          <el-text>{{ groupChallengeDetail.CategoryName === "" ? $t('GroupDetail.Uncategorized') : groupChallengeDetail.CategoryName }}</el-text>
        </div>
        <div v-if="groupChallengeDetail.Attachment.ID !== 0" style="margin-top: 5px;">
          <el-text>{{ $t('GroupDetail.Attachment') }}: </el-text>
          <el-link
            @click="DownloadAttachment(groupChallengeDetail.Attachment)"
            style="vertical-align: unset"
            :underline="false"
          >
            {{ groupChallengeDetail.Attachment.FileName }}
          </el-link>
        </div>
        <div
          v-if="checkContainerInUse(groupChallengeDetail.ID)"
          v-for="portMap in userGroupContainer.PortMaps"
          style="margin-top: 15px">
          <el-text>
            {{ portMap.PortProtocol }} -> {{ portMap.Link }}
          </el-text>
        </div>
        <el-progress
          v-if="checkContainerInUse(groupChallengeDetail.ID)"
          :percentage="(userGroupContainer.RemainingTime - Date.now()) / 36000"
          style="margin-top: 15px"
          >
          <el-countdown
            :value="userGroupContainer.RemainingTime"
            :finish="this.GetGroupContainerListByUser"
            value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
          />
        </el-progress>
        <div class="operations">
          <el-button
            @click="ResetuserGroupChallenge(groupChallengeDetail.ID)"
            >
            {{ $t('GroupDetail.Reset') }}
          </el-button>
          <el-button
            @click="CreateContainerByUser(createGroupContainerByUserData)"
            v-if="!checkContainerInUse(groupChallengeDetail.ID) && groupChallengeDetail.ImageID !== 0"
            >
            {{ $t('GroupDetail.Create') }}
          </el-button>
          <el-button
            @click="DestroyContainerByUser(destroyGroupContainerByUserData)"
            v-if="checkContainerInUse(groupChallengeDetail.ID)"
            >
            {{ $t('GroupDetail.Destroy') }}
          </el-button>
          <el-button
            @click="DelayContainerByUser(delayGroupContainerByUserData)"
            v-if="checkContainerInUse(groupChallengeDetail.ID)"
            >
            {{ $t('GroupDetail.Delay') }}
          </el-button>
        </div>
        <div style="display:flex; margin-top: 15px" v-if="checkContainerInUse(groupChallengeDetail.ID) || groupChallengeDetail.ImageID === 0">
          <el-input v-model="checkFlagData.Flag" :placeholder="$t('GroupDetail.InputFlag')"/>
          <el-button @click="CheckGroupChallengeFlag">{{ $t('GroupDetail.Submit') }}</el-button>
        </div>
      </el-dialog>
      <el-dialog
        v-model="userGroupContainerListVisible"
        :title="$t('GroupDetail.ContainerList')"
        width="600"
        >
        <el-text v-if="userGroupContainerList.length === 0" size="large">{{ $t('GroupDetail.NoContainerIsRunning') }}</el-text>
        <el-card
          v-for="userGroupContainer in userGroupContainerList"
          style="margin: 20px 0 20px 0;"
          >
          <h2 style="color: var(--el-text-color-primary)">
          {{ getContainerChallenge(userGroupContainer.GroupChallengeID).Group.Name }}
          /
          {{ getContainerChallenge(userGroupContainer.GroupChallengeID) === undefined ? "" : getContainerChallenge(userGroupContainer.GroupChallengeID).Title }}
          </h2>
          <div
            v-for="portMap in userGroupContainer.PortMaps"
            style="margin-top: 15px">
            <el-text>
              {{ portMap.PortProtocol }} -> {{ portMap.Link }}
            </el-text>
          </div>
          <el-progress
            :percentage="(userGroupContainer.RemainingTime - Date.now()) / 36000"
            style="margin-top: 15px"
            >
            <el-countdown
              :value="userGroupContainer.RemainingTime"
              :finish="this.GetGroupContainerListByUser"
              value-style="font-size: var(--el-font-size-base); color: var(--el-text-color-regular); line-height: 2;"
            />
          </el-progress>
          <div class="operations">
            <el-button
              @click="DestroyContainerByUser(userGroupContainer)"
              >
              {{ $t('GroupDetail.Destroy') }}
            </el-button>
            <el-button
              @click="DelayContainerByUser(userGroupContainer)"
              >
              {{ $t('GroupDetail.Delay') }}
            </el-button>
          </div>
        </el-card>
      </el-dialog>
      <el-dialog
        v-model="userGroupChallengeVisible"
        :title="$t('GroupDetail.SolvedChallenges')"
        width="70%"
      >
        <el-scrollbar style="height: 60vh;">
          <el-row>
            <el-col v-for="groupChallenge in userGroupChallengeList" :span="12">
              <div class="groupChallenge" @click="OpenGroupChallengeDetail(groupChallenge)">
                <h2 style="color: var(--el-text-color-primary)">{{ groupChallenge.Title }}<el-text v-if="checkChallengeSolved(groupChallenge.ID)">{{ $t('GroupDetail.Solved') }}</el-text></h2>
                <el-text truncated>{{ groupChallenge.Description }}</el-text>
              </div>
            </el-col>
          </el-row>
        </el-scrollbar>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import groupApi from "@/api/group.js"
import categoryApi from "@/api/category.js"
import groupChallengeApi from "@/api/groupChallenge.js"
import userGroupContainerApi from "@/api/userGroupContainer.js"
import userGroupChallengeApi from "@/api/userGroupChallenge.js"
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
      group: {
        "ID": 0,
      },
      userGroupContainerList: [],
      userGroupContainerListVisible: false,
      userGroupContainer: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "GroupChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      },
      categoryList: [],
      totalGroupChallengeList: [],
      groupChallengeQueryCategoryID: 0,
      groupChallengeQueryTitle: "",
      groupChallengeList: [],
      groupChallengeDetailVisible: false,
      groupChallengeDetail: {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryName": "",
        "ImageID": 0,
        "Attachment": {}
      },
      createGroupContainerByUserData: {
        "ID": 0,
      },
      destroyGroupContainerByUserData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      },
      delayGroupContainerByUserData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      },
      userGroupChallengeVisible: false,
      userGroupChallengeList: [],
      resetUserGroupChallengeData: {
        "GroupChallengeID": 0
      },
      checkFlagData: {
        "GroupChallengeID": 0,
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
    GetGroup() {
      groupApi.GetGroup(this.group).then(res => {
        if (res.code === 0) {
          this.group = res.data
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
    async GetGroupContainerListByUser() {
      return userGroupContainerApi.GetGroupContainerListByUser().then(res => {
        if (res.code === 0) {
          this.userGroupContainerList = res.data == null ? [] : res.data
          this.userGroupContainerList = this.userGroupContainerList.map(container => {
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
    GetGroupChallengeList() {
      groupChallengeApi.GetGroupChallengeList({}).then(res => {
        if (res.code === 0) {
          this.totalGroupChallengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
      groupChallengeApi.GetGroupChallengeList({
        "GroupID": this.group.ID,
        "CategoryID": this.groupChallengeQueryCategoryID,
        "Title": this.groupChallengeQueryTitle
      }).then(res => {
        if (res.code === 0) {
          this.groupChallengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenGroupChallengeDetail(groupChallenge) {
      this.groupChallengeDetailVisible = true
      this.groupChallengeDetail = {
        "ID": groupChallenge.ID,
        "Title": groupChallenge.Title,
        "Description": groupChallenge.Description,
        "CategoryName": groupChallenge.Category.Name,
        "ImageID": groupChallenge.ImageID,
        "Attachment": groupChallenge.Attachment
      }
      this.createGroupContainerByUserData = {
        "ID": groupChallenge.ID,
      }
      this.checkFlagData = {
        "GroupChallengeID": groupChallenge.ID,
      }
      if (this.checkContainerInUse(groupChallenge.ID)) {
        this.getContainerInUse(groupChallenge.ID)
        this.destroyGroupContainerByUserData = {
          "DockerNodeID": this.userGroupContainer.DockerNodeID,
          "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
        }
        this.delayGroupContainerByUserData = {
          "DockerNodeID": this.userGroupContainer.DockerNodeID,
          "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
        }
        this.checkFlagData = {
          "GroupChallengeID": this.groupChallengeDetail.ID,
          "DockerNodeID": this.userGroupContainer.DockerNodeID,
          "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
        }
      }
      this.resetUserGroupChallengeData = {
        "GroupChallengeID": groupChallenge.ID
      }
    },
    ClearGroupChallengeDetail() {
      this.groupChallengeDetail = {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryName": "",
        "ImageID": 0,
        "Attachment": {}
      }
      this.createGroupContainerByUserData = {
        "ID": 0,
      }
      this.destroyGroupContainerByUserData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      }
      this.delayGroupContainerByUserData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": ""
      }
      this.userGroupContainer = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "GroupChallengeID": 0,
        "RemainingTime": 0,
        "PortMaps": []
      }
      this.resetUserGroupChallengeData = {
        "GroupChallengeID": 0
      }
      this.checkFlagData = {
        "GroupChallengeID": 0,
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
        "Flag": "",
      }
    },
    CreateContainerByUser(createGroupContainerByUserData) {
      userGroupContainerApi.CreateGroupContainerByUser(createGroupContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetGroupContainerListByUser()
        if (this.checkContainerInUse(this.groupChallengeDetail.ID)) {
          this.getContainerInUse(this.groupChallengeDetail.ID)
          this.destroyGroupContainerByUserData = {
            "DockerNodeID": this.userGroupContainer.DockerNodeID,
            "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
          }
          this.delayGroupContainerByUserData = {
            "DockerNodeID": this.userGroupContainer.DockerNodeID,
            "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
          }
          this.checkFlagData = {
            "GroupChallengeID": this.groupChallengeDetail.ID,
            "DockerNodeID": this.userGroupContainer.DockerNodeID,
            "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
          }
        }
      }).catch(error => {
        console.log(error)
      })
    },
    DestroyContainerByUser(destroyGroupContainerByUserData) {
      userGroupContainerApi.DestroyGroupContainerByUser(destroyGroupContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetGroupContainerListByUser()
        this.destroyGroupContainerByUserData = {
          "DockerNodeID": 0,
          "DockerNodeContainerID": ""
        }
        this.delayGroupContainerByUserData = {
          "DockerNodeID": 0,
          "DockerNodeContainerID": ""
        }
        this.userGroupContainer = {
          "DockerNodeContainerID": "",
          "DockerNodeID": 0,
          "GroupChallengeID": 0,
          "RemainingTime": 0,
          "PortMaps": []
        }
        this.checkFlagData = {
          "GroupChallengeID": 0,
          "DockerNodeID": 0,
          "DockerNodeContainerID": "",
          "Flag": "",
        }
      }).catch(error => {
        console.log(error)
      })
    },
    DelayContainerByUser(delayGroupContainerByUserData) {
      userGroupContainerApi.DelayGroupContainerByUser(delayGroupContainerByUserData).then(async res => {
        if (res.code === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
          })
        } else {
          ElMessage.error(res.msg)
        }
        await this.GetGroupContainerListByUser()
        if (this.checkContainerInUse(this.groupChallengeDetail.ID)) {
          this.getContainerInUse(this.groupChallengeDetail.ID)
          this.destroyGroupContainerByUserData = {
            "DockerNodeID": this.userGroupContainer.DockerNodeID,
            "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
          }
          this.delayGroupContainerByUserData = {
            "DockerNodeID": this.userGroupContainer.DockerNodeID,
            "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
          }
          this.checkFlagData = {
            "GroupChallengeID": this.groupChallengeDetail.ID,
            "DockerNodeID": this.userGroupContainer.DockerNodeID,
            "DockerNodeContainerID": this.userGroupContainer.DockerNodeContainerID
          }
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetuserGroupChallengeList() {
      userGroupChallengeApi.GetUserGroupChallengeList().then(res => {
        if (res.code === 0) {
          this.userGroupChallengeList = res.data == null ? [] : res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ResetuserGroupChallenge() {
      userGroupChallengeApi.ResetUserGroupChallenge(this.resetUserGroupChallengeData).then(res => {
        if (res.code === 0) {
          this.GetuserGroupChallengeList()
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
    CheckGroupChallengeFlag() {
      userGroupChallengeApi.CheckGroupChallengeFlag(this.checkFlagData).then(res => {
        if (res.code === 0) {
          this.GetuserGroupChallengeList()
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
            ElMessage.error(this.t('GroupDetail.DownloadAttachmentFail'))
          }
        }).catch(error => {
          console.log(error)
        })
      }
    },
    checkContainerInUse(groupChallengeID) {
      return this.userGroupContainerList.some(container => container.GroupChallengeID === groupChallengeID)
    },
    getContainerInUse(groupChallengeID) {
      this.userGroupContainer = this.userGroupContainerList.find(container => container.GroupChallengeID === groupChallengeID)
    },
    getContainerChallenge(groupChallengeID) {
      return this.totalGroupChallengeList.find(groupChallenge => groupChallenge.ID === groupChallengeID)
    },
    checkChallengeSolved(groupChallengeID) {
      return this.userGroupChallengeList.some(userGroupChallenge => userGroupChallenge.ID === groupChallengeID)
    },
    getCategory(categoryName) {
      return this.categoryList.find(category => category.Name === categoryName)
    }
  },
  async mounted() {
    this.group.ID = Number(this.$route.params.id)
    this.GetGroup()
    await this.GetGroupContainerListByUser()
    await this.GetCategoryList()
    const category = this.getCategory(this.$route.params.category)
    this.groupChallengeQueryCategoryID = category === undefined ? 0 : category.ID
    this.GetGroupChallengeList()
    this.GetuserGroupChallengeList()
  },
  async beforeRouteUpdate(to, from, next) {
    this.group.ID = Number(to.params.id)
    this.GetGroup()
    await this.GetGroupContainerListByUser()
    await this.GetCategoryList()
    const category = this.getCategory(to.params.category)
    this.groupChallengeQueryCategoryID = category === undefined ? 0 : category.ID
    this.groupChallengeQueryTitle = ""
    this.GetGroupChallengeList()
    this.GetuserGroupChallengeList()
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

.groupChallenge {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
  cursor: pointer;
}
.groupChallenge h2{
  margin: 20px;
}
.groupChallenge .el-text{
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
