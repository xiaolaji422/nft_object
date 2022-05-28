package grpc

import (
	"context"
	"errors"

	"github.com/xiaolaji422/golink/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewWs(appid ...string) *webSocket {
	var appid_end = "422"
	if len(appid) > 0 {
		appid_end = appid[0]
	}
	return &webSocket{Appid: appid_end}
}

type webSocket struct {
	Appid string
}

func (w *webSocket) Conn() (*grpc.ClientConn, error) {
	return grpc.Dial("127.0.0.1:9558", grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// 给指定人员发消息
func (w *webSocket) SendMsg(ctx context.Context, userid, msg string) error {
	conn, err := w.Conn()
	defer conn.Close()
	if err != nil {
		return err
	}
	req := &pb.SendMsgReq{
		Appid:   w.Appid,
		Userid:  userid,
		Message: "I come from grpc client",
	}

	proxy := pb.NewMessageClient(conn)
	rep, err := proxy.SendMsg(context.Background(), req) // 调用目标地址为前面启动的服务监听的地址
	if err != nil {
		return err
	}
	if rep.Code != 0 {
		return errors.New(rep.Msg)
	}
	return nil
}

// 给应用下所有人发送消息
func (w *webSocket) SendAll(ctx context.Context, userid []string, msg string) (int, error) {
	conn, err := w.Conn()
	defer conn.Close()
	if err != nil {
		return 0, err
	}
	req := &pb.SendMsgAllReq{
		Appid:   w.Appid,
		Userid:  userid,
		Message: msg,
	}

	proxy := pb.NewMessageClient(conn)
	rep, err := proxy.SendAll(context.Background(), req) // 调用目标地址为前面启动的服务监听的地址
	if err != nil {
		return 0, err
	}
	if rep.Code != 0 {
		return 0, errors.New(rep.Msg)
	}
	return int(rep.Count), nil
}
