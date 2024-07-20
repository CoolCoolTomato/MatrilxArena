<template>
  <el-container>
    <el-main>
      <el-scrollbar style="width: 80%; left: 10%;">
        <el-affix :offset="100">
          <div style="display: flex; margin: 0 20px 20px 20px;">
            <el-input :placeholder="$t('Challenge.FindChallenge')"/>
            <el-button type="primary" style="margin-left: 10px;">
              {{ $t('Challenge.Find') }}
              <el-icon style="margin-left: 10px">
                <Search fill="var(var(--el-button-text-color))"/>
              </el-icon>
            </el-button>
          </div>
        </el-affix>
        <el-row>
          <el-col v-for="group in groupList" :span="8">
            <div class="group">
              <h2 style="color: var(--el-text-color-primary)">{{ group.Name }}</h2>
              <el-text truncated>{{ group.Description }}</el-text>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
    </el-main>
  </el-container>
</template>
<script>
import groupApi from "@/api/group.js"
import { ElMessage } from "element-plus"
import {useI18n} from "vue-i18n"
import Search from "@/components/icons/Search.vue"

export default {
  components: { Search },
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      groupList: [],
    }
  },
  methods: {
    GetGroupList() {
      groupApi.GetGroupList().then(res => {
        if (res.code === 0) {
          this.groupList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
  },
  mounted() {
    this.GetGroupList()
  }
}
</script>
<style scoped>
.group {
  height: 200px;
  margin: 20px;
  border: 1px solid var(--el-border-color);
  cursor: pointer;
}
.group h2{
  margin: 20px;
}
.group .el-text{
  margin: 20px;
  width: calc(100% - 40px);
}
</style>
