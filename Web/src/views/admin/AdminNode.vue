<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">{{ $t('AdminNode.DockerNodeManager') }}</h2>
    </el-header>
    <el-main>
      <el-scrollbar>
        <el-button
          style="margin: 10px"
          @click="createDockerNodeFormVisible = true"
          type="primary"
          >
          {{ $t('AdminNode.Add') }}
        </el-button>
        <el-table
          :data="dockerNodeList"
          table-layout="fixed"
        >
          <el-table-column prop="Host" :label="$t('AdminNode.Host')"/>
          <el-table-column prop="Port" :label="$t('AdminNode.Port')"/>
          <el-table-column fixed="right" :label="$t('AdminNode.Operations')" width="300px">
            <template #default=scope>
              <el-button
                @click="OpenDockerNodeDetail(scope.row)"
                >
                {{ $t('AdminNode.Detail') }}
              </el-button>
              <el-button
                @click="OpenUpdateDockerNodeForm(scope.row)"
                >
                {{ $t('AdminNode.Update') }}
              </el-button>
              <el-button
                @click="OpenDeleteDockerNodeForm(scope.row)"
                >
                {{ $t('AdminNode.Delete') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-scrollbar>
      <el-dialog
        v-model="createDockerNodeFormVisible"
        :title="$t('AdminNode.CreateNode')"
        width="500"
        @close="ClearCreateDockerNodeForm"
        >
        <el-form :model=createDockerNodeData>
          <el-form-item :label="$t('AdminNode.Host')" :label-width="labelWidth">
            <el-input v-model="createDockerNodeData.Host"/>
          </el-form-item>
          <el-form-item :label="$t('AdminNode.Port')" :label-width="labelWidth">
            <el-input v-model="createDockerNodeData.Port"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateDockerNode"
            type="primary"
            >
            {{ $t('AdminNode.Submit') }}
          </el-button>
          <el-button
            @click="createDockerNodeFormVisible = false"
            >
            {{ $t('AdminNode.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateDockerNodeFormVisible"
        :title="$t('AdminNode.UpdateNode')"
        width="500"
        @close="ClearUpdateDockerNodeForm"
        >
        <el-form :model=updateDockerNodeData>
          <el-form-item :label="$t('AdminNode.Host')" :label-width="labelWidth">
            <el-input v-model="updateDockerNodeData.Host"/>
          </el-form-item>
          <el-form-item :label="$t('AdminNode.Port')" :label-width="labelWidth">
            <el-input v-model="updateDockerNodeData.Port"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateDockerNode"
            type="primary"
            >
            {{ $t('AdminNode.Submit') }}
          </el-button>
          <el-button
            @click="updateDockerNodeFormVisible = false"
            >
            {{ $t('AdminNode.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteDockerNodeFormVisible"
        :title="$t('AdminNode.DeleteNode')"
        width="500"
        @close="ClearDeleteDockerNodeForm"
        >
        <el-text>{{ $t('AdminNode.AreYouConfirmToDeleteTheDockerNode') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteDockerNode"
            type="primary"
            >
            {{ $t('AdminNode.Confirm') }}
          </el-button>
          <el-button
            @click="deleteDockerNodeFormVisible = false"
            >
            {{ $t('AdminNode.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import dockerNodeApi from "@/api/dockerNode.js"
import { ElMessage } from 'element-plus'
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      labelWidth: 100,
      dockerNodeList: [],
      createDockerNodeFormVisible: false,
      createDockerNodeData: {
        "Host": "",
        "Port": "",
      },
      updateDockerNodeFormVisible: false,
      updateDockerNodeData: {
        "ID": 0,
        "Host": "",
        "Port": "",
      },
      deleteDockerNodeFormVisible: false,
      deleteDockerNodeData: {
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
      dockerNodeApi.CreateDockerNode(this.createDockerNodeData).then(res => {
        if (res.code === 0) {
          this.createDockerNodeFormVisible = false
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
      this.createDockerNodeData = {
        "Host": "",
        "Port": "",
      }
    },
    UpdateDockerNode() {
      dockerNodeApi.UpdateDockerNode(this.updateDockerNodeData).then(res => {
        if (res.code === 0) {
          this.updateDockerNodeFormVisible = false
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
      this.updateDockerNodeData = {
        "ID": row.ID,
        "Host": row.Host,
        "Port": row.Port,
      }
      this.updateDockerNodeFormVisible = true
    },
    ClearUpdateDockerNodeForm() {
      this.updateDockerNodeData = {
        "ID": 0,
        "Host": "",
        "Port": "",
      }
    },
    DeleteDockerNode() {
      dockerNodeApi.DeleteDockerNode(this.deleteDockerNodeData).then(res => {
        if (res.code === 0) {
          this.deleteDockerNodeFormVisible = false
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
      this.deleteDockerNodeData = {
        "ID": row.ID,
      }
      this.deleteDockerNodeFormVisible = true
    },
    ClearDeleteDockerNodeForm() {
      this.deleteDockerNodeData = {
        "ID": 0,
      }
    },
  },
  mounted() {
    this.GetDockerNodeList()
  },
}
</script>
