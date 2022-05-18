<template>
    <div  :class="$style.warpper">
        <div :class="$style.form">
            <div v-for="(item,idx) in formData" :key="idx" :class="$style.item">
                <div :class="$style['item-ctn']">
                    <span >账号</span>
                    <el-input v-model="item.admin" placeholder="请输入账号" />
                </div>
                 <div :class="$style['item-ctn']">
                    <span >Cookie</span>
                    <el-input v-model="item.cookie" placeholder="请输入Cookie" />
                </div>
            </div>
            
        </div>
        <div :class="$style.footer">
                 <el-button :class="$style['footer-btn']" type="primary" @click="saveCookie"> 保 存 </el-button>
        </div>
    </div>    

</template>
<script lang="ts" setup>
import {ref} from "vue"


const formData  = ref([
    {admin:"",cookie:""}
])

const saveCookie = ()=>{
    console.log(formData.value,"saveCookie")
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