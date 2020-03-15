package builtin

import (
	"github.com/arkgo/ark"
)

var (
	Found  = ark.Result(1, "_found", "不存在", false)
	Error  = ark.Result(2, "_error", "系统错误", false)
	Failed = ark.Result(3, "_error", "参数错误", false)
	Denied = ark.Result(4, "_denied", "拒绝访问", false)

	MappingEmpty = ark.Result(11, "_mapping_empty", "%s不可为空", false)
	MappingError = ark.Result(12, "_mapping_error", "%s不是有效的值", false)

	AuthEmpty = ark.Result(21, "_auth_empty", "%s未登录", false)
	AuthError = ark.Result(22, "_auth_error", "%s无效登录", false)

	ItemEmpty = ark.Result(31, "_item_empty", "%s记录不存在", false)
	ItemError = ark.Result(32, "_item_error", "%s记录不存在", false)
)
