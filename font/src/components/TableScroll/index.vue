<template>
    <div  class="custom_scroll " ref="customScrollRef" >
      <div class= "custom_scroll_body " :style="`width:${scrollWidth}`"></div>
    </div>
</template>

<script lang="ts" setup>
import {defineProps,onMounted,ref,defineEmits, watch} from 'vue'

const props = defineProps({
    scrollWidth:{
        type:String,
        default:"100%",
    }

})
const emit = defineEmits(["change"])
const customScrollRef =ref(null)
const scrool = function (val){
      let scrollLeft = val.target.scrollLeft   
      emit("change",{scrollLeft:scrollLeft}) 
}
onMounted(()=>{
      setTimeout(()=>{
        customScrollRef.value.addEventListener("scroll",scrool)
      },0)
})
</script>
<style lang="scss">
// 滚动条
.custom_scroll{
  position: sticky;
  bottom:12px;
  height: 20px;
  width: calc(100%);
  z-index: 999;
  background-color:#fff ;
   overflow-x:auto ;
  .custom_scroll_body{
    height: 1px;
  }
}
.custom_scroll::-webkit-scrollbar {
    width:12px;
}
/* 滚动槽 */
.custom_scroll::-webkit-scrollbar-track {
  border-radius:10px;
}
/* 滚动条滑块 */
.custom_scroll::-webkit-scrollbar-thumb {
  border-radius:10px;
  background:rgba(144, 147, 153,0.5);
  // opacity: 0.3;
}

.custom_scroll::-webkit-scrollbar-thumb:hover {
  background:rgba(144, 147, 153,0.8);
  // opacity: 0.5;
}
</style>