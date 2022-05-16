package redis

import (
	"nft_object/library/redis"
	"time"
)

type INoticeRedis interface {

	//  设置人员最新的id
	SetAdminNewerTime(login_name string, t time.Time) error
	// 设置人员最新的id
	GetAdminNewerTime(login_name string) (time.Time, error)
}

var NoticeRedisImpl = func() INoticeRedis {
	return &notice{
		redis_key:            "nft:notice:newer:",
		admin_redis_key:      "admin:",
		admin_read_redis_key: "admin:read:",
		sys_redis_key:        "sys_newer",
	}
}

type notice struct {
	redis_key            string
	admin_redis_key      string
	admin_read_redis_key string
	sys_redis_key        string
}

// 获取当前人最新的公告号
// 1 获取正在预览的
// 2 获取最终的
//
// func (r notice) GetNewerTime(login_name string) (time.Time, error) {
// 	var (
// 		admin_newer  = time.Time{} // 当前用户的最新
// 		is_new_admin = false       //是否是新用户
// 	)

// 	// 获取自己本身的
// 	res, err := redis.GetVar(r.handleRedisKey(r.admin_redis_key + login_name))
// 	if err != nil {
// 		return admin_newer, err
// 	}
// 	if res.IsEmpty() || res.Time().Unix() <= 0 {
// 		is_new_admin = true // 新用户
// 	}

// 	if v := res.Time(); v.Unix() > 0 {
// 		admin_newer = v
// 	}

// 	sys_newer, err := r.GetSysNewerId()
// 	if err != nil {
// 		return 0, err
// 	}

// 	if sys_newer > admin_newer {
// 		admin_newer = sys_newer
// 		if is_new_admin {
// 			admin_newer = sys_newer - 1
// 		}
// 		// 系统值
// 		err = r.SetAdminNewerId(login_name, admin_newer)
// 	}
// 	return admin_newer, err
// }

func (r notice) SetAdminNewerTime(login_name string, t time.Time) error {
	ot, err := r.GetAdminNewerTime(login_name)
	if err != nil {
		return err
	}
	if t.After(ot) {
		return redis.Set(r.handleRedisKey(r.admin_redis_key+login_name), t)
	}
	return err

}

func (r notice) GetAdminNewerTime(login_name string) (time.Time, error) {
	res, err := redis.GetVar(r.handleRedisKey(r.admin_redis_key + login_name))
	if err != nil {
		return time.Time{}, err
	}
	return res.Time(), err
}

// // 设置最新的Id(永久)
// func (r notice) SetSysNewerTime(t time.Time) error {

// 	res, err := redis.GetVar(r.handleRedisKey(r.sys_redis_key))
// 	if err != nil {
// 		return err
// 	}
// 	if v := res.Time(); v.Before(t) {
// 		err = redis.Set(r.handleRedisKey(r.sys_redis_key), t)
// 		return err
// 	}
// 	return err
// }

// func (r notice) GetSysNewerId() (int, error) {
// 	res, err := redis.GetVar(r.handleRedisKey(r.sys_redis_key))
// 	return res.Int(), err
// }

func (r notice) handleRedisKey(key string) string {
	return r.redis_key + key
}
