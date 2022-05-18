<template>
    <div  :class="$style.warpper">
        <div :class="$style.form">
            <div :class="$style['item-ctn']">
                <span :class="$style.title">商品搜索：</span>
                <el-select
                    style="width: calc(100% - 120px);"
                    v-model="goodsId"
                    multiple
                    filterable
                    remote
                    reserve-keyword
                    placeholder="请输入商品搜索"
                    :remote-method="remoteMethod"
                    :loading="loading"
                >
                    <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                    />
                </el-select>
            </div>
            <div :class="$style.name">
            <span :class="$style.title">搜索结果：</span><span>搜搜撒大苏打</span>
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
            <div :class="$style['content-form']"></div>
        </div>
    </div>    
</template>
<script lang="ts" setup>
import { ref } from 'vue';
const goodsId = ref()


const remoteMethod = (val:string)=>{
    console.log(val,"remoteMethod")
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
        width: 100%;
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
         height: 100%;
    }
}
</style>