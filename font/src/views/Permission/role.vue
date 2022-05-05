<template>
  <div class="app-container">
      <div class="roleList flex">
        <el-card class="box-card" style="width:500px">
          <template #header>
            <div class="card-header">
              <span class="addTit">角色</span>
            <i class="el-icon-plus addRole" @click="dialogFormVisible = true"></i>
            </div>
          </template>
          <ul class="ul-role-list">
            <li v-for="(item, index) in roleList" :class="[active.id == item.id && 'active']" :key="index"
                @click="handleRight(item.id, item.name, true)">
                <RoleName  :item="item" :index="index" @changeName="handleNameChange"></RoleName>
            </li>
          </ul>
        </el-card>
        <el-card class="box-card btnbox" style="margin-left: 40px;">
           <template #header>
            <div class="card-header">
              <span class="title addTit">权限设置 <span v-if="active.name" style="color:red">({{active.name}})</span> </span>              
              <el-button type="primary" class="submitRight" @click.native="handleSubmitRight()">提交</el-button>
            </div>
          </template>
          <SelectApi  v-model="roleApi"></SelectApi>
        </el-card>
      </div>
      <el-dialog destroy-on-close title="编辑角色" v-model="dialogFormVisible" @close="dialogClosed" size="tiny">
        <el-form :model="form" :rules="rules" ref="form">
          <el-form-item label="角色名称" :label-width="formLabelWidth" prop="name">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click.native="dialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click.native="handleSubmitAddRole">确 定</el-button>
        </div>
      </el-dialog>
    
  </div>
</template>

<script >
  import permissionApi from '@/api/permission'
  import RoleName from './components/RoleName.vue'
  import SelectApi from './components/select_apis.vue'
  import {permissionStore} from '@/store/modules/permission'

  export default {
    name: 'Role',
    components:{
      RoleName,
      SelectApi
    },
    data() {
      return {
        form: {
          name: ''
        },
        rules: {
          name: [{required: true, message: '请输入角色名称', trigger: 'blur'}]
        },
        // api 分组
        apiGroupData:[],
        // 所有api接口
        apiAllData:[],
        roleApi:[],
        checked: false,
        roleList: [], // 角色信息列表
        roleListObj: {},
        checkval: {},  // 角色的权限列表
        allRightList: [], // 全部的权限列表
        dialogFormVisible: false,
        formLabelWidth: '100px',
        active: {id: '', name: ''}, // 当前选择的角色
        cur_role_index_temp: null,
        cur_role_val_temp: null,
        token: {},
      }
    },
    mounted() {
      this.getAllRoleRight();
    },
    created(){
    },
    methods: {
      handleNameChange(item, index) {
        this.handleSubmitEditRole(item.id, item.name, index);
      },
      async getAllRoleRight() {
        const roleData = await permissionStore().getRoleData()
      
        this.roleList = roleData
        if (this.roleList && this.roleList.length) {
          this.active = {
            id: this.roleList[0].id,
            name: this.roleList[0].name
          };
        }else{
          return 
        }
        for (let item of this.roleList) {
          let role_item = {
            "id": item.id,
            "name": item.name,
            "isEdit": false
          }
          this.roleListObj[item.id] = role_item;
        }
        let id = this.roleList[0].id;
          this.getRoleRight(id);
      },
      getRoleRight(id) {
        permissionApi.getRoleRights({id: id}).then(res => {
          if(res.data && res.data.length){
             this.roleApi = res.data
          }else{
            this.roleApi =[]
          }
        })
      },
      handleResetState() {
        for (let index in this.roleListObj) {
          let item = this.roleListObj[index]
          if (item.isEdit)
            item.isEdit = false;
        }
      },
      handleSubmitAddRole() {
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.dialogFormVisible = false;

            let params = {
              name: this.form.name
            }
            permissionApi.addRoleStore(params).then(res => {
              this.$message.success('成功添加角色名称！');
              permissionStore().clearRole()
              this.getAllRoleRight();
            })
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      handleSubmitEditRole(id, name, index) {
        if (!id) return
        let params = {
          id: id,
          name: name
        }

        permissionApi.updateRole(params).then(res => {
          this.$message.success('成功修改角色名称！');
           permissionStore().clearRole()
          // this.$store.dispatch("systemApi/clearRole")
          this.roleList[index].name = name;
          this.getAllRoleRight();
          this.$forceUpdate();
        })
      },
      handleRight(id, name, isReset) {
        if (this.cur_role_index_temp !== null || this.cur_role_val_temp !== null)
          this.roleList[this.cur_role_index_temp].name = this.cur_role_val_temp;
        this.active = {
          id: id,
          name: name
        }
        for (let i in this.checkval) {
          this.checkval[i] = false
        }
        if (isReset) {
          this.handleResetState()
        }
        this.getRoleRight(id);
      },
      // 更新权限
      handleSubmitRight() {
        let params = {};
        params.roleId = this.active.id;
        params.name = this.active.name;
        params.apis = this.roleApi;
        params.isRights = 1;
        permissionApi.updateRoleApi(params).then(res => {
          this.$message.success('成功编辑权限！');
        })
      },
      dialogClosed() {
        this.$refs.form.resetFields()
      },
   
    },
  }
</script>

<style lang="scss" scoped>
.el-card__header{padding: 18px 20px;
    border-bottom: 1px solid #ebeef5;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;}
  .roleList .el-card {margin-bottom: 20px;}
  .roleList .el-card .addTit{font-weight: bold;}
  .roleList .el-card .addRole{float: right;line-height: 20px;cursor: pointer;}
  .roleList .el-card .addRole:hover{color: #20A0FF;}
  .roleList .ic-check i{float: right; line-height: 40px;}
  .roleList .submitRight {float: right; position: relative;top: -8px;}
  .roleList .ic-check:hover, .roleList .ic-close:hover{color: #20a0ff;}
  .roleList .active .ic-check:hover, .active .ic-close:hover{color: #fff;}
  .roleList .el-input--small .el-input__inner{margin-top: 5px;margin-left: -12px;}
  .ul-role-list {padding: 6px 0;}
  .ul-role-list li{line-height: 40px;padding: 0 20px;}
  .ul-role-list li:hover{background-color: #eff2f7;cursor: pointer;}
  .ul-role-list li.active{background-color: #20a0ff;color: #fff;}
  .ul-role-list li .editBtn{float: right;line-height: 40px;color: #fff;}
  .ul-role-list li:hover .editBtn{color: #8492A6;}
  .ul-role-list li.active .editBtn{color: #20a0ff;}
  .ul-role-list li.active:hover .editBtn{color: #fff;}
  .roleList .right-title{
    text-align: left;
    font-weight: bold;
  }
</style>
<style>
  .rolename-list .el-card__body {
    padding: 0;
    text-align: left;
  }
</style>
