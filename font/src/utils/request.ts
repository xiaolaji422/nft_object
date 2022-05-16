import { useLayoutStore } from "@/store/modules/layout";
import axios from "axios";
import router from '@/router'
import{local} from "./storage"
import { AxiosResponse } from "axios";
import { ElLoading, ElMessage, ElNotification } from "element-plus";
import { useApptore } from '@/store/modules/app'
import { flatMap } from "lodash";
import { URL } from "url";
// 序列化参数
// @ts-ignore
function objectToFormData(obj, rootName?, ignoreList?) {
  const result: any = {};

  function recurse(src, prop) {
    const { toString } = Object.prototype;
    if (toString.call(src) == "[object Object]") {
      let isEmpty = true;
      for (const p in src) {
        isEmpty = false;
        if (src[p]) {
          recurse(src[p], prop ? `${prop}[${p}]` : p);
        }
      }
      if (isEmpty && prop) {
        result[prop] = {};
      }
    } else if (toString.call(src) == "[object Array]") {
      const len = src.length;
      if (len > 0) {
        src.forEach((item: any, index: Number) => {
          recurse(item, prop ? `${prop}[${index}]` : index);
        });
      }
    } else {
      result[prop] = src;
    }
  }

  recurse(obj, "");

  return result;
}

// 获取参数
function getUrlArgs() {
  const args: any = {};
  let query = window.location.href;
  if (query) {
    query = query.split("?")[1] || "";
  } else {
    query = "";
  }
  const pairs = query.split("&");
  for (let i = 0; i < pairs.length; i++) {
    const pos = pairs[i].indexOf("=");
    if (pos == -1) continue;
    const argname = pairs[i].substring(0, pos);
    const value = pairs[i].substring(pos + 1);
    args[argname] = value;
  }
  return args;
}

// 设置限频
function setLimit(config:any){
  if (!config.requestLimit) return true;
  let limit_time = config.limitTimes ?? 1
  const data = {
    actionTime: +new Date(),
    waitTime: limit_time *1000
  }
  sessionStorage.setItem(btoa(config.limitKey), JSON.stringify(data))
}
//  获取是否限频
function getLimit(config:any) {
  if (!config.requestLimit) return true;
  try {
   
    let now = +new Date()
    let { actionTime  = 0, waitTime = 0 } = JSON.parse(sessionStorage.getItem(btoa(config.limitKey) ) as string) || {}
    if (now - actionTime <= waitTime) {
      let msg = `点击太快啦！请在 ${((waitTime - now + actionTime) / 1000)}s 后再试`
      return msg
    }
    return true
  }catch(error){
    return "请联系管理员：l10001"
  }
}
let loading: { close(): void };
// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: import.meta.env.VITE_BASE_URL as string,
  withCredentials: true,
  timeout: 50000, // 请求超时时间
  headers: {'X-Requested-With': 'XMLHttpRequest'},
  requestLimit: false,   // 是否开启限频
  limitKey: "",  // 限频的附件字段
  limitTimes: 1,  // 限频时长
});

// 异常拦截处理器
const errorHandler = (error: { message: string }) => {
  loading.close();
  console.log(`err${error}`);
  ElNotification({
    title: "请求失败",
    message: error.message,
    type: "error",
  });
  return Promise.reject(error);
};
// request interceptor
request.interceptors.request.use((config: any) => {
  config.headers['Auth-Sign'] =local.get("sign_xiaolaji")
  //  处理限频的参数
  if(config.limitKey){
    if(config.params && config.params[config.limitKey]) {
      config.limitKey = config.url + config.params[config.limitKey]
    }else if (config.data && config.data[config.limitKey]){
      config.limitKey = config.url +config.data[config.limitKey]
    }
  }
  // 是否需要序列化数据
  const processData = config.processData;
  if (processData) {
    delete config.params.processData;
    if (config.data) {
      config.params = config.data;
      delete config.data;
    }
    config.params = objectToFormData(config.params);
  }
  if (config.method.toUpperCase() == 'POST') {
    const FormData = new window.FormData();
    if (config.params) {
      config.data = config.data || config.params;
      delete config.params;
    }

    for (const key in config.data) {
      if(config.data[key]!=null &&config.data[key]!=undefined ){
        FormData.append(key, config.data[key]);
      }
    }
    config.data = FormData;
  }

  //  判断限频
  let check_limit =  getLimit( config) 
  if (check_limit=== true){
    return config;
  }else{
    return Promise.reject({msg:check_limit});
  }
}, (error) => {
  // do something with request error
  return Promise.reject(error);
});
// response interceptor
request.interceptors.response.use((response: AxiosResponse<IResponse>) => {
 
  const res: any = response.data;
    if (res.code == 4001) {
      router.replace({path: '/ErrorPage/401'});
      return Promise.reject('没有权限');
    }else if (res.code == 1005) {
      // 暂时关闭登录;
      router.replace({path: '/Login'});
      return Promise.reject('未登录');
    }else if (res.code == 1004) {
      router.replace({path: '/Login'});
      // window.location.href =  url;
      return Promise.reject('未登陆');
    }else if(res.code > 100) {
      if(res.msg && res.msg.length){
        ElMessage({
          message: res.msg || '网络错误，请稍后再试',
          type: 'error',
          duration: 5 * 1000,
        });
      }
      return Promise.reject(res || '网络错误，请稍后再试');
    }else if (res.code ==100){
      console.log(res,"10000")
      if (res && res.data && res.data.sign){
        local.set("sign_xiaolaji",res.data.sign)
      }
    }
  // 限频管理
  if(response.config && response.config.requestLimit){
    setLimit(response.config)
  }
    return res;
}, (error) => {
  ElMessage({
    message: error.msg || error.message,
    type: 'error',
    duration: 5 * 1000,
  });
  return Promise.reject(error)
})

export default request;
