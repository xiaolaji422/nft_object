import { defineStore } from "pinia";
import { stringify } from "querystring";

/**
 * 页面搜索条件缓存
 * 仅缓存最新的table查询条件
 * 切换之后自动覆盖  覆盖的参数  pageName
 * 
 */
export const cachePageStore = defineStore({
    // id: 必须的，在所有 Store 中唯一
  id: "cachePageStore",
  state:()=>({
    pageInfo:{},
    pageName:"",
  }),
  actions:{
    setPageInfo(pageName,pageInfo) {
        this.pageName = pageName
        this.pageInfo = pageInfo
    },
    getPageInfo(pageName) {
        return new Promise((resolve, reject) => {
            if(this.pageName!=pageName){
                resolve({})
            }
            resolve(this.pageInfo)
        })
    }
    
}

})