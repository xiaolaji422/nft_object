<template>
    <div >
        <div v-for="(item,index) in apiGroupData" :key="`group_${index}`" >
            <div v-if="item.enabled">
                <div  class="mb15 title flex">
                    <p :span="24">{{item.name}} </p> <p class="ml-20 select_all"   @click="selectGroup(item.id)" >全选</p>
                </div>
                <el-checkbox-group v-model="selectData" class="flex flex-warp" >
                    <div v-for="(aV,aK) in apiData"  :key="`api_${aK}`">
                        <div v-if="aV.enabled ==1 && aV.group_id == item.id" class="ml-20" style="border-bottom: 1px #F9FAFC solid" >
                            <el-checkbox   :label="aV.id">{{aV.name}}</el-checkbox>
                        </div>
                    </div>
                </el-checkbox-group>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import {  ref, reactive, onMounted,defineEmits,watch} from 'vue'
import {permissionStore} from '@/store/modules/permission'
const props = defineProps({
  modelValue: {
      type: Array,
      default: [],
    },
})
let apiGroupData = ref([])
let apiData = ref([])
let roleApi =ref([])
const selectData = ref([])
const initData =  async function(){
    let apiGroupDataRes = await  permissionStore().getApiGroupData() 
    apiGroupData.value = apiGroupDataRes
    let apiDataRes =await  permissionStore().getApiData() 
    apiData.value = apiDataRes
}
const emit = defineEmits(["update:modelValue"])

watch(()=>selectData,(newValues) => {
  emit("update:modelValue",newValues.value)
},
{ deep:true})

watch(()=>props.modelValue,(newValues) => {
    console.log(newValues,"props.modelValue")
  selectData.value = newValues??[]
})

const selectGroup= function(group_id){
    apiData.value.forEach(item => {
        if(item.group_id == group_id){
            if(!selectData.value.includes(item.id)){
                 selectData.value.push(item.id)
            }
        }
    });
}
onMounted (()=>{
    initData()
})
</script>
<style lang="scss" scoped>
    .title{
        position: relative;
        margin-top:10px;
    }

    .row{
        display: flex;
    }

    .ml-20{
        margin-left:20px ;
    }

    .flex-warp{
        flex-wrap: wrap;
    }

    .select_all{
        font-size: 14px;
        color:#409EFF;
        user-select: none;
        cursor: pointer;
    }
</style>