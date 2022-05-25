<template>
  <div class="app-container">
        <audio ref="waringRef"  >
            <source src="@/assets/mp3/waring.mp3" type="audio/mpeg">
        </audio>
        <div> 最新公告（单个平台最新1条公告，若单个平台5分钟连续多条则多条都告警展示）</div>
        <el-table cell-class-name="newstable" :show-header="false" :data="tableData" style="padding:10px;" @cell-click="handleCellClick">
           <el-table-column prop="notice_time" label="公告时间" width="200"/>
           <el-table-column prop="notice_time" label="产品名称" width="100">
                <template #default="{row}">
                    {{ getPlatName(row.platform) }}
                </template>
            </el-table-column>
            <el-table-column prop="name" label="内容" class-name="notice-name"/>
        </el-table>
        <div style="margin-top:40px">
        <div style="padding: 12px 0;">历史公告</div>
            <app-table
                @cell-click="handleCellClick"
                :config="tableConfig"
                ref="appTableRef"
                tableName="userTable"
            >
            <template #platform >
                <el-table-column prop="notice_time" label="产品名称" width="200">
                    <template #default="{row}">
                        {{ getPlatName(row.platform) }}
                    </template>
                </el-table-column>
            </template>
            <template #name >
                <el-table-column prop="name" label="内容" class-name="notice-name"/>
            </template>
        </app-table>
        </div>
  </div>
</template>
<script lang="ts" setup>
import noticeApi from '@/api/notice'
import { $notificatioSuccess,$notificatioWarn } from "@/utils/utils";
import { defineComponent, ref, reactive, onMounted } from 'vue'
import {getMapData} from '@/utils/statusMap'
import { flatMap } from 'lodash';
import {WarnningPlay} from "@/utils/audioPlay"
import {playerStore} from "@/store/modules/palyser"
import {websocketStore} from "@/store/modules/websocket"
import socketUtil from "@/utils/websocket"
  //  faq配置列表
const name =  'NoticeTable'
const tableData = ref<any>([])

const paltFormData = ref([])
const waringRef = ref<any>()

const getPlatName = (platform:number)=>{
    var info = paltFormData.value.find((item)=>item.value == platform)
     if (info && info.text){
       return info.text
   }
    return ""
}
const {Play} = playerStore()
const  socket= websocketStore()
const getTableData = async()=>{
    try{
        socket.sendMsg({"appid":"422","userid":"lajipeng"})
        const {data} = await noticeApi.queryNotice({})
        if (data && data.data ){
            tableData.value = data.data
            // 获取成功
            if (data.is_warn){
                Play("有新公告")
            }
            
           
        }
    }catch(err){}
}

const refreshTableData = ()=>{
    getTableData()
}

const handleCellClick = (event:any)=>{
    window.open(event?.notice_url)
}

const appTableRef = ref()

onMounted(async ()=>{
    refreshTableData()
    paltFormData.value = getMapData("notice_type")
    window.autoPlayMp3= setInterval(getTableData,10000)
})

const tableConfig = reactive({
      tableColumn: [{
          label: '公告时间',
          width: 200,
          prop:"notice_time",
          align: 'center'
        },
        {
          slot:"platform"
        },
        {
           slot:"name"
        }],
      tableSearchBtnName: { reset: false},
      tableListApi: (params:any) => noticeApi.queryHistoryNotice(params),
      tableListParams: {},
      tableDeleteApi: (id: number) => ()=>{},
      tableDeleteParams: {},
      tableHeaderRight: true
    })

</script>
<style lang="scss" scoped>
.app-container{
    margin: 20px;
}
:deep .el-table__row:hover{.notice-name{color:#1890FF;}}
:deep  .notice-name {
    cursor: pointer;
    user-select: none;
 }
 :deep .newstable .el-table .el-table__cell{
     border-bottom: 0;
     padding:6px 0 ;
 }
</style>