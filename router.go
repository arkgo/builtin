package builtin

import (
	"github.com/arkgo/ark"
)

func init() {

	ark.Sites.Register("_doc_", ark.Router{
		Uri:  "/_doc_",
		Name: "系统文档", Desc: "系统文档",
		Action: func(ctx *ark.Http) {
			ctx.Data["cryptos"] = ark.Cryptos()
			ctx.Data["results"] = ark.Results(ctx.Lang())
			ctx.Data["routers"] = ark.Routers(ctx.Site)
			ctx.View("_doc_")
		},
	})

}
