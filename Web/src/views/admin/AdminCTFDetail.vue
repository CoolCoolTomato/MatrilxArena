<template>
  <el-container>
    <el-header>
      <div style="display: flex; align-items: center;">
        <h2 style="color: var(--el-text-color-primary);">{{ $t('AdminCTFDetail.CTFDetail') }}</h2>
        <div style="flex-grow: 1;"></div>
        <div style="margin-right: 50px;">
          <el-button
            style="margin: 10px;"
            @click="goBack"
          >
            {{ $t('AdminCTFDetail.Back') }}
          </el-button>
        </div>
      </div>
    </el-header>
    <el-main>
      <el-tabs
        v-model="activeTab"
        class="el-tabs-custom"
        >
        <el-tab-pane :label="$t('AdminCTFDetail.Challenges')" name="challenges">
          <div style="height: 50px; display: flex; align-items: center;">
            <el-button
              @click="OpenCreateCTFChallengeForm"
              type="primary"
            >
              {{ $t('AdminCTFDetail.Add') }}
            </el-button>
            <el-input
              v-model="ctfChallengeQueryTitle"
              style="width: 560px; margin: 10px;"
              :placeholder="$t('AdminCTFDetail.FindChallenges')"
            >
              <template #prepend>
                <el-select
                  v-model="ctfChallengeQueryCategoryID"
                  filterable
                  style="width: 110px;"
                  :placeholder="$t('AdminCTFDetail.Select')"
                >
                  <el-option
                    v-for="category in categoryList"
                    :key="category.ID"
                    :label="category.ID === 0 ? $t('AdminCTFDetail.All') : category.Name"
                    :value="category.ID"
                  ></el-option>
                </el-select>
              </template>
              <template #append>
                <el-button
                  @click="GetCTFChallengeList"
                  style="width: 100px;"
                >
                  {{ $t('AdminCTFDetail.Find') }}
                  <el-icon style="margin-left: 10px">
                    <Search fill="var(var(--el-button-text-color))"/>
                  </el-icon>
                </el-button>
              </template>
            </el-input>
          </div>
          <el-table
            :data="ctfChallengeList"
            table-layout="fixed"
            height="calc(100% - 50px)"
            >
            <el-table-column prop="Title" :label="$t('AdminCTFDetail.Title')"/>
            <el-table-column prop="Description" :label="$t('AdminCTFDetail.Description')"/>
            <el-table-column :label="$t('AdminCTFDetail.Category')">
              <template  #default=scope>
                {{ scope.row.Category.Name === "" ? $t('AdminCTFDetail.Null') : scope.row.Category.Name}}
              </template>
            </el-table-column>
            <el-table-column :label="$t('AdminCTFDetail.Image')">
              <template  #default=scope>
                {{ scope.row.Image.Remark === "" ? $t('AdminCTFDetail.Null') : scope.row.Image.Remark}}
              </template>
            </el-table-column>
            <el-table-column fixed="right" :label="$t('AdminCTFDetail.Operations')" width="300px">
              <template #default=scope>
                <el-button
                  @click="OpenCTFChallengeDetail(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Detail') }}
                </el-button>
                <el-button
                  @click="OpenUpdateCTFChallengeForm(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Update') }}
                </el-button>
                <el-button
                  @click="OpenDeleteChallengeForm(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Delete') }}
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane :label="$t('AdminCTFDetail.Users')" name="users">
          <div style="height: 50px; display: flex; align-items: center;">
            <el-button
              @click="OpenAddCTFUserForm"
              type="primary"
              >
              {{ $t('AdminCTFDetail.Add') }}
            </el-button>
            <el-input
              v-model="ctfUserQueryUsername"
              style="width: 450px; margin: 10px;"
              :placeholder="$t('AdminCTFDetail.FindUsers')"
            >
              <template #append>
                <el-button
                  @click="GetCTFUserList"
                  style="width: 100px;"
                >
                  {{ $t('AdminCTFDetail.Find') }}
                  <el-icon style="margin-left: 10px">
                    <Search fill="var(var(--el-button-text-color))"/>
                  </el-icon>
                </el-button>
              </template>
            </el-input>
          </div>
          <el-table
            :data="ctfUserList"
            table-layout="fixed"
            height="calc(100% - 50px)"
            >
            <el-table-column prop="Username" :label="$t('AdminCTFDetail.Username')"/>
            <el-table-column prop="Email" :label="$t('AdminCTFDetail.Email')"/>
            <el-table-column  :label="$t('AdminCTFDetail.SolvedChallenges')">
              <template #default="scope">
                {{ scope.row.CTFChallenges.length }}
              </template>
            </el-table-column>
            <el-table-column fixed="right" :label="$t('AdminCTFDetail.Operations')" width="230">
              <template #default=scope>
                <el-button
                  @click="OpenCTFUserDetail(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Detail') }}
                </el-button>
                <el-button
                  @click="OpenRemoveCTFUserForm(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Remove') }}
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane :label="$t('AdminCTFDetail.Teams')" name="teams">
          <div style="height: 50px; display: flex; align-items: center;">
            <el-button
              @click="OpenCreateCTFTeamForm"
              type="primary"
              >
              {{ $t('AdminCTFDetail.Add') }}
            </el-button>
            <el-input
              v-model="ctfTeamQueryName"
              style="width: 450px; margin: 10px;"
              :placeholder="$t('AdminCTFDetail.FindTeams')"
            >
              <template #append>
                <el-button
                  @click="GetCTFTeamList"
                  style="width: 100px;"
                >
                  {{ $t('AdminCTFDetail.Find') }}
                  <el-icon style="margin-left: 10px">
                    <Search fill="var(var(--el-button-text-color))"/>
                  </el-icon>
                </el-button>
              </template>
            </el-input>
          </div>
          <el-table
            :data="ctfTeamList"
            table-layout="fixed"
            height="calc(100% - 50px)"
          >
            <el-table-column prop="Name" :label="$t('AdminCTFDetail.Name')"/>
            <el-table-column :label="$t('AdminCTFDetail.Captain')">
              <template #default="scope">
                {{ scope.row.User.Username }}
              </template>
            </el-table-column>
            <el-table-column :label="$t('AdminCTFDetail.Users')">
              <template #default="scope">
                {{ scope.row.Users.length }}
              </template>
            </el-table-column>
            <el-table-column :label="$t('AdminCTFDetail.SolvedChallenges')">
              <template #default="scope">
                {{ scope.row.CTFChallenges.length }}
              </template>
            </el-table-column>
            <el-table-column fixed="right" :label="$t('AdminCTFDetail.Operations')" width="300px">
              <template #default=scope>
                <el-button
                  @click="OpenCTFTeamDetail(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Detail') }}
                </el-button>
                <el-button
                  @click="OpenUpdateCTFTeamForm(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Update') }}
                </el-button>
                <el-button
                  @click="OpenDeleteCTFTeamForm(scope.row)"
                  >
                  {{ $t('AdminCTFDetail.Delete') }}
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
      <el-dialog
        v-model="ctfChallengeDetailVisible"
        :title="$t('AdminCTFDetail.ChallengeDetail')"
        width="900"
        @close="ClearCTFChallengeDetail"
        >
        <el-card>
          <h2>{{ $t('AdminCTFDetail.Challenge') }}</h2>
          <el-row>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Title') }}: {{ ctfChallengeDetail.Title }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Description') }}: {{ ctfChallengeDetail.Description }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Category') }}: {{ ctfChallengeDetail.Category.Name === "" ? $t('AdminCTFDetail.Null') : ctfChallengeDetail.Category.Name }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Image') }}: {{ ctfChallengeDetail.Image.Remark === "" ? $t('AdminCTFDetail.Null') : ctfChallengeDetail.Image.Remark }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.RepoTags') }}: {{ ctfChallengeDetail.Image.RepoTags === "" ? $t('AdminCTFDetail.Null') : ctfChallengeDetail.Image.RepoTags }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Attachment') }}: {{ ctfChallengeDetail.Attachment.FileName === "" ? $t('AdminCTFDetail.Null') : ctfChallengeDetail.Attachment.FileName }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.SpecifiedPorts') }}: {{ ctfChallengeDetail.SpecifiedPorts }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Commands') }}: {{ ctfChallengeDetail.Commands }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Score') }}: {{ ctfChallengeDetail.Score }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Flag') }}: {{ ctfChallengeDetail.Flag }}</p>
            </el-col>
          </el-row>
          <h2>{{ $t('AdminCTFDetail.Solver') }}</h2>
          <div>
            <el-row v-if="ctfChallengeDetail.Users.length !== 0">
              <el-col
                :span="4"
                v-for="user in ctfChallengeDetail.Users"
              >
                <div style="display: flex; flex-direction: column; align-items: center; margin: 10px">
                  <el-avatar size="default">
                    {{ user.Username[0] }}
                  </el-avatar>
                  <el-text>{{ user.Username }}</el-text>
                </div>
              </el-col>
            </el-row>
            <el-text v-else>
              {{ $t('AdminCTFDetail.Null') }}
            </el-text>
          </div>
        </el-card>
        <template #footer>
          <el-button @click="ctfChallengeDetailVisible = false">{{ $t('AdminCTFDetail.Close') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="createCTFChallengeFormVisible"
        :title="$t('AdminCTFDetail.CreateChallenge')"
        width="600"
        @close="ClearCreateCTFChallengeForm"
        >
        <el-form :model=createCTFChallengeData>
          <el-form-item :label="$t('AdminCTFDetail.Title')" :label-width="labelWidth">
            <el-input v-model="createCTFChallengeData.Title"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Description')" :label-width="labelWidth">
            <el-input v-model="createCTFChallengeData.Description"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Category')" :label-width="labelWidth">
            <el-select
              v-model="createCTFChallengeData.CategoryID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
              >
              <el-option
                v-for="category in categoryList"
                :key="category.ID"
                :label="category.Name"
                :value="category.ID"
                ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Image')" :label-width="labelWidth">
            <el-select
              v-model="createCTFChallengeData.ImageID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
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
          <el-form-item :label="$t('AdminCTFDetail.Attachment')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-select
                style="width: 50%"
                v-model="createCTFChallengeData.AttachmentID"
                filterable
                :placeholder="$t('AdminCTFDetail.Select')"
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
                  <el-button style="margin-left: 20px">{{ $t('AdminCTFDetail.SelectFile') }}</el-button>
                </template>
                <div style="display: inline-flex">
                  <el-button
                    @click="UploadAttachment"
                    style="margin-left: 10px"
                    type="primary"
                    >
                    {{ $t('AdminCTFDetail.Upload') }}
                  </el-button>
                </div>
              </el-upload>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.SpecifiedPorts')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                v-model="createCTFChallengePort"
              >
                <template #append>
                  <el-select
                    v-model="createCTFChallengeProtocol"
                    :placeholder="$t('AdminCTFDetail.Select')"
                    style="width: 100px"
                  >
                    <el-option label="tcp" value="tcp"/>
                    <el-option label="udp" value="udp"/>
                  </el-select>
                </template>
              </el-input>
              <el-button
                style="margin-left: 20px;"
                @click="AddCreateCTFChallengeSpecifiedPort"
              >
                {{ $t('AdminCTFDetail.Add') }}
              </el-button>
            </div>
            <div
              v-if="createCTFChallengeData.SpecifiedPorts.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
              >
              <div
                v-for="(port, index) in createCTFChallengeData.SpecifiedPorts"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ port }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveCreateCTFChallengeSpecifiedPort(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  {{ $t('AdminCTFDetail.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Commands')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.createCTFChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddCreateCTFChallengeCommand"
              >
                {{ $t('AdminCTFDetail.Add') }}
              </el-button>
            </div>
            <div
              v-if="createCTFChallengeData.Commands.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
            >
              <div
                v-for="(command, index) in createCTFChallengeData.Commands"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ command }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveCreateCTFChallengeCommand(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  {{ $t('AdminCTFDetail.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Score')" :label-width="labelWidth">
            <el-input-number v-model="createCTFChallengeData.Score"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Flag')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                style="width: 245px"
                v-if="createCTFChallengeFlagType === 'auto'"
                :placeholder="$t('AdminCTFDetail.Auto')"
                disabled
              />
              <el-input
                style="width: 245px"
                v-if="createCTFChallengeFlagType === 'specify'"
                v-model="createCTFChallengeData.Flag"
              />
              <div style="flex-grow: 1" v-if="createCTFChallengeFlagType !== ''"></div>
              <el-radio-group v-model="createCTFChallengeFlagType">
                <el-radio value="auto" style="margin: 0" border>{{ $t('AdminCTFDetail.Auto') }}</el-radio>
                <el-radio value="specify" style="margin: 0" border>{{ $t('AdminCTFDetail.Specify') }}</el-radio>
              </el-radio-group>
            </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateCTFChallenge"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Submit') }}
          </el-button>
          <el-button
            @click="createCTFChallengeFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateCTFChallengeFormVisible"
        :title="$t('AdminCTFDetail.UpdateChallenge')"
        width="600"
        @close="ClearUpdateCTFChallengeForm"
        >
        <el-form :model=updateCTFChallengeData>
          <el-form-item :label="$t('AdminCTFDetail.Title')" :label-width="labelWidth">
            <el-input v-model="updateCTFChallengeData.Title"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Description')" :label-width="labelWidth">
            <el-input v-model="updateCTFChallengeData.Description"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Category')" :label-width="labelWidth">
            <el-select
              v-model="updateCTFChallengeData.CategoryID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
              >
              <el-option
                v-for="category in categoryList"
                :key="category.ID"
                :label="category.Name"
                :value="category.ID"
                ></el-option>
              <template v-if="!categoryList.some(category => category.ID === updateCTFChallengeData.CategoryID)" slot="empty">
                <el-option
                  :label="$t('AdminCTFDetail.DeletedCategory')"
                  :value="updateCTFChallengeData.CategoryID"
                  disabled
                />
              </template>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Image')" :label-width="labelWidth">
            <el-select
              v-model="updateCTFChallengeData.ImageID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
              >
              <el-option
                v-for="image in imageList"
                :key="image.ID"
                :label="image.Remark"
                :value="image.ID"
                >
              </el-option>
              <template v-if="!imageList.some(image => image.ID === updateCTFChallengeData.ImageID)" slot="empty">
                <el-option
                  :label="$t('AdminCTFDetail.DeletedImage')"
                  :value="updateCTFChallengeData.ImageID"
                  disabled
                />
              </template>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Attachment')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-select
                style="width: 50%"
                v-model="updateCTFChallengeData.AttachmentID"
                filterable
                :placeholder="$t('AdminCTFDetail.Select')"
                >
                <el-option
                  v-for="attachment in attachmentList"
                  :key="attachment.ID"
                  :label="attachment.FileName"
                  :value="attachment.ID"
                >
                </el-option>
                <template v-if="!attachmentList.some(attachment => attachment.ID === updateCTFChallengeData.AttachmentID)" slot="empty">
                  <el-option
                    :label="$t('AdminCTFDetail.DeletedFile')"
                    :value="updateCTFChallengeData.AttachmentID"
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
                  <el-button style="margin-left: 20px">{{ $t('AdminCTFDetail.SelectFile') }}</el-button>
                </template>
                <div style="display: inline-flex">
                  <el-button
                    @click="UploadAttachment"
                    style="margin-left: 10px"
                    type="primary"
                    >
                    {{ $t('AdminCTFDetail.Upload') }}
                  </el-button>
                </div>
              </el-upload>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.SpecifiedPorts')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                v-model="updateCTFChallengePort"
              >
                <template #append>
                  <el-select
                    v-model="updateCTFChallengeProtocol"
                    :placeholder="$t('AdminCTFDetail.Select')"
                    style="width: 100px"
                  >
                    <el-option label="tcp" value="tcp"/>
                    <el-option label="udp" value="udp"/>
                  </el-select>
                </template>
              </el-input>
              <el-button
                style="margin-left: 20px;"
                @click="AddUpdateCTFChallengeSpecifiedPort"
              >
                {{ $t('AdminCTFDetail.Add') }}
              </el-button>
            </div>
            <div
              v-if="updateCTFChallengeData.SpecifiedPorts.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
              >
              <div
                v-for="(port, index) in updateCTFChallengeData.SpecifiedPorts"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ port }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveUpdateCTFChallengeSpecifiedPort(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  {{ $t('AdminCTFDetail.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Commands')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input v-model="this.updateCTFChallengeCommand" />
              <el-button
                style="margin-left: 20px;"
                @click="AddUpdateCTFChallengeCommand"
              >
                {{ $t('AdminCTFDetail.Add') }}
              </el-button>
            </div>
            <div
              v-if="updateCTFChallengeData.Commands.length > 0"
              style="position: relative; width: 100%; border: var(--el-border); margin-top: 20px;"
            >
              <div
                v-for="(command, index) in updateCTFChallengeData.Commands"
                :key="index"
                style="display: flex; align-items: center; width: 100%; position: relative"
              >
                <el-text style="margin: 3px 3px 3px 10px;" size="small">{{ command }}</el-text>
                <div style="flex-grow: 1"></div>
                <el-button
                  size="small"
                  @click="RemoveUpdateCTFChallengeCommand(index)"
                  style="margin: 3px 10px 3px 3px;"
                >
                  {{ $t('AdminCTFDetail.Delete') }}
                </el-button>
              </div>
            </div>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Score')" :label-width="labelWidth">
            <el-input-number v-model="createCTFChallengeData.Score"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Flag')" :label-width="labelWidth">
            <div style="display: flex; width: 100%;">
              <el-input
                style="width: 245px"
                v-if="updateCTFChallengeFlagType === 'auto'"
                :placeholder="$t('AdminCTFDetail.Auto')"
                disabled
              />
              <el-input
                style="width: 245px"
                v-if="updateCTFChallengeFlagType === 'specify'"
                v-model="updateCTFChallengeData.Flag"
              />
              <div style="flex-grow: 1" v-if="updateCTFChallengeFlagType !== ''"></div>
              <el-radio-group v-model="updateCTFChallengeFlagType">
                <el-radio value="auto" style="margin: 0" border>{{ $t('AdminCTFDetail.Auto') }}</el-radio>
                <el-radio value="specify" style="margin: 0" border>{{ $t('AdminCTFDetail.Specify') }}</el-radio>
              </el-radio-group>
            </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateCTFChallenge"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Submit') }}
          </el-button>
          <el-button
            @click="updateCTFChallengeFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteCTFChallengeFormVisible"
        :title="$t('AdminCTFDetail.DeleteChallenge')"
        width="500"
        @close="ClearDeleteCTFChallengeForm"
        >
        <el-text>{{ $t('AdminCTFDetail.AreYouConfirmToDeleteTheChallenge') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteCTFChallenge"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Confirm') }}
          </el-button>
          <el-button
            @click="deleteCTFChallengeFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>

      <el-dialog
        v-model="ctfUserDetailVisible"
        :title="$t('AdminCTFDetail.UserDetail')"
        width="900"
        @close="ClearCTFUserDetail"
        >
        <el-card>
          <h2>{{ $t('AdminCTFDetail.User') }}</h2>
          <el-row>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Username') }}: {{ ctfUserDetail.Username }}</p>
            </el-col>
            <el-col :span="12">
              <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Email') }}: {{ ctfUserDetail.Email }}</p>
            </el-col>
          </el-row>
          <h2>{{ $t('AdminCTFDetail.SolvedChallenges') }}</h2>
          <div>
            <el-row v-if="ctfUserDetail.CTFChallenges.length !== 0">
              <el-col
                :span="6"
                v-for="ctfChallenge in ctfUserDetail.CTFChallenges"
              >
                <div
                  style="margin: 10px;
                  height: 50px;
                  display: flex;
                  align-items: center;
                  border: solid 1px var(--el-border-color);
                  border-left: solid 5px var(--el-border-color);
                  position: relative"
                >
                  <el-text style="margin-left: 20px">{{ ctfChallenge.Title }}</el-text>
                  <el-icon size="25" style="right: 20px; position: absolute">
                    <True fill="var(--el-color-primary)"/>
                  </el-icon>
                </div>
              </el-col>
            </el-row>
            <el-text v-else>
              {{ $t('AdminCTFDetail.Null') }}
            </el-text>
          </div>
        </el-card>
        <template #footer>
          <el-button @click="ctfUserDetailVisible = false">{{ $t('AdminCTFDetail.Close') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="addCTFUserFormVisible"
        :title="$t('AdminCTFDetail.AddUser')"
        width="500"
        @close="ClearAddCTFUserForm"
        >
        <el-form :model="addCTFUserData">
          <el-form-item>
            <el-select
              v-model="addCTFUserData.UserID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
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
            @click="AddCTFUser"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Submit') }}
          </el-button>
          <el-button
            @click="addCTFUserFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="removeCTFUserFormVisible"
        :title="$t('AdminCTFDetail.RemoveUser')"
        width="500"
        @close="ClearRemoveCTFUserForm"
        >
        <el-text>{{ $t('AdminCTFDetail.AreYouConfirmToRemoveTheUser') }}</el-text>
        <template #footer>
          <el-button
            @click="RemoveCTFUser"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Confirm') }}
          </el-button>
          <el-button
            @click="removeCTFUserFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>

      <el-dialog
        v-model="ctfTeamDetailVisible"
        :title = "$t('AdminCTFDetail.TeamDetail')"
        width="900"
        @close="ClearCTFTeamDetail"
      >
        <el-card>
          <h2>{{ $t('AdminCTFDetail.Team') }}</h2>
          <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Name') }}: {{ ctfTeamDetail.Name }}</p>
          <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Captain') }}: {{ ctfTeamDetail.User }}</p>
          <p style="word-break: break-all;">{{ $t('AdminCTFDetail.Description') }}: {{ ctfTeamDetail.Description }}</p>
          <div style="display: flex; align-items: center;">
            <h2>{{ $t('AdminCTFDetail.Users') }}</h2>
            <el-button
              @click="OpenAddCTFTeamUserForm(ctfTeamDetail.ID)"
              size="small"
              type="primary"
              style="margin-left: 20px;"
            >
              {{ $t('AdminCTFDetail.Add') }}
            </el-button>
          </div>
          <div>
            <el-table
              v-if="ctfTeamDetail.Users.length !== 0"
              :data="ctfTeamDetail.Users"
              table-layout="fixed"
            >
              <el-table-column prop="Username" :label="$t('AdminCTFDetail.Username')"/>
              <el-table-column prop="Email" :label="$t('AdminCTFDetail.Email')"/>
              <el-table-column fixed="right" :label="$t('AdminCTFDetail.Operations')" width="130">
                <template #default=scope>
                  <el-button
                    @click="OpenRemoveCTFTeamUserForm(ctfTeamDetail.ID, scope.row.ID)"
                    size="small"
                  >
                    {{ $t('AdminCTFDetail.Remove') }}
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-text v-else>
              {{ $t('AdminCTFDetail.Null') }}
            </el-text>
          </div>
          <h2>{{ $t('AdminCTFDetail.SolvedChallenges') }}</h2>
          <div>
            <el-row v-if="ctfTeamDetail.CTFChallenges.length !== 0">
              <el-col
                :span="6"
                v-for="ctfChallenge in ctfTeamDetail.CTFChallenges"
              >
                <div
                  style="margin: 10px;
                  height: 50px;
                  display: flex;
                  align-items: center;
                  border: solid 1px var(--el-border-color);
                  border-left: solid 5px var(--el-border-color);
                  position: relative"
                >
                  <el-text style="margin-left: 20px">{{ ctfChallenge.Title }}</el-text>
                  <el-icon size="25" style="right: 20px; position: absolute">
                    <True fill="var(--el-color-primary)"/>
                  </el-icon>
                </div>
              </el-col>
            </el-row>
            <el-text v-else>
              {{ $t('AdminCTFDetail.Null') }}
            </el-text>
          </div>
        </el-card>
        <template #footer>
          <el-button @click="ctfTeamDetailVisible = false">{{ $t('AdminCTFDetail.Close') }}</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="createCTFTeamFormVisible"
        :title="$t('AdminCTFDetail.CreateCTFTeam')"
        width="500"
        @close="ClearCreateCTFTeamForm"
      >
        <el-form :model="createCTFTeamData">
          <el-form-item :label="$t('AdminCTFDetail.Name')" :label-width="labelWidth">
            <el-input v-model="createCTFTeamData.Name"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Captain')" :label-width="labelWidth">
            <el-select
              v-model="createCTFTeamData.UserID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
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
          <el-form-item :label="$t('AdminCTFDetail.Description')" :label-width="labelWidth">
            <el-input v-model="createCTFTeamData.Description"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateCTFTeam"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Submit') }}
          </el-button>
          <el-button
            @click="createCTFTeamFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateCTFTeamFormVisible"
        :title="$t('AdminCTFDetail.UpdateCTFTeam')"
        width="500"
        @close="ClearUpdateCTFTeamForm"
      >
        <el-form :model="updateCTFTeamData">
          <el-form-item :label="$t('AdminCTFDetail.Name')" :label-width="labelWidth">
            <el-input v-model="updateCTFTeamData.Name"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCTFDetail.Captain')" :label-width="labelWidth">
            <el-select
              v-model="updateCTFTeamData.UserID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
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
          <el-form-item :label="$t('AdminCTFDetail.Description')" :label-width="labelWidth">
            <el-input v-model="updateCTFTeamData.Description"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateCTFTeam"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Submit') }}
          </el-button>
          <el-button
            @click="updateCTFTeamFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteCTFTeamFormVisible"
        :title="$t('AdminCTFDetail.DeleteCTFTeam')"
        width="500"
        @close="ClearDeleteCTFTeamForm"
        >
        <el-text>{{ $t('AdminCTFDetail.AreYouConfirmToDeleteTheCTFTeam') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteCTFTeam"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Confirm') }}
          </el-button>
          <el-button
            @click="deleteCTFTeamFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>

      <el-dialog
        v-model="addCTFTeamUserFormVisible"
        :title="$t('AdminCTFDetail.AddUser')"
        width="500"
        @close="ClearAddCTFTeamUserForm"
        >
        <el-form :model="addCTFTeamUserData">
          <el-form-item>
            <el-select
              v-model="addCTFTeamUserData.UserID"
              filterable
              :placeholder="$t('AdminCTFDetail.Select')"
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
            @click="AddCTFTeamUser"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Submit') }}
          </el-button>
          <el-button
            @click="addCTFTeamUserFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="removeCTFTeamUserFormVisible"
        :title="$t('AdminCTFDetail.RemoveUser')"
        width="500"
        @close="ClearRemoveCTFTeamUserForm"
        >
        <el-text>{{ $t('AdminCTFDetail.AreYouConfirmToRemoveTheUser') }}</el-text>
        <template #footer>
          <el-button
            @click="RemoveCTFTeamUser"
            type="primary"
            >
            {{ $t('AdminCTFDetail.Confirm') }}
          </el-button>
          <el-button
            @click="removeCTFTeamUserFormVisible = false"
            >
            {{ $t('AdminCTFDetail.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import ctfApi from "@/api/ctf.js"
import ctfChallengeApi from "@/api/ctfChallenge.js"
import ctfUserApi from "@/api/ctfUser.js"
import categoryApi from "@/api/category.js"
import imageApi from "@/api/image.js"
import attachmentApi from "@/api/attachment.js"
import userApi from "@/api/user.js"
import ctfTeamApi from "@/api/ctfTeam.js"
import ctfTeamUserApi from "@/api/ctfTeamUser.js"
import Search from "@/components/icons/Search.vue"
import True from "@/components/icons/True.vue"
import {ElMessage} from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: { Search, True },
  data() {
    return {
      activeTab: "challenges",
      ctf: {
        "ID": 0,
      },
      labelWidth: 120,
      categoryList: [],
      imageList: [],
      attachmentList: [],
      ctfChallengeQueryCategoryID: 0,
      ctfChallengeQueryTitle: "",
      ctfChallengeList: [],
      ctfChallengeDetailVisible: false,
      ctfChallengeDetail: {
        "Title": "",
        "Description": "",
        "Category": {},
        "Image": {},
        "Attachment": {},
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "Score": 0,
        "Users": [],
        "CTFTeams": []
      },
      createCTFChallengeFormVisible: false,
      createCTFChallengeData: {
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "Score": 0,
        "CTFID": 0
      },
      createCTFChallengePort: "",
      createCTFChallengeProtocol: "",
      createCTFChallengeCommand: "",
      createCTFChallengeFlagType: "",
      updateCTFChallengeFormVisible: false,
      updateCTFChallengeData: {
        "ID": 0,
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "Score": 0,
        "CTFID": 0
      },
      updateCTFChallengePort: "",
      updateCTFChallengeProtocol: "",
      updateCTFChallengeCommand: "",
      updateCTFChallengeFlagType: "",
      deleteCTFChallengeFormVisible: false,
      deleteCTFChallengeData: {
        "ID": 0,
      },
      uploadAttachmentData: {
        "fileName": "",
        "file": null
      },

      userList: [],
      ctfUserQueryUsername: "",
      ctfUserList: [],
      ctfUserDetailVisible: false,
      ctfUserDetail: {
        "Username": "",
        "Email": "",
        "CTFChallenges": []
      },
      addCTFUserFormVisible: false,
      addCTFUserData: {
        "CTFID": 0,
        "UserID": null
      },
      removeCTFUserFormVisible: false,
      removeCTFUserData: {
        "CTFID": 0,
        "UserID": null
      },

      ctfTeamList: [],
      ctfTeamQueryName: "",
      ctfTeamDetailVisible: false,
      ctfTeamDetail: {
        "ID": 0,
        "Name": "",
        "User": "",
        "Description": "",
        "Users": [],
        "CTFChallenges": []
      },
      createCTFTeamFormVisible: false,
      createCTFTeamData: {
        "Name": "",
        "Description": "",
        "UserID": null,
        "CTFID": 0
      },
      updateCTFTeamFormVisible: false,
      updateCTFTeamData: {
        "ID": 0,
        "Name": "",
        "Description": "",
        "UserID": null,
        "CTFID": 0
      },
      deleteCTFTeamFormVisible: false,
      deleteCTFTeamData: {
        "ID": 0
      },

      addCTFTeamUserFormVisible: false,
      addCTFTeamUserData: {
        "CTFTeamID": 0,
        "UserID": null
      },
      removeCTFTeamUserFormVisible: false,
      removeCTFTeamUserData: {
        "CTFTeamID": 0,
        "UserID": null
      },
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1);
    },
    GetCTFChallengeCategoryList() {
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
      imageApi.GetImageList({}).then(res => {
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
      attachmentApi.GetAttachmentList({}).then(res => {
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
    GetCTF() {
      ctfApi.GetCTF(this.ctf).then(res => {
        if (res.code === 0) {
          this.ctf = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetCTFChallengeList() {
      ctfChallengeApi.GetCTFChallengeList({
        "CategoryID": this.ctfChallengeQueryCategoryID,
        "Title": this.ctfChallengeQueryTitle,
        "CTFID": this.ctf.ID
      }).then(res => {
        if (res.code === 0) {
          this.ctfChallengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCTFChallengeDetail(row) {
      this.ctfChallengeDetail = row
      this.ctfChallengeDetailVisible = true
    },
    ClearCTFChallengeDetail() {
      this.ctfChallengeDetail = {
        "Title": "",
        "Description": "",
        "Category": {},
        "Image": {},
        "Attachment": {},
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "Score": 0,
        "Users": []
      }
    },
    AddCreateCTFChallengeSpecifiedPort() {
      if (this.createCTFChallengePort === "") {
        ElMessage.error(this.t('AdminCTFDetail.PleaseInputPort'))
        return
      }
      if (this.createCTFChallengeProtocol === "") {
        ElMessage.error(this.t('AdminCTFDetail.PleaseInputProtocol'))
        return
      }
      let portProtocol = this.createCTFChallengePort + "/" + this.createCTFChallengeProtocol
      if (this.createCTFChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error(this.t('AdminCTFDetail.DuplicateSpecifiedPort'))
        return
      }
      this.createCTFChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveCreateCTFChallengeSpecifiedPort(index) {
      this.createCTFChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddCreateCTFChallengeCommand() {
      if (this.createCTFChallengeCommand === "") {
        ElMessage.error(this.t('AdminCTFDetail.PleaseInputCommand'))
        return
      }
      if (this.createCTFChallengeData.Commands.includes(this.createCTFChallengeCommand)) {
        ElMessage.error(this.t('AdminCTFDetail.DuplicateCommand'))
        return
      }
      this.createCTFChallengeData.Commands.push(this.createCTFChallengeCommand)
    },
    RemoveCreateCTFChallengeCommand(index) {
      this.createCTFChallengeData.Commands.splice(index, 1);
    },
    CreateCTFChallenge() {
      if (this.createCTFChallengeFlagType === "auto") {
        this.createCTFChallengeData.Flag = "auto"
      }
      ctfChallengeApi.CreateCTFChallenge(this.createCTFChallengeData).then(res => {
        if (res.code === 0) {
          this.createCTFChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCreateCTFChallengeForm() {
      this.createCTFChallengeData.CTFID = this.ctf.ID
      this.createCTFChallengeFormVisible = true
    },
    ClearCreateCTFChallengeForm() {
      this.createCTFChallengeData = {
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "ImageID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "Score": 0,
        "CTFID": 0
      }
      this.createCTFChallengePort = ""
      this.createCTFChallengeProtocol = ""
      this.createCTFChallengeCommand = ""
      this.createCTFChallengeFlagType = ""
      this.ClearUploadAttachmentForm()
    },
    AddUpdateCTFChallengeSpecifiedPort() {
      if (this.updateCTFChallengePort === "") {
        ElMessage.error(this.t('AdminCTFDetail.PleaseInputPort'))
        return
      }
      if (this.updateCTFChallengeProtocol === "") {
        ElMessage.error(this.t('AdminCTFDetail.PleaseInputProtocol'))
        return
      }
      let portProtocol = this.updateCTFChallengePort + "/" + this.updateCTFChallengeProtocol
      if (this.updateCTFChallengeData.SpecifiedPorts.includes(portProtocol)) {
        ElMessage.error(this.t('AdminCTFDetail.DuplicateSpecifiedPort'))
        return
      }
      this.updateCTFChallengeData.SpecifiedPorts.push(portProtocol)
    },
    RemoveUpdateCTFChallengeSpecifiedPort(index) {
      this.updateCTFChallengeData.SpecifiedPorts.splice(index, 1);
    },
    AddUpdateCTFChallengeCommand() {
      if (this.updateCTFChallengeCommand === "") {
        ElMessage.error(this.t('AdminCTFDetail.PleaseInputCommand'))
        return
      }
      if (this.updateCTFChallengeData.Commands.includes(this.updateCTFChallengeCommand)) {
        ElMessage.error(this.t('AdminCTFDetail.DuplicateCommand'))
        return
      }
      this.updateCTFChallengeData.Commands.push(this.updateCTFChallengeCommand)
    },
    RemoveUpdateCTFChallengeCommand(index) {
      this.updateCTFChallengeData.Commands.splice(index, 1);
    },
    UpdateCTFChallenge() {
      if (this.updateCTFChallengeFlagType === "auto") {
        this.updateCTFChallengeData.Flag = "auto"
      }
      ctfChallengeApi.UpdateCTFChallenge(this.updateCTFChallengeData).then(res => {
        if (res.code === 0) {
          this.updateCTFChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetCTFChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateCTFChallengeForm(row) {
      this.updateCTFChallengeData = {
        "ID": row.ID,
        "Title": row.Title,
        "Description": row.Description,
        "CategoryID": row.CategoryID,
        "ImageID": row.ImageID,
        "AttachmentID": row.AttachmentID,
        "SpecifiedPorts": row.SpecifiedPorts === null ? [] : [...row.SpecifiedPorts],
        "Commands": row.Commands === null ? [] : [...row.Commands],
        "Flag": row.Flag,
        "Score": row.Score,
        "CTFID": row.CTFID
      }
      this.updateCTFChallengeFlagType = row.Flag === "auto" ? "auto" : "specify"
      this.updateCTFChallengeFormVisible = true
    },
    ClearUpdateCTFChallengeForm() {
      this.updateCTFChallengeData = {
        "ID": 0,
        "ImageID": 0,
        "Title": "",
        "Description": "",
        "CategoryID": 0,
        "AttachmentID": 0,
        "SpecifiedPorts": [],
        "Commands": [],
        "Flag": "",
        "Score": 0,
        "CTFID": 0
      }
      this.updateCTFChallengePort = ""
      this.updateCTFChallengeProtocol = ""
      this.updateCTFChallengeCommand = ""
      this.updateCTFChallengeFlagType = ""
      this.ClearUploadAttachmentForm()
    },
    DeleteCTFChallenge() {
      ctfChallengeApi.DeleteCTFChallenge(this.deleteCTFChallengeData).then(res => {
        if (res.code === 0) {
          this.deleteCTFChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetCTFChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteChallengeForm(row) {
      this.deleteCTFChallengeData = {
        "ID": row.ID
      }
      this.deleteCTFChallengeFormVisible = true
    },
    ClearDeleteCTFChallengeForm() {
      this.deleteCTFChallengeData = {
        "ID": 0
      }
    },
    UploadAttachment() {
      attachmentApi.UploadAttachment(this.uploadAttachmentData.fileName, this.uploadAttachmentData.file).then(res => {
        if (res.code === 0) {
          if (this.createCTFChallengeFormVisible) {
            this.createCTFChallengeData.AttachmentID = res.data.ID
          } else if (this.updateCTFChallengeFormVisible) {
            this.updateCTFChallengeData.AttachmentID = res.data.ID
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
      userApi.GetUserList({}).then(res => {
        if (res.code === 0) {
          this.userList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    GetCTFUserList() {
      ctfUserApi.GetCTFUserList({
        "CTFID": this.ctf.ID,
        "Username": this.ctfUserQueryUsername
      }).then(res => {
        if (res.code === 0) {
          this.ctfUserList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCTFUserDetail(row) {
      this.ctfUserDetail = row
      this.ctfUserDetailVisible = true
    },
    ClearCTFUserDetail() {
      this.ctfUserDetail = {
        "Username": "",
        "Email": "",
        "CTFChallenges": []
      }
    },
    AddCTFUser() {
      ctfUserApi.AddCTFUser(this.addCTFUserData).then(res => {
        if (res.code === 0) {
          this.addCTFUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetCTF()
          this.GetCTFUserList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenAddCTFUserForm() {
      this.addCTFUserData.CTFID = this.ctf.ID
      this.addCTFUserFormVisible = true
    },
    ClearAddCTFUserForm() {
      this.addCTFUserData = {
        "CTFID": 0,
        "UserID": null
      }
    },
    RemoveCTFUser() {
      ctfUserApi.RemoveCTFUser(this.removeCTFUserData).then(res => {
        if (res.code === 0) {
          this.removeCTFUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetCTF()
          this.GetCTFUserList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenRemoveCTFUserForm(row) {
      this.removeCTFUserData.CTFID = this.ctf.ID
      this.removeCTFUserData.UserID = row.ID
      this.removeCTFUserFormVisible = true
    },
    ClearRemoveCTFUserForm() {
      this.removeCTFUserData = {
        "CTFID": 0,
        "UserID": null
      }
    },

    OpenCTFTeamDetail(row) {
      this.ctfTeamDetail = {
        "ID": row.ID,
        "Name": row.Name,
        "User": row.User.Username,
        "Description": row.Description,
        "Users": row.Users,
        "CTFChallenges": row.CTFChallenges
      }
      this.ctfTeamDetailVisible = true
    },
    ClearCTFTeamDetail() {
      this.ctfTeamDetail = {
        "CTFTeamID": 0,
        "Name": "",
        "User": "",
        "Description": "",
        "Users": [],
        "CTFChallenges": []
      }
    },
    async GetCTFTeamList(){
      return ctfTeamApi.GetCTFTeamList({
        "Name": this.ctfTeamQueryName,
        "CTFID": this.ctf.ID
      }).then(res => {
        if (res.code === 0) {
          this.ctfTeamList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateCTFTeam(){
      ctfTeamApi.CreateCTFTeam(this.createCTFTeamData).then(res => {
        if (res.code === 0) {
          this.createCTFTeamFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFTeamList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCreateCTFTeamForm() {
      this.createCTFTeamData.CTFID = this.ctf.ID
      this.createCTFTeamFormVisible = true
    },
    ClearCreateCTFTeamForm() {
      this.createCTFTeamData = {
        "Name": "",
        "Description": "",
        "UserID": null,
        "CTFID": 0
      }
    },
    UpdateCTFTeam(){
      ctfTeamApi.UpdateCTFTeam(this.updateCTFTeamData).then(res => {
        if (res.code === 0) {
          this.updateCTFTeamFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFTeamList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateCTFTeamForm(row) {
      this.updateCTFTeamData = {
        "ID": row.ID,
        "Name": row.Name,
        "Description": row.Description,
        "UserID": row.UserID,
        "CTFID": row.CTFID
      }
      this.updateCTFTeamFormVisible = true
    },
    ClearUpdateCTFTeamForm() {
      this.updateCTFTeamData = {
        "ID": 0,
        "Name": "",
        "Description": "",
        "UserID": null,
        "CTFID": 0
      }
    },
    DeleteCTFTeam() {
      ctfTeamApi.DeleteCTFTeam(this.deleteCTFTeamData).then(res => {
        if (res.code === 0) {
          this.deleteCTFTeamFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetCTFTeamList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteCTFTeamForm(row) {
      this.deleteCTFTeamData = {
        "ID": row.ID,
      }
      this.deleteCTFTeamFormVisible = true
    },
    ClearDeleteCTFTeamForm() {
      this.deleteCTFTeamData = {
        "ID": 0,
      }
    },

    AddCTFTeamUser() {
      ctfTeamUserApi.AddCTFTeamUser(this.addCTFTeamUserData).then(async res => {
        if (res.code === 0) {
          this.addCTFTeamUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          await this.GetCTFTeamList()
          this.ctfTeamDetail = this.ctfTeamList.find(ctfTeam => ctfTeam.ID === this.ctfTeamDetail.ID)
          this.ctfTeamDetail.User = this.ctfTeamDetail.User.Username
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenAddCTFTeamUserForm(teamID) {
      this.addCTFTeamUserData.CTFTeamID = teamID
      this.addCTFTeamUserFormVisible = true
    },
    ClearAddCTFTeamUserForm() {
      this.addCTFTeamUserData = {
        "CTFTeamID": 0,
        "UserID": null
      }
    },
    RemoveCTFTeamUser() {
      ctfTeamUserApi.RemoveCTFTeamUser(this.removeCTFTeamUserData).then(async res => {
        if (res.code === 0) {
          this.removeCTFTeamUserFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFTeamList()
          await this.GetCTFTeamList()
          this.ctfTeamDetail = this.ctfTeamList.find(ctfTeam => ctfTeam.ID === this.ctfTeamDetail.ID)
          this.ctfTeamDetail.User = this.ctfTeamDetail.User.Username
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenRemoveCTFTeamUserForm(teamID, userID) {
      this.removeCTFTeamUserData = {
        "CTFTeamID": teamID,
        "UserID": userID
      }
      this.removeCTFTeamUserFormVisible = true
    },
    ClearRemoveCTFTeamUserForm() {
      this.removeCTFTeamUserData = {
        "CTFTeamID": 0,
        "UserID": null
      }
    },

  },
  mounted() {
    this.ctf.ID = Number(this.$route.params.id)
    this.GetCTF()
    this.GetCTFChallengeCategoryList()
    this.GetImageList()
    this.GetAttachmentList()
    this.GetCTFChallengeList()
    this.GetUserList()
    this.GetCTFUserList()
    this.GetCTFTeamList()
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
