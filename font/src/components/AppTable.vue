<template>
  <div class="app-table-wrapper">
    <div class="table-header-wrapper">
      <el-form
        @submit.native.prevent
        v-if="appConfig.tableSearch.length !== 0"
        :inline="true"
        :model="tableSearchModel"
        ref="tableSearchRef"
        class="app-table-search"
      >
        <el-form-item
          v-for="(item, index) in appConfig.tableSearch"
          :key="index"
          :label="item.label"
          :required="item.required"
        >
          <!-- 输入框 -->
          <el-input
            v-if="item.type === 'text'"
            v-model="tableSearchModel[item.key]"
            :type="item.inputType ? item.inputType : 'text'"
            :name="item.key"
            :placeholder="item.placeholder"
            :clearable="item.clearable"
            :class="item.class"
            :disabled="loading"
            @keyup.enter.native="onSearchSubmit(true)"
          > 
            <template #suffix>
              <i class="el-icon-search icon-search-style" @click="onSearchSubmit(true)"></i>
            </template>
          </el-input>
          <!-- 下拉框 -->
          <el-select
            v-else-if="item.type === 'select'"
            v-model="tableSearchModel[item.key]"
            :filterable="item.filterable"
            :name="item.key"
            :placeholder="item.placeholder"
            :clearable="item.clearable"
            :class="item.class"
            :disabled="loading"
            @change="onSearchSubmit(true)"
          >
            <el-option
              v-for="(option, indexItem) in item.options"
              :key="indexItem"
              :label="option.label"
              :value="option.value"
              :disabled="loading"
            />
          </el-select>
          <!-- 日期时间选择器 -->
          <el-date-picker
            v-else-if="item.type === 'daterange'"
            v-model="tableSearchModel[item.key]"
            :align="item.align"
            :range-separator="item.rangeSeparator || '至'"
            :start-placeholder="item.startPlaceholder || '开始时间'"
            :end-placeholder="item.endPlaceholder || '结束时间'"
            :name="item.key"
            :picker-options="item.pickerOptions"
            :class="item.class"
            :disabled="loading"
            type="daterange"
          />
          <!-- 日期时间范围 -->
          <el-date-picker
            v-else-if="item.type === 'datetimerange'"
            v-model="tableSearchModel[item.key]"
            type="datetimerange"
            :disabled="loading"
            :class="item.class"
            @change="onSearchSubmit(true)"
            value-format="YYYY-MM-DD HH:mm:ss"
            :range-separator='item.rangeSeparator || "至"'
            :start-placeholder='item.startPlaceholder || "开始时间"'
            :end-placeholder='item.endPlaceholder || "结束时间"'
          >
          </el-date-picker>
          <!-- 日期选择器 -->
          <el-date-picker
            v-else-if="item.type === 'date'"
            v-model="tableSearchModel[item.key]"
            :name="item.key"
            :value-format="item.format"
            :placeholder="item.placeholder"
            :class="item.class"
            :disabled="loading"
            type="date"
          />
        </el-form-item>
        <slot name="header-form"  />
        <el-form-item :class="{'table-header-right':true}">
          <el-button
            v-if="appConfig.tableSearchBtnName.reset"
            :loading="loading"
            icon="el-icon-refresh-right"
            type="primary"
            @click="onSearchReset"
          >
            {{ appConfig.tableSearchBtnName.reset }}
          </el-button>
          <slot name="header-action" :loading="loading" />
        </el-form-item>
      </el-form>
      <div class="table-header-actions">
        <!-- <slot name="header-action" :loading="loading" /> -->
      </div>
    </div>

    <el-table
      v-loading="loading"
      v-bind="appConfig.tableAttr"
      :data="tableData"
      :border ="false"
      :stripe ="false"
      :header-cell-style="{background:'rgba(0, 0, 0, 0.05)'}"
      ref="tableRef"
      element-loading-text="加载数据中..."
      element-loading-spinner="el-icon-loading"
      @selection-change="handleSelectionChange"
      @sort-change="changeSort"
    >
      <template v-for="(item, index) in appConfig.tableColumn" :key="index">
        <!--多选框-->
        <el-table-column
          v-if="item.type && item.type === 'selection'"
          :width="item.width ? item.width : null"
          :align="item.align"
          type="selection"
        />
        <!--序列号-->
        <el-table-column
          v-else-if="item.type && item.type === 'index'"
          :label="item.label"
          :width="item.width ? item.width : null"
          :align="item.align"
          :fixed="item.fixed"
          type="index"
        />
        <!--自定义列插槽-->
        <slot v-else-if="item.slot" :name="item.slot" />
        <!--表格数据渲染-->
        <el-table-column
          v-else
          :label="item.label"
          :sortable="item.sortable"
          :width="item.width ? item.width : null"
          :min-width="item.minWidth ? item.minWidth : null"
          :align="item.align"
          :fixed="item.fixed"
        >
          <template v-slot="scope">
            <template v-if="item.action">
              <slot name="action-before" :scope="scope.row" />
              <template v-for="(act, idx) in item.action">
                <el-button
                  v-if="act.type === 'editBox'"
                  :key="idx"
                  type="text"
                  size="small"
                  class="font-normal"
                  :class="act.class"
                  @click="handleEdit(scope.row)"
                >
                  编辑
                </el-button>
                <slot name="action-middle" v-if='idx == 0' :scope="scope.row"/>
                <el-button
                  class="font-normal"
                  :key="idx"
                  type="text"
                  size="small"
                  v-if="act.type === 'view'"
                  :class="act.class"
                  @click="handleView(scope.row)"
                >
                  查看
                </el-button>

                <el-button
                  v-if="act.type === 'delete'"
                  :key="idx"
                  type="text"
                  size="small"
                  class="font-normal"
                  :class="act.class"
                  @click="handleDelete(scope.row)"
                >
                  删除
                </el-button>

              </template>
              <slot name="action-after" :scope="scope.row" />
            </template>
            <template v-else-if="item.dateTimeFormat">{{ scope.row[item.prop] }}</template>
            <template v-else>{{ filterVal(scope.row[item.prop]) }}</template>
          </template>
        </el-table-column>
      </template>
      <el-table-column  v-if="autoColumn"></el-table-column>
    </el-table>
    <!-- <div  class="custom_scroll " ref="customScrollRef" >
      <div class= "custom_scroll_body " :style="`width:${scrollWidth}`"></div>
    </div> -->
    <TableScroll @change="scrollChange" :scrollWidth="scrollWidth"></TableScroll>
    <el-pagination
      v-if="appConfig.isPagination"
      v-bind="appConfig.pagination"
      :current-page="pagination.page"
      :page-count="pagination.pageCount"
      :page-size="pagination.limit"
      :total="pagination.totalCount"
      :disabled="loading"
      class="app-pagination"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive, watch, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { forEach, merge, throttle } from 'lodash'
