// WebSocket注册ts示例:

import {notify,confirm} from "@/utils/notify";


const heartCheck = {
        timeout: 60000,//60秒
        timeoutObj: null,
        reset: function(){
            clearTimeout(this.timeoutObj);
            return this;
        },
        start: function(){
            this.timeoutObj = setTimeout(function(){
                //这里发送一个心跳，后端收到后，返回一个心跳消息，
                //onmessage拿到返回的心跳就说明连接正常
                ws.send("HeartBeat");
            }, this.timeout)
        }
    }
class Socket {
    ws: any;
    timeoutObj: any;
    //注册ws
    registerWebSorkt(){
        // let url = 'ws://127.0.0.1:9559'
        let url = 'ws://39.108.59.113:9559'
        this.ws = new WebSocket(url);
        this.ws.onmessage = this.loginSuccCall;//WebSorkt通知
        this.ws.onopen  = this.onOpen
        this.ws.onerror = this.websocketOnError();//WebSorkt异常
        this.ws.onclose = this.websocketClose();//WebSorkt关闭
        this.onheartbeat()
        // this.ws.onheartbeat = this.onheartbeat
        // setTimeout(()=>{this.overTime(this.ws)},1000*60*10);//10分钟二维码登录过期
    }

    //关闭ws
    overTime(ws: WebSocket){
        console.log("主动关闭websocket========");
        ws.close();
    }

    send(data: any): void {
        if (!this.ws || this.ws.readyState !== this.ws.OPEN) {
          throw new Error('没有连接到服务器，无法推送')
        }
        data = JSON.stringify(data)
        this.ws.send(data)
    }

    onOpen(event:any){
        // this.onheartbeat()
        console.log("that.websocket: connection success",event);
    }

    onheartbeat() {
        var self = this
        this.timeoutObj = setInterval(function(){
            //这里发送一个心跳，后端收到后，返回一个心跳消息，
            //onmessage拿到返回的心跳就说明连接正常
            self.ws.send("ping");
        }, 3000)
      };

    //ws通知监听
    loginSuccCall(wsData: any){ 
         if(wsData && wsData.data&& wsData.data.length){
             try {
                const res = JSON.parse(wsData.data)
                if(!res ){
                    return 
                }
                
                res.code = parseInt(res.code)
                if (res.code == 101){
                //    ping success
                }else if (res.code == 0){
                    console.log("websocket get a message:",res)
                    notify.success(res.data)
                    // alert(res.data)
                }else {
                    notify.error(res.msg??"websocket错误")
                }
             } catch (error) {
                console.log("onmessage error:",wsData)
                notify.error(wsData.data)
             }
         }
    }

    //ws异常监听
    websocketOnError(){ 
        console.log("websocket发生异常,3秒后重连========");
    }

    //ws关闭监听
    websocketClose(){
        console.log("websocket关闭========");
    }
}

const socket =new Socket()

export default {
    socket
}