import permissionApi from "@/api/permission"
import { defineStore } from "pinia";
import { local,session  } from "@/utils/storage"
export const permissionStore = defineStore({
    // id: 必须的，在所有 Store 中唯一
  id: "permissionStore",
  state:()=>({
    apiData: [],
    apiGroupData: [],
    roleData: [],
  }),
  actions:{
    getApiData() {
        return new Promise((resolve, reject) => {
            if (this.apiData && this.apiData.length) resolve(this.apiData)
            let localData = session.get("SET_API")
            if (localData && localData.length) {
                this.apiData = localData
                return resolve(localData)
            }
            permissionApi.getApiAll({}).then(response => {
                const { data } = response
                this.apiData = data
                session.set("SET_API",data)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },
    clearApi() {
        this.apiData = []
        session.remove('SET_API')
    },
    getApiGroupData() {
        return new Promise((resolve, reject) => {
            // store 中获取
            if (this.apiGroupData && this.apiGroupData.length) resolve(this.apiGroupData)
                // 本地获取
            let localData = session.get("SET_API_GROUP")
            if (localData && localData.length) {
                this.apiGroupData = localData
                resolve(localData)
                return
            }
            permissionApi.getApiGroupAll({}).then(response => {
                const { data } = response
                this.apiGroupData = data
                session.set("SET_API_GROUP",data)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },
    clearApiGroup() {
        this.apiGroupData = []
        session.remove('SET_API_GROUP')
    },
    // 角色信息
    getRoleData() {
        return new Promise((resolve, reject) => {
            // store 中获取
            if (this.roleData && this.roleData.length) resolve(this.roleData)
                // // 本地获取
            let localData = session.get("SET_ROLE")
            if (localData && localData.length) {
                this.roleData = localData
                resolve(localData)
                return
            }
            permissionApi.getAllRights({}).then(response => {
                const { data } = response
                this.roleData = data
                session.set("SET_ROLE",data)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },
    clearRole() {
        this.roleData = []
        session.remove('SET_ROLE')
    }
}

})