import { useApptore } from '@/store/modules/app'

import {cachePageStore} from '@/store/modules/cache-page'
import TableScroll from '@/components/TableScroll/index.vue'

export default defineComponent({
  name: 'AppTable',
  props: {
    config: {
      require: true,
      type: Object,
      default: () => {}
    },
    tableName:{
      require: true,
      type: String,
      default: "table",
    },
  },
  components:{
    TableScroll,
  },
  setup(props, { emit }) {
    const autoColumn = ref(true)
    const loading = ref(false)
    const tableData = ref([])
    const selectionRow = ref([])
    const store = useApptore()
    let sortParams = {}
    let appConfig = reactive(merge({
      getTableList: true,
      isPagination: true,
      isDelete: true,
      tableAttr: {
        stripe: true,
        border: true,
        fit: true,
        highlightCurrentRow: true
      },
      tableColumn: [],
      tableListApi: null,
      tableListParams: {}, // tableListApi 所需要的参数
      tableDeleteApi: null,
      tableDeleteParams: {},
      pagination: {
        pageSizes: [10, 20, 30, 40, 50, 100],
        layout: 'total, sizes, prev, pager, next, jumper',
        background: true
      },
      tableSearch: [],
      tableSearchBtnName: {
        reset: '重置'
      }
    }, props.config))

     // 手动实现滚动条
    const scrollWidth = ref("100%")
    const tableRef = ref(null)
    const scorllDom = ref(null)
    const customScrollRef =ref(null)
    
    onMounted(()=>{
      setTimeout(()=>{
        if(scorllDom.value == null){
          scorllDom.value = tableRef.value.$el.querySelector(".el-table__body-wrapper")
        }
        scrollWidth.value = tableRef.value.bodyWidth
      },200)
      window.onresize = () => {
        return (() => {
          scrollWidth.value = tableRef.value.bodyWidth
        })()
      }
    })

    for (let i =0;i<appConfig.tableColumn.length;i++){
      // 是否需要加入空格行
      let item = appConfig.tableColumn[i]
      if  ((!item.slot) && !(item.action && item.fiexd) &&  (!item.width) ) {
        autoColumn.value = false;
        break;
      }
    }
    watch(props.config, (a, b) => {
      // 搜索列表
      b.isNewTableSearch && (appConfig.tableSearch = b.tableSearch)
    })
    
    const tableSearchModel = reactive({})
    const pagination = reactive({
      page: 1,
      pageCount: 1,
      limit: 10,
      totalCount: 1
    })
    // 滚动条事件
    const scrollChange = function (val){
      let scorllLeft = val.scrollLeft
      scorllDom.value.scrollTo(scorllLeft,0)
    }
    // 全局对象类型  IObjType
    const handleEdit = (row: IObjType) => {
      emit('handleEdit', row)
    }
    const handleView = row => {
      emit('handleView', row)
    }
    const handleDelete = row => {
      if (!appConfig.isDelete) {
        return
      }
      ElMessageBox.confirm(
        '确认删除？',
        'Warning',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          title:"删除确认"
        }
      )
        .then(() => {
          if (row.confTitleId) {
            row.id = row.confTitleId
          }
          appConfig.tableDeleteApi({id:row.id}).then(data => {
            if (data.code === 0) {
              ElMessage.success('删除成功！')
              // 删除之后再当前页
              onSearchSubmit(false)
            } else {
              ElMessage.error(data.msg)
            }
          })
        })
    }
    /**
     * 搜索条件初始化
     */
    const initSearch = () => {
      appConfig.tableSearch.forEach(item => {
        tableSearchModel[item.key] = item.value ? item.value : ''
      })
      onSearchSubmit(false)
    }
    /**
     * 获取列表数据
     */
    const getTableListData = params => {
      if (!appConfig.getTableList) return
      loading.value = true
      appConfig.tableListApi(params).then((data: IObjType) => {
        if (data?.code === 0) {
          tableData.value = data?.data?.items
          pagination.pageCount = data?.data?.pages
          pagination.totalCount = data?.data?.total
        } else {
          
          ElMessage.error(data.msg)
        }
        loading.value = false
        emit('actionTableData', data)
      }).catch(() => {
        loading.value = false
      })
    }

    /**
     * 搜索按钮触发
    
     */
    const onSearchSubmit = throttle(async (refresh) => {
      pagination.page = 1
      onsubmit(refresh)
    }, 1000)

    // 缓存搜索条件
    const handleCacheSearchInfo = async (refresh)=>{
      let storeInfo = await cachePageStore().getPageInfo(props.tableName)
      if(storeInfo && !refresh){
        if(storeInfo.page)  pagination.page = storeInfo.page;
        if(storeInfo.limit)  pagination.limit = storeInfo.limit;
        if(storeInfo.tableSearchModel){
          for (var i in storeInfo.tableSearchModel){
           tableSearchModel[i] = storeInfo.tableSearchModel[i]
          }
        }
        if(storeInfo.sortParams)  sortParams = storeInfo.sortParams
        if(storeInfo.tableListParams)  appConfig.tableListParam = storeInfo.tableListParams
      }
      
    }
    /** 真实提交
     * 所有的搜索条件都在这里触发
     * 节流
     */
    const onsubmit = async (refresh)=>{
      // 看一下是否要缓存
      await handleCacheSearchInfo(refresh)
      // 搜集所有搜索条件
      const params = merge(
        appConfig.tableListParams,
        {
          page: pagination.page,
          limit: pagination.limit
        },
        tableSearchModel,
        sortParams
      )
      // 保存最终搜索信息到vuex
      let  cacheInfo={
        page: pagination.page,
        limit: pagination.limit,
        tableSearchModel:tableSearchModel,
        sortParams:sortParams,
        tableListParams:appConfig.tableListParams,
      }
      cachePageStore().setPageInfo(props.tableName,cacheInfo)

      getTableListData(params)
    }
    /**
     * 分页 - 当前页码
     * @param val
     */
    const handleCurrentChange = val => {
      // todo
      if (val !== null) {
        pagination.page = val
        // 直接提交
        onsubmit(true)
      }
    }
    /**
     * 分页 - 当前条数
     * @param val
     */
    const handleSizeChange = val => {
      pagination.limit = val
      onSearchSubmit(true)
    }
    /**
     * 过滤值
     * @param val
     * @returns {string|*}
     */
    const filterVal = val => {
      if (val === null || val === '' || val === undefined) {
        return '-'
      } else {
        return val
      }
    }
    /**
     * 重置搜索
     */
    const onSearchReset = () => {
      initSearch()
    }
    /**
     * 勾选列表回调
     * @param val
     */
    const handleSelectionChange = val => {
      selectionRow.value = val
      emit('selection-change', val)
    }

    initSearch()

    /**
     * table 排序 没有获取到prop
     */
    const changeSort = ({ column, prop, order }: any) => {
      let item = appConfig.tableColumn.find((item: IObjType) => column && item?.label === column.label)

      sortParams = {
        _order_property: item?.prop,
        _order_sort: order === 'ascending' ? 'asc' : order ? 'desc' : ''
      }
      onSearchSubmit(true)
    }
    
    return {
      autoColumn,
      loading,
      appConfig,
      tableSearchModel,
      tableData,
      pagination,
      selectionRow,
      handleEdit,
      handleView,
      handleDelete,
      onSearchSubmit,
      handleCurrentChange,
      handleSizeChange,
      onSearchReset,
      filterVal,
      handleSelectionChange,
      changeSort,
      tableRef,
      scrollWidth,
      customScrollRef,
      scrollChange,
    }
  }
})
</script>

