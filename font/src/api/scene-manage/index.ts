import request from "@/utils/request";
export function getTableList(params){
    return request({
        url:'/admin/scene_conf/items',
        method:'get',
        params
    })
}

export function getItemDetail(params){
    return request({
        url:'/admin/scene_conf/detail',
        method:'get',
        params
    })
}

export function addNewScene(params){
    return request({
        url:'/admin/scene_conf/add',
        method:'post',
        params,
        requestLimit:true,
    })
}

export function editTableItem(params){
    return request({
        url:'/admin/scene_conf/edit',
        method:'post',
        params,
        requestLimit:true
    })
}
export function deleteTableItem(params){
    return request({
        url:'/admin/scene_conf/delete',
        method:'post',
        params
    })
}
export function sceneShow(params){
    return request({
        url:'/admin/scene_conf/show',
        method:'post',
        params,
        requestLimit:true,
        limitTimes:2,
        limitKey:"id",
    })
}
export function sceneUnshow(params){
    return request({
        url:'/admin/scene_conf/unshow',
        method:'post',
        params,
        requestLimit:true,
        limitTimes:2,
        limitKey:"id",
    })
}