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
func HashSetVar(keys, field string, val interface{}) error {
	_, err := g.Redis().Do("HSet", keys, field, val)
	return err
}

func ListRPush(keys, val interface{}) error {
	_, err := g.Redis().Do("RPUSH", keys, val)
	return err
}

func ListLPush(keys, val interface{}) error {
	_, err := g.Redis().Do("LPUSH", keys, val)
	return err
}

func HashGetVar(keys, field string) (*gvar.Var, error) {
	return g.Redis().DoVar("HGet", keys, field)
}
func HashDel(keys, field string) error {
	_, err := g.Redis().Do("HDEL", keys, field)
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
