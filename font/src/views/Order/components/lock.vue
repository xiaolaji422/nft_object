<template>
    <div  :class="$style.warpper">
        <div :class="$style.form">
            <div :class="$style['item-ctn']">
                <span :class="$style.title">商品搜索：</span>
                <el-select
                    style="width: calc(100% - 120px);"
                    v-model="formData.album_name"
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
            <span :class="$style.title">搜索结果：</span><span>{{formData.album_name}}</span>
            </div>
        </div>
        <div :class="$style.content">
            <div :class="$style['content-table']">
                <el-table :data="tableData" style="width: 100%;" height="95%"  row-class-name="table-class">
                    <el-table-column prop="gNum" label="编号"  />
                    <el-table-column prop="priceCny" label="价格"  />
                    <el-table-column prop="gStatus" label="状态(1次/s)" />
                </el-table>
            </div>
            <div :class="$style['content-form']">
                <el-form>
                    <el-form-item label="账号">
                        <el-select v-model="formData.account_id" class="m-2" placeholder="选择账号">
                            <el-option
                            v-for="item in account"
                            :key="item.id"
                            :label="item.account"
                            :value="item.id"
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
                            <el-col :span="8" v-if="!(formData.id || islock)"><el-button size="small" @click="lockGoods" type="primary">自动锁单</el-button></el-col>
                            <el-col :span="8" v-else><el-button size="small" @click="unlockGoods" type="primary">取消锁单</el-button></el-col>
                        </el-row>
                    
                </el-form>
                 
            </div>
        </div>
    </div>    
</template>
<script lang="ts" setup>
import { ref,onMounted,defineProps} from 'vue';
import nftApi from "@/api/nft/index"
import lockApi from "@/api/account_lock/index"
import {notify,confirm} from "@/utils/notify";
const prop = defineProps({
        account:{
            type:Array,
            default:[],
        },
        lock:{
            type:Object,
            default:{},
        },
})
const goodsId = ref()

const formData = ref({
    id:"",
    account_id:"",
    min:"",
    max:"",
    album_id:"",
    album_name:"",
})
onMounted(()=>{
    remoteMethod("")
    if(prop.lock){
        let info = prop.lock
        formData.value.id = info.id??""
        formData.value.account_id = info.account_id??""
        formData.value.min = info.min??""
        formData.value.max = info.max??""
        formData.value.album_id = info.album_id??""
        formData.value.album_name = info.album_name??""
    }
    if (formData.value.album_id){
        productDetail(formData.value.album_id)
    }
})
const islock=ref(false)
const loading = ref(false)
const remoteMethod =async (val:string)=>{
  const res = await  nftApi.serachAlbum({keyword:val})
  if (res && res.data && res.data.items){
      goodsData.value = res.data.items
  }else{
      goodsData.value = []
  }
}

const tableData = ref([])

const changeGoods = (val:any)=>{
    formData.value.album_id = val.id
    formData.value.album_name = val.name
    productDetail(val.original_id)
    
}

const productDetail =async (id:any)=>{
     const res = await nftApi.albumDetail({album_id:id})
     if (res && res.data){
         tableData.value = res.data
     }else{
         tableData.value = []
     }
}
const lockGoods = async()=>{
    
    const res = await lockApi.Save(formData.value)
    console.log(res,"lockGoods",formData.value)
    if(res){
        islock.value = true
        notify.success("锁单成功")
    }
}

const handleChcke = ()=>{
    
}
const unlockGoods = async()=>{
    notify.info("开发中")
    islock.value = false
}
const goodsData = ref([])

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