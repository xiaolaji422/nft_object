import { ElMessage } from 'element-plus';
import { defineStore } from "pinia";

// 全局共用的状态
export const useApptore = defineStore({
  id: "App",
  state: ()=> ({
  }),
  getters: {
  },
  actions: {
    /**
     * 设置按钮操作执行等待时间
     * @param key 按钮操作的 key
     * @param s 操作在 s 秒内不能再次触发
     */
    setActionWaitMSStatus(key: string, s: Number) {
      if (!key) return
      if (s === +s) {
        s = s * 1000 // 转成 ms
      } else {
        s = 1000 // 默认 1s
      }
      const data = {
        actionTime: +new Date(),
        waitTime: s
      }
      sessionStorage.setItem(key, JSON.stringify(data))
    },
    /**
     * 判断当前操作是否可以执行
     * @param key 按钮操作的 key
     */
    getActionWaitMSStatus(key: string) {
      let status = true

      return new Promise((resolve, reject) => {
        if (!key) {
          resolve(status)
        } else {
          let now = +new Date()
          try {
            let { actionTime  = 0, waitTime = 0 } = JSON.parse(sessionStorage.getItem(key) as string) || {}
            if (now - actionTime <= waitTime) {
              ElMessage.warning(`请在 ${((waitTime - now + actionTime) / 1000)}s 后执行此操作`)
              status = false
            }
            resolve(status)
          } catch (error) {
            
            ElMessage.error(error ?? '时间存储数据异常')
            reject(error)
          }
        }
      })
    }
  }
});