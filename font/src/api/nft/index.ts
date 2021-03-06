import request from '@/utils/request';

export default {
    // 商品搜索
    serachAlbum(data:any) {
        return request({
            url: '/admin/album/search',
            method: 'GET',
            params: {...data },
        });
    },
     // 商品搜索
     albumDetail(data:any) {
        return request({
            url: '/admin/album/price_list',
            method: 'GET',
            params: {...data },
        });
    },
};