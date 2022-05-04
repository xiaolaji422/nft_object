import request from '@/utils/request'

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

// 保存产品信息
export function saveBaseInfo(params) {
    return request({
        url: '/admin/product_info/extend',
        method: 'post',
        params: params
    })
}

// 获取产品信息
export function getBaseInfo(params) {
    return request({
        url: '/admin/product_info/ext_detail',
        method: 'get',
        params: params
    })
}

// 保存专区分类
export function saveZoneCategory(params) {
    return request({
        url: '/admin/product_info/set_sdk',
        method: 'get',
        params: params
    })
}

// 同步至现网
export function launchOnline(params) {
    return request({
        url: '/admin/product_info/online_source',
        method: 'post',
        params: params
    })
}

// 上传图片
export function uploadIcons(params) {
    return request({
        url: '/admin/icon_conf/uploads',
        method: 'post',
        params: params
    })
}

// 手游 AI 产品名称
export function getAIProName(params = {}) {
    return request({
        url: 'http://kfwp.cm.com/xiaoyanAi/getClassL1/',
        method: 'get',
        params
    })
}