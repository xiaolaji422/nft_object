const dataMap=new Map( )
dataMap.set("status", [
    { text: "启用", value: 1 },
    { text: "禁用", value: 0 },
])

dataMap.set("sex", [
    { text: "男", value: 1 },
    { text: "女", value: 0 },
])

dataMap.set("notice_type", [
    { value: 1, text: "ibox",icon:"" },
])

export function getMapData(properties:string) {
    if (dataMap.has(properties)){
        return dataMap.get(properties)
    }else{
        return null
    }
}