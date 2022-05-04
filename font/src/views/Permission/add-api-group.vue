<template>
  <el-dialog
    destroy-on-close
    append-to-body
    title="添加分组"
    v-model="showDialog"
    @close="closeDialog"
  >
    <el-form
      ref="formref"
      :model="form"
      :rules="rules"
      status-icon
      label-position="left"
      label-width="100px"
      style="width: 400px; margin-left: 80px"
    >
      <el-form-item label="分组名称" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入分组名称"
          class="input-x2"
        />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="closeDialog"> 取消 </el-button>
      <el-button type="primary" @click="subInfo"> 确认 </el-button>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import permissionApi from '@/api/permission'
import { defineComponent ,ref, reactive, onMounted} from 'vue'
import {permissionStore} from '@/store/modules/permission'
import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
const components = defineComponent([
])

const emit = defineEmits(["refresh"])
const showDialog = ref(false)
const groupData = ref([])
const form  =reactive({
  id:"",
  name:"",
}) 
const statusData = ref([])
// 校验规则
const rules = {
    name: [{required: true, message: '请输入分组名称', trigger: 'blur'}],
}
onMounted(async () =>{
})

const show = function(info){
  if(!info) info = {}
  form.name = info.name ?? "",
  form.id = info.id ?? "",
  showDialog.value = true
}
defineExpose({
  show
})
const  formref = ref(null)
const  subInfo = ()=> {
  formref.value.validate((valid) => {
      if(valid){
        if(form.id){
          //  修改
          permissionApi.editApiGroup(form).then(res=>{
              $notificatioSuccess("操作成功")
              permissionStore().clearApi()
              closeDialog()
          }).catch(err=>{
            
          })
        }else{
          permissionApi.addApiGroup(form).then(res=>{
              $notificatioSuccess("操作成功")
              permissionStore().clearApi()
              closeDialog()
          }).catch(err=>{

          })
        }
          
      }else{
          return false
      }
  })
}

const closeDialog = function(){
  showDialog.value = false
  emit("refresh",true)
}
</script>



