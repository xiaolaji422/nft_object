<template>
  <el-dialog
    append-to-body
    destroy-on-close
    title="添加角色"
    v-model="showDialog"
    @close="closeDialog"
    width="40%"
    top="30vh"
    :close-on-click-modal="false"
  >
  <div class="select-body">
    选择角色：<el-select
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
  </div>
    
    <div slot="footer" class="dialog-footer">
      <el-button @click="closeDialog"> 取消 </el-button>
      <el-button type="primary" @click="subInfo"> 确认 </el-button>
    </div>
  </el-dialog>
</template>
<script lang="ts" setup>
import permissionApi from '@/api/permission'
import { defineComponent,defineProps, onMounted,ref,reactive,defineExpose,defineEmits} from 'vue'
import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
import {permissionStore} from '@/store/modules/permission'


const roleData = ref([])
const initData = async function() {
  let roleDataRes = await  permissionStore().getRoleData();
  roleData.value = roleDataRes;
}
const showDialog = ref(false)
const admin_ids = ref([])
const emit = defineEmits(["refresh"])
const show = function(ids){
  admin_ids.value = ids
  showDialog.value = true
}

defineExpose({
  show
})

const add_role_id = ref()
const closeDialog = function(refresh){
  showDialog.value = false
  if(refresh == true){
    emit("refresh",true)
  }
  
}

const  addRoleFormRef = ref()
const   subInfo = ()=> {
  var params ={
    role_id:add_role_id.value,
    admin_ids:admin_ids.value
  }
  if(admin_ids.value == undefined  || admin_ids.value.length <=0){
    $notificatioWarn("操作对象不能为空")
    return
  }
 
  permissionApi.addUserRole(params).then(res=>{
      $notificatioSuccess("操作成功")
      closeDialog(true)
  },err=>{
  })
    
}
onMounted(async () =>{
  initData()
})



</script>
<style lang="scss" scoped>
 .select-body{
   margin:40px;
 }
 .input-xl{
   width:400px;
 }
</style>