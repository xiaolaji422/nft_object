import request from '@/utils/request'

export function getTableList(params) {
    return request({
        url: '/admin/api/items',
        method: 'get',
        params: params
    })
}

export function deleteTableItem(id) {
    return request({
        url: '/admin/api/items',
        method: 'get',
        params: id
    })
}