import request from '@/utils/request'


export default {
    authAdminInfo(data?) {
        return request({
        url: '/admin/authUser/userInfo',
        method: 'GET',
        params: { ...data },
        });
    },
};