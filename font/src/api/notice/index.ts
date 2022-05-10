import request from '@/utils/request';

export default {
    // 获取所有的接口列表
    queryNotice(data:any) {
        return request({
            url: '/admin/notice/query_notice',
            method: 'GET',
            params: {...data },
        });
    },
    queryHistoryNotice(data:any) {
        return request({
            url: '/admin/notice/query_his_notice',
            method: 'GET',
            params: {...data },
        });
    }
};