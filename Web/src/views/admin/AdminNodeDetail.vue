<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary);">Node Detail</h2>
    </el-header>
    <el-main>
      <el-descriptions size="large">
        <el-descriptions-item label="Host" width="130px">{{dockerNode.Host}}</el-descriptions-item>
        <el-descriptions-item label="Port" width="130px">{{dockerNode.Port}}</el-descriptions-item>
      </el-descriptions>
      <el-tabs
        v-model="activeTab"
        class="el-tabs-custom"
        >
        <el-tab-pane label="Image" name="image">
          <el-table
            :data="imageList"
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
                  >
                  Detail
                </el-button>
                <el-button
                  >
                  Remove
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane label="Container" name="container">
          <el-table
            :data="containerList"
            table-layout="fixed"
            height="99%"
            >
            <el-table-column prop="Names" label="Names"/>
            <el-table-column prop="Image" label="Image"/>
            <el-table-column prop="State" label="State"/>
            <el-table-column fixed="right" label="Operations" width="200px">
              <template #default=scope>
                <el-button
                  >
                  Detail
                </el-button>
                <el-button
                  >
                  Remove
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>
<script>
import dockerNodeApi from "@/api/dockerNode.js";
import {ElMessage} from "element-plus";
export default {
  data() {
    return {
      activeTab: "image",
      dockerNode: {
        ID: 0,
      },
      getImageListFromDockerNodeData: {
        DockerNodeID: 0,
      },
      imageList: [],
      getContainerListFromDockerNodeData: {
        DockerNodeID: 0,
      },
      containerList: [],

    }
  },
  methods: {
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
    GetImageListFromDockerNode() {
      dockerNodeApi.GetImageListFromDockerNode(this.getImageListFromDockerNodeData).then(res => {
        if (res.code === 0) {
          this.imageList = res.data
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
          this.containerList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    formatDate(timestamp) {
      const date = new Date(timestamp * 1000);
      const year = date.getFullYear();
      const month = (date.getMonth() + 1).toString().padStart(2, '0');
      const day = date.getDate().toString().padStart(2, '0');
      return `${year}-${month}-${day}`;
    },
  },
  mounted() {
    this.dockerNode.ID = Number(this.$route.params.id)
    this.getImageListFromDockerNodeData.DockerNodeID = Number(this.$route.params.id)
    this.getContainerListFromDockerNodeData.DockerNodeID = Number(this.$route.params.id)
    this.GetDockerNode()
    this.GetImageListFromDockerNode()
    this.GetContainerListFromDockerNode()
  }
}
</script>
<style>
.el-tabs-custom{
  height: 90%;
  overflow-y: hidden;
}
.el-tabs-custom .el-tabs__content{
  height: 90%;
}
.el-tabs-custom .el-tabs__content .el-tab-pane{
  height: 100%;
}
</style>
