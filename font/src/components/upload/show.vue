<template>
  <div class="flex flex-wrap">
  <div
    v-for="(item,index) in images"
    :key ="index"
    class="img-model"
    @mouseover="hover = index"
    @mouseleave="hover = -1"
    :style="`height:${height};width:${width}`"
    >
    <el-image class="img item_center" :src="item" :preview-src-list="images" :style="`height:${height};width:${width};`">
      <div slot="error" class="image-slot border_alt ">
        <div style="font-size:12px;">{{altStr}}</div>
      </div>
    </el-image>
    
      <div v-if="isDelete " :class="`op-content ${hover==index ? 'op-content_hover' : ''}`" @click.stop="">
        <i class="el-icon-delete padding10 cursor_pointer"  @click="handleDeleteImg(item)" >
        </i>
      </div>
  </div>
  </div>
</template>
<script lang="ts" setup>
import { defineComponent,defineProps,ref,defineEmits, onMounted, watch} from 'vue'
const name= "showImage"
const prop = defineProps({
    // 展示的图片的路径
    url: {
      type:[String,Array],
      default:"",
      require:true,
    },
    // 是否展示覆盖的标签
    isShow: {
      type: Boolean,
      default: true,
    },
    // 是否可以删除
    isDelete:{
      type: Boolean,
      default: true,
    },
    height: {
      type: String,
      default: "200px",
    },
    width: {
      type: String,
      default: "200px",
    },
    altStr: {
      type: String,
      default: "图片还未同步外网，请稍候",
    },
})

const emit =defineEmits(["edit","handleDeleteImg"])

const images = ref([])

watch(()=>prop.url,(newValues) => {
  setImages()
},
 { deep: true })

const setImages = ()=>{
    if(typeof(prop.url) == 'string'){
    images.value = [prop.url]
    }else if(typeof(prop.url) == 'object'){
        images.value = Object.assign([],prop.url)
    }
    console.log(images.value)
}

const editImage = ()=> {
    emit("edit");
}

const hover = ref(-1)
const dialogVisible = ref(false)
const dialogImageUrl = ref(-1)
const handleDeleteImg = (item)=>{
    // 图片删除
    emit("handleDeleteImg",item);
}


const showImag =(item,index)=>{
    dialogVisible.value = true;
    dialogImageUrl.value = index
}
const showBack = (step)=>{
    dialogImageUrl.value= (dialogImageUrl.value + step)%(images.value.length);
}
</script>
<style lang="scss" scoped>
.img-model {
    position: relative;
    margin-right: 10px;
    margin-top: 10px;
}
.img{
  border-radius:6px;
 
}
// 显示
.dialog_img{
  width: 100%;
  min-height: calc(50vh);
  justify-content: center;
  align-items: center;
  display: flex;
}
.item_center{
  justify-content: center;
  align-items: center;
  display: flex;
  user-select: none;
}

.icon_back{
  cursor: pointer;
  font-size:24px;
  font-weight: 600;
}
.border_alt{
   border: 1px solid #000000;
   text-align:center;
}
.op-content {
    width: 24px;
    height: 24px;
    position: absolute;
    background-color: rgba(245, 245, 245, 0.8);
    top: 0;
    right: 0;
    z-index: -1;
    font-weight: 800;
    color: rgba(105,105,105,1);
    font-size: 16px;
    display: flex;
    justify-content: center;
    align-items: center;
    border-bottom-left-radius: 6px;
}
.op-content_hover {
    z-index: 10;
}
.padding10 {
    padding: 0 10px;
}
</style>
