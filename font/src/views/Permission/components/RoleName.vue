<!-- 角色名称展示&修改 -->
<template>
    <div>
        <div v-if="isInput" class="flex">
             <el-input ref="name" placeholder="请输入内容" @blur="handleBlur" v-model="newName"></el-input>
             <el-button type="success" icon="el-icon-edit" circle @click.stop="handleSubmit"></el-button>
        </div>
        <div style="font-size:16px;" v-else>
            <span class="role-name">{{role.name}}</span>
            <i class="el-icon-edit editBtn" @click.stop="handleChange"></i>
        </div>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                role:{
                    name:'',
                    id:'',
                },
                isInput: false,
                newName:'',
                timer:null
            }
        },
        props:{
            'item':{
                default:{
                    name:'',
                    id:''
                }
            },
            'index':{
                type: Number,
                default:0
            }
        },
        mounted() {
            this.role = this.item;
        },
        methods: {
            handleBlur(e){
                this.timer = setTimeout(()=>{
                    this.isInput = false;
                    clearTimeout(this.timer);
                    this.timer = null;
                },200);
            },
            handleChange(){
                this.isInput = !this.isInput;
                this.$nextTick( () => {
                    this.$refs.name.$refs.input.focus();
                    this.newName = this.item.name;
                })
            },
            handleSubmit(e){
                let data = {
                    id:this.role.id,
                    name:this.newName
                };
                if(data.name === this.role.name) {
                  return
                }
                if(!data.name) {
                  this.$message.error('角色名不能为空')
                  return;
                }
                this.$emit('changeName',data,this.index);
                this.isInput = !this.isInput;
            }
        }
    }
</script>
<style scoped>
    .ul-role-list li .editBtn{float: right;line-height: 40px;color: #fff;}
    .ul-role-list li:hover .editBtn{color: #8492A6;}
    .ul-role-list li.active .editBtn{color: #20a0ff;}
    .ul-role-list li.active:hover .editBtn {color: #fff;}
    .roleList li.active .ic-check,.roleList .ic-check:hover, .roleList .ic-close:hover{color: #20a0ff;}
    .roleList .active .ic-check:hover, .active .ic-close:hover{color: #20a0ff;}
</style>
