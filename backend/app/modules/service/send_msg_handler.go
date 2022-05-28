package service

import (
	"context"
	"encoding/json"
	"nft_object/app/repo/grpc"
)

// 业务层代码
// 公告的接口

// 公告业务对外提供的服务
var SendMsgImpl = func() ISendMsg {
	return &send_msg{}
}

type WsMessage struct {
	Code  int         `json:"code"`  // 返回体结构
	Msg   string      `json:"msg"`   // 消息提示
	Data  interface{} `json:"data"`  // 数据
	Scene string      `json:"scene"` // 场景
}

func NewMessage(scene string, data ...interface{}) *WsMessage {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return &WsMessage{
		Code:  0,
		Msg:   "ok",
		Data:  d,
		Scene: scene,
	}
}

type ISendMsg interface {
	//  获取最新公告
	SendMsg(ctx context.Context, userid string, msg *WsMessage) error
	// 历史公告
	SendAll(ctx context.Context, userid []string, msg *WsMessage) (int, error)
}

type send_msg struct{}

func (s *send_msg) SendMsg(ctx context.Context, userid string, msg *WsMessage) error {
	msg_str, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return grpc.NewWs().SendMsg(ctx, userid, string(msg_str))
}

func (s *send_msg) SendAll(ctx context.Context, userid []string, msg *WsMessage) (int, error) {
	msg_str, err := json.Marshal(msg)
	if err != nil {
		return 0, err
	}
	return grpc.NewWs().SendAll(ctx, userid, string(msg_str))
}
