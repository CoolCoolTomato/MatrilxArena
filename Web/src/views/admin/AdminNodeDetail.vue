<template>
  <el-container>
    <el-header style="display: flex">
      <h2 style="color: var(--el-text-color-primary);">
        {{ $t('AdminNodeDetail.NodeDetail') }} /
        {{ $t('AdminNodeDetail.Host') }}: {{dockerNode.Host}} |
        {{ $t('AdminNodeDetail.Port') }}: {{dockerNode.Port}}
      </h2>
      <div style="flex-grow: 1;" />
      <div style="display: flex; align-items: center">
        <el-button
          @click="pullImageFormVisible = true"
          type="primary"
          >
          {{ $t('AdminNodeDetail.PullImage') }}
        </el-button>
        <el-button
          @click="goBack"
          >
          {{ $t('AdminNodeDetail.Back') }}
        </el-button>
      </div>
    </el-header>
    <el-main>
      <el-tabs
        v-model="activeTab"
        class="el-tabs-custom"
        >
        <el-tab-pane :label="$t('AdminNodeDetail.Image')" name="image">
          <el-scrollbar>
            <el-table
              :data="dockerNodeImageList"
              table-layout="fixed"
              height="99%"
              >
              <el-table-column prop="RepoTags" :label="$t('AdminNodeDetail.RepoTags')"/>
              <el-table-column :label="$t('AdminNodeDetail.Size')">
                <template #default="scope">
                  {{ (scope.row.Size / 1000000).toFixed(2) }} MB
                </template>
              </el-table-column>
              <el-table-column :label="$t('AdminNodeDetail.Created')">
                <template #default="scope">
                  {{ formatDate(scope.row.Created) }}
                </template>
              </el-table-column>
              <el-table-column fixed="right" :label="$t('AdminNodeDetail.Operations')" width="200px">
                <template #default=scope>
                  <el-button
                    size="small"
                    @click="OpenImageDetail(scope.row)"
                    >
                    {{ $t('AdminNodeDetail.Detail') }}
                  </el-button>
                  <el-button
                    size="small"
                    @click="OpenRemoveImageForm(scope.row)"
                    >
                    {{ $t('AdminNodeDetail.Remove') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        <el-tab-pane :label="$t('AdminNodeDetail.Container')" name="container">
          <el-scrollbar>
            <el-table
              :data="dockerNodeContainerList"
              table-layout="fixed"
              height="99%"
              >
              <el-table-column prop="Names" :label="$t('AdminNodeDetail.Names')"/>
              <el-table-column prop="Image" :label="$t('AdminNodeDetail.Image')"/>
              <el-table-column prop="State" :label="$t('AdminNodeDetail.State')"/>
              <el-table-column fixed="right" :label="$t('AdminNodeDetail.Operations')" width="320px">
                <template #default=scope>
                  <el-button
                    size="small"
                    @click="OpenContainerDetail(scope.row)"
                    >
                    {{ $t('AdminNodeDetail.Detail') }}
                  </el-button>
                  <el-button
                    size="small"
                    @click="OpenStartContainerForm(scope.row)"
                    >
                    {{ $t('AdminNodeDetail.Start') }}
                  </el-button>
                  <el-button
                    size="small"
                    @click="OpenStopContainerForm(scope.row)"
                    >
                    {{ $t('AdminNodeDetail.Stop') }}
                  </el-button>
                  <el-button
                    size="small"
                    @click="OpenRemoveContainerForm(scope.row)"
                    >
                    {{ $t('AdminNodeDetail.Remove') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
      </el-tabs>
      <el-dialog
        v-model="pullImageFormVisible"
        :title="$t('AdminNodeDetail.PullImage')"
        width="1000px"
        style="max-height: 620px;"
        @close="ClearPullImageForm"
        >
        <el-row>
          <el-col :span="15">
            <el-table
              ref="pullImageForm"
              :data="imageList"
              max-height="500px"
              @selection-change="handleSelectionChange"
              >
              <el-table-column type="selection" width="50" />
              <el-table-column prop="Remark" :label="$t('AdminNodeDetail.Remark')"/>
              <el-table-column prop="RepoTags" :label="$t('AdminNodeDetail.RepoTags')"/>
              <el-table-column :label="$t('AdminNodeDetail.Repository')">
                <template #default="scope">
                    {{ formatRepository(scope.row.Repository) }}
                </template>
              </el-table-column>
              <el-table-column :label="$t('AdminNodeDetail.Existence')">
                <template #default="scope">
                  {{ checkImageExistence(scope.row.RepoTags) }}
                </template>
              </el-table-column>
            </el-table>
          </el-col>
          <el-col :span="9">
            <el-table
              :data="pullImageDataList"
              max-height="500px"
              >
              <el-table-column prop="RepoTags" :label="$t('AdminNodeDetail.SelectedImage')"/>
              <el-table-column prop="Status" :label="$t('AdminNodeDetail.Status')"/>
            </el-table>
          </el-col>
        </el-row>
        <template #footer>
          <el-button
            @click="PullImageFromDockerNode"
            type="primary"
            >
          {{ $t('AdminNodeDetail.Pull') }}
          </el-button>
          <el-button
            @click="pullImageFormVisible = false"
            >
          {{ $t('AdminNodeDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="imageDetailVisible"
        :title="$t('AdminNodeDetail.ImageDetail')"
        width="1000"
        @close="ClearImageDetail"
        >
        <el-card>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.ID') }}: {{ imageDetail.Id }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.RepoTags') }}: {{ imageDetail.RepoTags }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Size') }}: {{ (imageDetail.Size / 1000000).toFixed(2) }} MB</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Created') }}: {{ formatDate(imageDetail.Created) }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.RepoDigests') }}: {{ imageDetail.RepoDigests }}</p>
        </el-card>
        <template #footer>
          <el-button @click="imageDetailVisible = false">{{ $t('AdminNodeDetail.Close') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="removeImageFormVisible"
        :title="$t('AdminNodeDetail.RemoveImage')"
        width="500"
        @close="ClearRemoveImageForm"
      >
        <el-text>{{ $t('AdminNodeDetail.AreYouConfirmToRemoveTheImage') }}</el-text>
        <template #footer>
          <el-button
            @click="RemoveImageFromDockerNode"
            type="primary"
            >
            {{ $t('AdminNodeDetail.Confirm') }}
          </el-button>
          <el-button
            @click="removeImageFormVisible = false"
            >
            {{ $t('AdminNodeDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="containerDetailVisible"
        :title="$t('AdminNodeDetail.ContainerDetail')"
        width="1000"
        @close="ClearContainerDetail"
        >
        <el-card>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.ID') }}: {{ containerDetail.Id }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Names') }}: {{ containerDetail.Names }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Created') }}: {{ formatDate(containerDetail.Created) }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Image') }}: {{ containerDetail.Image }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.ImageID') }}: {{ containerDetail.ImageID }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Ports') }}: {{ containerDetail.Ports }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.State') }}: {{ containerDetail.State }}</p>
          <p style="word-break: break-all;">{{ $t('AdminNodeDetail.Status') }}: {{ containerDetail.Status }}</p>
        </el-card>
        <template #footer>
          <el-button @click="containerDetailVisible = false">{{ $t('AdminNodeDetail.Close') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="startContainerFormVisible"
        :title="$t('AdminNodeDetail.StartContainer')"
        width="500"
        @close="ClearStartContainerForm"
        >
        <el-text>{{ $t('AdminNodeDetail.AreYouConfirmToStartTheContainer') }}</el-text>
        <template #footer>
          <el-button
            @click="StartContainerFromDockerNode"
            type="primary"
            >
            {{ $t('AdminNodeDetail.Confirm') }}
          </el-button>
          <el-button
            @click="startContainerFormVisible = false"
            >
            {{ $t('AdminNodeDetail.Cancel') }}
          </el-button>
          </template>
      </el-dialog>
      <el-dialog
        v-model="stopContainerFormVisible"
        :title="$t('AdminNodeDetail.StopContainer')"
        width="500"
        @close="ClearStopContainerForm"
        >
        <el-text>{{ $t('AdminNodeDetail.AreYouConfirmToStopTheContainer') }}</el-text>
        <template #footer>
          <el-button
            @click="StopContainerFromDockerNode"
            type="primary"
            >
            {{ $t('AdminNodeDetail.Confirm') }}
          </el-button>
          <el-button
            @click="stopContainerFormVisible = false"
            >
            {{ $t('AdminNodeDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="removeContainerFormVisible"
        :title="$t('AdminNodeDetail.RemoveContainer')"
        width="500"
        @close="ClearRemoveContainerForm"
        >
        <el-text>{{ $t('AdminNodeDetail.AreYouConfirmToRemoveTheContainer') }}</el-text>
        <template #footer>
          <el-button
            @click="RemoveContainerFromDockerNode"
            type="primary"
            >
            {{ $t('AdminNodeDetail.Confirm') }}
          </el-button>
          <el-button
            @click="removeContainerFormVisible = false"
            >
            {{ $t('AdminNodeDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import dockerNodeApi from "@/api/dockerNode.js"
import imageApi from "@/api/image.js"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      activeTab: "image",
      dockerNode: {
        "ID": 0,
      },
      getImageListFromDockerNodeData: {
        "DockerNodeID": 0,
      },
      dockerNodeImageList: [],
      imageList: [],
      inPull: false,
      pullImageFormVisible: false,
      pullImageDataList: [],
      imageDetailVisible: false,
      imageDetail: {
        "Id": "",
        "Created": 0,
        "RepoDigests": "",
        "RepoTags": "",
        "Size": ""
      },
      removeImageFormVisible: false,
      removeImageData: {
        "DockerNodeID": 0,
        "DockerNodeImageID": "",
      },
      getContainerListFromDockerNodeData: {
        "DockerNodeID": 0,
      },
      dockerNodeContainerList: [],
      containerDetailVisible: false,
      containerDetail: {
        "Id": "",
        "Names": [],
        "Created": 0,
        "Image": "",
        "ImageID": "",
        "Ports": [],
        "State": "",
        "Status": "",
      },
      startContainerFormVisible: false,
      startContainerData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
      },
      stopContainerFormVisible: false,
      stopContainerData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
      },
      removeContainerFormVisible: false,
      removeContainerData: {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
      },
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1);
    },
    GetDockerNode() {
      dockerNodeApi.GetDockerNode(this.dockerNode).then(res => {
        if (res.code === 0) {
          this.dockerNode = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetImageList() {
      imageApi.GetImageList().then(res => {
        if (res.code === 0) {
          this.imageList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetImageListFromDockerNode() {
      dockerNodeApi.GetImageListFromDockerNode(this.getImageListFromDockerNodeData).then(res => {
        if (res.code === 0) {
          this.dockerNodeImageList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetContainerListFromDockerNode() {
      dockerNodeApi.GetContainerListFromDockerNode(this.getContainerListFromDockerNodeData).then(res => {
        if (res.code === 0) {
          this.dockerNodeContainerList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    async PullImageFromDockerNode() {
      this.inPull = true;
      const pullPromises = this.pullImageDataList.map(async data => {
        try {
          data.Status = this.t('pulling')
          const res = await dockerNodeApi.PullImageFromDockerNode({
            "DockerNodeID": data.DockerNodeID,
            "ImageID": data.ImageID
          });
          if (res.code === 0) {
            data.Status = this.t('pulled')
          }
        } catch (error) {
          console.log(error)
        }
      });
      await Promise.all(pullPromises)
      this.GetImageListFromDockerNode()
      this.inPull = false
    },
    ClearPullImageForm() {
      if (!this.inPull){
        this.$refs.pullImageForm.clearSelection()
        this.pullImageDataList = []
      }
    },
    OpenImageDetail(row){
      this.imageDetail = row
      this.imageDetailVisible = true
    },
    ClearImageDetail(){
      this.imageDetail = {
        "Id": "",
        "Created": 0,
        "RepoDigests": "",
        "RepoTags": "",
        "Size": ""
      }
    },
    RemoveImageFromDockerNode() {
      dockerNodeApi.RemoveImageFromDockerNode(this.removeImageData).then(res => {
        if (res.code === 0) {
          this.removeImageFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetImageListFromDockerNode()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenRemoveImageForm(row)  {
      this.removeImageData = {
        "DockerNodeID": this.dockerNode.ID,
        "DockerNodeImageID": row.Id,
      }
      this.removeImageFormVisible = true
    },
    ClearRemoveImageForm() {
      this.removeImageData = {
        "DockerNodeID": 0,
        "DockerNodeImageID": "",
      }
    },
    OpenContainerDetail(row) {
      this.containerDetail = row
      this.containerDetailVisible = true
    },
    ClearContainerDetail() {
      this.containerDetail = {
        "Id": "",
        "Names": [],
        "Created": 0,
        "Image": "",
        "ImageID": "",
        "Ports": [],
        "State": "",
        "Status": "",
      }
    },
    StartContainerFromDockerNode() {
      dockerNodeApi.StartContainerFromDockerNode(this.startContainerData).then(res => {
        if (res.code === 0) {
          this.startContainerFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetContainerListFromDockerNode()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenStartContainerForm(row) {
      this.startContainerData = {
        "DockerNodeID": this.dockerNode.ID,
        "DockerNodeContainerID": row.Id,
      }
      this.startContainerFormVisible = true
    },
    ClearStartContainerForm() {
      this.startContainerData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
      }
    },
    StopContainerFromDockerNode() {
      dockerNodeApi.StopContainerFromDockerNode(this.stopContainerData).then(res => {
        if (res.code === 0) {
          this.stopContainerFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetContainerListFromDockerNode()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenStopContainerForm(row) {
      this.stopContainerData = {
        "DockerNodeID": this.dockerNode.ID,
        "DockerNodeContainerID": row.Id,
      }
      this.stopContainerFormVisible = true
    },
    ClearStopContainerForm() {
      this.stopContainerData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
      }
    },
    RemoveContainerFromDockerNode() {
      dockerNodeApi.RemoveContainerFromDockerNode(this.removeContainerData).then(res => {
        if (res.code === 0) {
          this.removeContainerFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetContainerListFromDockerNode()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenRemoveContainerForm(row) {
      this.removeContainerData = {
        "DockerNodeID": this.dockerNode.ID,
        "DockerNodeContainerID": row.Id,
      }
      this.removeContainerFormVisible = true
    },
    ClearRemoveContainerForm() {
      this.removeContainerData = {
        "DockerNodeID": 0,
        "DockerNodeContainerID": "",
      }
    },
    handleSelectionChange(selection) {
      if (!this.inPull){
        this.pullImageDataList = selection.map(item => ({
          "DockerNodeID": this.dockerNode.ID,
          "ImageID": item.ID,
          "RepoTags": item.RepoTags,
          "Status": "waiting",
        }));
      }
    },
    checkImageExistence(repoTags) {
      return this.dockerNodeImageList.some(image => image.RepoTags.some(tag => repoTags.includes(tag))) ? this.t('AdminNodeDetail.Exists') : this.t('AdminNodeDetail.NotExists')
    },
    formatDate(timestamp) {
      const date = new Date(timestamp * 1000)
      const year = date.getFullYear()
      const month = (date.getMonth() + 1).toString().padStart(2, '0')
      const day = date.getDate().toString().padStart(2, '0')
      return `${year}-${month}-${day}`
    },
    formatRepository(repository) {
      return repository === '' ? this.t('AdminNodeDetail.NotSpecified') : repository
    },
  },
  mounted() {
    this.dockerNode.ID = Number(this.$route.params.id)
    this.getImageListFromDockerNodeData.DockerNodeID = Number(this.$route.params.id)
    this.getContainerListFromDockerNodeData.DockerNodeID = Number(this.$route.params.id)
    this.GetDockerNode()
    this.GetImageList()
    this.GetImageListFromDockerNode()
    this.GetContainerListFromDockerNode()
  }
}
</script>
<style>
.el-tabs-custom{
  overflow-y: hidden;
  height: 100%;
}
.el-tabs-custom .el-tabs__content{
  height: 90%;
}
.el-tabs-custom .el-tabs__content .el-tab-pane{
  height: 100%;
}
</style>
<style scoped>

</style>
