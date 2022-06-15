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
        appid:"422nft",
        userid:""
    }),
    actions:{
        async sendMsg (data:Object){
            if(!this.ws){
                console.log("我还未注册","sendMsg",data)
                this.reigster(this.login_name)
            }else{
                data.appid = this.appid
                data.userid = this.userid+"nft"
                this.ws.send(data)
            }
       },
       async reigster(userid:any){
           
           this.userid = userid
           if(!this.ws){
                this.ws = socketUtil.socket
                this.ws.registerWebSorkt()
                var self = this
                setTimeout(() => {
                    self.sendMsg({userid:userid,appid:this.appid})
                }, 1000);
           }else{
               this.sendMsg({userid:userid,appid:this.appid})
           }
       }
    }
})