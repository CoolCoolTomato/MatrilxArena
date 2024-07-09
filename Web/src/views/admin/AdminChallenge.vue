<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">{{ $t('AdminChallenge.ChallengeManager') }}</h2>
    </el-header>
    <el-main>
      <el-scrollbar>
        <el-button
          style="margin: 10px"
          @click="createChallengeFormVisible = true"
          type="primary"
          >
          {{ $t('AdminChallenge.Add') }}
        </el-button>
        <el-table
          :data="challengeList"
          table-layout="fixed"
          >
          <el-table-column prop="Title" :label="$t('AdminChallenge.Title')"/>
          <el-table-column prop="Description" :label="$t('AdminChallenge.Description')"/>
          <el-table-column :label="$t('AdminChallenge.Image')">
            <template  #default=scope>
              {{ scope.row.Image.Remark === "" ? $t('AdminChallenge.Null') : scope.row.Image.Remark}}
            </template>
          </el-table-column>
          <el-table-column fixed="right" :label="$t('AdminChallenge.Operations')" width="300px">
            <template #default=scope>
              <el-button
                @click="OpenChallengeDetail(scope.row)"
                >
                {{ $t('AdminChallenge.Detail') }}
              </el-button>
              <el-button
                @click="OpenUpdateChallengeForm(scope.row)"
                >
                {{ $t('AdminChallenge.Update') }}
              </el-button>
              <el-button
                @click="OpenDeleteChallengeForm(scope.row)"
                >
                {{ $t('AdminChallenge.Delete') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-scrollbar>
      <el-dialog
        v-model="challengeDetailVisible"
        :title="$t('AdminChallenge.ChallengeDetail')"
        width="500"
        @close="ClearChallengeDetail"
        >
        <el-card>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Title') }}: {{ challengeDetail.Title }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Description') }}: {{ challengeDetail.Description }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Category') }}: {{ challengeDetail.Category.Name === "" ? $t('AdminChallenge.Null') : challengeDetail.Category.Name }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Image') }}: {{ challengeDetail.Image.Remark === "" ? $t('AdminChallenge.Null') : challengeDetail.Image.Remark }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.RepoTags') }}: {{ challengeDetail.Image.RepoTags === "" ? $t('AdminChallenge.Null') : challengeDetail.Image.RepoTags }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Attachment') }}: {{ challengeDetail.Attachment.FileName === "" ? $t('AdminChallenge.Null') : challengeDetail.Attachment.FileName }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.SpecifiedPorts') }}: {{ challengeDetail.SpecifiedPorts }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Commands') }}: {{ challengeDetail.Commands }}</p>
          <p style="word-break: break-all;">{{ $t('AdminChallenge.Flag') }}: {{ challengeDetail.Flag }}</p>
        </el-card>
        <template #footer>
          <el-button @click="challengeDetailVisible = false">{{ $t('AdminChallenge.Close') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="createChallengeFormVisible"
        :title="$t('AdminChallenge.CreateChallenge')"
        width="600"
        @close="ClearCreateChallengeForm"
        >
        <el-form :model=createChallengeData>
          <el-form-item :label="$t('AdminChallenge.Title')" :label-width="labelWidth">
            <el-input v-model="createChallengeData.Title"/>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Description')" :label-width="labelWidth">
            <el-input v-model="createChallengeData.Description"/>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Category')" :label-width="labelWidth">
            <el-select
              v-model="createChallengeData.CategoryID"
              filterable
              :placeholder="$t('AdminChallenge.Select')"
              >
              <el-option
                v-for="category in categoryList"
                :key="category.ID"
                :label="category.Name"
                :value="category.ID"
                ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Image')" :label-width="labelWidth">
            <el-select
              v-model="createChallengeData.ImageID"
              filterable
              :placeholder="$t('AdminChallenge.Select')"
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
          <el-form-item :label="$t('AdminChallenge.Attachment')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-select
                style="width: 50%"
                v-model="createChallengeData.AttachmentID"
                filterable
                :placeholder="$t('AdminChallenge.Select')"
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
                  <el-button style="margin-left: 20px">{{ $t('AdminChallenge.SelectFile') }}</el-button>
                </template>
                <div style="display: inline-flex">
                  <el-button
                    @click="UploadAttachment"
                    style="margin-left: 10px"
                    type="primary"
                    >
                    {{ $t('AdminChallenge.Upload') }}
                  </el-button>
                </div>
              </el-upload>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.SpecifiedPorts')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                v-model="createChallengePort"
              >
                <template #append>
                  <el-select
                    v-model="createChallengeProtocol"
                    :placeholder="$t('AdminChallenge.Select')"
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
                {{ $t('AdminChallenge.Add') }}
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
                  {{ $t('AdminChallenge.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Commands')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.createChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddCreateChallengeCommand"
              >
                {{ $t('AdminChallenge.Add') }}
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
                  {{ $t('AdminChallenge.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Flag')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                style="width: 245px"
                v-if="createChallengeFlagType === 'auto'"
                :placeholder="$t('AdminChallenge.Auto')"
                disabled
              />
              <el-input
                style="width: 245px"
                v-if="createChallengeFlagType === 'specify'"
                v-model="createChallengeData.Flag"
              />
              <div style="flex-grow: 1" v-if="createChallengeFlagType !== ''"></div>
              <el-radio-group v-model="createChallengeFlagType">
                <el-radio value="auto" style="margin: 0" border>{{ $t('AdminChallenge.Auto') }}</el-radio>
                <el-radio value="specify" style="margin: 0" border>{{ $t('AdminChallenge.Specify') }}</el-radio>
              </el-radio-group>
            </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateChallenge"
            type="primary"
            >
            {{ $t('AdminChallenge.Submit') }}
          </el-button>
          <el-button
            @click="createChallengeFormVisible = false"
            >
            {{ $t('AdminChallenge.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateChallengeFormVisible"
        :title="$t('AdminChallenge.UpdateChallenge')"
        width="600"
        @close="ClearUpdateChallengeForm"
        >
        <el-form :model=updateChallengeData>
          <el-form-item :label="$t('AdminChallenge.Title')" :label-width="labelWidth">
            <el-input v-model="updateChallengeData.Title"/>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Description')" :label-width="labelWidth">
            <el-input v-model="updateChallengeData.Description"/>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Category')" :label-width="labelWidth">
            <el-select
              v-model="updateChallengeData.CategoryID"
              filterable
              :placeholder="$t('AdminChallenge.Select')"
              >
              <el-option
                v-for="category in categoryList"
                :key="category.ID"
                :label="category.Name"
                :value="category.ID"
                ></el-option>
              <template v-if="!categoryList.some(category => category.ID === updateChallengeData.CategoryID)" slot="empty">
                <el-option
                  :label="$t('AdminChallenge.DeletedCategory')"
                  :value="updateChallengeData.CategoryID"
                  disabled
                />
              </template>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Image')" :label-width="labelWidth">
            <el-select
              v-model="updateChallengeData.ImageID"
              filterable
              :placeholder="$t('AdminChallenge.Select')"
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
                  :label="$t('AdminChallenge.DeletedImage')"
                  :value="updateChallengeData.ImageID"
                  disabled
                />
              </template>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Attachment')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-select
                style="width: 50%"
                v-model="updateChallengeData.AttachmentID"
                filterable
                :placeholder="$t('AdminChallenge.Select')"
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
                    :label="$t('AdminChallenge.DeletedFile')"
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
                  <el-button style="margin-left: 20px">{{ $t('AdminChallenge.SelectFile') }}</el-button>
                </template>
                <div style="display: inline-flex">
                  <el-button
                    @click="UploadAttachment"
                    style="margin-left: 10px"
                    type="primary"
                    >
                    {{ $t('AdminChallenge.Upload') }}
                  </el-button>
                </div>
              </el-upload>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.SpecifiedPorts')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                v-model="updateChallengePort"
              >
                <template #append>
                  <el-select
                    v-model="updateChallengeProtocol"
                    :placeholder="$t('AdminChallenge.Select')"
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
                {{ $t('AdminChallenge.Add') }}
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
                  {{ $t('AdminChallenge.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Commands')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.updateChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddUpdateChallengeCommand"
              >
                {{ $t('AdminChallenge.Add') }}
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
                  {{ $t('AdminChallenge.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminChallenge.Flag')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                style="width: 245px"
                v-if="updateChallengeFlagType === 'auto'"
                :placeholder="$t('AdminChallenge.Auto')"
                disabled
              />
              <el-input
                style="width: 245px"
                v-if="updateChallengeFlagType === 'specify'"
                v-model="updateChallengeData.Flag"
              />
              <div style="flex-grow: 1" v-if="updateChallengeFlagType !== ''"></div>
              <el-radio-group v-model="updateChallengeFlagType">
                <el-radio value="auto" style="margin: 0" border>{{ $t('AdminChallenge.Auto') }}</el-radio>
                <el-radio value="specify" style="margin: 0" border>{{ $t('AdminChallenge.Specify') }}</el-radio>
              </el-radio-group>
            </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateChallenge"
            type="primary"
            >
            {{ $t('AdminChallenge.Submit') }}
          </el-button>
          <el-button
            @click="updateChallengeFormVisible = false"
            >
            {{ $t('AdminChallenge.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteChallengeFormVisible"
        :title="$t('AdminChallenge.DeleteChallenge')"
        width="500"
        @close="ClearDeleteChallengeForm"
        >
        <el-text>{{ $t('AdminChallenge.AreYouConfirmToDeleteTheChallenge') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteChallenge"
            type="primary"
            >
            {{ $t('AdminChallenge.Confirm') }}
          </el-button>
          <el-button
            @click="deleteChallengeFormVisible = false"
            >
            {{ $t('AdminChallenge.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import challengeApi from "@/api/challenge.js"
import categoryApi from "@/api/category.js"
import imageApi from "@/api/image.js"
import attachmentApi from "@/api/attachment.js"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      labelWidth: 120,
      categoryList: [],
      imageList: [],
      attachmentList: [],
      challengeList: [],
      challengeDetailVisible: false,
      challengeDetail: {
        "Title": "",
        "Description": "",
        "Category": {},
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
        "CategoryID": 0,
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
        "CategoryID": 0,
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
    GetChallengeCategoryList() {
      categoryApi.GetCategoryList().then(res => {
        if (res.code === 0) {
          this.categoryList = res.data
          this.categoryList.push({
            "ID": 0,
            "Name": "null",
            "Order": 0
          })
          this.categoryList.sort((a, b) => a.Order - b.Order)
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
        "Category": {},
        "Image": {},
        "Attachment": {},
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      }
    },
    AddCreateChallengeSpecifiedPort() {
      if (this.createChallengePort === "") {
        ElMessage.error(this.t('AdminChallenge.PleaseInputPort'))
        return
      }
      if (this.createChallengeProtocol === "") {
        ElMessage.error(this.t('AdminChallenge.PleaseInputProtocol'))
        return
      }
      let portProtocol = this.createChallengePort + "/" + this.createChallengeProtocol
      if (this.createChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error(this.t('AdminChallenge.DuplicateSpecifiedPort'))
        return
      }
      this.createChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveCreateChallengeSpecifiedPort(index) {
      this.createChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddCreateChallengeCommand() {
      if (this.createChallengeCommand === "") {
        ElMessage.error(this.t('AdminChallenge.PleaseInputCommand'))
        return
      }
      if (this.createChallengeData.Commands.includes(this.createChallengeCommand)) {
        ElMessage.error(this.t('AdminChallenge.DuplicateCommand'))
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
        "CategoryID": 0,
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
        ElMessage.error(this.t('AdminChallenge.PleaseInputPort'))
        return
      }
      if (this.updateChallengeProtocol === "") {
        ElMessage.error(this.t('AdminChallenge.PleaseInputProtocol'))
        return
      }
      let portProtocol = this.updateChallengePort + "/" + this.updateChallengeProtocol
      if (this.updateChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error(this.t('AdminChallenge.DuplicateSpecifiedPort'))
        return
      }
      this.updateChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveUpdateChallengeSpecifiedPort(index) {
      this.updateChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddUpdateChallengeCommand() {
      if (this.updateChallengeCommand === "") {
        ElMessage.error(this.t('AdminChallenge.PleaseInputCommand'))
        return
      }
      if (this.updateChallengeData.Commands.includes(this.updateChallengeCommand)) {
        ElMessage.error(this.t('AdminChallenge.DuplicateCommand'))
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
        "CategoryID": row.CategoryID,
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
        "CategoryID": 0,
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
    this.GetChallengeCategoryList()
    this.GetImageList()
    this.GetAttachmentList()
    this.GetChallengeList()
  }
}
</script>
