import request from "@/utils/request";

export function getTableList(params:any){
    return request({
        url:'/admin/product_info/items',
        method:'get',
        params,
     
    })
}

export function getItemDetail(params:any){
    return request({
        url:'/admin/product_info/detail',
        method:'get',
        params:{id:params.id}
    })
}

export function saveItemDetail(data:any){
    return request({
        url:"/admin/product_info/init",
        method:"post",
        data
    })
}

export function prodInfoConfig(data:any){
    return request({
        url:"/admin/product_info/open",
        method:"post",
        data
    })
}