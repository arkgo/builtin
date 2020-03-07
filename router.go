package builtin

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {

	ark.Sites.Router("_doc_", Map{
		"name": "系统文档", "text": "系统文档",
		"action": func(ctx *ark.Http) {
			ctx.Data["cryptos"] = ark.Cryptos()
			ctx.Data["results"] = ark.Results(ctx.Lang())
			ctx.Data["routers"] = ark.Routers(ctx.Site)
			ctx.View("_doc_")
		},
	})

}
