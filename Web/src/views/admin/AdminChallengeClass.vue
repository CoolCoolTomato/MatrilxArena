<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">ChallengeClass Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createChallengeClassFormVisible = true">Add</el-button>
      <el-switch
        style="margin: 10px"
        @change="switchAllowSort"
        v-model="allowSort"
        active-text="Allow sort"
      />
      <el-table
        id="challengeClassListTable"
        :data="challengeClassList"
        table-layout="fixed"
        row-key="ID"
      >
        <el-table-column prop="Name" label="Name"/>
        <el-table-column prop="Order" label="Order"/>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateChallengeClassForm(scope.row)"
            >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteChallengeClassForm(scope.row)"
            >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createChallengeClassFormVisible"
        title="Create ChallengeClass"
        width="500"
        @close="ClearCreateChallengeClassForm"
      >
        <el-form :model=createChallengeClassData>
          <el-form-item label="Name" :label-width="labelWidth">
            <el-input v-model="createChallengeClassData.Name"/>
          </el-form-item>
          <el-form-item label="Order" :label-width="labelWidth">
            <el-input v-model.number="createChallengeClassData.Order" type="number"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createChallengeClassFormVisible = false">Cancel</el-button>
          <el-button @click="CreateChallengeClass">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateChallengeClassFormVisible"
        title="Update ChallengeClass"
        width="500"
        @close="ClearUpdateChallengeClassForm"
      >
        <el-form :model=updateChallengeClassData>
          <el-form-item label="Name" :label-width="labelWidth">
            <el-input v-model="updateChallengeClassData.Name"/>
          </el-form-item>
          <el-form-item label="Order" :label-width="labelWidth">
            <el-input v-model.number="updateChallengeClassData.Order" type="number"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateChallengeClassFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateChallengeClass">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteChallengeClassFormVisible"
        title="Delete ChallengeClass"
        width="500"
        @close="ClearDeleteChallengeClassForm"
      >
        <el-text>Are you confirm to delete the challengeClass?</el-text>
        <template #footer>
          <el-button @click="deleteChallengeClassFormVisible = false">Cancel</el-button>
          <el-button @click="DeleteChallengeClass">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import challengeClassApi from "@/api/challengeClass.js"
import { ElMessage } from "element-plus";
import Sortable from 'sortablejs';

export default {
  data() {
    return {
      labelWidth: 100,
      challengeClassList: [],
      createChallengeClassFormVisible: false,
      createChallengeClassData: {
        "Name": "",
        "Order": 0,
      },
      updateChallengeClassFormVisible: false,
      updateChallengeClassData: {
        "Name": "",
        "Order": 0,
      },
      deleteChallengeClassFormVisible: false,
      deleteChallengeClassData: {
        "ID": 0,
      },
      sortable: null,
      allowSort: false,
    }
  },
  methods: {
    GetChallengeClassList() {
      challengeClassApi.GetChallengeClassList().then(res => {
        if (res.code === 0) {
          this.challengeClassList = res.data
          this.challengeClassList.sort((a, b) => a.Order - b.Order)
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateChallengeClass() {
      challengeClassApi.CreateChallengeClass(this.createChallengeClassData).then(res => {
        if (res.code === 0) {
          this.createChallengeClassFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetChallengeClassList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateChallengeClassForm() {
      this.createChallengeClassData = {
        "Name": "",
        "Order": 0,
      }
    },
    UpdateChallengeClass() {
      challengeClassApi.UpdateChallengeClass(this.updateChallengeClassData).then(res => {
        if (res.code === 0) {
          this.updateChallengeClassFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetChallengeClassList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateChallengeClassForm(row) {
      this.updateChallengeClassData = {
        "ID": row.ID,
        "Name": row.Name,
        "Order": row.Order,
      }
      this.updateChallengeClassFormVisible = true
    },
    ClearUpdateChallengeClassForm() {
      this.updateChallengeClassData = {
        "ID": 0,
        "Name": "",
        "Order": 0,
      }
    },
    DeleteChallengeClass() {
      challengeClassApi.DeleteChallengeClass(this.deleteChallengeClassData).then(res => {
        if (res.code === 0) {
          this.deleteChallengeClassFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetChallengeClassList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteChallengeClassForm(row) {
      this.deleteChallengeClassData = {
        "ID": row.ID,
      }
      this.deleteChallengeClassFormVisible = true
    },
    ClearDeleteChallengeClassForm() {
      this.deleteChallengeClassData = {
        "ID": 0,
      }
    },
    initSortable() {
      const challengeClassListTable = document.querySelector("#challengeClassListTable tbody")
      this.sortable = new Sortable(challengeClassListTable, {
        disabled: !this.allowSort,
        animation: 100,
        onEnd: (evt) => {
          const newIndex = evt.newIndex
          const oldIndex = evt.oldIndex
          const movedItem = this.challengeClassList.splice(oldIndex, 1)[0]
          this.challengeClassList.splice(newIndex, 0, movedItem)
          this.updateOrder()
        }
      })
    },
    switchAllowSort(){
      this.sortable.option('disabled', !this.allowSort)
    },
    updateOrder() {
      const updatePromises = [];
      let total = 0
      let success = 0
      let fail = 0
      this.challengeClassList.forEach((challengeClass, index) => {
        if (challengeClass.Order !== index + 1) {
          total += 1
          challengeClass.Order = index + 1
          const updatePromise = challengeClassApi.UpdateChallengeClass(challengeClass).then(res => {
              if (res.code === 0) {
                success += 1
              } else {
                fail += 1
              }
              return res
            }).catch(error => {
              console.log(error)
              throw error
            })
          updatePromises.push(updatePromise)
        }
      })
      Promise.all(updatePromises).then(() => {
        if (fail === 0) {
          ElMessage({
            message: `Update ${total} items, ${success} success, ${fail} fail`,
            type: 'success',
          })
        } else {
          ElMessage({
            message: `Update ${total} items, ${success} success, ${fail} fail`,
            type: 'warning',
          })
        }
        this.GetChallengeClassList()
      }).catch(error => {
        console.log(error)
      })
    },
  },
  mounted() {
    this.GetChallengeClassList()
    this.initSortable()
  }
}
</script>
<style scoped>
::v-deep(#challengeClassListTable .el-table__row:hover) {
  cursor: pointer;
}
</style>