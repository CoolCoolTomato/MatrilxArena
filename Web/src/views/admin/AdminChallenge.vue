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
        <el-table-column label="Image">
          <template  #default=scope>
            {{ scope.row.Image.Remark === "" ? "null" : scope.row.Image.Remark}}
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenChallengeDetail(scope.row)"
              >
              Detail
            </el-button>
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
        v-model="challengeDetailVisible"
        title="Challenge Detail"
        width="500"
        @close="ClearChallengeDetail"
        >
        <el-card>
          <p style="word-break: break-all;">Title: {{ challengeDetail.Title }}</p>
          <p style="word-break: break-all;">Description: {{ challengeDetail.Description }}</p>
          <p style="word-break: break-all;">ChallengeClass: {{ challengeDetail.ChallengeClass.Name === "" ? "null" : challengeDetail.ChallengeClass.Name }}</p>
          <p style="word-break: break-all;">Image: {{ challengeDetail.Image.Remark === "" ? "null" : challengeDetail.Image.Remark }}</p>
          <p style="word-break: break-all;">RepoTags: {{ challengeDetail.Image.RepoTags === "" ? "null" : challengeDetail.Image.RepoTags }}</p>
          <p style="word-break: break-all;">Attachment: {{ challengeDetail.Attachment.FileName === "" ? "null" : challengeDetail.Attachment.FileName }}</p>
          <p style="word-break: break-all;">SpecifiedPorts: {{ challengeDetail.SpecifiedPorts }}</p>
          <p style="word-break: break-all;">Commands: {{ challengeDetail.Commands }}</p>
          <p style="word-break: break-all;">Flag: {{ challengeDetail.Flag }}</p>
        </el-card>
        <template #footer>
          <el-button @click="challengeDetailVisible = false">Close</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="createChallengeFormVisible"
        title="Create Challenge"
        width="600"
        @close="ClearCreateChallengeForm"
        >
        <el-form :model=createChallengeData>
          <el-form-item label="Title" :label-width="labelWidth">
            <el-input v-model="createChallengeData.Title"/>
          </el-form-item>
          <el-form-item label="Description" :label-width="labelWidth">
            <el-input v-model="createChallengeData.Description"/>
          </el-form-item>
          <el-form-item label="ChallengeClass" :label-width="labelWidth">
            <el-select
              v-model="createChallengeData.ChallengeClassID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="challengeClass in challengeClassList"
                :key="challengeClass.ID"
                :label="challengeClass.Name"
                :value="challengeClass.ID"
                ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Image" :label-width="labelWidth">
            <el-select
              v-model="createChallengeData.ImageID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="image in imageList"
                :key="image.ID"
                :label="image.Remark"
                :value="image.ID"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Attachment" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-select
                style="width: 50%"
                v-model="createChallengeData.AttachmentID"
                filterable
                placeholder="Select"
                >
                <el-option
                  v-for="attachment in attachmentList"
                  :key="attachment.ID"
                  :label="attachment.FileName"
                  :value="attachment.ID"
                >
                </el-option>
              </el-select>
              <el-upload
                style="width: 50%;"
                ref="upload"
                :limit="1"
                :auto-upload="false"
                :on-change="handleFileChange"
                :on-exceed="handleExceed"
              >
                <template #trigger>
                  <el-button style="margin-left: 20px">Select</el-button>
                </template>
                <div style="display: inline-flex">
                  <el-button @click="UploadAttachment" style="margin-left: 10px">Upload</el-button>
                </div>
              </el-upload>
            </div>
          </el-form-item>
          <el-form-item label="SpecifiedPorts" :label-width="labelWidth">
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
            <div
              v-if="createChallengeData.SpecifiedPorts.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
              >
              <div
                v-for="(port, index) in createChallengeData.SpecifiedPorts"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ port }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveCreateChallengeSpecifiedPort(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Commands" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.createChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddCreateChallengeCommand"
              >
                Add
              </el-button>
            </div>
            <div
              v-if="createChallengeData.Commands.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
            >
              <div
                v-for="(command, index) in createChallengeData.Commands"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ command }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveCreateChallengeCommand(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Flag" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                style="width: 270px"
                v-if="createChallengeFlagType === 'auto'"
                placeholder="auto"
                disabled
              />
              <el-input
                style="width: 270px"
                v-if="createChallengeFlagType === 'specify'"
                v-model="createChallengeData.Flag"
              />
              <div style="flex-grow: 1" v-if="createChallengeFlagType !== ''"></div>
              <el-radio-group v-model="createChallengeFlagType">
                <el-radio value="auto" style="margin: 0" border>auto</el-radio>
                <el-radio value="specify" style="margin: 0" border>specify</el-radio>
              </el-radio-group>
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
        width="600"
        @close="ClearUpdateChallengeForm"
        >
        <el-form :model=updateChallengeData>
          <el-form-item label="Title" :label-width="labelWidth">
            <el-input v-model="updateChallengeData.Title"/>
          </el-form-item>
          <el-form-item label="Description" :label-width="labelWidth">
            <el-input v-model="updateChallengeData.Description"/>
          </el-form-item>
          <el-form-item label="ChallengeClass" :label-width="labelWidth">
            <el-select
              v-model="updateChallengeData.ChallengeClassID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="challengeClass in challengeClassList"
                :key="challengeClass.ID"
                :label="challengeClass.Name"
                :value="challengeClass.ID"
                ></el-option>
              <template v-if="!challengeClassList.some(challengeClass => challengeClass.ID === updateChallengeData.ChallengeClassID)" slot="empty">
                <el-option
                  label="Deleted challengeClass"
                  :value="updateChallengeData.ChallengeClassID"
                  disabled
                />
              </template>
            </el-select>
          </el-form-item>
          <el-form-item label="Image" :label-width="labelWidth">
            <el-select
              v-model="updateChallengeData.ImageID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="image in imageList"
                :key="image.ID"
                :label="image.Remark"
                :value="image.ID"
                >
              </el-option>
              <template v-if="!imageList.some(image => image.ID === updateChallengeData.ImageID)" slot="empty">
                <el-option
                  label="Deleted image"
                  :value="updateChallengeData.ImageID"
                  disabled
                />
              </template>
            </el-select>
          </el-form-item>
          <el-form-item label="Attachment" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-select
                style="width: 50%"
                v-model="updateChallengeData.AttachmentID"
                filterable
                placeholder="Select"
                >
                <el-option
                  v-for="attachment in attachmentList"
                  :key="attachment.ID"
                  :label="attachment.FileName"
                  :value="attachment.ID"
                >
                </el-option>
                <template v-if="!attachmentList.some(attachment => attachment.ID === updateChallengeData.AttachmentID)" slot="empty">
                  <el-option
                    label="Deleted file"
                    :value="updateChallengeData.AttachmentID"
                    disabled
                  />
                </template>
              </el-select>
              <el-upload
                style="width: 50%;"
                ref="upload"
                :limit="1"
                :auto-upload="false"
                :on-change="handleFileChange"
                :on-exceed="handleExceed"
              >
                <template #trigger>
                  <el-button style="margin-left: 20px">Select</el-button>
                </template>
                <div style="display: inline-flex">
                  <el-button @click="UploadAttachment" style="margin-left: 10px">Upload</el-button>
                </div>
              </el-upload>
            </div>
          </el-form-item>
          <el-form-item label="SpecifiedPorts" :label-width="labelWidth">
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
            <div
              v-if="updateChallengeData.SpecifiedPorts.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
              >
              <div
                v-for="(port, index) in updateChallengeData.SpecifiedPorts"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ port }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveUpdateChallengeSpecifiedPort(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Commands" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.updateChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddUpdateChallengeCommand"
              >
                Add
              </el-button>
            </div>
            <div
              v-if="updateChallengeData.Commands.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
            >
              <div
                v-for="(command, index) in updateChallengeData.Commands"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ command }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveUpdateChallengeCommand(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  Delete
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Flag" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                style="width: 270px"
                v-if="updateChallengeFlagType === 'auto'"
                placeholder="auto"
                disabled
              />
              <el-input
                style="width: 270px"
                v-if="updateChallengeFlagType === 'specify'"
                v-model="updateChallengeData.Flag"
              />
              <div style="flex-grow: 1" v-if="updateChallengeFlagType !== ''"></div>
              <el-radio-group v-model="updateChallengeFlagType">
                <el-radio value="auto" style="margin: 0" border>auto</el-radio>
                <el-radio value="specify" style="margin: 0" border>specify</el-radio>
              </el-radio-group>
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
import challengeClassApi from "@/api/challengeClass.js";
import imageApi from "@/api/image.js";
import attachmentApi from "@/api/attachment.js";
import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      labelWidth: 120,
      challengeClassList: [],
      imageList: [],
      attachmentList: [],
      challengeList: [],
      challengeDetailVisible: false,
      challengeDetail: {
        "Title": "",
        "Description": "",
        "ChallengeClass": {},
        "Image": {},
        "Attachment": {},
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      },
      createChallengeFormVisible: false,
      createChallengeData: {
        "Title": "",
        "Description": "",
        "ChallengeClassID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      },
      createChallengePort: "",
      createChallengeProtocol: "",
      createChallengeCommand: "",
      createChallengeFlagType: "",
      updateChallengeFormVisible: false,
      updateChallengeData: {
        "ID": 0,
        "Title": "",
        "Description": "",
        "ChallengeClassID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      },
      updateChallengePort: "",
      updateChallengeProtocol: "",
      updateChallengeCommand: "",
      updateChallengeFlagType: "",
      deleteChallengeFormVisible: false,
      deleteChallengeData: {
        "ID": 0,
      },
      uploadAttachmentData: {
        "fileName": "",
        "file": null
      }
    }
  },
  methods: {
    GetChallengeCLassList() {
      challengeClassApi.GetChallengeClassList().then(res => {
        if (res.code === 0) {
          this.challengeClassList = res.data
          this.challengeClassList.push({
            "ID": 0,
            "Name": "null",
            "Order": 0
          })
          this.challengeClassList.sort((a, b) => a.Order - b.Order)
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
          this.imageList.push({
            "ID": 0,
            "Remark": "null"
          })
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetAttachmentList() {
      attachmentApi.GetAttachmentList().then(res => {
        if (res.code === 0) {
          this.attachmentList = res.data
          this.attachmentList.push({
            "ID": 0,
            "FileName": "null"
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
    OpenChallengeDetail(row) {
      this.challengeDetail = row
      this.challengeDetailVisible = true
    },
    ClearChallengeDetail() {
      this.challengeDetail = {
        "Title": "",
        "Description": "",
        "ChallengeClass": {},
        "Image": {},
        "Attachment": {},
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      }
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
      if (this.createChallengeFlagType === "auto") {
        this.createChallengeData.Flag = "auto"
      }
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
        "Title": "",
        "Description": "",
        "ChallengeClassID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      }
      this.createChallengePort = ""
      this.createChallengeProtocol = ""
      this.createChallengeCommand = ""
      this.createChallengeFlagType = ""
      this.ClearUploadAttachmentForm()
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
      if (this.updateChallengeFlagType === "auto") {
        this.updateChallengeData.Flag = "auto"
      }
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
        "Title": row.Title,
        "Description": row.Description,
        "ChallengeClassID": row.ChallengeClassID,
        "ImageID": row.ImageID,
        "AttachmentID": row.AttachmentID,
        "SpecifiedPorts": row.SpecifiedPorts === null ? [] : [...row.SpecifiedPorts],
        "Commands": row.Commands === null ? [] : [...row.Commands],
        "Flag": row.Flag
      }
      this.updateChallengeFlagType = row.Flag === "auto" ? "auto" : "specify"
      this.updateChallengeFormVisible = true
    },
    ClearUpdateChallengeForm() {
      this.updateChallengeData = {
        "ID": 0,
        "ImageID": 0,
        "Title": "",
        "Description": "",
        "ChallengeClassID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      }
      this.updateChallengePort = ""
      this.updateChallengeProtocol = ""
      this.updateChallengeCommand = ""
      this.updateChallengeFlagType = ""
      this.ClearUploadAttachmentForm()
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
    UploadAttachment() {
      attachmentApi.UploadAttachment(this.uploadAttachmentData.fileName, this.uploadAttachmentData.file).then(res => {
        if (res.code === 0) {
          if (this.createChallengeFormVisible) {
            this.createChallengeData.AttachmentID = res.data.ID
          } else if (this.updateChallengeFormVisible) {
            this.updateChallengeData.AttachmentID = res.data.ID
          }
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetAttachmentList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearUploadAttachmentForm() {
      this.uploadAttachmentData = {
        "fileName": "",
        "file": null
      }
      this.$refs.upload.clearFiles();
    },
    handleFileChange(file) {
      if (file) {
        this.uploadAttachmentData.fileName = file.name;
        this.uploadAttachmentData.file = file.raw;
      } else {
        this.uploadAttachmentData.fileName = '';
        this.uploadAttachmentData.file = null;
      }
    },
    handleExceed(files) {
      const uploadInstance = this.$refs.upload;
      uploadInstance.clearFiles();
      const file = files[0];
      uploadInstance.handleStart(file);
    }
  },
  mounted() {
    this.GetChallengeCLassList()
    this.GetImageList()
    this.GetAttachmentList()
    this.GetChallengeList()
  }
}
</script>
