package throttle

// ThrottlerInterface interface
type ThrottlerInterface interface {
	// Check 检查存储桶中是否还有令牌，或者是否在分配的时间限制内使用了太多令牌。在每次检查期间，
	// 如果成功，将根据 $cost 参数来减少可用令牌的数量
	// key      储存桶的key
	// capacity 储存桶中持有的令牌数量
	// seconds  储存桶完全填满的秒数
	// costs    此操作将会花费的令牌数量，默认是1
	Check(key string, capacity, seconds, cost int) bool
	GetTokenTime() int
}
