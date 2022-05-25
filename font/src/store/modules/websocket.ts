import { defineStore } from 'pinia'
import audioMp3 from "@/assets/mp3/waring.mp3"
import {notify,confirm} from "@/utils/notify";
import socketUtil from "@/utils/websocket"
import { time } from 'console';
import { Timer } from '@element-plus/icons';


export const websocketStore = defineStore({
    id:"websockte",
    state: ()=> ({
        ws:null,
        appid:422,
        login_name:""
    }),
    actions:{
        async sendMsg (data:any){
            if(!this.ws){
                console.log("我还未注册","sendMsg",data)
                this.reigster(this.login_name)
            }else{
                this.ws.send(data)
            }
       },
       async reigster(login_name:any){
           console.log(login_name,"register")
           this.login_name = login_name
            if(!this.ws){
                this.ws = socketUtil.socket
                this.ws.registerWebSorkt()
            }
       }
    }
})