import request from '@/utils/request';

export default {
    // 获取所有的接口列表
    getAllApi(data) {
        return request({
            url: '/admin/api/items',
            method: 'GET',
            params: {...data },
        });
    },
    // 获取全部权限值、全部角色
    getAllRights(data) {
        return request({
            url: '/admin/role/all',
            method: 'GET',
            params: {...data },
        });
    },
    // 获取全部权限值、全部角色
    geRoleList(data) {
        return request({
            url: '/admin/role/items',
            method: 'GET',
            params: {...data },
        });
    },

    // 新建对应角色
    addRoleStore(data) {
        return request({
            url: '/admin/role/add',
            method: 'POST',
            data: {
                ...data
            },
        });
    },
    // 修改对应角色
    updateRole(data) {
        return request({
            url: '/admin/role/update',
            method: 'POST',
            data: {...data, },
        });
    },
    //---------------管理者
    // 获取管理者列表
    getActivityList(params) {
        return request({
            url: '/admin/admin/items',
            method: 'GET',
            params: params,
        });
    },
    // 获取用户角色列表
    getUserRoleList(data) {
        return request({
            url: '/admin/admin/roles',
            method: 'GET',
            params: {...data },
        });
    },
    // 获取用户接口列表
    getUserApiArray(data) {
        return request({
            url: '/admin/admin/apiArray',
            method: 'GET',
            params: {...data },
        });
    },
    // 添加用户角色
    addUserRole(data) {
        return request({
            url: '/admin/adminRole/addRole',
            method: 'POST',
            data: {...data },
        });
    },
    // 添加用户角色
    delUserRole(data) {
        return request({
            url: '/admin/adminRole/delRole',
            method: 'POST',
            data: {...data },
        });
    },
    // 添加用户定制接口
    addUserApis(data) {
        return request({
            url: '/admin/adminApi/addApis',
            method: 'POST',
            data: {...data },
        });
    },
    // 启用/禁用用户角色
    enableAdminRole(data) {
        return request({
            url: '/admin/adminRole/enable',
            method: 'POST',
            data: {...data },
        });
    },
    // 启用/禁用  用户定制权限
    enableAdminApi(data) {
        return request({
            url: '/admin/adminApi/enable',
            method: 'POST',
            data: {...data },
        });
    },
    // 获取人员的接口列表
    getAdminApis(data) {
        return request({
            url: '/admin/admin/apis',
            method: 'GET',
            params: {...data },
        });
    },
    // 获取全部角色
    getAllRightsCfg(data) {
        return request({
            url: '/admin/role/getAllRightsCfg',
            method: 'GET',
            params: {...data },
        });
    },

    // 给用户赋权
    updateStaff(data) {
        return request({
            url: '/admin/staff/update',
            method: 'POST',
            data: {...data },
        });
    },

    // -------------------
    // 修改权限
    updateRoleApi(data) {
        return request({
            url: '/admin/role/setRoleApis',
            method: 'POST',
            data: {...data },
        });
    },
    // 获取对应角色的权限值
    getRoleRights(data) {
        return request({
            url: '/admin/role/getRoleApis',
            method: 'GET',
            params: {...data },
        });
    },

    //      接口api
    getApiList(params) {
        return request({
            url: '/admin/api/items',
            method: 'GET',
            params: params,
        });
    },
    getApiAll(data) {
        return request({
            url: '/admin/api/all',
            method: 'GET',
            params: {...data },
        });
    },
    getApiDetail(data) {
        return request({
            url: '/admin/api/info',
            method: 'GET',
            params: {...data },
        });
    },
    addApi(data) {
        return request({
            url: '/admin/api/add',
            method: 'POST',
            data: {...data },
        });
    },
    editApi(data) {
        return request({
            url: '/admin/api/edit',
            method: 'POST',
            data: {...data },
        });
    },
    enableApi(data) {
        return request({
            url: '/admin/api/enable',
            method: 'POST',
            data: {...data },
        });
    },
    //      接口分组api
    getApiGroupList(params) {
        return request({
            url: '/admin/apiGroup/items',
            method: 'GET',
            params: params,
        });
    },
    getApiGroupAll(data) {
        return request({
            url: '/admin/apiGroup/all',
            method: 'GET',
            params: {...data },
        });
    },
    getApiGroupDetail(data) {
        return request({
            url: '/admin/apiGroup/info',
            method: 'GET',
            params: {...data },
        });
    },
    addApiGroup(data) {
        return request({
            url: '/admin/apiGroup/add',
            method: 'POST',
            data: {...data },
        });
    },
    editApiGroup(data) {
        return request({
            url: '/admin/apiGroup/edit',
            method: 'POST',
            data: {...data },
        });
    },
    enabledApiGroup(data) {
        return request({
            url: '/admin/apiGroup/enabled',
            method: 'POST',
            data: {...data },
        });
    },
};