package auth

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"nft_object/library/redis"
	"nft_object/statusCode"
	"time"

	"github.com/gogf/gf/util/gconv"
)

//  sign 能够查询到的信息
type signInfo struct {
	TimeOut   time.Time // 过期时间
	LoginName string    // 登陆人
}

type AdminInfo struct {
	Sign      string    `json:"sign"`       // 当前有效的sign
	LoginName string    `json:"login_name"` // 当前登陆人
	IsAdmin   int       `json:"is_admin"`   // 是否是超级管理员
	Apis      []string  `json:"apis"`       // 拥有的权限列表
	TimeOut   time.Time `json:"time_out"`   // 过期时间
}

func NewAuth() *auth {
	return &auth{
		redis_key:  "admin:login:info:",
		login_code: statusCode.ERROR_NO_LOGIN,
		error_code: statusCode.ERROR,
	}
}

type auth struct {
	redis_key  string
	login_code int
	error_code int
}

// 设置登录信息
func (a *auth) SetLoginInfo(data AdminInfo, duration time.Duration) (string, error) {
	var (
		timeoutAt            = time.Now().Add(duration)
		sign_ori_str         = fmt.Sprintf("%d%s", timeoutAt, data.LoginName)
		sign                 = Md5Encrypt(sign_ori_str)
		sign_redis_key       = a.handleRedisKey(sign)
		admin_info_redis_key = a.handleRedisKey(data.LoginName)
		signData             = signInfo{
			TimeOut:   timeoutAt,
			LoginName: data.LoginName,
		}
	)

	data.Sign = sign
	data.TimeOut = timeoutAt
	// 先设置sign redis
	sign_data_str, err := json.Marshal(signData)
	if err != nil {
		return "", err
	}
	err = redis.SetTimeOut(sign_redis_key, sign_data_str, duration+time.Minute*1)
	if err != nil {
		return "", err
	}
	// 设置用户信息
	data_str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	err = redis.SetTimeOut(admin_info_redis_key, data_str, duration+time.Minute*1)

	return sign, err
}

func (a *auth) handleRedisKey(key string) string {
	return a.redis_key + key
}

// return  int  errorCode
func (a *auth) CheckSign(sign string) (*AdminInfo, int, error) {
	// 校验sign
	res, err := redis.Get(a.handleRedisKey(sign))
	if err != nil {
		return nil, 0, err
	}

	var signData = signInfo{}

	res_byt := gconv.Bytes(res)
	if len(res_byt) == 0 {
		return nil, a.login_code, errors.New("请重新登录")
	}
	if !json.Valid(res_byt) {
		return nil, a.login_code, errors.New("请重新登录:登录无效")
	}
	err = json.Unmarshal(gconv.Bytes(res), &signData)
	if err != nil {
		return nil, a.login_code, err
	}
	//校验过期时间
	if time.Since(signData.TimeOut) > time.Second*0 {
		// 当前时间大于过期时间
		return nil, a.login_code, errors.New("当前登录已过期")
	}

	return a.checkAdminInfo(signData.LoginName, sign)
}

// 获取toke
func (a *auth) checkAdminInfo(loginName, sign string) (*AdminInfo, int, error) {
	res, err := redis.Get(a.handleRedisKey(loginName))
	if err != nil {
		return nil, 0, err
	}

	var adminInfo = AdminInfo{}

	res_byt := gconv.Bytes(res)
	if len(res_byt) == 0 {
		return nil, a.login_code, errors.New("请重新登录")
	}
	if !json.Valid(res_byt) {
		return nil, a.login_code, errors.New("请重新登录:登录无效")
	}
	err = json.Unmarshal(gconv.Bytes(res), &adminInfo)
	if err != nil {
		return nil, a.login_code, err
	}

	if adminInfo.Sign != sign {
		return nil, a.login_code, errors.New("请重新登录:多账号登陆")
	}

	//校验过期时间
	if time.Since(adminInfo.TimeOut) > time.Second*0 {
		// 当前时间大于过期时间
		return nil, a.login_code, errors.New("当前登录已过期")
	}
	return &adminInfo, 0, nil
}

// md5
func Md5Encrypt(data string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}
