package builtin

import (
	"encoding/base64"

	"github.com/arkgo/ark"
	. "github.com/arkgo/asset"
	"github.com/arkgo/asset/util"
)

func init() {

	ark.Register("percent", ark.Crypto{
		Name: "百分比处理", Desc: "百分比处理",
		Encode: func(value Any, config Var) Any {
			//text -> text
			if vv, ok := value.(float64); ok {
				return ark.Precision(vv/100, 4)
			}
			return nil
		},
		Decode: func(value Any, config Var) Any {
			//data -> text
			if vv, ok := value.(float64); ok {
				return ark.Precision(vv*100, 4)
			}
			return nil
		},
	}, false)

	ark.Register("text", ark.Crypto{
		Name: "文本加密", Desc: "文本加密，自定义字符表的base64编码，字典：" + ark.TextAlphabet(),
		Alias: []string{"text", "string"},
		Encode: func(value Any, config Var) Any {
			text := util.ToString(value)
			return ark.Encrypt(text)
		},
		Decode: func(value Any, config Var) Any {
			text := util.ToString(value)
			return ark.Decrypt(text)
		},
	})
	ark.Register("texts", ark.Crypto{
		Name: "文本数组加密", Desc: "文本数组加密，自定义字符表的base64编码，字典：" + ark.TextAlphabet(),
		Alias: []string{"texts", "strings"},
		Encode: func(value Any, config Var) Any {
			if vv, ok := value.([]string); ok {
				return ark.Encrypts(vv)
			}
			return nil
		},
		Decode: func(value Any, config Var) Any {
			text := util.ToString(value)
			return ark.Decrypts(text)
		},
	})

	ark.Register("hash", ark.Crypto{
		Name: "数字加密", Desc: "数字加密",
		Alias: []string{"hash", "number", "digit"},
		Encode: func(value Any, config Var) Any {
			if vv, ok := value.(int64); ok {
				return ark.Enhash(vv)
			}
			return nil
		},
		Decode: func(value Any, config Var) Any {
			if vv, ok := value.(string); ok {
				return ark.Dehash(vv)
			}
			return nil
		},
	})

	ark.Register("hashs", ark.Crypto{
		Name: "数字数组加密", Desc: "数字数组加密",
		Alias: []string{"hashs", "numbers", "digits"},
		Encode: func(value Any, config Var) Any {
			if vv, ok := value.([]int64); ok {
				return ark.Enhashs(vv)
			}
			return nil
		},
		Decode: func(value Any, config Var) Any {
			if vv, ok := value.(string); ok {
				return ark.Dehashs(vv)
			}
			return nil
		},
	})

	ark.Register("base64", ark.Crypto{
		Name: "BASE64加解密", Desc: "BASE64加解密",
		Encode: func(value Any, config Var) Any {
			text := util.ToString(value)
			return base64.StdEncoding.EncodeToString([]byte(text))
		},
		Decode: func(value Any, config Var) Any {
			text := util.ToString(value)
			bytes, err := base64.StdEncoding.DecodeString(text)
			if err == nil {
				return string(bytes)
			}
			return nil
		},
	}, false)

}
