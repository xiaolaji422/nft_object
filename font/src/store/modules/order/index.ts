import { defineStore } from 'pinia'
import audioMp3 from "@/assets/mp3/waring.mp3"
import {notify,confirm} from "@/utils/notify";

export const playerStore = defineStore({
    id:"order",
    state: ()=> ({
        cookies:[],
    }),

    actions:{
        
    }
})