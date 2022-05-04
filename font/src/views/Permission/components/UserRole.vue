<template>
  <el-dialog
    append-to-body
    destroy-on-close
    title="设置角色"
    v-model="showDialog"
    @close="closeDialog"
    width="65%"
    :close-on-click-modal="false"
  >
    <el-card class="box-card">
      <template #header>
        <div class="card-header header-flex" >
          <p>授权用户：<span style="color: red">{{ userinfo.full_name }}</span></p>
          <div>
            <el-select
              v-model="add_role_id"
              clearable
              placeholder="请选择"
              class="input-xl"
            >
              <el-option
                v-for="item in roleData"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
            <el-button class="button" type="text" @click="addRole">新增角色</el-button>
          </div>
        </div>
      </template>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="role_name" label="角色名称" width="180" />
        <el-table-column prop="modified_time" label="最后操作时间" width="180" />
        <el-table-column prop="modified_user" label="最后操作人" width="180" />
        <el-table-column prop="enabled" label="状态" >
          <template v-slot="{row}">
              <el-switch
                class="switch sm"
                v-model="row.enabled"
                :active-value="1"
                :inactive-value="0"
                active-text="启用"
                inactive-text="禁用"
                @change="enableAdminRole(row.id, row.enabled)"
              ></el-switch>
            </template>
        </el-table-column>
        <el-table-column prop="modified_user" label="最后操作人" width="180" >
          <template v-slot="{row}">
            <el-button class="button" type="text" @click="delRole(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
       
    </el-card>
  </el-dialog>
</template>
<script lang="ts" setup>
import permissionApi from "@/api/permission";
import { defineComponent,defineProps, onMounted,ref,reactive,defineExpose,defineEmits} from 'vue'
import AppTable from '@/components/AppTable.vue'
import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
import {permissionStore} from '@/store/modules/permission'
const components = defineComponent([
  AppTable,
])

const emit = defineEmits(["refresh"])
const showDialog = ref(false)
const add_role_id = ref(null)   // 新增角色的id
const tableData = ref([])
const roleData = ref([])
const userinfo  =ref()

onMounted(async () =>{
  initData()
})
const initData = async function() {
    let roleDataRes = await  permissionStore().getRoleData();
    roleData.value = roleDataRes;
  }
const enableAdminRole = function(id, enabled) {
  permissionApi.enableAdminRole({
    id: id,
    enabled: enabled,
  }).then(res => {
      $notificatioSuccess("操作成功")
    }
  ).catch(err=>{

  });
}

const delRole = (id) =>{
      var params = { id: id }
      permissionApi.delUserRole(params).then(res=> {
        $notificatioSuccess("操作成功")
        getTableData(userinfo.value.id)
      },err=>{
      });
    }
const getTableData = (admin_id)=>{
  permissionApi.getUserRoleList({admin_id:admin_id}).then(res=>{
     tableData.value = res.data??[]
  })
} 

const closeDialog = function(){
  add_role_id.value = ""
  showDialog.value = false
  emit("refresh",true)
}
const addRole = function(){
  let params = {
    role_id:add_role_id.value,
    admin_ids: userinfo.value.id
  }
  permissionApi.addUserRole(params).then(res=>{
      $notificatioSuccess("操作成功")
      closeDialog()
  },err=>{
      // this.$message.error("操作失败")
  })
}

const show = function(info){
  userinfo.value = info
  showDialog.value = true
  getTableData(userinfo.value.id)
}
defineExpose({
  show
})

</script>

<style lang="scss" scoped>
 .header-flex{
   justify-content: space-between;
   display: flex;
 }
</style>