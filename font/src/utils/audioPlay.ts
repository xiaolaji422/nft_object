import audioMp3 from "@/assets/mp3/waring.mp3"
import {notify,confirm} from "@/utils/notify";

export function WarnningPlay() {
   if (!window.audioRef){
      confirm.warning('确认开启音乐提示吗').then((res)=>{
         if(res){
            window.audioRef = new Audio(audioMp3);
         }
      })
   }
   window.audioRef.play()
}


export function WarnningStop() {
   if (!window.audioRef){
      confirm.warning('确认开启音乐提示吗').then((res)=>{
         if(res){
            window.audioRef = new Audio(audioMp3);
            window.audioRef.load()
         }
      })
   }else{
      try {
         window.audioRef.load()
      } catch (error) {
         console.log(error)
      }
   }
}