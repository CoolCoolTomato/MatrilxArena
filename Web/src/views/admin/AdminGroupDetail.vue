<template>
  <el-header>
    <div style="display: flex; align-items: center;">
      <h2 style="color: var(--el-text-color-primary);">{{ $t('AdminGroupDetail.GroupDetail') }}</h2>
      <div style="flex-grow: 1;"></div>
      <div style="margin-right: 50px;">
        <el-button
          style="margin: 10px;"
          @click="goBack"
        >
          {{ $t('AdminGroupDetail.Back') }}
        </el-button>
      </div>
    </div>
  </el-header>
  <el-main>
    <el-tabs
      v-model="activeTab"
      class="el-tabs-custom"
      >
      <el-tab-pane :label="$t('AdminGroupDetail.Challenges')" name="challenges">
        <div style="height: 50px; display: flex; align-items: center;">
          <el-button
            @click="OpenCreateGroupChallengeForm"
            type="primary"
            >
            {{ $t('AdminGroupDetail.Add') }}
          </el-button>
        </div>
        <el-table
          :data="groupChallengeQueryList"
          table-layout="fixed"
          height="calc(100% - 50px)"
          >
          <el-table-column prop="Title" :label="$t('AdminGroupDetail.Title')"/>
          <el-table-column prop="Description" :label="$t('AdminGroupDetail.Description')"/>
          <el-table-column :label="$t('AdminGroupDetail.Category')">
            <template  #default=scope>
              {{ scope.row.Category.Name === "" ? $t('AdminGroupDetail.Null') : scope.row.Category.Name}}
            </template>
          </el-table-column>
          <el-table-column :label="$t('AdminGroupDetail.Image')">
            <template  #default=scope>
              {{ scope.row.Image.Remark === "" ? $t('AdminGroupDetail.Null') : scope.row.Image.Remark}}
            </template>
          </el-table-column>
          <el-table-column fixed="right" :label="$t('AdminGroupDetail.Operations')" width="300px">
            <template #default=scope>
              <el-button
                @click="OpenGroupChallengeDetail(scope.row)"
                >
                {{ $t('AdminGroupDetail.Detail') }}
              </el-button>
              <el-button
                @click="OpenUpdateGroupChallengeForm(scope.row)"
                >
                {{ $t('AdminGroupDetail.Update') }}
              </el-button>
              <el-button
                @click="OpenDeleteChallengeForm(scope.row)"
                >
                {{ $t('AdminGroupDetail.Delete') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane :label="$t('AdminGroupDetail.Users')" name="users">
        <div style="height: 50px; display: flex; align-items: center;">
          <el-button
            @click="OpenAddGroupUserForm"
            type="primary"
            >
            {{ $t('AdminGroupDetail.Add') }}
          </el-button>
        </div>
        <el-table
          :data="group.Users"
          table-layout="fixed"
          height="calc(100% - 50px)"
          >
          <el-table-column prop="Username" :label="$t('AdminGroupDetail.Username')"/>
          <el-table-column prop="Email" :label="$t('AdminGroupDetail.Email')"/>
          <el-table-column  :label="$t('AdminGroupDetail.SolvedChallenges')">
            <template #default="scope">
              {{ scope.row.GroupChallenges.length }}
            </template>
          </el-table-column>
          <el-table-column fixed="right" :label="$t('AdminGroupDetail.Operations')" width="230">
            <template #default=scope>
              <el-button
                @click="OpenGroupUserDetailForm(scope.row)"
                >
                {{ $t('AdminGroupDetail.Detail') }}
              </el-button>
              <el-button
                @click="OpenRemoveGroupUserForm(scope.row)"
                >
                {{ $t('AdminGroupDetail.Remove') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
    <el-dialog
      v-model="groupChallengeDetailVisible"
      :title="$t('AdminGroupDetail.ChallengeDetail')"
      width="500"
      @close="ClearGroupChallengeDetail"
      >
      <el-card>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Title') }}: {{ groupChallengeDetail.Title }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Description') }}: {{ groupChallengeDetail.Description }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Category') }}: {{ groupChallengeDetail.Category.Name === "" ? $t('AdminGroupDetail.Null') : groupChallengeDetail.Category.Name }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Image') }}: {{ groupChallengeDetail.Image.Remark === "" ? $t('AdminGroupDetail.Null') : groupChallengeDetail.Image.Remark }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.RepoTags') }}: {{ groupChallengeDetail.Image.RepoTags === "" ? $t('AdminGroupDetail.Null') : groupChallengeDetail.Image.RepoTags }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Attachment') }}: {{ groupChallengeDetail.Attachment.FileName === "" ? $t('AdminGroupDetail.Null') : groupChallengeDetail.Attachment.FileName }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.SpecifiedPorts') }}: {{ groupChallengeDetail.SpecifiedPorts }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Commands') }}: {{ groupChallengeDetail.Commands }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Flag') }}: {{ groupChallengeDetail.Flag }}</p>
      </el-card>
      <template #footer>
        <el-button @click="groupChallengeDetailVisible = false">{{ $t('AdminGroupDetail.Close') }}</el-button>
      </template>
    </el-dialog>
    <el-dialog
      v-model="createGroupChallengeFormVisible"
      :title="$t('AdminGroupDetail.CreateChallenge')"
      width="600"
      @close="ClearCreateGroupChallengeForm"
      >
      <el-form :model=createGroupChallengeData>
        <el-form-item :label="$t('AdminGroupDetail.Title')" :label-width="labelWidth">
          <el-input v-model="createGroupChallengeData.Title"/>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Description')" :label-width="labelWidth">
          <el-input v-model="createGroupChallengeData.Description"/>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Category')" :label-width="labelWidth">
          <el-select
            v-model="createGroupChallengeData.CategoryID"
            filterable
            :placeholder="$t('AdminGroupDetail.Select')"
            >
            <el-option
              v-for="category in categoryList"
              :key="category.ID"
              :label="category.Name"
              :value="category.ID"
              ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Image')" :label-width="labelWidth">
          <el-select
            v-model="createGroupChallengeData.ImageID"
            filterable
            :placeholder="$t('AdminGroupDetail.Select')"
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
        <el-form-item :label="$t('AdminGroupDetail.Attachment')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-select
              style="width: 50%"
              v-model="createGroupChallengeData.AttachmentID"
              filterable
              :placeholder="$t('AdminGroupDetail.Select')"
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
                <el-button style="margin-left: 20px">{{ $t('AdminGroupDetail.SelectFile') }}</el-button>
              </template>
              <div style="display: inline-flex">
                <el-button
                  @click="UploadAttachment"
                  style="margin-left: 10px"
                  type="primary"
                  >
                  {{ $t('AdminGroupDetail.Upload') }}
                </el-button>
              </div>
            </el-upload>
          </div>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.SpecifiedPorts')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-input
              v-model="createGroupChallengePort"
            >
              <template #append>
                <el-select
                  v-model="createGroupChallengeProtocol"
                  :placeholder="$t('AdminGroupDetail.Select')"
                  style="width: 100px"
                >
                  <el-option label="tcp" value="tcp"/>
                  <el-option label="udp" value="udp"/>
                </el-select>
              </template>
            </el-input>
            <el-button
              style="margin-left: 20px;"
              @click="AddCreateGroupChallengeSpecifiedPort"
            >
              {{ $t('AdminGroupDetail.Add') }}
            </el-button>
          </div>
          <div
            v-if="createGroupChallengeData.SpecifiedPorts.length > 0"
            style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
            >
            <div
              v-for="(port, index) in createGroupChallengeData.SpecifiedPorts"
              :key="index"
              style="display: flex; align-items: center; width: 100%; position: relative"
            >
              <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ port }}</el-text>
              <div style="flex-grow: 1"></div>
              <el-button
                size="small"
                @click="RemoveCreateGroupChallengeSpecifiedPort(index)"
                style="margin: 3px 10px 3px 3px;"
              >
                {{ $t('AdminGroupDetail.Delete') }}
              </el-button>
            </div>
          </div>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Commands')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-input v-model="this.createGroupChallengeCommand" />
            <el-button
              style="margin-left: 20px;"
              @click="AddCreateGroupChallengeCommand"
            >
              {{ $t('AdminGroupDetail.Add') }}
            </el-button>
          </div>
          <div
            v-if="createGroupChallengeData.Commands.length > 0"
            style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
          >
            <div
              v-for="(command, index) in createGroupChallengeData.Commands"
              :key="index"
              style="display: flex; align-items: center; width: 100%; position: relative"
            >
              <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ command }}</el-text>
              <div style="flex-grow: 1"></div>
              <el-button
                size="small"
                @click="RemoveCreateGroupChallengeCommand(index)"
                style="margin: 3px 10px 3px 3px;"
              >
                {{ $t('AdminGroupDetail.Delete') }}
              </el-button>
            </div>
          </div>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Flag')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-input
              style="width: 245px"
              v-if="createGroupChallengeFlagType === 'auto'"
              :placeholder="$t('AdminGroupDetail.Auto')"
              disabled
            />
            <el-input
              style="width: 245px"
              v-if="createGroupChallengeFlagType === 'specify'"
              v-model="createGroupChallengeData.Flag"
            />
            <div style="flex-grow: 1" v-if="createGroupChallengeFlagType !== ''"></div>
            <el-radio-group v-model="createGroupChallengeFlagType">
              <el-radio value="auto" style="margin: 0" border>{{ $t('AdminGroupDetail.Auto') }}</el-radio>
              <el-radio value="specify" style="margin: 0" border>{{ $t('AdminGroupDetail.Specify') }}</el-radio>
            </el-radio-group>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          @click="CreateGroupChallenge"
          type="primary"
          >
          {{ $t('AdminGroupDetail.Submit') }}
        </el-button>
        <el-button
          @click="createGroupChallengeFormVisible = false"
          >
          {{ $t('AdminGroupDetail.Cancel') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog
      v-model="updateGroupChallengeFormVisible"
      :title="$t('AdminGroupDetail.UpdateChallenge')"
      width="600"
      @close="ClearUpdateGroupChallengeForm"
      >
      <el-form :model=updateGroupChallengeData>
        <el-form-item :label="$t('AdminGroupDetail.Title')" :label-width="labelWidth">
          <el-input v-model="updateGroupChallengeData.Title"/>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Description')" :label-width="labelWidth">
          <el-input v-model="updateGroupChallengeData.Description"/>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Category')" :label-width="labelWidth">
          <el-select
            v-model="updateGroupChallengeData.CategoryID"
            filterable
            :placeholder="$t('AdminGroupDetail.Select')"
            >
            <el-option
              v-for="category in categoryList"
              :key="category.ID"
              :label="category.Name"
              :value="category.ID"
              ></el-option>
            <template v-if="!categoryList.some(category => category.ID === updateGroupChallengeData.CategoryID)" slot="empty">
              <el-option
                :label="$t('AdminGroupDetail.DeletedCategory')"
                :value="updateGroupChallengeData.CategoryID"
                disabled
              />
            </template>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Image')" :label-width="labelWidth">
          <el-select
            v-model="updateGroupChallengeData.ImageID"
            filterable
            :placeholder="$t('AdminGroupDetail.Select')"
            >
            <el-option
              v-for="image in imageList"
              :key="image.ID"
              :label="image.Remark"
              :value="image.ID"
              >
            </el-option>
            <template v-if="!imageList.some(image => image.ID === updateGroupChallengeData.ImageID)" slot="empty">
              <el-option
                :label="$t('AdminGroupDetail.DeletedImage')"
                :value="updateGroupChallengeData.ImageID"
                disabled
              />
            </template>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Attachment')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-select
              style="width: 50%"
              v-model="updateGroupChallengeData.AttachmentID"
              filterable
              :placeholder="$t('AdminGroupDetail.Select')"
              >
              <el-option
                v-for="attachment in attachmentList"
                :key="attachment.ID"
                :label="attachment.FileName"
                :value="attachment.ID"
              >
              </el-option>
              <template v-if="!attachmentList.some(attachment => attachment.ID === updateGroupChallengeData.AttachmentID)" slot="empty">
                <el-option
                  :label="$t('AdminGroupDetail.DeletedFile')"
                  :value="updateGroupChallengeData.AttachmentID"
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
                <el-button style="margin-left: 20px">{{ $t('AdminGroupDetail.SelectFile') }}</el-button>
              </template>
              <div style="display: inline-flex">
                <el-button
                  @click="UploadAttachment"
                  style="margin-left: 10px"
                  type="primary"
                  >
                  {{ $t('AdminGroupDetail.Upload') }}
                </el-button>
              </div>
            </el-upload>
          </div>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.SpecifiedPorts')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-input
              v-model="updateGroupChallengePort"
            >
              <template #append>
                <el-select
                  v-model="updateGroupChallengeProtocol"
                  :placeholder="$t('AdminGroupDetail.Select')"
                  style="width: 100px"
                >
                  <el-option label="tcp" value="tcp"/>
                  <el-option label="udp" value="udp"/>
                </el-select>
              </template>
            </el-input>
            <el-button
              style="margin-left: 20px;"
              @click="AddUpdateGroupChallengeSpecifiedPort"
            >
              {{ $t('AdminGroupDetail.Add') }}
            </el-button>
          </div>
          <div
            v-if="updateGroupChallengeData.SpecifiedPorts.length > 0"
            style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
            >
            <div
              v-for="(port, index) in updateGroupChallengeData.SpecifiedPorts"
              :key="index"
              style="display: flex; align-items: center; width: 100%; position: relative"
            >
              <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ port }}</el-text>
              <div style="flex-grow: 1"></div>
              <el-button
                size="small"
                @click="RemoveUpdateGroupChallengeSpecifiedPort(index)"
                style="margin: 3px 10px 3px 3px;"
              >
                {{ $t('AdminGroupDetail.Delete') }}
              </el-button>
            </div>
          </div>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Commands')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-input v-model="this.updateGroupChallengeCommand" />
            <el-button
              style="margin-left: 20px;"
              @click="AddUpdateGroupChallengeCommand"
            >
              {{ $t('AdminGroupDetail.Add') }}
            </el-button>
          </div>
          <div
            v-if="updateGroupChallengeData.Commands.length > 0"
            style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
          >
            <div
              v-for="(command, index) in updateGroupChallengeData.Commands"
              :key="index"
              style="display: flex; align-items: center; width: 100%; position: relative"
            >
              <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ command }}</el-text>
              <div style="flex-grow: 1"></div>
              <el-button
                size="small"
                @click="RemoveUpdateGroupChallengeCommand(index)"
                style="margin: 3px 10px 3px 3px;"
              >
                {{ $t('AdminGroupDetail.Delete') }}
              </el-button>
            </div>
          </div>
        </el-form-item>
        <el-form-item :label="$t('AdminGroupDetail.Flag')" :label-width="labelWidth">
          <div style="display: flex; width: 100%;">
            <el-input
              style="width: 245px"
              v-if="updateGroupChallengeFlagType === 'auto'"
              :placeholder="$t('AdminGroupDetail.Auto')"
              disabled
            />
            <el-input
              style="width: 245px"
              v-if="updateGroupChallengeFlagType === 'specify'"
              v-model="updateGroupChallengeData.Flag"
            />
            <div style="flex-grow: 1" v-if="updateGroupChallengeFlagType !== ''"></div>
            <el-radio-group v-model="updateGroupChallengeFlagType">
              <el-radio value="auto" style="margin: 0" border>{{ $t('AdminGroupDetail.Auto') }}</el-radio>
              <el-radio value="specify" style="margin: 0" border>{{ $t('AdminGroupDetail.Specify') }}</el-radio>
            </el-radio-group>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          @click="UpdateGroupChallenge"
          type="primary"
          >
          {{ $t('AdminGroupDetail.Submit') }}
        </el-button>
        <el-button
          @click="updateGroupChallengeFormVisible = false"
          >
          {{ $t('AdminGroupDetail.Cancel') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog
      v-model="deleteGroupChallengeFormVisible"
      :title="$t('AdminGroupDetail.DeleteChallenge')"
      width="500"
      @close="ClearDeleteGroupChallengeForm"
      >
      <el-text>{{ $t('AdminGroupDetail.AreYouConfirmToDeleteTheChallenge') }}</el-text>
      <template #footer>
        <el-button
          @click="DeleteGroupChallenge"
          type="primary"
          >
          {{ $t('AdminGroupDetail.Confirm') }}
        </el-button>
        <el-button
          @click="deleteGroupChallengeFormVisible = false"
          >
          {{ $t('AdminGroupDetail.Cancel') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog
      v-model="addGroupUserFormVisible"
      :title="$t('AdminGroupDetail.AddUser')"
      width="500"
      @close="ClearAddGroupUserForm"
      >
      <el-form :model="addGroupUserData">
        <el-form-item>
          <el-select
            v-model="addGroupUserData.UserID"
            filterable
            :placeholder="$t('AdminGroupDetail.Select')"
            >
            <el-option
              v-for="user in userList"
              :key="user.ID"
              :label="user.Username"
              :value="user.ID"
              >
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          @click="AddGroupUser"
          type="primary"
          >
          {{ $t('AdminGroupDetail.Submit') }}
        </el-button>
        <el-button
          @click="addGroupUserFormVisible = false"
          >
          {{ $t('AdminGroupDetail.Cancel') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog
      v-model="removeGroupUserFormVisible"
      :title="$t('AdminGroupDetail.RemoveUser')"
      width="500"
      @close="ClearRemoveGroupUserForm"
      >
      <el-text>{{ $t('AdminGroupDetail.AreYouConfirmToRemoveTheUser') }}</el-text>
      <template #footer>
        <el-button
          @click="RemoveGroupUser"
          type="primary"
          >
          {{ $t('AdminGroupDetail.Confirm') }}
        </el-button>
        <el-button
          @click="removeGroupUserFormVisible = false"
          >
          {{ $t('AdminGroupDetail.Cancel') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog
      v-model="groupUserDetailVisible"
      :title="$t('AdminGroupDetail.UserDetail')"
      width="500"
      @close="ClearGroupUserDetail"
      >
      <el-card>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Username') }}: {{ groupUserDetail.Username }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.Email') }}: {{ groupUserDetail.Email }}</p>
        <p style="word-break: break-all;">{{ $t('AdminGroupDetail.SolvedChallenges') }}: </p>
        <div v-for="groupChallenge in groupUserDetail.GroupChallenges">
          <el-text>{{ groupChallenge.Title }}</el-text>
        </div>
      </el-card>
      <template #footer>
        <el-button @click="groupUserDetailVisible = false">{{ $t('AdminGroupDetail.Close') }}</el-button>
      </template>
    </el-dialog>
  </el-main>
</template>
<script>
import groupApi from "@/api/group.js"
import groupChallengeApi from "@/api/groupChallenge.js"
import categoryApi from "@/api/category.js"
import imageApi from "@/api/image.js"
import attachmentApi from "@/api/attachment.js"
import userApi from "@/api/user.js"
import {ElMessage} from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      activeTab: "challenges",
      group: {
        "ID": 0,
      },
      labelWidth: 120,
      categoryList: [],
      imageList: [],
      attachmentList: [],
      groupChallengeQueryCategoryID: 0,
      groupChallengeQueryTitle: "",
      groupChallengeQueryList: [],
      groupChallengeDetailVisible: false,
      groupChallengeDetail: {
        "Title": "",
        "Description": "",
        "Category": {},
        "Image": {},
        "Attachment": {},
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": ""
      },
      createGroupChallengeFormVisible: false,
      createGroupChallengeData: {
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "GroupID": 0
      },
      createGroupChallengePort: "",
      createGroupChallengeProtocol: "",
      createGroupChallengeCommand: "",
      createGroupChallengeFlagType: "",
      updateGroupChallengeFormVisible: false,
      updateGroupChallengeData: {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "GroupID": 0
      },
      updateGroupChallengePort: "",
      updateGroupChallengeProtocol: "",
      updateGroupChallengeCommand: "",
      updateGroupChallengeFlagType: "",
      deleteGroupChallengeFormVisible: false,
      deleteGroupChallengeData: {
        "ID": 0,
      },
      uploadAttachmentData: {
        "fileName": "",
        "file": null
      },
      userList: [],
      groupUserDetailVisible: false,
      groupUserDetail: {
        "Username": "",
        "Email": "",
        "GroupChallenges": []
      },
      addGroupUserFormVisible: false,
      addGroupUserData: {
        "GroupID": 0,
        "UserID": null
      },
      removeGroupUserFormVisible: false,
      removeGroupUserData: {
        "GroupID": 0,
        "UserID": null
      },
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1);
    },
    GetGroupChallengeCategoryList() {
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
    GetGroup() {
      groupApi.GetGroup(this.group).then(res => {
        if (res.code === 0) {
          this.group = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetGroupChallengeList() {
      groupChallengeApi.GetGroupChallengeListByQuery({
        "CategoryID": this.groupChallengeQueryCategoryID,
        "Title": this.groupChallengeQueryTitle,
        "GroupID": this.group.ID
      }).then(res => {
        if (res.code === 0) {
          this.groupChallengeQueryList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenGroupChallengeDetail(row) {
      this.groupChallengeDetail = row
      this.groupChallengeDetailVisible = true
    },
    ClearGroupChallengeDetail() {
      this.groupChallengeDetail = {
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
    AddCreateGroupChallengeSpecifiedPort() {
      if (this.createGroupChallengePort === "") {
        ElMessage.error(this.t('AdminGroupDetail.PleaseInputPort'))
        return
      }
      if (this.createGroupChallengeProtocol === "") {
        ElMessage.error(this.t('AdminGroupDetail.PleaseInputProtocol'))
        return
      }
      let portProtocol = this.createGroupChallengePort + "/" + this.createGroupChallengeProtocol
      if (this.createGroupChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error(this.t('AdminGroupDetail.DuplicateSpecifiedPort'))
        return
      }
      this.createGroupChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveCreateGroupChallengeSpecifiedPort(index) {
      this.createGroupChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddCreateGroupChallengeCommand() {
      if (this.createGroupChallengeCommand === "") {
        ElMessage.error(this.t('AdminGroupDetail.PleaseInputCommand'))
        return
      }
      if (this.createGroupChallengeData.Commands.includes(this.createGroupChallengeCommand)) {
        ElMessage.error(this.t('AdminGroupDetail.DuplicateCommand'))
        return
      }
      this.createGroupChallengeData.Commands.push(this.createGroupChallengeCommand)
    },
    RemoveCreateGroupChallengeCommand(index) {
      this.createGroupChallengeData.Commands.splice(index, 1);
    },
    CreateGroupChallenge() {
      if (this.createGroupChallengeFlagType === "auto") {
        this.createGroupChallengeData.Flag = "auto"
      }
      groupChallengeApi.CreateGroupChallenge(this.createGroupChallengeData).then(res => {
        if (res.code === 0) {
          this.createGroupChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetGroupChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCreateGroupChallengeForm() {
      this.createGroupChallengeData.GroupID = this.group.ID
      this.createGroupChallengeFormVisible = true
    },
    ClearCreateGroupChallengeForm() {
      this.createGroupChallengeData = {
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "GroupID": 0
      }
      this.createGroupChallengePort = ""
      this.createGroupChallengeProtocol = ""
      this.createGroupChallengeCommand = ""
      this.createGroupChallengeFlagType = ""
      this.ClearUploadAttachmentForm()
    },
    AddUpdateGroupChallengeSpecifiedPort() {
      if (this.updateGroupChallengePort === "") {
        ElMessage.error(this.t('AdminGroupDetail.PleaseInputPort'))
        return
      }
      if (this.updateGroupChallengeProtocol === "") {
        ElMessage.error(this.t('AdminGroupDetail.PleaseInputProtocol'))
        return
      }
      let portProtocol = this.updateGroupChallengePort + "/" + this.updateGroupChallengeProtocol
      if (this.updateGroupChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error(this.t('AdminGroupDetail.DuplicateSpecifiedPort'))
        return
      }
      this.updateGroupChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveUpdateGroupChallengeSpecifiedPort(index) {
      this.updateGroupChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddUpdateGroupChallengeCommand() {
      if (this.updateGroupChallengeCommand === "") {
        ElMessage.error(this.t('AdminGroupDetail.PleaseInputCommand'))
        return
      }
      if (this.updateGroupChallengeData.Commands.includes(this.updateGroupChallengeCommand)) {
        ElMessage.error(this.t('AdminGroupDetail.DuplicateCommand'))
        return
      }
      this.updateGroupChallengeData.Commands.push(this.updateGroupChallengeCommand)
    },
    RemoveUpdateGroupChallengeCommand(index) {
      this.updateGroupChallengeData.Commands.splice(index, 1);
    },
    UpdateGroupChallenge() {
      if (this.updateGroupChallengeFlagType === "auto") {
        this.updateGroupChallengeData.Flag = "auto"
      }
      groupChallengeApi.UpdateGroupChallenge(this.updateGroupChallengeData).then(res => {
        if (res.code === 0) {
          this.updateGroupChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetGroupChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateGroupChallengeForm(row) {
      this.updateGroupChallengeData = {
        "ID": row.ID,
        "Title": row.Title,
        "Description": row.Description,
        "CategoryID": row.CategoryID,
        "ImageID": row.ImageID,
        "AttachmentID": row.AttachmentID,
        "SpecifiedPorts": row.SpecifiedPorts === null ? [] : [...row.SpecifiedPorts],
        "Commands": row.Commands === null ? [] : [...row.Commands],
        "Flag": row.Flag,
        "GroupID": row.GroupID
      }
      this.updateGroupChallengeFlagType = row.Flag === "auto" ? "auto" : "specify"
      this.updateGroupChallengeFormVisible = true
    },
    ClearUpdateGroupChallengeForm() {
      this.updateGroupChallengeData = {
        "ID": 0,
        "ImageID": 0,
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "GroupID": 0
      }
      this.updateGroupChallengePort = ""
      this.updateGroupChallengeProtocol = ""
      this.updateGroupChallengeCommand = ""
      this.updateGroupChallengeFlagType = ""
      this.ClearUploadAttachmentForm()
    },
    DeleteGroupChallenge() {
      groupChallengeApi.DeleteGroupChallenge(this.deleteGroupChallengeData).then(res => {
        if (res.code === 0) {
          this.deleteGroupChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetGroupChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteChallengeForm(row) {
      this.deleteGroupChallengeData = {
        "ID": row.ID,
      }
      this.deleteGroupChallengeFormVisible = true
    },
    ClearDeleteGroupChallengeForm() {
      this.deleteGroupChallengeData = {
        "ID": 0,
      }
    },
    UploadAttachment() {
      attachmentApi.UploadAttachment(this.uploadAttachmentData.fileName, this.uploadAttachmentData.file).then(res => {
        if (res.code === 0) {
          if (this.createGroupChallengeFormVisible) {
            this.createGroupChallengeData.AttachmentID = res.data.ID
          } else if (this.updateGroupChallengeFormVisible) {
            this.updateGroupChallengeData.AttachmentID = res.data.ID
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
        this.uploadAttachmentData.fileName = file.name
        this.uploadAttachmentData.file = file.raw
      } else {
        this.uploadAttachmentData.fileName = ''
        this.uploadAttachmentData.file = null
      }
    },
    handleExceed(files) {
      const uploadInstance = this.$refs.upload
      uploadInstance.clearFiles()
      const file = files[0]
      uploadInstance.handleStart(file)
    },

    GetUserList() {
      userApi.GetUserList().then(res => {
        if (res.code === 0) {
          this.userList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenGroupUserDetailForm(row) {
      this.groupUserDetail = {
        "Username": row.Username,
        "Email": row.Email,
        "GroupChallenges": row.GroupChallenges
      }
      this.groupUserDetailVisible = true
    },
    ClearGroupUserDetail() {
      this.groupUserDetail = {
        "Username": "",
        "Email": "",
        "GroupChallenges": []
      }
    },
    AddGroupUser() {
      groupApi.AddGroupUser(this.addGroupUserData).then(res => {
        if (res.code === 0) {
          this.addGroupUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetGroup()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenAddGroupUserForm() {
      this.addGroupUserData.GroupID = this.group.ID
      this.addGroupUserFormVisible = true
    },
    ClearAddGroupUserForm() {
      this.addGroupUserData = {
        "GroupID": 0,
        "UserID": null
      }
    },
    RemoveGroupUser() {
      groupApi.RemoveGroupUser(this.removeGroupUserData).then(res => {
        if (res.code === 0) {
          this.removeGroupUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetGroup()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenRemoveGroupUserForm(row) {
      this.removeGroupUserData.GroupID = this.group.ID
      this.removeGroupUserData.UserID = row.ID
      this.removeGroupUserFormVisible = true
    },
    ClearRemoveGroupUserForm() {
      this.removeGroupUserData = {
        "GroupID": 0,
        "UserID": null
      }
    },
  },
  mounted() {
    this.group.ID = Number(this.$route.params.id)
    this.GetGroup()
    this.GetGroupChallengeCategoryList()
    this.GetImageList()
    this.GetAttachmentList()
    this.GetGroupChallengeList()
    this.GetUserList()

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
