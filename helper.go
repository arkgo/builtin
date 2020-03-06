package builtin

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {

	ark.Helper("hashid", Map{
		"name": "hashid", "text": "hashid",
		"action": func(val int64) string {
			return ark.Enhash(val)
		},
	}, false)

	ark.Helper("percent", Map{
		"name": "百分比", "text": "百分比",
		"action": func(val float64) string {
			return fmt.Sprintf("%.2f", val*float64(100))
		},
	}, false)

	ark.Helper("round", Map{
		"name": "四舍六入", "text": "四舍六入",
		"action": func(val float64, precisions ...Any) string {
			precision := 2

			if len(precisions) > 0 {
				if vv, ok := precisions[0].(int); ok {
					precision = vv
				} else if vv, ok := precisions[0].(int64); ok {
					precision = int(vv)
				} else if vv, ok := precisions[0].(string); ok {
					if ii, ee := strconv.ParseInt(vv, 10, 64); ee == nil {
						precision = int(ii)
					}
				}
			}

			if precision > 0 {
				format := fmt.Sprintf("%%0.%vf", precision)
				return fmt.Sprintf(format, val)
			}

			return fmt.Sprintf("%f", val)
		},
	}, false)

	ark.Helper("raw", Map{
		"name": "原始输出", "text": "原始输出",
		"action": func(sss Any) template.HTML {
			if sss != nil {
				return template.HTML(fmt.Sprintf("%v", sss))
			}
			return template.HTML("")
		},
	}, false)

	ark.Helper("html", Map{
		"name": "输出HTML", "text": "输出HTML，和raw一个意思",
		"action": func(sss Any) template.HTML {
			if sss != nil {
				return template.HTML(fmt.Sprintf("%v", sss))
			}
			return template.HTML("")
		},
	}, false)

	ark.Helper("attr", Map{
		"name": "HTML属性输出", "text": "HTML属性输出，因为GO模板会自动转义html",
		"action": func(text Any) template.HTMLAttr {
			if text != nil {
				return template.HTMLAttr(fmt.Sprintf("%v", text))
			}
			return template.HTMLAttr("")
		},
	}, false)

	ark.Helper("url", Map{
		"name": "Url输出", "text": "Url输出，因为GO模板会自动转义html",
		"action": func(text Any) template.URL {
			if text != nil {
				return template.URL(fmt.Sprintf("%v", text))
			}
			return template.URL("")
		},
	}, false)

	ark.Helper("join", Map{
		"name": "数组join输出", "text": "数组join输出",
		"action": func(a Any, s string) template.HTML {
			strs := []string{}

			if a != nil {

				switch vv := a.(type) {
				case []string:
					for _, v := range vv {
						strs = append(strs, v)
					}
				case []Any:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int8:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int16:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int32:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []int64:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []float32:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				case []float64:
					for _, v := range vv {
						strs = append(strs, fmt.Sprintf("%v", v))
					}
				}
			}

			html := strings.Join(strs, s)
			return template.HTML(html)
		},
	}, false)

	ark.Helper("json", Map{
		"name": "json输出", "text": "json输出",
		"action": func(data Any) template.HTML {
			if data != nil {
				bytes, err := ark.Marshal(data)
				if err == nil {
					return template.HTML(string(bytes))
				}
			}
			return template.HTML("")
		},
	}, false)

	ark.Helper("mapping", Map{
		"name": "生成Map", "text": "生成Map",
		"action": func(args ...Any) Map {
			m := Map{}

			k := ""
			for i, v := range args {
				if (i+1)%2 == 1 {
					k = fmt.Sprintf("%v", v)
				} else {
					m[k] = v
				}
			}

			return m
		},
	}, false)

	ark.Helper("now", Map{
		"name": "当前时间", "text": "当前时间",
		"action": func() time.Time {
			return time.Now()
		},
	}, false)

	ark.Helper("in", Map{
		"name": "判断某个值是否在数组", "text": "判断某个值是否在数组",
		"action": func(val Any, arrs ...Any) bool {

			strVal := fmt.Sprintf("%v", val)
			strArr := []string{}

			if len(arrs) > 1 {
				for _, vv := range arrs {
					strArr = append(strArr, fmt.Sprintf("%v", vv))
				}
			} else {
				switch vv := arrs[0].(type) {
				case []Any:
					{
						for _, v := range vv {
							strArr = append(strArr, fmt.Sprintf("%v", v))
						}
					}
				case []string:
					for _, v := range vv {
						strArr = append(strArr, v)
					}
				case []int:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int8:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int16:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int32:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				case []int64:
					for _, v := range vv {
						strArr = append(strArr, fmt.Sprintf("%v", v))
					}
				default:
					strArr = append(strArr, fmt.Sprintf("%v", vv))
				}
			}

			for _, v := range strArr {
				if v == strVal {
					return true
				}
			}

			return false
		},
	}, false)

	ark.Helper("out", Map{
		"name": "输出输组某下标元素", "text": "输出输组某下标元素",
		"action": func(arr Any, i int) string {

			strArr := []string{}

			switch vv := arr.(type) {
			case []string:
				for _, v := range vv {
					strArr = append(strArr, v)
				}
			case []int:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int8:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int16:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int32:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			case []int64:
				for _, v := range vv {
					strArr = append(strArr, fmt.Sprintf("%v", v))
				}
			}

			if len(strArr) > i {
				return strArr[i]
			}

			return ""
		},
	}, false)

}
