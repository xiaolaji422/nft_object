<template>
    <div class='layout flex h-screen'>
        <div class='layout-sidebar-mask fixed w-screen h-screen bg-black bg-opacity-25 z-20' :class='{"hidden": getMenubar.status !== 2 }' @click='changeCollapsed' />
        <div
            v-if='getSetting.mode === "vertical" || getMenubar.isPhone'
            class='layout-sidebar flex flex-col h-screen transition-width duration-200 shadow'
            :class='{ 
                "w-64 width": getMenubar.status === 0 || getMenubar.status === 2, 
                "w-0": getMenubar.status === 3, 
                "w-16": getMenubar.status === 1, 
                "absolute z-30": getMenubar.status === 2 || getMenubar.status === 3, 
            }'
        >
            <div class='layout-sidebar-logo flex h-12 relative items-center shadow-lg'>
                <img style='width: 36px;height: 36px;' :src='icon'>
                <span v-if='getMenubar.status === 0 || getMenubar.status === 2' class='pl-2'>NFT 辅助系统</span>
            </div>
            <div class='layout-sidebar-menubar flex flex-1 overflow-hidden'>
                <el-scrollbar wrap-class='scrollbar-wrapper' style="height:100%">
                    <layout-menubar />
                </el-scrollbar>
            </div>

        <div class="login-out" @click="handleLoginOut">
            <i class="el-icon-circle-close"></i> 
            <span style="margin-left:10px;font-size: 14px;">{{userInfo.name}}</span>
        </div>
        </div>
        <div class='layout-main flex flex-1 flex-col overflow-x-hidden overflow-y-auto'>
            <div class='layout-main-navbar flex justify-between items-center h-12 shadow-sm overflow-hidden relative z-10'>
                <layout-navbar />
            </div>
         
            <div class='layout-main-content flex-1 overflow-hidden'>
                <layout-content />
            </div>
           
        </div>
    </div>
</template>

<script lang='ts'>
import { defineComponent, onMounted } from 'vue'
import LayoutContent from '@/layout/components/content.vue'
import LayoutMenubar from '@/layout/components/menubar.vue'
import LayoutNavbar from '@/layout/components/navbar.vue'
import LayoutTags from '@/layout/components/tags.vue'
import LayoutSideSetting from '@/layout/components/sideSetting.vue'
import { throttle } from '@/utils/tools'
import { useLayoutStore } from '@/store/modules/layout'
import { playerStore } from '@/store/modules/palyser'
import icon from '@/assets/img/icon.png'
import { confirm ,notify} from '@/utils/notify';
import loginApi from '@/api/login'
import router from '@/router'
export default defineComponent ({
    name: 'Layout',
    components: {
        LayoutContent,
        LayoutMenubar,
        LayoutNavbar,
        LayoutTags,
        LayoutSideSetting
    },
    setup() {
        const {Loadding} = playerStore()
        const { changeDeviceWidth, changeCollapsed, getMenubar, getSetting,getUserInfo } = useLayoutStore()
        onMounted(async() => {
            changeDeviceWidth()
            const throttleFn = throttle(300)
            let throttleF = async function() {
                await throttleFn()
                changeDeviceWidth()
            }
            Loadding()
            window.addEventListener('resize', throttleF, true)
        })
        
        const handleLoginOut =async ()=>{
                const isOk = await confirm.warning('您确定要退出登录系统吗');
                if (!isOk) {
                    return;
                }
                const result: any = await loginApi.loginOut({});
                if (result !== false) {
                    notify.success('退出登录成功');
                    try{
                        clearInterval(window.autoPlayMp3)
                        setTimeout(() => {
                        router.push({
                            path:"/Login"
                        })
                    }, 500);
                    }catch(err){
                        console.log(err,"clearInterval")
                     }
                    
                }
        }
      
        return {
            getMenubar,
            getSetting,
            changeCollapsed,
            icon,
            handleLoginOut,
            // userInfo,
            userInfo: getUserInfo,
        }
    }
})
</script>

<style lang='postcss' scoped>
    ::v-deep(.layout-sidebar-sidesetting .el-drawer__header) {
        margin-bottom: 0;
    }

    ::v-deep(.el-menu--horizontal>.el-menu-item) {
        height: 48px;
    }

    ::v-deep(.el-menu--horizontal>.el-sub-menu .el-sub-menu__title) {
        height: 48px;
        line-height: 48px;
    }
    .layout-sidebar-menubar {
        background-color: #213968!important;
    }

    .login-out{
        color:#ffffff;
        position: fixed;
        bottom: 10px;
        width: 240px;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>