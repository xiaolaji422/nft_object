import { defineStore } from "pinia";


export function initCategoryList() {
  return new Array(6).fill(0).map((_, i) => {
    return {
      sort: i,
      title: '',
      desc: ''
    }
  })
}
// defineStore 调用后返回一个函数，调用该函数获得 Store 实体
export const useProductZoneStore = defineStore({
  // id: 必须的，在所有 Store 中唯一
  id: "productZone",
  // state: 返回对象的函数
  state: ()=> ({
    reloadPageStatus: false, // 上传文件图片更新页面
    product_code: '',
    categoryList: initCategoryList(),
    aiList: []
  }),
  getters: {
    getProductCode(): string {
        return this.product_code
    },
    getCategoryList(): any {
      return this.categoryList
    },
    getReloadPageStatus(): boolean {
      return this.reloadPageStatus
    },
    getAIList(): any {
      return this.aiList
    }
  },
  actions: {
    setProductCode(code: string) {
      this.product_code = code
    },
    setCategoryList(val: any, delIndex: number) {
      
      if (delIndex >= 0) {
        this.categoryList.splice(delIndex, 1)
        return
      }
      if (Array.isArray(val) && val.length) {
        this.categoryList = [
          ...val
        ]
      } else {
        let len = this.categoryList.length
        this.categoryList.push({
          sort: ++len,
          ...val
        })
      }
    },
    setReloadPageStatus(status: boolean) {
      this.reloadPageStatus = status
    },
    setAIList(val: any) {
      this.aiList = val
    }
  }
});