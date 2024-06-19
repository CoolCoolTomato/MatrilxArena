<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">Challenge Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createChallengeFormVisible = true">Add</el-button>
      <el-table
        :data="challengeList"
        table-layout="fixed"
        >
        <el-table-column prop="Title" label="Title"/>
        <el-table-column prop="Description" label="Description"/>
        <el-table-column prop="Image.RepoTags" label="Image"/>
        <el-table-column prop="Attachment" label="Attachment"/>
        <el-table-column prop="SpecifiedPorts" label="SpecifiedPorts"/>
        <el-table-column prop="Commands" label="Commands"/>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateChallengeForm(scope.row)"
              >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteChallengeForm(scope.row)"
              >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createChallengeFormVisible"
        title="Create Challenge"
        width="500"
        @close="ClearCreateChallengeForm"
        >
        <el-form :model=createChallengeData>
          <el-form-item label="Title">
            <el-input v-model="createChallengeData.Title"/>
          </el-form-item>
          <el-form-item label="Description">
            <el-input v-model="createChallengeData.Description"/>
          </el-form-item>
          <el-form-item label="Image">
            <el-select
              v-model="createChallengeData.ImageID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="image in imageList"
                :key="image.RepoTags"
                :label="image.RepoTags"
                :value="image.ID"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Attachment">
            <el-input v-model="createChallengeData.Attachment"/>
          </el-form-item>
          <el-form-item label="SpecifiedPorts">
            <div style="display: flex; width: 100%;">
              <el-input
                v-model="createChallengePort"
              >
                <template #append>
                  <el-select
                    v-model="createChallengeProtocol"
                    placeholder="Select"
                    style="width: 100px"
                  >
                    <el-option label="tcp" value="tcp"/>
                    <el-option label="udp" value="udp"/>
                  </el-select>
                </template>
              </el-input>
              <el-button
                style="margin-left: 20px;"
                @click="AddCreateChallengeSpecifiedPort"
              >
                Add
              </el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <div style="position: relative; width: 100%;">
              <div
                v-for="(port, index) in createChallengeData.SpecifiedPorts"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
                >
                <el-text>{{ port }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveCreateChallengeSpecifiedPort(index)"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Commands">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.createChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddCreateChallengeCommand"
              >
                Add
              </el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <div style="position: relative; width: 100%;">
              <div
                v-for="(command, index) in createChallengeData.Commands"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
                >
                <el-text>{{ command }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveCreateChallengeCommand(index)"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createChallengeFormVisible = false">Cancel</el-button>
          <el-button @click="CreateChallenge">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateChallengeFormVisible"
        title="Update Challenge"
        width="500"
        @close="ClearUpdateChallengeForm"
        >
        <el-form :model=updateChallengeData>
          <el-form-item label="Title">
            <el-input v-model="updateChallengeData.Title"/>
          </el-form-item>
          <el-form-item label="Description">
            <el-input v-model="updateChallengeData.Description"/>
          </el-form-item>
          <el-form-item label="Image">
            <el-select
              v-model="updateChallengeData.ImageID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="image in imageList"
                :key="image.RepoTags"
                :label="image.RepoTags"
                :value="image.ID"
                >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Attachment">
            <el-input v-model="updateChallengeData.Attachment"/>
          </el-form-item>
          <el-form-item label="SpecifiedPorts">
            <div style="display: flex; width: 100%;">
              <el-input
                v-model="updateChallengePort"
              >
                <template #append>
                  <el-select
                    v-model="updateChallengeProtocol"
                    placeholder="Select"
                    style="width: 100px"
                  >
                    <el-option label="tcp" value="tcp"/>
                    <el-option label="udp" value="udp"/>
                  </el-select>
                </template>
              </el-input>
              <el-button
                style="margin-left: 20px;"
                @click="AddUpdateChallengeSpecifiedPort"
              >
                Add
              </el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <div style="position: relative; width: 100%;">
              <div
                v-for="(port, index) in updateChallengeData.SpecifiedPorts"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
                >
                <el-text>{{ port }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveUpdateChallengeSpecifiedPort(index)"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Commands">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.updateChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddUpdateChallengeCommand"
              >
                Add
              </el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <div style="position: relative; width: 100%;">
              <div
                v-for="(command, index) in updateChallengeData.Commands"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
                >
                <el-text>{{ command }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveUpdateChallengeCommand(index)"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateChallengeFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateChallenge">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteChallengeFormVisible"
        title="Delete Challenge"
        width="500"
        @close="ClearDeleteChallengeForm"
        >
        <el-text>Are you confirm to delete the challenge?</el-text>
        <template #footer>
        <el-button @click="deleteChallengeFormVisible = false">Cancel</el-button>
        <el-button @click="DeleteChallenge">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import challengeApi from "@/api/challenge.js"
import imageApi from "@/api/image.js";
import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      imageList: [],
      challengeList: [],
      createChallengeFormVisible: false,
      createChallengeData: {
        "ImageID": null,
        "Title": "",
        "Description": "",
        "Attachment": "",
        "SpecifiedPorts": [],
        "Commands": []
      },
      createChallengePort: "",
      createChallengeProtocol: "",
      createChallengeCommand: "",
      updateChallengeFormVisible: false,
      updateChallengeData: {
        "ID": 0,
        "ImageID": null,
        "Title": "",
        "Description": "",
        "Attachment": "",
        "SpecifiedPorts": [],
        "Commands": []
      },
      updateChallengePort: "",
      updateChallengeProtocol: "",
      updateChallengeCommand: "",
      deleteChallengeFormVisible: false,
      deleteChallengeData: {
        "ID": 0,
      }
    }
  },
  methods: {
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
    AddCreateChallengeSpecifiedPort() {
      if (this.createChallengePort === "") {
        ElMessage.error("Please input port")
        return
      }
      if (this.createChallengeProtocol === "") {
        ElMessage.error("Please input protocol")
        return
      }
      let portProtocol = this.createChallengePort + "/" + this.createChallengeProtocol
      if (this.createChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error("Duplicate specifiedPort")
        return
      }
      this.createChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveCreateChallengeSpecifiedPort(index) {
      this.createChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddCreateChallengeCommand() {
      if (this.createChallengeCommand === "") {
        ElMessage.error("Please input command")
        return
      }
      if (this.createChallengeData.Commands.includes(this.createChallengeCommand)) {
        ElMessage.error("Duplicate command")
        return
      }
      this.createChallengeData.Commands.push(this.createChallengeCommand)
    },
    RemoveCreateChallengeCommand(index) {
      this.createChallengeData.Commands.splice(index, 1);
    },
    CreateChallenge() {
      challengeApi.CreateChallenge(this.createChallengeData).then(res => {
        if (res.code === 0) {
          this.createChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateChallengeForm() {
      this.createChallengeData = {
        "ImageID": null,
        "Title": "",
        "Description": "",
        "Attachment": "",
        "SpecifiedPorts": [],
        "Commands": []
      }
      this.createChallengePort = ""
      this.createChallengeProtocol = ""
      this.createChallengeCommand = ""
    },
    AddUpdateChallengeSpecifiedPort() {
      if (this.updateChallengePort === "") {
        ElMessage.error("Please input port")
        return
      }
      if (this.updateChallengeProtocol === "") {
        ElMessage.error("Please input protocol")
        return
      }
      let portProtocol = this.updateChallengePort + "/" + this.updateChallengeProtocol
      if (this.updateChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error("Duplicate specifiedPort")
        return
      }
      this.updateChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveUpdateChallengeSpecifiedPort(index) {
      this.updateChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddUpdateChallengeCommand() {
      if (this.updateChallengeCommand === "") {
        ElMessage.error("Please input command")
        return
      }
      if (this.updateChallengeData.Commands.includes(this.updateChallengeCommand)) {
        ElMessage.error("Duplicate command")
        return
      }
      this.updateChallengeData.Commands.push(this.updateChallengeCommand)
    },
    RemoveUpdateChallengeCommand(index) {
      this.updateChallengeData.Commands.splice(index, 1);
    },
    UpdateChallenge() {
      challengeApi.UpdateChallenge(this.updateChallengeData).then(res => {
        if (res.code === 0) {
          this.updateChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateChallengeForm(row) {
      this.updateChallengeData = {
        "ID": row.ID,
        "ImageID": row.ImageID,
        "Title": row.Title,
        "Description": row.Description,
        "Attachment": row.Attachment,
        "SpecifiedPorts": row.SpecifiedPorts === null ? [] : row.SpecifiedPorts,
        "Commands": row.Commands === null ? [] : row.Commands
      }
      this.updateChallengeFormVisible = true
    },
    ClearUpdateChallengeForm() {
      this.updateChallengeData = {
        "ID": 0,
        "ImageID": null,
        "Title": "",
        "Description": "",
        "Attachment": "",
        "SpecifiedPorts": [],
        "Commands": []
      }
      this.updateChallengePort = ""
      this.updateChallengeProtocol = ""
      this.updateChallengeCommand = ""
    },
    DeleteChallenge() {
      challengeApi.DeleteChallenge(this.deleteChallengeData).then(res => {
        if (res.code === 0) {
          this.deleteChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteChallengeForm(row) {
      this.deleteChallengeData = {
        "ID": row.ID,
      }
      this.deleteChallengeFormVisible = true
    },
    ClearDeleteChallengeForm() {
      this.deleteChallengeData = {
        "ID": 0,
      }
    },
  },
  mounted() {
    this.GetImageList()
    this.GetChallengeList()
  }
}
</script>
