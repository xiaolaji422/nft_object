import { AxiosRequestConfig } from "axios";

declare module 'axios' {
  export interface AxiosRequestConfig  {
    requestLimit?: boolean,   // 是否开启限频
    limitKey?: String,
    limitTimes?: Number,
  }
}