<template>
    <div  :class="$style.warpper">
        <div :class="$style.form">
            <div v-for="(item,idx) in formData" :key="idx" :class="$style.item">
                <div :class="$style['item-ctn']">
                    <span >账号</span>
                    <el-input v-model="item.account" placeholder="请输入账号" />
                </div>
                 <div :class="$style['item-ctn']">
                    <span >Cookie</span>
                    <el-input v-model="item.info" placeholder="请输入Cookie" />
                </div>
                <div >
                 <el-button :class="$style['footer-btn']" type="primary" @click="saveAccount(item)"> 保 存 </el-button>
                 <!-- <el-button :class="$style['footer-btn']" type="primary" @click="delAccount(item)"> 删 除 </el-button> -->
                </div>
            </div>
            
        </div>
        <div :class="$style.footer">
                 <el-button :class="$style['footer-btn']" type="warning" @click="addAccount"> 新 增 </el-button>
        </div>
    </div>    

</template>
<script lang="ts" setup>
import {ref,onMounted} from "vue"
import accountApi from "@/api/account"
import {notify,confirm} from "@/utils/notify";
const formData  = ref([{account:"",info:""}])

const saveAccount = async (info:any)=>{
    console.log(info,"saveAccount")
    const res = accountApi.Save(info)
    notify.success("保存成功")
    console.log(res,saveAccount)
}
onMounted(()=>{
    listAccount()
})

const addAccount = ()=>{
    if(formData.value.length ==6){
        notify.warning("最多持有6个账号")
        return
    }
    formData.value.push({account:"",info:""})

}


const listAccount = async()=>{
    const res =await accountApi.List({})
    if (res && res.data){
        formData.value = res.data
    } else{
         formData.value =[{account:"",info:""}]
    }
}

</script>
<style lang="scss" module>

.warpper{
    margin: 10px;
}
.form{
    padding: 10px;
    width: 100%;
    min-height: 400px;
}
.footer {
    position: absolute;
    bottom: 0;
    width: 100%;
    display: flex;
    justify-content: flex-end;
    &-btn{
        margin-right: 20px;
    }
}
.item{
    display: flex;
    justify-content: space-around;
    align-items: center;
    height: 80px;
    &-ctn{
        width: 40%;
        // flex: 1;
        flex-wrap: nowrap;
        display: flex;
        justify-content: center;
        align-items: center;
        span{
            text-align: center;
            font-size: 14px;
            width:60px;
        }
    }
}


</style>