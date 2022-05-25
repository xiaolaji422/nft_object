import { defineStore } from 'pinia'
import audioMp3 from "@/assets/mp3/waring.mp3"
import {notify,confirm} from "@/utils/notify";

export const playerStore = defineStore({
    id:"palyer",
    state: ()=> ({
        open:false,
        player:null
    }),

    actions:{
        Loadding(){
            console.log("Loadding","player")
            if (this.player && this.player!=null){
                try {
                    this.player.load()
                 } catch (error) {
                    console.log(error)
                 }
             }else{
                try {
                    this.player = new Audio(audioMp3);
                 } catch (error) {
                    console.log(error)
                 }
             }
        },
        Play(info:string){
            if(!this.player){
                notify.warning("暂未开启音乐通知:"+info)
            }else{
                try {
                    this.player.play() 
                } catch (error) {
                    notify.success("暂未开启音乐通知:"+info)
                }
            }
        }, 
    }
})