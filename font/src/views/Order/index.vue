<template>
    <div  :class="$style.warpper">
        <div :class="$style['lock-body']">
            <div :class="$style['lock-item']" v-for="item in lockData" :key="item">
                <lock :account="accountData" :lock="item" @refresh="listLock()"></lock>
            </div>
        </div>
        
    </div>    

</template>
<script lang="ts" setup>
import {ref,defineComponent,onMounted} from "vue"
import Lock from "./components/lock.vue"
import accountApi from "@/api/account"
import {notify,confirm} from "@/utils/notify";
import lockApi from "@/api/account_lock/index"
defineComponent({
    Lock,
}) 

onMounted(()=>{
    listAccount()
    listLock()
})

const accountData=ref([])
const lockData=ref([{},{},{},{}])

const listAccount = async()=>{
    const res =await accountApi.List({})
    if (res && res.data){
        accountData.value = res.data
    } else{
        accountData.value =[]
    }
}

const listLock = async()=>{
    const res =await lockApi.List({})
    if (res && res.data && res.data.length){
        for (let index = 0; index < res.data.length; index++) {
            const element = res.data[index];
            if( index >3 ){
                return
            }
            lockData.value[index] = element
        }
    } else{
        lockData.value = [{},{},{},{}]
    }
}


</script>
<style lang="scss" module>

.warpper{
    margin: 10px;
    
}


.lock-body{
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
    padding:10px
}   

.lock-item{
    width: 45%;
    min-height: 200px;
    background-color: #FAFAFA;
    border-radius: 10px;
    border: 1px solid rgba($color: #f0f0f0, $alpha: 0.2);
    margin-bottom:10px
}


</style>