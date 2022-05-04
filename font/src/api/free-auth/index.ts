import request from '@/utils/request'
//  免权限的接口

//产品专区列表
export function getTabletList(params = {}) {
    return request({
        url: '/admin/product_info/ext_items',
        method: 'get',
        params
    })
}

//  产品下拉列表
export function getProdOptions(params = {}) {
    return request({
        url: '/admin/product_info/get_prod_options',
        method: 'get',
        params
    })
}

//  产品sdk下拉列表
export function getSDKOptions(params = {}) {
    return request({
        url: '/admin/product_info/get_sdk_options',
        method: 'get',
        params
    })
}


//  产品sdk下拉列表
export function getCSSColor(params = {}) {
    return request({
        url: '/admin/css_conf/get_color',
        method: 'get',
        params
    })
}

//  产品aiCode下拉列表
export function getAiMap(params = {}) {
    return request({
        url: '/admin/product_info/get_ai_map',
        method: 'get',
        params
    })
}

// 获取产品信息
export function getBaseInfo(params) {
    return request({
        url: '/admin/product_info/detail_code',
        method: 'get',
        params: params
    })
}