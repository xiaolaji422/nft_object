import request from '@/utils/request';

export default {
    // 登出
    loginOut(data:any) {
        return request({
            url: '/admin/auth/loginOut',
            method: 'POST',
            params: {...data },
        });
    },
    // 登录
    login(data:any) {
        return request({
            url: '/admin/auth/login',
            method: 'POST',
            params: {...data },
        });
    },
    // 注册
    register(data:any) {
        return request({
            url: '/admin/auth/register',
            method: 'POST',
            params: {...data },
        });
    },

};