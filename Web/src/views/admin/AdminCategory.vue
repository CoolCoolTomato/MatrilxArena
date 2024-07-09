<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">{{ $t('AdminCategory.CategoryManager') }}</h2>
    </el-header>
    <el-main>
      <el-scrollbar>
        <el-button
          style="margin: 10px"
          @click="createCategoryFormVisible = true"
          type="primary"
          >
          {{ $t('AdminCategory.Add') }}
        </el-button>
        <el-switch
          style="margin: 10px"
          @change="switchAllowSort"
          v-model="allowSort"
          :active-text="$t('AdminCategory.AllowSort')"
        />
        <el-table
          id="categoryListTable"
          :data="categoryList"
          table-layout="fixed"
          row-key="ID"
        >
          <el-table-column prop="Name" :label="$t('AdminCategory.Name')"/>
          <el-table-column prop="Order" :label="$t('AdminCategory.Order')"/>
          <el-table-column fixed="right" :label="$t('AdminCategory.Operations')" width="300px">
            <template #default=scope>
              <el-button
                @click="OpenUpdateCategoryForm(scope.row)"
              >
                {{ $t('AdminCategory.Update') }}
              </el-button>
              <el-button
                @click="OpenDeleteCategoryForm(scope.row)"
              >
                {{ $t('AdminCategory.Delete') }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-scrollbar>
      <el-dialog
        v-model="createCategoryFormVisible"
        :title="$t('AdminCategory.CreateCategory')"
        width="500"
        @close="ClearCreateCategoryForm"
      >
        <el-form :model=createCategoryData>
          <el-form-item :label="$t('AdminCategory.Name')" :label-width="labelWidth">
            <el-input v-model="createCategoryData.Name"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCategory.Order')" :label-width="labelWidth">
            <el-input v-model.number="createCategoryData.Order" type="number"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateCategory"
            type="primary"
            >
            {{ $t('AdminCategory.Submit') }}
          </el-button>
          <el-button
            @click="createCategoryFormVisible = false"
            >
            {{ $t('AdminCategory.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateCategoryFormVisible"
        :title="$t('AdminCategory.UpdateCategory')"
        width="500"
        @close="ClearUpdateCategoryForm"
      >
        <el-form :model=updateCategoryData>
          <el-form-item :label="$t('AdminCategory.Name')" :label-width="labelWidth">
            <el-input v-model="updateCategoryData.Name"/>
          </el-form-item>
          <el-form-item :label="$t('AdminCategory.Order')" :label-width="labelWidth">
            <el-input v-model.number="updateCategoryData.Order" type="number"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateCategory"
            type="primary"
            >
            {{ $t('AdminCategory.Submit') }}
          </el-button>
          <el-button
            @click="updateCategoryFormVisible = false"
            >
            {{ $t('AdminCategory.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteCategoryFormVisible"
        :title="$t('AdminCategory.DeleteCategory')"
        width="500"
        @close="ClearDeleteCategoryForm"
      >
        <el-text>{{ $t('AdminCategory.AreYouConfirmToDeleteTheCategory') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteCategory"
            type="primary"
            >
            {{ $t('AdminCategory.Confirm') }}
          </el-button>
          <el-button
            @click="deleteCategoryFormVisible = false"
            >
            {{ $t('AdminCategory.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import categoryApi from "@/api/category.js"
import { ElMessage } from "element-plus"
import Sortable from 'sortablejs'
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      labelWidth: 100,
      categoryList: [],
      createCategoryFormVisible: false,
      createCategoryData: {
        "Name": "",
        "Order": 0,
      },
      updateCategoryFormVisible: false,
      updateCategoryData: {
        "Name": "",
        "Order": 0,
      },
      deleteCategoryFormVisible: false,
      deleteCategoryData: {
        "ID": 0,
      },
      sortable: null,
      allowSort: false,
    }
  },
  methods: {
    GetCategoryList() {
      categoryApi.GetCategoryList().then(res => {
        if (res.code === 0) {
          this.categoryList = res.data
          this.categoryList.sort((a, b) => a.Order - b.Order)
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateCategory() {
      categoryApi.CreateCategory(this.createCategoryData).then(res => {
        if (res.code === 0) {
          this.createCategoryFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCategoryList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateCategoryForm() {
      this.createCategoryData = {
        "Name": "",
        "Order": 0,
      }
    },
    UpdateCategory() {
      categoryApi.UpdateCategory(this.updateCategoryData).then(res => {
        if (res.code === 0) {
          this.updateCategoryFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCategoryList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateCategoryForm(row) {
      this.updateCategoryData = {
        "ID": row.ID,
        "Name": row.Name,
        "Order": row.Order,
      }
      this.updateCategoryFormVisible = true
    },
    ClearUpdateCategoryForm() {
      this.updateCategoryData = {
        "ID": 0,
        "Name": "",
        "Order": 0,
      }
    },
    DeleteCategory() {
      categoryApi.DeleteCategory(this.deleteCategoryData).then(res => {
        if (res.code === 0) {
          this.deleteCategoryFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetCategoryList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteCategoryForm(row) {
      this.deleteCategoryData = {
        "ID": row.ID,
      }
      this.deleteCategoryFormVisible = true
    },
    ClearDeleteCategoryForm() {
      this.deleteCategoryData = {
        "ID": 0,
      }
    },
    initSortable() {
      const categoryListTable = document.querySelector("#categoryListTable tbody")
      this.sortable = new Sortable(categoryListTable, {
        disabled: !this.allowSort,
        animation: 100,
        onEnd: (evt) => {
          const newIndex = evt.newIndex
          const oldIndex = evt.oldIndex
          const movedItem = this.categoryList.splice(oldIndex, 1)[0]
          this.categoryList.splice(newIndex, 0, movedItem)
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
      this.categoryList.forEach((category, index) => {
        if (category.Order !== index + 1) {
          total += 1
          category.Order = index + 1
          const updatePromise = categoryApi.UpdateCategory(category).then(res => {
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
      if (total === 0) {
        return
      }
      Promise.all(updatePromises).then(() => {
        if (fail === 0) {
          ElMessage({
            message: `${this.t('AdminCategory.Update')} ${total} ${this.t('AdminCategory.items')}, ${success} ${this.t('AdminCategory.success')}, ${fail} ${this.t('AdminCategory.fail')}`,
            type: 'success',
          })
        } else {
          ElMessage({
            message: `${this.t('AdminCategory.Update')} ${total} ${this.t('AdminCategory.items')}, ${success} ${this.t('AdminCategory.success')}, ${fail} ${this.t('AdminCategory.fail')}`,
            type: 'warning',
          })
        }
        this.GetCategoryList()
      }).catch(error => {
        console.log(error)
      })
    },
  },
  mounted() {
    this.GetCategoryList()
    this.initSortable()
  }
}
</script>
<style scoped>
::v-deep(#categoryListTable .el-table__row:hover) {
  cursor: pointer;
}
</style>
