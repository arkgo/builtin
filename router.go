package builtin

import (
	"fmt"
	"time"

	"github.com/arkgo/ark"
	. "github.com/arkgo/asset"
	"github.com/arkgo/asset/util"
)

func init() {

	ark.Sites.Register("_doc_", ark.Router{
		Uri: "/_doc_", Name: "系统文档", Desc: "系统文档",
		Args: Vars{
			"p": ark.Define("string", false, "path"),
		},
		Setting: Map{"passport": true},
		Action: func(ctx *ark.Http) {
			ctx.Data["cryptos"] = ark.Cryptos()
			ctx.Data["results"] = ark.Results(ctx.Lang())
			ctx.Data["routers"] = ark.Routers(ctx.Site)

			path := "/"
			if vv, ok := ctx.Args["p"].(string); ok && vv != "" {
				path = vv
			}

			client := fmt.Sprintf("1234567890/ios/12.1.0/ios.test.com/1/%d/", time.Now().Unix())
			sign := util.Md5(client + path)
			client = client + sign

			ctx.Data["client"] = Map{
				"client":  client,
				"example": ark.Encrypt(client),
				"sign":    sign, "path": path,
			}

			ctx.View("_doc_")
		},
	})

}