<style lang="scss" scoped>

.app-table-wrapper {
  width: 98%;
  .table-header-wrapper {
    display: flex;
    justify-content: space-between;
    .app-table-search {
      :deep .el-form-item:last-child {
        margin-right: 0;
      }
    }
  }
  .app-pagination {
    text-align: right;
    margin-top: 20px;
  }
  :deep .el-table__body-wrapper::-webkit-scrollbar{
    width: 0;
  }
}
.table-header-wrapper{
  .el-form{
    width: 100%;
    .table-header-right{
      float:right;
    }
  }
}
.icon-search-style{
  cursor: pointer;
  font-size: 18px;
  position: relative;
  right: 10px;
}

// 滚动条
.custom_scroll{
  position: sticky;
  bottom:12px;
  height: 20px;
  width: calc(100%);
  z-index: 999;
  background-color:#fff ;
   overflow-x:auto ;
   margin-top:-20px ;
  .custom_scroll_body{
    height: 1px;
  }
}
.custom_scroll::-webkit-scrollbar {
    width:12px;
}
/* 滚动槽 */
.custom_scroll::-webkit-scrollbar-track {
  border-radius:10px;
}
/* 滚动条滑块 */
.custom_scroll::-webkit-scrollbar-thumb {
  border-radius:10px;
  background:rgba(144, 147, 153,0.5);
  // opacity: 0.3;
}

.custom_scroll::-webkit-scrollbar-thumb:hover {
  background:rgba(144, 147, 153,0.8);
  // opacity: 0.5;
}
</style>
