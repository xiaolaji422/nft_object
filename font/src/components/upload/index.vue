<template>
    <div>
        <div >
            <div  class="flex flex-wrap ">
                <!-- 图片展示 -->
                <show v-show="imgList && imgList.length"  
                :height="`${height}px`"  
                :width="`${width}px`" 
                :isDelete="!disabled"   
                @handleDeleteImg="handleDeleteImg"  
                :url="imgList" >
                </show>

                <!-- 图片上传 -->
                <el-upload
                name="data[Upload][file]"
                v-if="(maxCount>imgList.length) && !disabled"
                class="avatar-uploader img-model"
                :action="uploadUrl"
                :show-file-list="false"
                :on-success="uploadSuccess"
                :before-upload="beforeUpload">
                <div class="max_size_span" :style="`height:${height}px;width:${width}px;`">
                    <div>
                        <i class="el-icon-plus avatar-uploader-icon "></i>
                        <div style="height:10px;line-height:10px;">{{`${warntitle} ${maxSize}M`}}</div>
                    </div>
                </div>
                </el-upload>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import show from "./show.vue"
import { $notificatioSuccess,$notificatioWarn,$notificatioError } from "@/utils/utils";
import {defineComponent,defineProps,defineEmits,ref, onMounted} from 'vue'
import {ElMessage,ElMessageBox} from 'element-plus'
const prop = defineProps({
        maxSize:{
            type:Number,
            default:6,
        },
        // 最大上传参数
        maxCount:{
            type:Number,
            default:1,
        },
        modelValue:{
            type:[Array,String],
            required:true,
        },
        width:{
            type:[Number,String],
            default:120,
        },
        height:{
            type:[Number,String],
            default:120,
        },
        disabled:{
            type:Boolean,
            default:false,
        },
})
const warntitle = ref("大小<")
const emit = defineEmits(["update:modelValue"])
const components = defineComponent({
    show
})
const uploadUrl =`http://csis.cm.com/script/imgupload?domain=van.oa.com&referer=http%3A%2F%2Fvan.oa.com%2Ffaq%2Fupload&internet=true`
const imgList = ref([])
const is_array = ref(false)
const saveImgList = ref([])
const inType = ref("string")


const setValue= () =>{
    if (!prop.modelValue) {
        imgList.value = []
        return
    }
    if (typeof prop.modelValue === 'string') {
        if (prop.modelValue) {
            imgList.value = [prop.modelValue]
        } else {
            imgList.value = []
        }
    } else if (typeof prop.modelValue === 'object') {
        inType.value="object"
        imgList.value = Object.assign([],prop.modelValue)
    }

    if(saveImgList.value.length<=0){
        saveImgList.value = Object.assign([],imgList.value) 
    }
}

onMounted(()=>{
    setValue()
})
        // 图片上传校验
const  beforeUpload = (file)=> {
    if(prop.disabled){
        return false;
    }
    const isJPG = file.type === 'image/jpeg';
    const isPNG = file.type === 'image/png';
    if (!isJPG && !isPNG) {
        $notificatioWarn("图片只能是 JPG或PNG 格式!")
        return false;
    }
    if (prop.maxSize>0 && (file.size / 1024 / 1024) > prop.maxSize) {
        // (file.size / 1024 / 1024) < this.maxSize;
        $notificatioWarn(`图片大小不能超过${prop.maxSize}MB!`)
        return false;
    }
    return true;
}
// 图片上传成功
const uploadSuccess = (response, file) =>{
    var res = response
    const picurl = response.match(/(http\:\/\/file\.service\.qq\.com.+?)\'/);
    if (picurl && picurl[1]) {
        picurl[1] = picurl[1].replace('http://', 'https://');
        imgList.value.push(picurl[1]) 
        saveImgList.value.push(picurl[1]) 
        handleChangeValue()
    }else {
        $notificatioError( '上传图片失败')
    }
}
// 删除图片
const   handleDeleteImg = (url)=>{
    let urlIndex = -1;
    for(var i =0;i<imgList.value.length;i++){
            if(imgList.value[i] == url) {
            urlIndex =i;
            break;
        }
    }
    if(urlIndex!==-1){
        ElMessageBox.confirm(
        '是否确定删除?',
        'Warning',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          title:"删除确认"
        }).then(() => {
            imgList.value.splice(urlIndex,1)
            saveImgList.value.splice(urlIndex,1)
            handleChangeValue()
            $notificatioSuccess('操作成功!')
        })
    }
}

// 内容改变
const  handleChangeValue = () => {
    let value = saveImgList.value
    if (inType.value == 'string'){
        value = saveImgList.value[0]
    }else if(inType.value == "object"){
         value = saveImgList.value
    }
    emit('update:modelValue', value)
}
</script>
<style lang="scss" >

.img-model {
    position: relative;
    margin-right: 10px;
    margin-top: 10px;
}
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 18px;
    color: #8c939d;
    text-align: center;
  }
  .avatar {
    width:100%;
    height: 100%;
     border-radius: 6px;
  }

  .show_image{
      position: relative;
  }
 
  .max_size_span{
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    font-size: 10px;
    color: #CCCCCC;
  }
</style>