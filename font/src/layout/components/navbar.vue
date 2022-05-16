<template>
    <div v-if='getSetting.mode === "vertical" || getMenubar.isPhone' class='flex items-center px-4 flex-wrap h-12 leading-12'>
        <span class='text-2xl cursor-pointer h-12 leading-12' :class='{ "el-icon-s-fold": !getMenubar.status, "el-icon-s-unfold": getMenubar.status }' @click='changeCollapsed' />
        <!-- 面包屑导航 -->
        <div class='px-4'>
            <el-breadcrumb separator='/'>
                <transition-group name='breadcrumb'>
                    <el-breadcrumb-item key='/' :to='{ path: "/" }'>主页</el-breadcrumb-item>
                    <el-breadcrumb-item v-for='(v, i) in data.breadcrumbList' :key='v.path' :to='i === data.breadcrumbList.length - 1 ? "" : v.path'>{{ v.title }}</el-breadcrumb-item>
                </transition-group>
            </el-breadcrumb>
        </div>
    </div>

    <div class='flex items-center flex-row-reverse px-4 min-width-32'>
         <!-- <div class="close-autio" >关闭铃声</div> -->
         <!-- -->
         <!-- <span style="font-size: 12px;" >铃声开启</span> -->
         <el-checkbox v-model="checked1"  @click="playMp3">铃声</el-checkbox> 
         <!-- <img :src="closeAudioIcon" class="menu-icon" alt=""  @click="closeAudio"> -->
    </div>
</template>

<script lang='ts'>
import { defineComponent, reactive, watch,ref ,onMounted} from 'vue'
import { useLayoutStore } from '@/store/modules/layout'
import { useRoute, RouteLocationNormalizedLoaded } from 'vue-router'
import Notice from '@/layout/components/notice.vue'
import Screenfull from '@/layout/components/screenfull.vue'
import Search from '@/layout/components/search.vue'
import LayoutMenubar from '@/layout/components/menubar.vue'
import icon from '@/assets/img/icon.png'
import {playerStore} from "@/store/modules/palyser"
// import {WarnningStop} from "@/utils/audioPlay"
import closeAudioIcon from '@/assets/img/closeAudio.png'
interface IBreadcrumbList {
    path: string
    title: string | symbol
}
const {Loadding,Play,player}  = playerStore()
// 面包屑导航
const breadcrumb = (route: RouteLocationNormalizedLoaded) => {
    const fn = () => {
        const breadcrumbList:Array<IBreadcrumbList> = []
        const notShowBreadcrumbList = ['Dashboard', 'RedirectPage'] // 不显示面包屑的导航
        if(route.matched[0] && (notShowBreadcrumbList.includes(route.matched[0].name as string))) return breadcrumbList
        route.matched.forEach(v => {
            const obj:IBreadcrumbList = {
                title: v.meta.title as string,
                path: v.path
            }
            breadcrumbList.push(obj)
        })
        // 同名删一个
        let map = {}, arr = []
        breadcrumbList.forEach(item => {
            if (!map[item.title]) {
                map[item.title] = true
                arr.push(item)
            }
        })
        return arr
    }
    let data = reactive({
        breadcrumbList: fn()
    })
    watch(() => route.path, () => data.breadcrumbList = fn())
    const closeAudio = ()=>{
        Loadding()
    }
    return { data ,closeAudio}
}

export default defineComponent ({
    name: 'LayoutNavbar',
    components: {
        Notice,
        Search,
        Screenfull,
        LayoutMenubar
    },
    setup() {
        const { getMenubar, getUserInfo, changeCollapsed, logout, getSetting } = useLayoutStore()
        const route = useRoute()
        const checked1 = ref<boolean>(true)
        onMounted(()=>{
            if(player == null){
                checked1.value = false
            }
        })
        const playMp3= ()=>{
            if (checked1.value == true){   
                Loadding()
            }
            // checked1.value = !checked1.value
        }
        return {
            getMenubar,
            userInfo: getUserInfo,
            changeCollapsed,
            logout,
            ...breadcrumb(route),
            getSetting,
            icon,
            closeAudioIcon,
            checked1,
            playMp3
        }
    }
})
</script>

<style lang='postcss' scoped>
.breadcrumb-enter-active,
.breadcrumb-leave-active {
    transition: all 0.5s;
}

.breadcrumb-enter-from,
.breadcrumb-leave-active {
    opacity: 0;
    transform: translateX(20px);
}

.breadcrumb-move {
    transition: all 0.5s;
}

.breadcrumb-leave-active {
    position: absolute;
}

.menu-icon {
    width: 20px;
    height: 20px;
    margin-right: 13px;
}
</style>