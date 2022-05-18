<template>
    <div  :class="$style.warpper">
        <div :class="$style['lock-body']">
            <div :class="$style['lock-item']" v-for="item in 4" :key="item">
                <lock></lock>
            </div>
        </div>
        
    </div>    

</template>
<script lang="ts" setup>
import {ref,defineComponent} from "vue"
import Lock from "./components/lock.vue"

const formData  = ref([
    {admin:"",cookie:""}
])
defineComponent({
    Lock,
})

const saveCookie = ()=>{
    console.log(formData.value,"saveCookie")
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