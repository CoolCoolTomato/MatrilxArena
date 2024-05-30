<template>
  <el-text size="large">Docker Node Manager</el-text>
  <br>
  <el-button style="margin: 10px" @click="CreateDockerNodeFormVisible = true">Add</el-button>
  <el-table
    :data="dockerNodeList"
    table-layout="fixed"
  >
    <el-table-column prop="Host" label="Host"/>
    <el-table-column prop="Port" label="Port"/>
    <el-table-column fixed="right" label="Operations" width="300px">
      <template #default=scope>
        <el-button
          @click="OpenDockerNodeDetail(scope.row)"
          >
          Detail
        </el-button>
        <el-button
          @click="OpenUpdateDockerNodeForm(scope.row)"
          >
          Update
        </el-button>
        <el-button
          @click="OpenDeleteDockerNodeForm(scope.row)"
          >
          Delete
        </el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-dialog
    v-model="CreateDockerNodeFormVisible"
    title="Create Node"
    width="500"
    @close="ClearCreateDockerNodeForm"
    >
    <el-form :model=CreateDockerNodeData>
      <el-form-item label="Host">
        <el-input v-model="CreateDockerNodeData.Host"/>
      </el-form-item>
      <el-form-item label="Port">
        <el-input v-model="CreateDockerNodeData.Port"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="CreateDockerNodeFormVisible = false">Cancel</el-button>
      <el-button @click="CreateDockerNode">Submit</el-button>
    </template>
  </el-dialog>
  <el-dialog
    v-model="UpdateDockerNodeFormVisible"
    title="Update Node"
    width="500"
    @close="ClearUpdateDockerNodeForm"
    >
    <el-form :model=UpdateDockerNodeData>
      <el-form-item label="Host">
        <el-input v-model="UpdateDockerNodeData.Host"/>
      </el-form-item>
      <el-form-item label="Port">
        <el-input v-model="UpdateDockerNodeData.Port"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="UpdateDockerNodeFormVisible = false">Cancel</el-button>
      <el-button @click="UpdateDockerNode">Submit</el-button>
    </template>
  </el-dialog>
  <el-dialog
    v-model="DeleteDockerNodeFormVisible"
    title="Delete Node"
    width="500"
    @close="ClearDeleteDockerNodeForm"
    >
    <el-text>Are you confirm to delete the docker node?</el-text>
    <template #footer>
      <el-button @click="DeleteDockerNodeFormVisible = false">Cancel</el-button>
      <el-button @click="DeleteDockerNode">Confirm</el-button>
    </template>
  </el-dialog>
</template>
<script>
import dockerNodeApi from "@/api/dockerNode.js"
import { ElMessage } from 'element-plus'

export default {
  data() {
    return {
      dockerNodeList: [],
      CreateDockerNodeFormVisible: false,
      CreateDockerNodeData: {
        "Host": "",
        "Port": "",
      },
      UpdateDockerNodeFormVisible: false,
      UpdateDockerNodeData: {
        "ID": 0,
        "Host": "",
        "Port": "",
      },
      DeleteDockerNodeFormVisible: false,
      DeleteDockerNodeData: {
        "ID": 0,
      }
    }
  },
  methods: {
    GetDockerNodeList() {
      dockerNodeApi.GetDockerNodeList().then(res => {
        if (res.code === 0) {
          this.dockerNodeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDockerNodeDetail(row) {
      this.$router.push(`/admin/node/${row.ID}`);
    },
    CreateDockerNode() {
      dockerNodeApi.CreateDockerNode(this.CreateDockerNodeData).then(res => {
        if (res.code === 0) {
          this.CreateDockerNodeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetDockerNodeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateDockerNodeForm() {
      this.CreateDockerNodeData = {
        "Host": "",
        "Port": "",
      }
    },
    UpdateDockerNode() {
      dockerNodeApi.UpdateDockerNode(this.UpdateDockerNodeData).then(res => {
        if (res.code === 0) {
          this.UpdateDockerNodeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetDockerNodeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateDockerNodeForm(row) {
      this.UpdateDockerNodeData = {
        "ID": row.ID,
        "Host": row.Host,
        "Port": row.Port,
      }
      this.UpdateDockerNodeFormVisible = true
    },
    ClearUpdateDockerNodeForm() {
      this.UpdateDockerNodeData = {
        "ID": 0,
        "Host": "",
        "Port": "",
      }
    },
    DeleteDockerNode() {
      dockerNodeApi.DeleteDockerNode(this.DeleteDockerNodeData).then(res => {
        if (res.code === 0) {
          this.DeleteDockerNodeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetDockerNodeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteDockerNodeForm(row) {
      this.DeleteDockerNodeData = {
        "ID": row.ID,
      }
      this.DeleteDockerNodeFormVisible = true
    },
    ClearDeleteDockerNodeForm() {
      this.DeleteDockerNodeData = {
        "ID": 0,
      }
    },
  },
  mounted() {
    this.GetDockerNodeList()
  },
}
</script>
