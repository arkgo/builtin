package builtin

import (
	"github.com/arkgo/ark"
)

var (
	Found  = ark.Result(1, ".found", "不存在", false)
	Error  = ark.Result(2, ".error", "系统错误", false)
	Failed = ark.Result(3, ".error", "参数错误", false)
	Denied = ark.Result(4, ".denied", "拒绝访问", false)

	MappingEmpty = ark.Result(11, ".mapping.empty", "%s不可为空", false)
	MappingError = ark.Result(12, ".mapping.error", "%s不是有效的值", false)

	AuthEmpty = ark.Result(21, ".auth.empty", "%s未登录", false)
	AuthError = ark.Result(22, ".auth.error", "%s无效登录", false)

	ItemEmpty = ark.Result(31, ".item.empty", "%s记录不存在", false)
	ItemError = ark.Result(32, ".item.error", "%s记录不存在", false)
)
