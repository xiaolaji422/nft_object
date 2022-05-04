import request from '@/utils/request';

export default {
    // 登出
    loginOut(data) {
        return request({
            url: '/admin/auth/loginOut',
            method: 'GET',
            params: {...data },
        });
    },
    // 登录
    login(data) {
        return request({
            url: '/admin/auth/login',
            method: 'POST',
            params: {...data },
        });
    },
    // 注册
    register(data) {
        return request({
            url: '/admin/auth/register',
            method: 'POST',
            params: {...data },
        });
    },

};