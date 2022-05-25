import request from '@/utils/request';

export default {
    // 商品搜索
    serachAblum(data:any) {
        return request({
            url: '/admin/album/search',
            method: 'GET',
            params: {...data },
        });
    },
};