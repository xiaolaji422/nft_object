<template>
    <div class='w-screen h-screen bg-gray-800'>
        <div class='layout-login' @keyup='enterSubmit'>
            <h2 class='text-2xl font-semibold text-gray-100 text-center mb-6'>NFT 辅助系统</h2>
            <h3 class='text-2xl font-semibold text-gray-100 text-center mb-6'>注册</h3>
            <el-form ref='ruleForm' label-position='right' label-width='80px' :model='form' :rules='rules'>
                <el-form-item class='mb-6 -ml-20' prop='name'>
                    <el-input v-model='form.name' placeholder='请输入用户名(英文和数字)' prefix-icon='el-icon-user' />
                </el-form-item>
                <el-form-item class='mb-6 -ml-20' prop='pwd'>
                    <el-input v-model='form.pwd' placeholder='请输入密码' prefix-icon='el-icon-lock' show-password />
                </el-form-item>
                 <el-form-item class='mb-6 -ml-20' prop='pwd2'>
                    <el-input v-model='form.pwd2' placeholder='请再次输入密码' prefix-icon='el-icon-lock' show-password />
                </el-form-item>
                <el-form-item class='mb-6 -ml-20'>
                    <el-button type='primary' class='w-full' @click='onSubmit'>注 册</el-button>
                </el-form-item>
                <div class='flex justify-between'>
                    <div></div>
                    <div style="color:white">返回<el-button class="signBtn" type='text' @click="login">登录</el-button></div>
                </div>
            </el-form>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { useLayoutStore } from '@/store/modules/layout'
import { ElNotification } from 'element-plus'
import { validate } from '@/utils/formExtend'
import { useRouter } from 'vue-router'

const formRender = () => {
    const { register } = useLayoutStore()
    let form = reactive({
        name: '',
        pwd: '',
        pwd2:"",
    })
    const ruleForm = ref(null)
    const enterSubmit = (e:KeyboardEvent) => {
        if(e.key === 'Enter') {
            onSubmit()
        }
    }
    const onSubmit = async() => {
        let { name, pwd } = form
        if(!await validate(ruleForm)) return
        register({ username: name, password: pwd })
       
    }
    let router =  useRouter()
    const login = ()=>{
        router.push({
          path: `/Login`,
        })
    }
    const rules = reactive({
        name: [
            {trigger: 'blur' },
            { validator: (rule: any, value: any, callback: (arg0?: Error|undefined) => void) => {
                if (!value) {
                    return callback(new Error('用户名不能为空'))
                }
               var  reg=/^[A-z0-9]*$/;
               if (!reg.test(value)){
                   return callback(new Error('请输入英文和字符串'))
               } 
                callback()
            }, trigger: 'blur' 
            }
        ],
        pwd: [
            { validator: (rule: any, value: any, callback: (arg0?: Error|undefined) => void) => {
                if (!value) {
                    return callback(new Error('密码不能为空'))
                }
                callback()
            }, trigger: 'blur' 
            }
        ],
        pwd2: [
            { validator: (rule: any, value: any, callback: (arg0?: Error|undefined) => void) => {
                if (!value) {
                    return callback(new Error('密码不能为空'))
                }else if(value != form.pwd){
                     return callback(new Error('两次密码输入不一致'))
                }
                callback()
            }, trigger: 'blur' 
            }
        ]
    })
    return {
        form, 
        onSubmit,
        enterSubmit,
        rules,
        ruleForm,
        login
    }
}
export default defineComponent({
    name: 'Login',
    setup() {
        return {
            labelCol: { span: 4 },
            wrapperCol: { span: 14 },
            ...formRender()
        }
    }
})
</script>

<style lang='postcss' scoped>
.layout-login {
    padding-top: 10%;
    width: 400px;
    margin: 0 auto;

    ::v-deep(.el-input__inner) {
        border: 1px solid hsla(0, 0%, 100%, 0.1);
        background: rgba(0, 0, 0, 0.1);
        border-radius: 5px;
        color: #ddd;
    }
}
.signBtn{
    font-size: 16px;
}

</style>