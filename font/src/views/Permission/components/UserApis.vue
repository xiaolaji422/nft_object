<template>
  <el-dialog
    destroy-on-close
    append-to-body
    v-if="showDialog"
    title="定制权限"
    v-model="showDialog"
    @close="closeDialog"
    width="60%"
    :close-on-click-modal="false"
  >
    <div class="head-flex">
        <p>授权用户：<span style="color:red">{{userinfo.full_name}}</span></p>
        <el-button type="primary" class="submitRight" @click="submitUserApis">提交</el-button>
    </div>
    <SelectApi :key="userinfo.id"  v-model="newAuths"></SelectApi>
  </el-dialog>
</template>
<script lang="ts" setup>
import SelectApi from './select_apis.vue'
import permissionApi from "@/api/permission";
import { defineComponent,ref,defineExpose,defineEmits} from 'vue'

import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
const name = "UserApis"
const components = defineComponent([
  SelectApi
])
const roleApi = ref([])
const showDialog = ref(false)
const userinfo = ref()
const emit = defineEmits(["refresh"])
const show = function(info){
  userinfo.value = info
  showDialog.value = true
  getAdminApis(userinfo.value.id)
}

// 原有权限
const oldAuths = ref([])
// 现在的权限
const newAuths = ref([])
// 获取现有权限
const getAdminApis = (admin_id)=>{
  permissionApi.getAdminApis({admin_id:admin_id}).then(res=>{
    oldAuths.value = Object.assign([],res.data??[]) 
    newAuths.value = Object.assign([],res.data??[]) 
    console.log(oldAuths.value,newAuths.value)
  })
}
// 提交权限设置
const submitUserApis = () => {
  console.log(oldAuths.value,newAuths.value)
    // 禁用的
    var disableApis = oldAuths.value.filter(function(v){ return newAuths.value.indexOf(v) == -1 })
    // 新增的
    var enableApis = newAuths.value.filter(function(v){ return oldAuths.value.indexOf(v) == -1 })
    // 禁用
    if (disableApis.length <=0 && enableApis.length <=0){
      $notificatioWarn("不存在权限修改")
    }
    if(disableApis && disableApis.length){
      var params={
          admin_ids:userinfo.value.id,
          apiIds:disableApis,
          enabled:0,
        }
      permissionApi.addUserApis(params).then(res=>{
         $notificatioSuccess("操作成功")
         closeDialog()
      })
    }
    // 启用
    if(enableApis && enableApis.length){
        var params={
          admin_ids:userinfo.value.id,
          apiIds:enableApis,
          enabled:1,
        }
      permissionApi.addUserApis(params).then(res=>{
        $notificatioSuccess("操作成功")
         closeDialog()
      })
    }
}
const closeDialog = function(){
  showDialog.value = false
  emit("refresh",true)
}

defineExpose({
  show
})
</script>
<style lang="scss" scoped>
.head-flex{
  display: flex;
  justify-content: space-between;
}
.submitRight{
  margin-right :20px;
}
</style>