<template>
  <div class="app-container">
    <app-table
      :config="tableConfig"
      ref="appTableRef"
      tableName="apiTable"
    >
    <template #is_enabeld >
      <el-table-column v-slot="{row}">
          <el-switch
          v-model="row.enabled"
            class="switch sm"
            active-color='#265BEDCC'
            inactive-color='#909399'
            active-text="启用"
            inactive-text="禁用"
            :active-value="1"
            :inactive-value="0"
            @click="stateChange(row)"
          >
        </el-switch>
      </el-table-column>
    </template>
      <template #header-action>  
          <el-button type="primary" @click="dialogShow('add')">添加接口</el-button>
          <el-button type="primary" @click="dialogShow('group')">管理分组</el-button>
      </template>
      <!--  操作按钮 -->
      <template #action-before="{scope}">
        <el-button type="text" @click="dialogShow('edit',scope)">查看/修改</el-button>
      </template>
  </app-table>
   <addApi  ref="addApiRef"  @refresh="refreshData"></addApi>
  </div>
</template>
<script lang="ts" setup>
  import addApi from './components/add-api.vue'
  import groupApi from './apiGroup.vue'
  import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
  import permissionApi from '@/api/permission'
  import { defineComponent ,ref, reactive, onMounted} from 'vue'
  import AppTable from '@/components/AppTable.vue'
  import {permissionStore} from '@/store/modules/permission'
  import { useRouter } from 'vue-router'

const components = defineComponent([
  addApi,
  groupApi,
  AppTable,
])

const tableConfig = reactive({
   tableColumn: [
        {
          label: '接口ID',
          width: 210,
          prop:"id",
          align: 'center'
        },
        {
          label: '接口名称',
          align: 'center',
          prop: 'name',
          width: 220
        },
        {
          label: '接口路由',
          align: 'center',
          prop: 'route',
          width: 166
        },
        {
          label: '接口限频',
          align: 'center',
          prop: 'limit',
          width: 166
        },
        {
          label: '请求方法',
          align: 'center',
          prop: 'methods',
          width: 166
        },
         {
          label: '最后操作时间',
          align: 'center',
          prop: 'modified_time',
          width: 166
        },{
          slot:"is_enabeld"
        },
        {
          width: 200,
          fixed: 'right',
          label: '操作',
          action: []
        }
       
      ],
      tableSearchBtnName: {
        reset: false
      },
      tableListApi: (params:any) => permissionApi.getApiList(params),
      tableListParams: {},
      tableDeleteApi: null,
      tableDeleteParams: {},
      tableSearch: [
        {
          label: '接口路由搜索',
          type: 'text',
          key: 'route',
          inputType: 'text',
          placeholder:"请输入接口路由搜索",
          clearable: true,
        },
      ],
      tableHeaderRight: true
})

const appTableRef = ref(null)

const  refreshData = () =>{
  appTableRef.value.onSearchSubmit()
}

const addApiRef = ref(null)
let router =  useRouter()
const dialogShow = function(dn,info){
    switch(dn){
      case "add":
       addApiRef.value.show(info)
        break
      case "edit":
        //  多人授权
        addApiRef.value.show(info)
        break
      case "group":
        // 跳转路由
        router.push({
          path: `/permission/systemApiGroup`,
        })
        break
    }
    return 
  }

const  initData = async function(){
  let dataRes = await permissionStore().getApiGroupData()
  // ApiGroupData.value = dataRes
  tableConfig.isNewTableSearch = true
  // tableConfig.tableSearch[1].options = dataRes
}
onMounted(async ()=>{
  initData()
})

const stateChange = (row)=>{
  var params={
      id:row.id,
      enabled:row.enabled
  }
    permissionApi.enableApi(params).then(res=>{
      $notificatioSuccess("操作成功")
      permissionStore().clearApi()
    },err=>{
      row.enabled = row.enabled?0:1
    })
}



</script>


<style lang="scss" scoped>
  .group-action-bar {
    display: flex;
    justify-content: space-between;
    padding-bottom: 20px;

    .date-input-search-btn {
      display: flex;

      .search-input {
        display: flex;
        margin-left: 10px;

        .search-btn {
          margin-left: 10px;
        }
      }
    }
  }

  .page-wrap {
    display: flex;
    justify-content: flex-end;
    margin-top: 20px;
  }
</style>
