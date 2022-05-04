<template>
  <div class="app-container">
    <app-table
      :config="tableConfig"
      ref="appTableRef"
      tableName="apiGroupTable"
    >
    <template #is_enabeld >
      <el-table-column v-slot="{row}" label="分组状态">
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
          <el-button type="primary" @click="dialogShow('add')">添加分组</el-button>
          <el-button type="primary" @click="dialogShow('apimanage')">管理接口</el-button>
      </template>
      <!--  操作按钮 -->
      <template #action-before="{scope}">
        <el-button type="text" @click="dialogShow('edit',scope)">查看/修改</el-button>
      </template>
  </app-table>
   <addApiGroup  ref="addApiGroupRef"  @refresh="refreshData"></addApiGroup>
  </div>
</template>
<script lang="ts" setup>
  import addApiGroup from './add-api-group.vue'
  import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
  import permissionApi from '@/api/permission'
  import { defineComponent ,ref, reactive, onMounted} from 'vue'
  import AppTable from '@/components/AppTable.vue'
  import {permissionStore} from '@/store/modules/permission'
  import { useRouter } from 'vue-router'

const components = defineComponent([
  addApiGroup,
  AppTable,
])

const tableConfig = reactive({
   tableColumn: [
        {
          label: '分组ID',
          prop:"id",
          align: 'center'
        },
        {
          label: '分组名称',
          align: 'center',
          prop: 'name',
        },
        {
          label: '创建时间',
          align: 'center',
          prop: 'create_time',
        },
        {
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
      tableListApi: (params:any) => permissionApi.getApiGroupList(params),
      tableListParams: {},
      tableDeleteApi: null,
      tableDeleteParams: {},
      tableSearch: [
         {
          label: '分组名称',
          type: 'text',
          key: 'name',
          inputType: 'text',
          placeholder:"请输入分组名称搜索",
          clearable: true,
        },
      ],
      tableHeaderRight: true
})

const appTableRef = ref(null)

const  refreshData = () =>{
  appTableRef.value.onSearchSubmit()
}

const addApiGroupRef = ref(null)
let router =  useRouter()
const dialogShow = function(dn,info){
    switch(dn){
      case "add":
       addApiGroupRef.value.show(info)
        break
      case "edit":
        //  多人授权
        addApiGroupRef.value.show(info)
        break
      case "apimanage":
        // 跳转路由
        router.push({
          path: `/permission/systemApi`,
        })
        break
    }
    return 
  }

const  initData = async function(){}
onMounted(async ()=>{
  initData()
})

const stateChange = (row)=>{
  var params={
      id:row.id,
      enabled:row.enabled
  }
    permissionApi.enabledApiGroup(params).then(res=>{
      $notificatioSuccess("操作成功")
      permissionStore().clearApiGroup()
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
