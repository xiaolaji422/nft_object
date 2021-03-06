import request from '@/utils/request';

export default {
    // 保存账号
    Save(data:any) {
        return request({
            url: '/admin/account/save',
            method: 'POST',
            params: {...data },
        });
    },
     // 账号列表
     List(data:any) {
        return request({
            url: '/admin/account/list',
            method: 'GET',
            params: {...data },
        });
    },
};