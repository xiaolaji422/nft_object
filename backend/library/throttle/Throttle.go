package throttle

import (
	"math"
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
)

// Throttle 限流类，参考：https://codeigniter.org.cn/user_guide/libraries/throttler.html
type Throttle struct {
	TokenTime int
	Prefix    string
}

// New new
func New() *Throttle {
	r := new(Throttle)
	r.TokenTime = 0
	r.Prefix = "throttler_"
	return r
}

// Check 检查存储桶中是否还有令牌，或者是否在分配的时间限制内使用了太多令牌。在每次检查期间，如果成功，将根据 $cost 参数来减少可用令牌的数量
// key      储存桶中持有的令牌数量
// capacity 储存桶中持有的令牌数量
// seconds  储存桶完全填满的秒数
// costs    此操作将会花费的令牌数量，默认是1
func (c *Throttle) Check(key string, capacity int, seconds int, costs ...int) bool {

	cost := 1
	if len(costs) >= 1 {
		cost = costs[0]
	}

	tokenName := c.Prefix + key
	retToken := c.Get(tokenName)

	// 如果令牌(token)数不存在，则初始化令牌数
	if retToken == "" {
		c.Set(tokenName, capacity-cost, seconds)
		c.Set(tokenName+"time", int(time.Now().Unix()), seconds)
		return true
	}

	// 当前令牌(token)数
	tokens, err := strconv.Atoi(retToken)
	if err != nil {
		tokens = 0
	}

	throttleTime, err := strconv.Atoi(c.Get(tokenName + "time"))
	if err != nil {
		throttleTime = 0
	}
	elapsed := int(time.Now().Unix()) - throttleTime

	// 每秒要添加回的令牌(token)数
	rate := capacity / seconds

	//多少秒才能获得新令牌。
	//我们必须至少等待1秒才能获取新令牌。
	//主要是为了允许开发者向用户报告而存储的。
	newTokenAvailable := 1 - elapsed
	if rate > 0 {
		newTokenAvailable = (1 / rate) - elapsed
	}
	c.TokenTime = int(math.Max(1, float64(newTokenAvailable)))

	//根据每秒的数字添加令牌
	//应该重新填充，然后检查容量
	//确保存储桶没有溢出。
	tokens += rate * elapsed

	if tokens > capacity {
		tokens = capacity
	}

	// 如果token数据大于0，则减少可用令牌（token)的数量
	if tokens > 0 {
		c.Set(tokenName, tokens-cost, seconds)
		c.Set(tokenName+"time", int(time.Now().Unix()), seconds)
		return true
	}
	return false
}

// GetTokenTime 获取token时间
func (c *Throttle) GetTokenTime() int {
	return c.TokenTime
}

// Get 获取缓存数据
func (c *Throttle) Get(key string) string {

	rediscachecfg := g.Config().GetString("redis.cache")
	if rediscachecfg == "" {
		// 如果没有配置redis，使用gcache
		s, err := gcache.Get(key)
		if err != nil {
			return ""
		}
		return gconv.String(s)
	} else {
		gredis := g.Redis("cache")
		s, err := gredis.Do("GET", key)
		if err != nil {
			return ""
		}
		return gconv.String(s)
	}
	return ""
}

// Set 设置缓存
// key      存储桶key.
// cost     此操作使用的令牌数
// seconds  存储桶完全充满所需的时间
func (c *Throttle) Set(key string, capacity int, seconds int) error {
	rediscachecfg := g.Config().GetString("redis.cache")
	if rediscachecfg == "" {
		// 如果没有配置redis，使用gcache
		gcache.Set(key, capacity, time.Second*time.Duration(seconds))
	} else {
		// 使用redis
		gredis := g.Redis("cache")
		_, err := gredis.Do("SET", key, capacity)
		if err != nil {
			return err
		}
		// 设置key的过期时间
		gredis.Do("EXPIRE", key, seconds)
	}
	return nil
}

/**
在请求hook使用示例（每秒并发访问2次）
if !throttle.New().Check(r.GetClientIp(), 120, 60) {
	r.Response.WriteStatus(http.StatusTooManyRequests, "No Access!!!")
	r.ExitAll()
}
*/
