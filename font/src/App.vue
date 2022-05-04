<template>
    <ElConfigProvider :locale='locale'>
        <router-view />
    </ElConfigProvider>
</template>

<script lang='ts'>
import { defineComponent, ref, Ref, watch } from 'vue'
import locale from 'element-plus/lib/locale/lang/zh-cn'
import { ElConfigProvider } from 'element-plus'
import { changeThemeDefaultColor } from '@/utils/changeThemeColor'
import { ITheme } from '@/type/config/theme'
import theme from '@/config/theme'
import { useLayoutStore } from '@/store/modules/layout'

export default defineComponent ({
    name: 'App',
    components: {
        ElConfigProvider
    },
    setup() {
        changeThemeDefaultColor()
        const { getSetting } = useLayoutStore()

        // 重新获取主题色
        const f = () => {
            let themeArray = theme()
            return getSetting.theme >= themeArray.length ? themeArray[0] : themeArray[getSetting.theme]
        }

        let themeStyle:Ref<ITheme> = ref(f())
        watch(() => getSetting.theme, () => themeStyle.value = f())
        watch(() => getSetting.color.primary, () => themeStyle.value = f())
        return {
            locale,
            themeStyle,
            getSetting,
        }
    }
})
</script>

<style lang='postcss'>
    .el-table {
        thead {
            color: #000!important;
            font-size: 18px;
            line-height: 25px;
            th {
                font-weight: 500;
            }
        }
    }
    .el-table__header-wrapper, .el-table__fixed-header-wrapper {
        thead {
            th {
                background-color: #F7F7FA!important;
            }
        }
    }
    .el-breadcrumb__item {
        .el-breadcrumb__inner {
            color: #666!important;
            font-weight: normal!important;
        }
    }
    .el-breadcrumb__item:last-child {
        .el-breadcrumb__inner {
            color: #000!important;
        }
    }
    .el-table__fixed-right-patch {
        width: 0!important;
    }

    .layout-sidebar-sidesetting > i {
        background-color: v-bind(themeStyle.sidebarActiveBg);
        color: v-bind(themeStyle.sidebarColor);
    }

    .layout-sidebar.width {
        width: 240px!important;
    }
    .layout-sidebar {
        background-color: v-bind(themeStyle.sidebarBg);

        .layout-sidebar-logo {
            background-color: #213968;
            color: v-bind(themeStyle.logoColor || themeStyle.sidebarColor);
            padding-left: 24px;
        }

        .el-menu {
            background-color:  #213968;
            border-right: 0;
            max-width: 240px;

            .el-menu {
                background-color: #1E335C;
            }
        }

        .el-sub-menu__title {
            color: v-bind(themeStyle.sidebarColor);
        }

        .el-menu-item {
            color: v-bind(themeStyle.sidebarColor);
        }

        .el-menu-item:focus,
        .el-menu-item:hover,
        .el-sub-menu__title:hover {
            background-color: transparent;
            color: v-bind(themeStyle.sidebarActiveColor);
        }

        .el-menu-item-group__title {
            padding: 0;
        }

        .el-menu-item.is-active {
            color:#1890FF;
            border-right: 3px solid v-bind(themeStyle.sidebarActiveBorderRightBg);
        }

        .el-sub-menu.is-active > .el-sub-menu__title,
        .el-sub-menu.is-active > .el-sub-menu__title > i {
            color: v-bind(themeStyle.sidebarActiveColor);
        }
    }

    .layout-main-navbar {
        background-color: v-bind(getSetting.mode === "vertical" || getSetting.isPhone ? themeStyle.navbarBg : themeStyle.sidebarBg);
        color: v-bind(getSetting.mode === "vertical" || getSetting.isPhone ? themeStyle.navbarColor : themeStyle.sidebarColor);

        .el-breadcrumb .el-breadcrumb__inner,
        .el-breadcrumb .el-breadcrumb__separator,
        .el-breadcrumb .el-breadcrumb__inner:hover,
        .el-breadcrumb .el-breadcrumb__separator:hover,
        .el-dropdown {
            color: v-bind(themeStyle.navbarColor);
        }

        .layout-sidebar-menubar {
            .el-menu {
                background-color: v-bind(themeStyle.sidebarBg);
                border-right: 0;
            }

            .el-menu--horizontal {
                & > .el-menu-item {
                    color: v-bind(themeStyle.sidebarColor);
                }

                & > .el-sub-menu .el-sub-menu__title {
                    color: v-bind(themeStyle.sidebarColor);
                }

                .el-menu-item:not(.is-disabled):focus,
                .el-menu-item:not(.is-disabled):hover {
                    background-color: v-bind(themeStyle.sidebarActiveBg);
                    color: v-bind(themeStyle.sidebarActiveColor);
                }

                & > .el-sub-menu .el-sub-menu__title:hover {
                    background-color: v-bind(themeStyle.sidebarActiveBg);
                    color: v-bind(themeStyle.sidebarActiveColor);
                }

                & > .el-menu-item-group__title {
                    padding: 0;
                }

                & > .el-menu-item.is-active {
                    background-color: v-bind(themeStyle.sidebarActiveBg);
                    color: v-bind(themeStyle.sidebarActiveColor) !important;
                    border-right: 3px solid v-bind(themeStyle.sidebarActiveBorderRightBg);
                }

                & > .el-sub-menu.is-active > .el-sub-menu__title,
                & > .el-sub-menu.is-active > .el-sub-menu__title > i {
                    color: v-bind(themeStyle.sidebarActiveColor);
                }
            }
        }
    }

    .layout-main-tags {
        background-color: v-bind(themeStyle.tagsbg);
        color: v-bind(themeStyle.tagsColor);

        .layout-tags-active {
            background-color: v-bind(themeStyle.tagsActiveBg);
            color: v-bind(themeStyle.tagsActiveColor);
        }
    }

    .layout-main-content {
        background-color: #fff;
    }

</style>
        <!-- background-color: v-bind(themeStyle.mainBg); -->
