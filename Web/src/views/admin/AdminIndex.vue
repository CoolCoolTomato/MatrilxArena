<template>
  <el-container>
    <el-main>
      <el-scrollbar>
        <el-row>
          <el-col :span="24">
            <el-card style="margin: 0 20px 20px 20px">
              <el-row>
                <el-col :span="6">
                  <el-statistic :title="$t('AdminIndex.NumberOfChallenges')" :value="challengeList.length" />
                </el-col>
                <el-col :span="6">
                  <el-statistic :title="$t('AdminIndex.NumberOfUsers')" :value="userList.length" />
                </el-col>
                <el-col :span="6">
                  <el-statistic :title="$t('AdminIndex.LiveDockerNode')" :value="dockerNodeList.length" />
                </el-col>
                <el-col :span="6">
                  <el-statistic :title="$t('AdminIndex.NumberOfImages')" :value="imageList.length" />
                </el-col>
              </el-row>
            </el-card>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-card>
              <el-text size="large">{{ $t('AdminIndex.ChallengesByCategory') }}</el-text>
              <div style="height: 350px;"></div>
            </el-card>
          </el-col>
          <el-col :span="10">
            <el-card>
              <el-text size="large">Hello World</el-text>
              <div style="height: 350px;"></div>
            </el-card>
          </el-col>
          <el-col :span="4">
            <el-card>
              <el-text size="large">Hello World</el-text>
              <div style="height: 350px;"></div>
            </el-card>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="18">
            <el-card style="height: calc(100% - 40px);"></el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <div style="display: flex; margin: 20px">
                <LanguageSelect />
              </div>
              <div style="display: flex; margin: 20px">
                <ThemeSelect />
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-scrollbar>
    </el-main>
  </el-container>
</template>

<script>
import dockerNodeApi from "@/api/dockerNode.js"
import imageApi from "@/api/image.js"
import categoryApi from "@/api/category.js"
import challengeApi from "@/api/challenge.js"
import userApi from "@/api/user.js"
import LanguageSelect from "@/components/LanguageSelect.vue"
import ThemeSelect from "@/components/ThemeSelect.vue"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: { LanguageSelect, ThemeSelect },
  data() {
    return {
      dockerNodeList: [],
      imageList: [],
      categoryList: [],
      challengeList: [],
      userList: [],
    }
  },
  methods: {
    async GetDockerNodeList() {
      return dockerNodeApi.GetDockerNodeList().then(res => {
        if (res.code === 0) {
          this.dockerNodeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    async GetImageList() {
      return imageApi.GetImageList().then(res => {
        if (res.code === 0) {
          this.imageList = res.data
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
    async GetChallengeList() {
      return challengeApi.GetChallengeList().then(res => {
        if (res.code === 0) {
          this.challengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    async GetUserList() {
      return userApi.GetUserList().then(res => {
        if (res.code === 0) {
          this.userList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
  },
  async mounted() {
    await this.GetDockerNodeList()
    await this.GetImageList()
    await this.GetCategoryList()
    await this.GetChallengeList()
    await this.GetUserList()
  },
}
</script>
<style scoped>
.el-card {
  margin: 20px;
}
.el-row {
  margin-right: 20px;
}
.el-col {
  text-align: center;
}
</style>
