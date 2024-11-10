<template>
  <el-container>
    <el-header>
      <div style="display: flex; align-items: center;">
        <h2 style="color: var(--el-text-color-primary);">{{ t('AdminCTF.CTFManager') }}</h2>
        <div style="flex-grow: 1;"></div>
        <div style="margin-right: 50px;">
          <el-input
            v-model="ctfQueryName"
            style="width: 450px; margin: 10px;"
            :placeholder="$t('AdminCTF.FindCTFs')"
          >
            <template #append>
              <el-button
                @click="GetCTFList"
                style="width: 100px;"
              >
                {{ $t('AdminCTF.Find') }}
                <el-icon style="margin-left: 10px">
                  <Search fill="var(var(--el-button-text-color))"/>
                </el-icon>
              </el-button>
            </template>
          </el-input>
          <el-button
            style="margin: 10px;"
            @click="createCTFFormVisible = true"
            type="primary"
          >
            {{ t('AdminCTF.Add') }}
          </el-button>
        </div>
      </div>
    </el-header>
    <el-main>
      <el-table
        :data="ctfList"
        table-layout="fixed"
        height="100%"
      >
        <el-table-column prop="Name" :label="t('AdminCTF.Name')"/>
        <el-table-column prop="Description" :label="t('AdminCTF.Description')"/>
        <el-table-column :label="t('AdminCTF.Public')">
          <template #default="scope">
            <el-text v-if="scope.row.Public">
              {{ t('AdminCTF.True') }}
            </el-text>
            <el-text v-else>
              {{ t('AdminCTF.False') }}
            </el-text>
          </template>
        </el-table-column>
        <el-table-column :label="t('AdminCTF.TeamSignIn')">
          <template #default="scope">
            <el-text v-if="scope.row.TeamSignIn">
              {{ t('AdminCTF.True') }}
            </el-text>
            <el-text v-else>
              {{ t('AdminCTF.False') }}
            </el-text>
          </template>
        </el-table-column>
        <el-table-column :label="t('AdminCTF.StartTime')">
          <template #default="scope">
            <el-text>
              {{ formatDateTime(scope.row.StartTime) }}
            </el-text>
          </template>
        </el-table-column>
        <el-table-column :label="t('AdminCTF.EndTime')">
          <template #default="scope">
            <el-text>
              {{ formatDateTime(scope.row.EndTime) }}
            </el-text>
          </template>
        </el-table-column>
        <el-table-column fixed="right" :label="t('AdminCTF.Operations')" width="300">
          <template #default=scope>
            <el-button
              @click="OpenCTFDetail(scope.row)"
            >
              {{ $t('AdminCTF.Detail') }}
            </el-button>
            <el-button
              @click="OpenUpdateCTFForm(scope.row)"
            >
              {{ t('AdminCTF.Update') }}
            </el-button>
            <el-button
              @click="OpenDeleteCTFForm(scope.row)"
            >
              {{ t('AdminCTF.Delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createCTFFormVisible"
        :title="t('AdminCTF.CreateCTF')"
        width="500"
        @close="ClearCreateCTFForm"
      >
        <el-form :model="createCTFData">
          <el-form-item :label="t('AdminCTF.Name')" :label-width="labelWidth">
            <el-input v-model="createCTFData.Name"/>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.Description')" :label-width="labelWidth">
            <el-input v-model="createCTFData.Description" type="textarea" autosize/>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.Public')" :label-width="labelWidth">
            <el-select v-model="createCTFData.Public">
              <el-option
                key="true"
                :label="t('AdminCTF.True')"
                :value="true"
                />
              <el-option
                key="false"
                :label="t('AdminCTF.False')"
                :value="false"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.TeamSignIn')" :label-width="labelWidth">
            <el-select v-model="createCTFData.TeamSignIn">
              <el-option
                key="true"
                :label="t('AdminCTF.True')"
                :value="true"
                />
              <el-option
                key="false"
                :label="t('AdminCTF.False')"
                :value="false"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.StartTime')" :label-width="labelWidth">
            <el-date-picker
              v-model="createCTFData.StartTime"
              type="datetime"
              :placeholder="t('AdminCTF.SelectTime')"
            />
          </el-form-item>
          <el-form-item :label="t('AdminCTF.EndTime')" :label-width="labelWidth">
            <el-date-picker
              v-model="createCTFData.EndTime"
              type="datetime"
              :placeholder="t('AdminCTF.SelectTime')"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateCTF"
            type="primary"
          >
            {{ t('AdminCTF.Submit') }}
          </el-button>
          <el-button
            @click="createCTFFormVisible = false"
          >
            {{ t('AdminCTF.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateCTFFormVisible"
        :title="t('AdminCTF.UpdateCTF')"
        width="500"
        @close="ClearUpdateCTFForm"
      >
        <el-form :model=updateCTFData>
          <el-form-item :label="t('AdminCTF.Name')" :label-width="labelWidth">
            <el-input v-model="updateCTFData.Name"/>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.Description')" :label-width="labelWidth">
            <el-input v-model="updateCTFData.Description" type="textarea" autosize/>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.Public')" :label-width="labelWidth">
            <el-select v-model="updateCTFData.Public">
              <el-option
                key="true"
                :label="t('AdminCTF.True')"
                :value="true"
              />
              <el-option
                key="false"
                :label="t('AdminCTF.False')"
                :value="false"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.TeamSignIn')" :label-width="labelWidth">
            <el-select v-model="updateCTFData.TeamSignIn">
              <el-option
                key="true"
                :label="t('AdminCTF.True')"
                :value="true"
                />
              <el-option
                key="false"
                :label="t('AdminCTF.False')"
                :value="false"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('AdminCTF.StartTime')" :label-width="labelWidth">
            <el-date-picker
              v-model="updateCTFData.StartTime"
              type="datetime"
              :placeholder="t('AdminCTF.SelectTime')"
            />
          </el-form-item>
          <el-form-item :label="t('AdminCTF.EndTime')" :label-width="labelWidth">
            <el-date-picker
              v-model="updateCTFData.EndTime"
              type="datetime"
              :placeholder="t('AdminCTF.SelectTime')"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateCTF"
            type="primary"
          >
            {{ t('AdminCTF.Submit') }}
          </el-button>
          <el-button
            @click="updateCTFFormVisible = false"
          >
            {{ t('AdminCTF.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteCTFFormVisible"
        :title="t('AdminCTF.DeleteCTF')"
        width="500"
        @close="ClearDeleteCTFForm"
      >
        <el-text>{{ t('AdminCTF.AreYouConfirmToDeleteTheCTF') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteCTF"
            type="primary"
          >
            {{ t('AdminCTF.Confirm') }}
          </el-button>
          <el-button
            @click="deleteCTFFormVisible = false"
          >
            {{ t('AdminCTF.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import ctfApi from "@/api/ctf.js"
import Search from "@/components/icons/Search.vue"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  components: { Search },
  data() {
    return {
      labelWidth: 100,
      ctfQueryName: "",
      ctfList: [],
      createCTFFormVisible: false,
      createCTFData: {
        "Name": "",
        "Description": "",
        "Public": false,
        "TeamSignIn": false,
        "StartTime": null,
        "EndTime": null
      },
      updateCTFFormVisible: false,
      updateCTFData: {
        "ID": 0,
        "Name": "",
        "Description": "",
        "Public": false,
        "TeamSignIn": false,
        "StartTime": null,
        "EndTime": null
      },
      deleteCTFFormVisible: false,
      deleteCTFData: {
        "ID": 0,
      }
    }
  },
  methods: {
    GetCTFList() {
      ctfApi.GetCTFList({
        "Name": this.ctfQueryName
      }).then(res => {
        if (res.code === 0) {
          this.ctfList = res.data
          this.ctfList.forEach(item => {
            item.StartTime = new Date(item.StartTime)
            item.EndTime = new Date(item.EndTime)
        });
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenCTFDetail(row) {
      this.$router.push(`/admin/ctf/${row.ID}`)
    },
    CreateCTF() {
      ctfApi.CreateCTF(this.createCTFData).then(res => {
        if (res.code === 0) {
          this.createCTFFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateCTFForm() {
      this.createCTFData = {
        "Name": "",
        "Description": "",
        "Public": false,
        "TeamSignIn": false,
        "StartTime": null,
        "EndTime": null
      }
    },
    UpdateCTF() {
      ctfApi.UpdateCTF(this.updateCTFData).then(res => {
        if (res.code === 0) {
          this.updateCTFFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateCTFForm(row) {
      this.updateCTFData = {
        "ID": row.ID,
        "Name": row.Name,
        "Description": row.Description,
        "Public": row.Public,
        "TeamSignIn": false,
        "StartTime": row.StartTime,
        "EndTime": row.EndTime
      }
      this.updateCTFFormVisible = true
    },
    ClearUpdateCTFForm() {
      this.updateCTFData = {
        "ID": 0,
        "Name": "",
        "Description": "",
        "Public": false,
        "TeamSignIn": false,
        "StartTime": null,
        "EndTime": null
      }
    },
    DeleteCTF() {
      ctfApi.DeleteCTF(this.deleteCTFData).then(res => {
        if (res.code === 0) {
          this.deleteCTFFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCTFList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteCTFForm(row) {
      this.deleteCTFData = {
        "ID": row.ID,
      }
      this.deleteCTFFormVisible = true
    },
    ClearDeleteCTFForm() {
      this.deleteCTFData = {
        "ID": 0,
      }
    },
    formatDateTime(dateTime) {
      const year = dateTime.getFullYear()
      const month = String(dateTime.getMonth() + 1).padStart(2, '0')
      const day = String(dateTime.getDate()).padStart(2, '0')
      const hours = String(dateTime.getHours()).padStart(2, '0')
      const minutes = String(dateTime.getMinutes()).padStart(2, '0')
      return `${year}-${month}-${day} ${hours}:${minutes}`
    },
  },
  mounted() {
    this.GetCTFList()
  }
}
</script>
