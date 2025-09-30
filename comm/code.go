package comm

import "github.com/zjutjh/jhgo/kit"

const CodeOK kit.Code = 0 // CodeOK 成功

// 系统错误码
const (
	CodeUnknowkitor            kit.Code = 10000 // 未知错误
	CodeThirdServiceError      kit.Code = 10001 // 三方服务错误
	CodeDatabaseError          kit.Code = 10002 // 数据库错误
	CodeRedisError             kit.Code = 10003 // Redis错误
	CodeMiddlewareServiceError kit.Code = 10004 // 中间件服务错误
)

// 业务通用错误码
const (
	CodeNotLogggedIn       kit.Code = 20000 // 用户未登录
	CodeLoginExpired       kit.Code = 20001 // 登录过期，请重新登录
	CodePermissionDenied   kit.Code = 20002 // 用户无权限
	CodeParamterInvalid    kit.Code = 20003 // 参数非法
	CodeDataParseError     kit.Code = 20004 // 数据解析异常
	CodeDataNotFound       kit.Code = 20005 // 数据不存在
	CodeDataConflict       kit.Code = 20006 // 数据冲突
	CodeServiceMaintenance kit.Code = 20007 // 系统维护中
	CodeTooFequently       kit.Code = 20008 // 操作过于频繁/未获得锁
)

// 业务错误码 从 30000 开始
const ()
