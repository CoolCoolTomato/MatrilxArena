<template>
  <el-container>
    <el-header style="display: flex">
      <h2 style="color: var(--el-text-color-primary);">Node Detail / Host: {{dockerNode.Host}} | Port: {{dockerNode.Port}}</h2>
      <div style="flex-grow: 1;" />
      <div style="display: flex; align-items: center">
        <el-button
          @click="pullImageFormVisible = true"
          >Pull image</el-button>
        <el-button
          @click="goBack"
          >
           Back
        </el-button>
      </div>
    </el-header>
    <el-main>
      <el-tabs
        v-model="activeTab"
        class="el-tabs-custom"
        >
        <el-tab-pane label="Image" name="image">
          <el-table
            :data="dockerNodeImageList"
            table-layout="fixed"
            height="99%"
            >
            <el-table-column prop="RepoTags" label="RepoTags"/>
            <el-table-column label="Size">
              <template #default="scope">
                {{ (scope.row.Size / 1000000).toFixed(2) }} MB
              </template>
            </el-table-column>
            <el-table-column label="Created">
              <template #default="scope">
                {{ formatDate(scope.row.Created) }}
              </template>
            </el-table-column>
            <el-table-column fixed="right" label="Operations" width="200px">
              <template #default=scope>
                <el-button
                  size="small"
                  @click="OpenImageDetail(scope.row)"
                  >
                  Detail
                </el-button>
                <el-button
                  size="small"
                  @click="OpenRemoveImageForm(scope.row)"
                  >
                  Remove
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane label="Container" name="container">
          <el-table
            :data="dockerNodeContainerList"
            table-layout="fixed"
            height="99%"
            >
            <el-table-column prop="Names" label="Names"/>
            <el-table-column prop="Image" label="Image"/>
            <el-table-column prop="State" label="State"/>
            <el-table-column fixed="right" label="Operations" width="320px">
              <template #default=scope>
                <el-button
                  size="small"
                  @click="OpenContainerDetail(scope.row)"
                  >
                  Detail
                </el-button>
                <el-button
                  size="small"
                  @click="OpenStartContainerForm(scope.row)"
                  >
                  Start
                  </el-button>
                <el-button
                  size="small"
                  @click="OpenStopContainerForm(scope.row)"
                  >
                  Stop
                </el-button>
                <el-button
                  size="small"
                  @click="OpenRemoveContainerForm(scope.row)"
                  >
                  Remove
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
      <el-dialog
        v-model="pullImageFormVisible"
        title="Pull Image"
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
              <el-table-column prop="Remark" label="Remark"/>
              <el-table-column prop="RepoTags" label="RepoTags"/>
              <el-table-column label="Repository">
                <template #default="scope">
                    {{ formatRepository(scope.row.Repository) }}
                </template>
              </el-table-column>
              <el-table-column label="Existence">
                <template #default="scope">
                  {{ checkImageExistence(scope.row.RepoTags) }}
                </template>
              </el-table-column>
            </el-table>
          </el-col>
          <el-col :span=9>
            <el-table
              :data="pullImageDataList"
              max-height="500px"
              >
              <el-table-column prop="RepoTags" label="Selected image"/>
              <el-table-column prop="Status" label="Status"/>
            </el-table>
          </el-col>
        </el-row>
        <template #footer>
          <el-button
            @click="pullImageFormVisible = false"
            >
            Cancel
          </el-button>
          <el-button
            @click="PullImageFromDockerNode"
            >
            Pull
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="imageDetailVisible"
        title="Image Detail"
        width="1000"
        @close="ClearImageDetail"
        >
        <el-card>
          <p style="word-break: break-all;">ID: {{ imageDetail.Id }}</p>
          <p style="word-break: break-all;">RepoTags: {{ imageDetail.RepoTags }}</p>
          <p style="word-break: break-all;">Size: {{ (imageDetail.Size / 1000000).toFixed(2) }} MB</p>
          <p style="word-break: break-all;">Created: {{ formatDate(imageDetail.Created) }}</p>
          <p style="word-break: break-all;">RepoDigests: {{ imageDetail.RepoDigests }}</p>
        </el-card>
        <template #footer>
          <el-button @click="imageDetailVisible = false">Close</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="removeImageFormVisible"
        title="Remove Image"
        width="500"
        @close="ClearRemoveImageForm"
      >
        <el-text>Are you confirm to remove the image?</el-text>
        <template #footer>
          <el-button @click="removeImageFormVisible = false">Cancel</el-button>
          <el-button @click="RemoveImageFromDockerNode">Confirm</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="containerDetailVisible"
        title="Container Detail"
        width="1000"
        @close="ClearContainerDetail"
        >
        <el-card>
          <p style="word-break: break-all;">ID: {{ containerDetail.Id }}</p>
          <p style="word-break: break-all;">Names: {{ containerDetail.Names }}</p>
          <p style="word-break: break-all;">Created: {{ formatDate(containerDetail.Created) }}</p>
          <p style="word-break: break-all;">Image: {{ containerDetail.Image }}</p>
          <p style="word-break: break-all;">ImageID: {{ containerDetail.ImageID }}</p>
          <p style="word-break: break-all;">Ports: {{ containerDetail.Ports }}</p>
          <p style="word-break: break-all;">State: {{ containerDetail.State }}</p>
          <p style="word-break: break-all;">Status: {{ containerDetail.Status }}</p>
        </el-card>
        <template #footer>
          <el-button @click="containerDetailVisible = false">Close</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="startContainerFormVisible"
        title="Start Container"
        width="500"
        @close="ClearStartContainerForm"
        >
        <el-text>Are you confirm to start the container?</el-text>
        <template #footer>
          <el-button @click="startContainerFormVisible = false">Cancel</el-button>
          <el-button @click="StartContainerFromDockerNode">Confirm</el-button>
          </template>
      </el-dialog>
      <el-dialog
        v-model="stopContainerFormVisible"
        title="Stop Container"
        width="500"
        @close="ClearStopContainerForm"
        >
        <el-text>Are you confirm to stop the container?</el-text>
        <template #footer>
        <el-button @click="stopContainerFormVisible = false">Cancel</el-button>
        <el-button @click="StopContainerFromDockerNode">Confirm</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="removeContainerFormVisible"
        title="Remove Container"
        width="500"
        @close="ClearRemoveContainerForm"
        >
        <el-text>Are you confirm to remove the container?</el-text>
        <template #footer>
        <el-button @click="removeContainerFormVisible = false">Cancel</el-button>
        <el-button @click="RemoveContainerFromDockerNode">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import dockerNodeApi from "@/api/dockerNode.js";
import imageApi from "@/api/image.js";
import {ElMessage} from "element-plus";
export default {
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
          data.Status = "pulling";
          const res = await dockerNodeApi.PullImageFromDockerNode({
            "DockerNodeID": data.DockerNodeID,
            "ImageID": data.ImageID
          });
          if (res.code === 0) {
            data.Status = "pulled";
          }
        } catch (error) {
          console.log(error);
        }
      });
      await Promise.all(pullPromises)
      this.GetImageListFromDockerNode()
      this.inPull = false;
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
      return this.dockerNodeImageList.some(image => image.RepoTags.some(tag => repoTags.includes(tag))) ? 'Exists' : 'Not Exists';
    },
    formatDate(timestamp) {
      const date = new Date(timestamp * 1000)
      const year = date.getFullYear()
      const month = (date.getMonth() + 1).toString().padStart(2, '0')
      const day = date.getDate().toString().padStart(2, '0')
      return `${year}-${month}-${day}`
    },
    formatRepository(repository) {
      return repository === '' ? 'Not Specified' : repository;
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
.el-menu-item-no-style{
  background-color: transparent !important;
  color: inherit !important;
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  cursor: default;
}
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
