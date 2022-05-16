<template>
    <div class='flex flex-center'>
        <el-row class='errPage-container'>
            <el-col :span='6'>
                
            </el-col>
            <el-col :span='12'>
                <img :src='errGif' width='313' height='428' alt='Girl has dropped her ice cream.'>
                <p class='  my-6'>暂无权限，请联系<span style="color:#1890FF;">QQ1185753000或251974419</span></p>
            </el-col>
            <el-col :span='6'> </el-col>
         
        </el-row>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import errGif from '@/assets/img/401.png'
import { useLayoutStore } from '@/store/modules/layout'
export default defineComponent({
    name: '401',
    setup() {
        const { color } = useLayoutStore().getSetting
        
        return {
            errGif: `${errGif}?${+new Date()}`,
            color
        }
    }
})
</script>

<style lang='postcss' scoped>
.errPage-container {
    padding-top: 100px;
    width: 800px;

    .pan-back-btn {
        background: #008489;
        color: #fff;
        border: none !important;
    }

    .pan-gif {
        margin: 0 auto;
        display: block;
    }

    .pan-img {
        display: block;
        margin: 0 auto;
        width: 100%;
    }

    .text-jumbo {
        font-size: 60px;
        font-weight: 700;
        color: #484848;
    }

    .list-unstyled {
        font-size: 14px;

        li {
            padding-bottom: 5px;
        }

        a {
            color: #008489;
            text-decoration: none;

            &:hover {
                text-decoration: underline;
            }
        }
    }

    .bullshit__return-home {
        display: block;
        float: left;
        width: 110px;
        height: 36px;
        background: v-bind(color.primary);
        border-radius: 50px;
        text-align: center;
        color: #fff;
        font-size: 14px;
        line-height: 36px;
        cursor: pointer;
        animation-name: slideUp;
        animation-duration: 0.5s;
        animation-delay: 0.3s;
        animation-fill-mode: forwards;
    }

    img {
        filter: brightness(95%);
    }
}
</style>