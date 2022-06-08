import request from '@/utils/request';

export default {
    // 创建锁单
    Save(data:any) {
        return request({
            url: '/admin/account_lock/save',
            method: 'POST',
            params: {...data },
        });
    },
     // 锁单列表
     List(data:any) {
        return request({
            url: '/admin/account_lock/list',
            method: 'GET',
            params: {...data },
        });
    },
     // 取消锁单
     Cancel(data:any) {
        return request({
            url: '/admin/account_lock/cancel',
            method: 'POST',
            params: {...data },
        });
    },
};