package comm

import "github.com/zjutjh/mygo/kit"

var CodeOK = kit.CodeOK

// 内置错误码
var (
	CodeServerError        = kit.CodeServerError
	CodeNotLoggedIn        = kit.CodeNotLoggedIn
	CodeLoginExpired       = kit.CodeLoginExpired
	CodePermissionDenied   = kit.CodePermissionDenied
	CodeParameterInvalid   = kit.CodeParameterInvalid
	CodeDataParseError     = kit.CodeDataParseError
	CodeDataNotFound       = kit.CodeDataNotFound
	CodeDataConflict       = kit.CodeDataConflict
	CodeServiceMaintenance = kit.CodeServiceMaintenance
	CodeTooFrequently      = kit.CodeTooFrequently
)

// 业务错误码 从 2003xx 开始
var ()
