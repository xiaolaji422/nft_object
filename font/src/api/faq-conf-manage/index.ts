import request from "@/utils/request";

//  列表
export function getTableList(params){
    return request({
        url:'/admin/faq_conf/items',
        method:'get',
        params:params
    })
}

// FAQ详情
export function detailFAQ(params){
    return request({
        url:'/admin/faq_conf/detail',
        method:'get',
        params:params
    })
}

// 新增FAQ
export function addFAQ(params){
    return request({
        url:'/admin/faq_conf/add',
        method:'post',
        params:params
    })
}

// 修改FAQ
export function editFAQ(params){
    return request({
        url:'/admin/faq_conf/edit',
        method:'post',
        params:params
    })
}

// 上线FAQ
export function showFAQ(params){
    return request({
        url:'/admin/faq_conf/show',
        method:'post',
        params:params
    })
}

// 下线FAQ
export function unshowFAQ(params){
    return request({
        url:"/admin/faq_conf/unshow",
        method:'post',
        params:params,
        requestLimit:true,          // 开启限频  默认false
        limitTimes:10,          // 限频时间  单位：秒
        limitKey:"id",          // 需要对单条记录限频的  此处可设置为id，默认为"",为""时会对整个接口限频
    })
}

//  删除
export function deleteFAQ(params){
    return request({
        url:'/admin/faq_conf/delete',
        method:'post',
        params:params
    })
}

//  faq归档下拉列表
export function getArchivecOptions(params = {}) {
    return request({
        url: '/admin/prod_archive/get_options',
        method: 'get',
        params
    })
}

//  faq下拉列表
export function getFAQOptions(params = {}) {
    return request({
        url: '/admin/faq_conf/get_options',
        method: 'get',
        params
    })
}
