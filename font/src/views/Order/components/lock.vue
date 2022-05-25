<template>
    <div  :class="$style.warpper">
        <div :class="$style.form">
            <div :class="$style['item-ctn']">
                <span :class="$style.title">商品搜索：</span>
                <el-select
                    style="width: calc(100% - 120px);"
                    v-model="formData.goodsName"
                    filterable
                    remote
                    reserve-keyword
                    placeholder="请输入商品搜索"
                    :remote-method="remoteMethod"
                    :loading="loading"
                    value-key="id"
                    @change ="changeGoods"
                >
                    <el-option
                    v-for="item in goodsData"
                    :key="item.id"
                    :label="item.name"
                    :value="item"
                    />
                </el-select>
            </div>
            <div :class="$style.name">
            <span :class="$style.title">搜索结果：</span><span>{{formData.goodsName}}</span>
            </div>
        </div>
        <div :class="$style.content">
            <div :class="$style['content-table']">
                <el-table :data="tableData" style="width: 100%;" height="95%"  row-class-name="table-class">
                    <el-table-column prop="name" label="编号"  />
                    <el-table-column prop="name" label="价格"  />
                    <el-table-column prop="name" label="状态(1次/s)" />
                </el-table>
            </div>
            <div :class="$style['content-form']">
                <el-form>
                    <el-form-item label="账号">
                        <el-select v-model="formData.adminInfo" class="m-2" placeholder="Select">
                            <el-option
                            v-for="item in adminData"
                            :key="item.account"
                            :label="item.account"
                            :value="item.account"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="价格">
                        <el-row :gutter="10">
                            <el-col :span="11"><el-input v-model="formData.min"/></el-col>
                            -
                            <el-col :span="11"><el-input v-model="formData.max"/></el-col>
                        </el-row>
                    </el-form-item>
                     <el-row :gutter="40">
                            <el-col :span="8" v-if="!islock"><el-button size="small" @click="lockGoods" type="primary">自动锁单</el-button></el-col>
                            <el-col :span="8" v-else><el-button size="small" @click="unlockGoods" type="primary">取消锁单</el-button></el-col>
                        </el-row>
                    
                </el-form>
                 
            </div>
        </div>
    </div>    
</template>
<script lang="ts" setup>
import { ref,onMounted } from 'vue';
import nftApi from "@/api/nft/index"
import {notify,confirm} from "@/utils/notify";
const goodsId = ref()

const formData = ref({
    accountInfo:{},
    min:"",
    max:"",
    goodsId:"",
    goodsName:"",
})
onMounted(()=>{
    remoteMethod("")
})
const islock=ref(false)
const loading = ref(false)
const remoteMethod =async (val:string)=>{
  const res = await  nftApi.serachAblum({keyword:val})
  console.log(res,"remo")
  if (res && res.data && res.data.items){
      goodsData.value = res.data.items
  }else{
      goodsData.value = []
  }
}

const tableData = [
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  {name: 'Tom'},
  
]

const changeGoods = (val:any)=>{
    console.log(val,"changeGoods",val)
    formData.value.goodsId = val.id
    formData.value.goodsName = val.name
}
const lockGoods = async()=>{
    notify.info("开发中")
    islock.value = true
}
const unlockGoods = async()=>{
    notify.info("开发中")
    islock.value = false
}
const goodsData = ref([])

const adminData=[
    {account:"12313",cookie:"12313"},
    {account:"12313",cookie:"12313"},
    {account:"12313",cookie:"12313"},
    {account:"12313",cookie:"12313"},
    {account:"12313",cookie:"12313"},
    {account:"12313",cookie:"12313"},
    {account:"12313",cookie:"12313"},
]
</script>
<style lang="scss" scoped>
:deep .el-table .el-table__cell{
    padding: 6px;
}

:deep .el-table thead{
    font-size:14px;
}

:deep .el-table tr{
    background-color:#FAFAFA ;
}
</style>
<style lang="scss" module>

.warpper{
    width:100%;
    height: 400px;
}
.name{
    margin:10px 0;
}
.form{
    padding: 10px;
    width: 100%;
    height: 110px;
    
}
.item{
    display: flex;
    justify-content: flex-start;
    align-items: center;
    height: 80px;
    &-ctn{
        width: 70%;
        flex-wrap: nowrap;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        
    }
}
.title{
    text-align: center;
    font-size: 14px;
}

.content{
    display: flex;
    height: calc(100% - 110px);
    border-radius: 6px;
    &-table{
        border-radius: 6px;
        width: 60%;
        background-color: #FAFAFA;
        height: 100%;
    }
    &-form{
        border-radius: 6px;
        width: 40%;
        background-color: #FAFAFA;
        // height: 100%;
        margin-top:20px;
        padding: 10px;

    }
}
</style>