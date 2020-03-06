package builtin

import (
	"encoding/base64"

	"github.com/arkgo/ark"
	"github.com/arkgo/asset/util"
	. "github.com/arkgo/base"
)

func init() {

	ark.Crypto("percent", Map{
		"name": "百分比处理", "text": "百分比处理",
		"encode": func(value Any) Any {
			//text -> text
			if vv, ok := value.(float64); ok {
				return ark.Precision(vv/100, 4)
			}
			return nil
		},
		"decode": func(value Any) Any {
			//data -> text
			if vv, ok := value.(float64); ok {
				return ark.Precision(vv*100, 4)
			}
			return nil
		},
	}, false)

	ark.Crypto("text", Map{
		"cryptos": []string{"text", "string"},
		"name":    "文本加密", "text": "文本加密，自定义字符表的base64编码，字典：" + ark.TextAlphabet(),
		"encode": func(value Any) Any {
			text := util.ToString(value)
			return ark.Encrypt(text)
		},
		"decode": func(value Any) Any {
			text := util.ToString(value)
			return ark.Decrypt(text)
		},
	})
	ark.Crypto("texts", Map{
		"cryptos": []string{"texts", "strings"},
		"name":    "文本数组加密", "text": "文本数组加密，自定义字符表的base64编码，字典：" + ark.TextAlphabet(),
		"encode": func(value Any) Any {
			if vv, ok := value.([]string); ok {
				return ark.Encrypts(vv)
			}
			return nil
		},
		"decode": func(value Any) Any {
			text := util.ToString(value)
			return ark.Decrypts(text)
		},
	})

	ark.Crypto("hash", Map{
		"cryptos": []string{"hash", "number", "digit"},
		"name":    "数字加密", "text": "数字加密",
		"encode": func(value Any) Any {
			if vv, ok := value.(int64); ok {
				return ark.Enhash(vv)
			}
			return nil
		},
		"decode": func(value Any) Any {
			if vv, ok := value.(string); ok {
				return ark.Dehash(vv)
			}
			return nil
		},
	})

	ark.Crypto("hashs", Map{
		"cryptos": []string{"hashs", "numbers", "digits"},
		"name":    "数字数组加密", "text": "数字数组加密",
		"encode": func(value Any) Any {
			if vv, ok := value.([]int64); ok {
				return ark.Enhashs(vv)
			}
			return nil
		},
		"decode": func(value Any) Any {
			if vv, ok := value.(string); ok {
				return ark.Dehashs(vv)
			}
			return nil
		},
	})

	ark.Crypto("base64", Map{
		"name": "BASE64加解密", "text": "BASE64加解密",
		"encode": func(value Any) Any {
			text := util.ToString(value)
			return base64.StdEncoding.EncodeToString([]byte(text))
		},
		"decode": func(value Any) Any {
			text := util.ToString(value)
			bytes, err := base64.StdEncoding.DecodeString(text)
			if err == nil {
				return string(bytes)
			}
			return nil
		},
	}, false)

}
