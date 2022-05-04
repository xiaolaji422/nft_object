package statusCode

// 统一定义系统状态码

const (
	// 成功
	SUCCESS = 0

	// 未知错误
	ERROR = 1000

	// 参数错误
	ERROR_PARAMS = 1003

	// 登录失败
	ERROR_LOGIN = 1004

	// 未登录
	ERROR_NO_LOGIN = 1005

	// 操作失败
	FAILURE = 2000

	// 操作insert db失败
	FAILURE_INSERT = 2001

	// 操作update db失败
	FAILURE_UPDATE = 2002

	// 操作delete db失败
	FAILURE_DELETE = 2003

	// 禁止访问
	FORBIDDEN = 4000

	// 禁止SN访问
	FORBIDDEN_SN = 4001

	// 禁止IP访问
	FORBIDDEN_IP = 4002

	// 禁止访问方法
	FORBIDDEN_METHOD = 4003

	// 无权限访问
	FORBIDDEN_NO_PERMISSION = 4004

	// 访问频繁
	FORBIDDEN_BUSY = 4005
)
