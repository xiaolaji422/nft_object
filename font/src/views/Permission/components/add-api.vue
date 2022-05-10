<template>
  <el-dialog
    destroy-on-close
    append-to-body
    :title="`${form.id?'修改':'新增'}接口`"
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
      <el-form-item label="接口分组" prop="group_id">
        <el-select
          v-model="form.group_id"
          clearable
          placeholder="请选择"
          style="width:300px;"
        >
          <el-option
           v-for="item in groupData"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="接口名称" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入接口名称"
          class="input-x3"
        />
      </el-form-item>
      <el-form-item label="接口路由" prop="route">
        <el-input
          v-model="form.route"
          placeholder="请输入接口路由,例如：/auth/login"
          class="input-x3"
        />
      </el-form-item>
      <el-form-item label="接口限频 /S" prop="limit">
        <el-input
          v-model="form.limit"
          placeholder="可输入限频次数/秒,例如：10"
          class="input-x3"
        /> <p>未填写或填写为 "0" 表示接口不限频</p>
      </el-form-item>
      <el-form-item label="请求方法" prop="methods">
         <el-radio v-model="form.methods" label="GET" >GET</el-radio>
        <el-radio v-model="form.methods" label="POST" >POST</el-radio>
      </el-form-item>
      <el-form-item label="接口状态" prop="status">
       <el-radio v-for="item in statusData"  v-model="form.enabled" :label="item.value" :key="item.value">{{item.text}}</el-radio>
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
import { getMapData } from "@/utils/statusMap"
const components = defineComponent([
])

const emit = defineEmits(["refresh"])
const showDialog = ref(false)
const groupData = ref([])
const form  =reactive({
  id:"",
  group_id:"",
  name:"",
  methods:"",
  route:"",
  limit:"",
  enabled:"",
})
const statusData = ref([])
// 校验规则
const rules = {
    name: [{required: true, message: '请输入api名称', trigger: 'blur'}],
    route: [{required: true, message: '请输入api路由', trigger: 'blur'}],
    group_id: [{required: true, message: '请选择接口分组', trigger: 'blur'}],
    limit: [{required: true, message: '请输入限频设置', trigger: 'blur'}],
    methods: [{required: true, message: '请确认请求方法', trigger: 'blur'}],
}
onMounted(async () =>{
  initData()
})
const initData = async function() {
    let roleDataRes = await  permissionStore().getApiGroupData();
    groupData.value = roleDataRes;
    let statusRes = getMapData("status")
    statusData.value = statusRes
}

const show = function(info){
  if(!info) info = {}
  form.id = info.id ?? "",
  form.group_id = info.group_id ?? "",
  form.name = info.name ?? "",
  form.methods = info.methods??"",
  form.route = info.route??"",
  form.limit = info.limit??"",
  form.enabled = info.enabled??"",
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
          permissionApi.editApi(form).then(res=>{
              $notificatioSuccess("操作成功")
              permissionStore().clearApi()
              closeDialog()
          }).catch(err=>{
            
          })
        }else{
          permissionApi.addApi(form).then(res=>{
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
