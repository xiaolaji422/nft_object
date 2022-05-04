<template>
  <div class="app-container">
    <!--查询-->
  <app-table
    :config="tableConfig"
    ref="appTableRef"
    @selection-change="getSelection"
    tableName="userTable"
  >
  <template #index >
      <el-table-column  type="selection" :width="55"></el-table-column>
  </template>
    <template #header-action>  
      <el-button type="primary" :disable="selectionIds.length <=0" @click="dialogShow('addUserRolesBatch')">批量设置角色</el-button>
    </template>
    <!--  操作按钮 -->
    <template #action-before="{scope}">
      <el-button type="text" @click="dialogShow('userRole',scope)">管理角色</el-button>
      <el-button type="text" @click="dialogShow('userApis',scope)">定制权限</el-button>
    </template>
  </app-table>
  <!-- 弹窗 -->
  <UserRole ref="userRoleRef"  @refresh="refreshData"> </UserRole>
  <UserApis ref="userApisRef" @refresh="refreshData"> </UserApis>
  <AddAdminRole  ref="addAdminRole"  @refresh="refreshData" ></AddAdminRole>
  </div>
</template>


<script lang="ts" setup>
import permissionApi from '@/api/permission'
import { defineComponent, ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
import AppTable from '@/components/AppTable.vue'
import { getTableList,showFAQ,unshowFAQ,deleteFAQ } from '@/api/faq-conf-manage/index'
import UserRole from "./components/UserRole.vue"
import UserApis from "./components/UserApis.vue"
import AddAdminRole from "./components/AddAdminRole.vue"
import {permissionStore} from '@/store/modules/permission'


  //  faq配置列表
  const name =  'FaqConfTable'
  const components = defineComponent([
    AppTable,
    UserRole,
    UserApis,
    AddAdminRole
    ])

  const tableConfig = reactive({
      tableColumn: [
        {
          slot:"index"
        },
        {
          label: '最近编辑时间',
          width: 210,
          prop:"update_time",
          align: 'center'
        },
        {
          label: '用户全称',
          align: 'center',
          prop: 'full_name',
          width: 220
        },
        {
          align: 'center',
          label: '登录名',
          prop: 'login_name',
          width: 166
        },
        {
          align: 'center',
          label: '系统角色',
          prop: 'role_name',
          width: 166
        },
        {
          align: 'center',
          label: '部门名',
          prop: 'department_name',
          width: 166
        },
        {
          align: 'center',
          label: '职位名称',
          prop: 'post_name',
          width: 166
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
      tableListApi: (params:any) => permissionApi.getActivityList(params),
      tableListParams: {},
      tableDeleteApi: (id: number) => deleteFAQ(id),
      tableDeleteParams: {},
      tableSearch: [
        {
          label: 'RTX',
          type: 'text',
          key: 'login_name',
          inputType: 'text',
          placeholder: '请输入RTX搜索',
          clearable: true,
        },
         {
          label: '部门',
          type: 'text',
          key: 'department_name',
          inputType: 'text',
          placeholder: '所属部门搜索',
          clearable: true
        },
      ],
      tableHeaderRight: true
    })
  const appTableRef = ref(null)
  let router = useRouter()
  const handleAdd = () => {
    router.push({
      //传递参数使用query的话，指定path或者name都行，但使用params的话，只能使用name指定
      path: `/faq-conf/index/modify`
    })

  }
const roleData = ref([])
const  initData = async function(){
  let roleDataRes = await permissionStore().getRoleData()
  roleData.value.push(roleDataRes) 
}
onMounted(async ()=>{
  initData()
})

const selectionIds = ref([])
const getSelection = function(val){
  let data = []
  if(val){
    val.forEach(item => {
      data.push(item.id)
    });
  }
  selectionIds.value = data
}
  const dialogName = ref("")
  let handleInfo = reactive({})
  const multipleSelection = ref([])
    // 编辑权限
  const  refreshData = () =>{
    appTableRef.value.onSearchSubmit()
  }

  const userRoleRef = ref(null)
  const userApisRef = ref(null)
  // 多人授权
  const addAdminRole = ref(null)

  const dialogShow = function(dn,info){
    switch(dn){
      case "userRole":
        userRoleRef.value.show(info)
        break
      case "addUserRolesBatch":
        //  多人授权
        if (selectionIds.value.length <= 0 ){
          $notificatioWarn("请选择要筛选的员工")
          return 
        }
        addAdminRole.value.show(selectionIds.value)
        break
      case "userApis":
        userApisRef.value.show(info)
        break
    }
    return 
  }
    

</script>
