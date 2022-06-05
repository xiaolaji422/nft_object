package redis

import (
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
)

// 设置
func SetTimeOut(keys string, data interface{}, timeout time.Duration) error {
	ttl := int(timeout.Seconds())
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.Send("SET", keys, data)
	conn.Send("EXPIRE", keys, ttl)
	err := conn.Flush()
	if err != nil {
		return err
	}
	_, err = conn.Receive()
	if err != nil {
		return err
	}
	_, err = conn.Receive()
	return err
}

func Set(keys string, data interface{}) error {
	_, err := g.Redis().Do("Set", keys, data)
	return err
}

// 获取
func Get(keys string) (interface{}, error) {
	return g.Redis().Do("Get", keys)
}

func GetVar(keys string) (*gvar.Var, error) {
	return g.Redis().DoVar("Get", keys)
}

func DelVar(keys string) (*gvar.Var, error) {
	return g.Redis().DoVar("Del", keys)
}
