package builtin

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/arkgo/ark"
	. "github.com/arkgo/asset"
)

func init() {

	ark.Register("hash", ark.Type{
		Name: "密文", Desc: "密文",
		Valid: func(value Any, config Var) bool {
			if value == nil {
				return false
			}
			switch v := value.(type) {
			case string:
				{
					if v == "" {
						return false
					}
				}
			}
			return true
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case string:
				if ark.Match(v, "hash") {
					return v
				} else {
					return Hash(v)
				}
			}
			return fmt.Sprintf("%v", value)
		},
	}, false)

	ark.Register("any", ark.Type{
		Name: "任意类型", Desc: "任意类型",
		Valid: func(value Any, config Var) bool {
			return true
		},
		Value: func(value Any, config Var) Any {
			return value
		},
	}, false)

	ark.Register("[any]", ark.Type{
		Name: "Anys类型", Desc: "Anys类型",
		Valid: func(value Any, config Var) bool {
			return true
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case Any:
				{
					return []Any{v}
				}
			case []Any:
				{
					return v
				}
			default:
			}
			return []Map{}
		},
	}, false)

	ark.Register("map", ark.Type{
		Name: "MAP类型", Desc: "MAP类型",
		Valid: func(value Any, config Var) bool {
			switch value.(type) {
			case Map:
				{
					return true
				}
			case []Map:
				{
					return true
				}
			default:
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case Map:
				{
					return v
				}
			case []Map:
				{
					if len(v) > 0 {
						return v[0]
					}
				}
			default:
			}
			return Map{}
		},
	}, false)

	ark.Register("[map]", ark.Type{
		Name: "MAPS类型", Desc: "MAPS类型",
		Valid: func(value Any, config Var) bool {
			switch value.(type) {
			case Map:
				{
					return true
				}
			case []Map:
				{
					return true
				}
			default:
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case Map:
				{
					return []Map{v}
				}
			case []Map:
				{
					return v
				}
			default:
			}
			return []Map{}
		},
	}, false)

	//---------- bool begin ----------------------------------
	ark.Register("bool", ark.Type{
		Name: "布尔型", Desc: "布尔型",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case bool:
				{
					return true
				}
			case string:
				{
					if v == "true" || v == "false" || v == "0" || v == "1" || v == "yes" || v == "no" {
						return true
					}
				}
			case int, int8, int16, int32, int64, float32, float64:
				{
					return true
				}
			default:
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case bool:
				{
					return v
				}
			case string:
				{
					if v == "true" || v == "1" || v == "yes" {
						return true
					} else {
						return false
					}
				}
			case int, int8, int16, int32, int64, float32, float64:
				{
					s := fmt.Sprintf("%v", v)
					if s == "0" {
						return false
					} else {
						return true
					}
				}
			default:

			}

			return false
		},
	}, false)

	ark.Register("[bool]", ark.Type{
		Name: "布尔型数组", Desc: "布尔型数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case bool:
				{
					return true
				}
			case []bool:
				{
					return true
				}
			case string:
				{

					if (strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}")) ||
						strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {

						//支持以下几种分隔符
						//"," " " ";"
						sp := ","
						if strings.Index(v, ",") >= 0 {
							sp = ","
						} else if strings.Index(v, " ") >= 0 {
							sp = " "
						} else if strings.Index(v, ";") >= 0 {
							sp = ";"
						} else {
							sp = ","
						}

						arr := strings.Split(v[1:len(v)-1], sp)

						for _, sv := range arr {
							sv = strings.TrimSpace(sv)
							if sv == "" {
								continue
							}
							if sv != "t" && sv != "T" && sv != "true" && sv != "TRUE" && sv != "1" &&
								sv != "f" && sv != "F" && sv != "FALSE" && sv != "false" && sv != "0" {
								return false
							}
						}

						return true

					} else {

						if v == "true" || v == "false" || v == "0" || v == "1" || v == "yes" || v == "no" || v == "t" || v == "f" {
							return true
						}
					}
				}
			case []string:
				{
					for _, s := range v {
						if !(s == "true" || s == "false" || s == "0" || s == "1" || s == "yes" || s == "no" || s == "t" || s == "f") {
							return false
						}
					}
					return true
				}
			default:

			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case bool:
				{
					return []bool{true}
				}
			case []bool:
				{
					return v
				}
			case string:
				{

					//支持postgres数组
					if (strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}")) ||
						(strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]")) {

						//支持以下几种分隔符
						//"," " " ";"
						sp := ","
						if strings.Index(v, ",") >= 0 {
							sp = ","
						} else if strings.Index(v, " ") >= 0 {
							sp = " "
						} else if strings.Index(v, ";") >= 0 {
							sp = ";"
						} else {
							sp = ","
						}

						arr := []bool{}

						strArr := strings.Split(v[1:len(v)-1], sp)
						for _, sv := range strArr {
							sv = strings.TrimSpace(sv)
							if sv == "" {
								continue
							}

							if sv == "t" || sv == "T" || sv == "true" || sv == "TRUE" || sv == "1" {
								arr = append(arr, true)
							} else {
								arr = append(arr, false)
							}
						}

						return arr

					} else {
						if v == "true" || v == "t" || v == "1" || v == "yes" {
							return []bool{true}
						} else {
							return []bool{false}
						}
					}
				}
			case []string:
				{
					vvvvv := []bool{}
					for _, s := range v {
						if s == "true" || s == "1" || s == "yes" {
							vvvvv = append(vvvvv, true)
						} else {
							vvvvv = append(vvvvv, false)
						}
					}
					return vvvvv
				}
			default:

			}

			return false
		},
	}, false)
	//----------- bool end ---------------

	//---------- int begin ----------------------------------

	ark.Register("int", ark.Type{
		Name: "整型", Desc: "整型",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case int, int32, int64, int8:
				{
					return true
				}
			case float32, float64:
				return true
			case string:
				{
					v = strings.TrimSpace(v)
					if _, e := strconv.ParseInt(v, 10, 64); e == nil {
						return true
					}
				}
			default:
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case int:
				{
					return int64(v)
				}
			case int8:
				{
					return int64(v)
				}
			case int16:
				{
					return int64(v)
				}
			case int32:
				{
					return int64(v)
				}
			case int64:
				{
					return int64(v)
				}
			case float32:
				return int64(v)
			case float64:
				return int64(v)
			case string:
				{
					v = strings.TrimSpace(v)
					if i, e := strconv.ParseInt(v, 10, 64); e == nil {
						return i
					}
				}
			default:

			}

			return int64(0)
		},
	}, false)

	ark.Register("[int]", ark.Type{
		Name: "整型数组", Desc: "整型数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case int, int8, int16, int32, int64:
				{
					return true
				}
			case []int, []int8, []int16, []int32, []int64:
				{
					return true
				}
			case float32, float64:
				return true
			case []float32, []float64:
				return true
			case []string:
				{

					if len(v) > 0 {
						for _, sv := range v {
							sv = strings.TrimSpace(sv)
							if _, e := strconv.ParseInt(sv, 10, 64); e != nil {
								return false
							}
						}
						return true
					}
				}
			case []Any:
				if len(v) > 0 {
					for _, av := range v {
						sv := strings.TrimSpace(fmt.Sprintf("%v", av))
						if _, e := strconv.ParseInt(sv, 10, 64); e != nil {
							return false
						}
					}
				}
				return true
			case string:
				{

					//此为postgresql数组返回的数组格式
					//{1,2,3,4,5,6,7,8,9}
					if (strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}")) ||
						strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {

						//支持以下几种分隔符
						//"," " " ";"
						sp := ","
						if strings.Index(v, ",") >= 0 {
							sp = ","
						} else if strings.Index(v, " ") >= 0 {
							sp = " "
						} else if strings.Index(v, ";") >= 0 {
							sp = ";"
						} else {
							sp = ","
						}

						arr := strings.Split(v[1:len(v)-1], sp)

						for _, sv := range arr {
							sv = strings.TrimSpace(sv)
							if sv != "" {
								if _, e := strconv.ParseInt(sv, 10, 64); e != nil {
									return false
								}
							}
						}

						return true
						/*
								//不再使用json解析，因为json解析大数字，18位数时，会有精度问题
							} else if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
								jv := []interface{}{}
								e := ark.Unmarshal([]byte(v), &jv)
								if e == nil {
									return true
								} else {
									return false
								}*/
					} else {

						v = strings.TrimSpace(v)
						if _, e := strconv.ParseInt(v, 10, 64); e == nil {
							return true
						} else {
							return false
						}
					}

				}

			default:
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case int:
				{
					return []int64{int64(v)}
				}
			case int8:
				{
					return []int64{int64(v)}
				}
			case int16:
				{
					return []int64{int64(v)}
				}
			case int32:
				{
					return []int64{int64(v)}
				}
			case int64:
				{
					return []int64{int64(v)}
				}

			case []int:
				{
					arr := []int64{}
					for _, iv := range v {
						arr = append(arr, int64(iv))
					}
					return arr
				}
			case []int8:
				{
					arr := []int64{}
					for _, iv := range v {
						arr = append(arr, int64(iv))
					}
					return arr
				}
			case []int16:
				{
					arr := []int64{}
					for _, iv := range v {
						arr = append(arr, int64(iv))
					}
					return arr
				}
			case []int32:
				{
					arr := []int64{}
					for _, iv := range v {
						arr = append(arr, int64(iv))
					}
					return arr
				}
			case []int64:
				{
					arr := []int64{}
					for _, iv := range v {
						arr = append(arr, int64(iv))
					}
					return arr
				}

			case float32:
				return []int64{int64(v)}
			case float64:
				return []int64{int64(v)}

			case []float32:
				arr := []int64{}
				for _, iv := range v {
					arr = append(arr, int64(iv))
				}
				return arr
			case []float64:
				arr := []int64{}
				for _, iv := range v {
					arr = append(arr, int64(iv))
				}
				return arr

			case []string:
				{
					arr := []int64{}
					for _, sv := range v {
						sv = strings.TrimSpace(sv)
						if iv, e := strconv.ParseInt(sv, 10, 64); e == nil {
							arr = append(arr, iv)
						}
					}

					return arr
				}
			case []Any:
				{
					arr := []int64{}
					for _, av := range v {
						sv := strings.TrimSpace(fmt.Sprintf("%v", av))
						if iv, e := strconv.ParseInt(sv, 10, 64); e == nil {
							arr = append(arr, int64(iv))
						}
					}

					return arr
				}
			case string:
				{

					if (strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}")) ||
						(strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]")) {

						//支持以下几种分隔符
						//"," " " ";"
						sp := ","
						if strings.Index(v, ",") >= 0 {
							sp = ","
						} else if strings.Index(v, " ") >= 0 {
							sp = " "
						} else if strings.Index(v, ";") >= 0 {
							sp = ";"
						} else {
							sp = ","
						}

						arr := []int64{}
						strArr := strings.Split(v[1:len(v)-1], sp)
						for _, sv := range strArr {
							sv = strings.TrimSpace(sv)
							if sv != "" {
								if iv, e := strconv.ParseInt(sv, 10, 64); e == nil {
									arr = append(arr, iv)
								}
							}
						}
						return arr
						/*
								//不再使用json转换，因为json的float在大数字18位长的时候，会有精度问题，
							} else if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
								jv := []interface{}{}
								e := ark.Unmarshal([]byte(v), &jv)

								if e == nil {

									arr := []int64{}
									//所以符合的类型,才写入数组
									//json回转,所有的数都是float64
									for _,anyVal := range jv {
										if newVal,ok := anyVal.(float64); ok {
											arr = append(arr, int64(newVal))
										}
									}

									return arr
								}
						*/
					} else {

						v = strings.TrimSpace(v)
						if vvv, e := strconv.ParseInt(v, 10, 64); e == nil {
							return []int64{vvv}
						}
					}

				}
			default:
			}

			return []int64{}
		},
	}, false)

	//---------- int end ----------------------------------

	//---------- string begin ----------------------------------

	ark.Register("string", ark.Type{
		Name: "字符串", Desc: "字符串",
		Valid: func(value Any, config Var) bool {
			switch v := value.(type) {
			case string:
				if v != "" {
					return true
				}
			case []byte:
				s := fmt.Sprintf("%s", v)
				if s != "" {
					return true
				}
			default:
				if value != nil {
					return true
				}
			}
			return false
		},
		Value: func(value Any, config Var) Any {
			return strings.TrimSpace(fmt.Sprintf("%v", value))
		},
	}, false)

	ark.Register("[string]", ark.Type{
		Name: "字符数组", Desc: "字符数组",
		Valid: func(value Any, config Var) bool {
			switch value.(type) {
			case []string:
				//要不要判断是否为空数组
				return true
			case []Any:
				//要不要判断是否为空数组
				return true
			case string:
				return true
				/*
					left, right := v[0:1], v[len(v)-1:len(v)]
					if left == "[" && right == "]" {
						return true
					} else if left == "{" && right == "}" {
						return true
					} else {
						return true
					}
				*/
			default:
				return false
			}
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case []string:

				//去空字串
				reals := []string{}
				for _, sv := range v {
					if sv != "" {
						reals = append(reals, sv)
					}
				}

				return reals
			case []Any:

				//去空字串
				reals := []string{}
				for _, av := range v {
					sv := fmt.Sprintf("%v", av)
					if sv != "" {
						reals = append(reals, sv)
					}
				}

				return reals
			case string:
				left, right := v[0:1], v[len(v)-1:len(v)]
				if strings.HasPrefix(v, `["`) && strings.HasSuffix(v, `"]`) {
					// ["a","b","c"]
					s := v[2 : len(v)-2] //去掉{""}
					if s != "" {
						return strings.Split(s, `","`)
					}
					return []string{}
				} else if left == "[" && right == "]" {

					s := v[1 : len(v)-1] //去掉[] {}
					if s != "" {
						return strings.Split(s, `,`)
					}
					return []string{}

				} else if strings.HasPrefix(v, `{"`) && strings.HasSuffix(v, `"}`) {
					//cockroach字串数组返回格式 {"a","b","c"}
					s := v[2 : len(v)-2] //去掉{""}
					if s != "" {
						return strings.Split(s, `","`)
					}
					return []string{}
				} else if left == "{" && right == "}" {
					//postgresl字符串
					s := v[1 : len(v)-1] //去掉[] {}
					if s != "" {
						return strings.Split(s, `,`)
					}
					return []string{}

				} else {
					if strings.Contains(v, "\n") {
						return strings.Split(v, "\n")
					} else {
						return []string{v}
					}
				}

			/*
				s := v[1:len(v)-1]	//去掉[] {}
				if s == "" {
					return []string{}
				} else {
					return strings.Split(s, ",")
				}
			*/
			default:
				return v
			}
		},
	}, false)

	ark.Register("[line]", ark.Type{
		Name: "字符分行数组", Desc: "字符分行数组",
		Valid: func(value Any, config Var) bool {
			switch value.(type) {
			case []string:
				return true
			case string:
				return true
			default:
				return false
			}
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case []string:

				//去空字串
				reals := []string{}
				for _, sv := range v {
					sv = strings.TrimSpace(sv)
					if sv != "" {
						reals = append(reals, sv)
					}
				}

				return reals
			case string:

				v = strings.Replace(v, "\r", "", -1)
				arr := strings.Split(v, "\n")

				//去空字串
				reals := []string{}
				for _, sv := range arr {
					sv = strings.TrimSpace(sv)
					if sv != "" {
						reals = append(reals, sv)
					}
				}

				return reals
			default:
				return []string{}
			}
		},
	}, false)
	//---------- string end ----------------------------------

	//---------- datetime begin ----------------------------------

	ark.Register("date", ark.Type{
		Name: "日期时间", Desc: "日期时间",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case time.Time:
				return true
			case *time.Time:
				return true
			case int64:
				return true
			case string:
				return ark.Match(v, "date")
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case int64:
				return time.Unix(v, 0)
			case string:
				lay := "2006-01-02"
				if len(v) == 8 {
					lay = "20060102"
				} else if len(v) == 10 {
					lay = "2006-01-02"
				} else {
					lay = "2006-01-02"
				}

				dt, err := time.Parse(lay, v)
				if err == nil {
					return dt
				} else {
					return v
				}
			}

			return value
		},
	}, false)

	ark.Register("[date]", ark.Type{
		Name: "日期时间数组", Desc: "日期时间数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case []time.Time:
				return true
			case *[]time.Time:
				return true
			case string:
				return ark.Match(v, "date")
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case []time.Time:
				return v
			case *[]time.Time:
				return v
			case string:
				lay := "2006-01-02 15:04:05"
				if len(v) == 8 {
					lay = "20060102"
				} else if len(v) == 10 {
					lay = "2006-01-02"
				} else {
					lay = "2006-01-02 15:04:05"
				}

				dt, err := time.Parse(lay, v)
				if err == nil {
					return []time.Time{dt}
				}
			}

			return value
		},
	}, false)

	ark.Register("datetime", ark.Type{
		Name: "日期时间", Desc: "日期时间",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case time.Time:
				return true
			case *time.Time:
				return true
			case string:
				return ark.Match(v, "datetime")
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case time.Time:
				return v
			case *time.Time:
				return v
			case string:
				lay := "2006-01-02 15:04:05"
				if len(v) == 8 {
					lay = "20060102"
				} else if len(v) == 10 {
					lay = "2006-01-02"
				} else {
					lay = "2006-01-02 15:04:05"
				}

				dt, err := time.ParseInLocation(lay, v, time.Local)
				if err == nil {
					return dt
				} else {
					return v
				}
			}
			return value
		},
	}, false)

	ark.Register("[datetime]", ark.Type{
		Name: "日期时间数组", Desc: "日期时间数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case []time.Time:
				return true
			case *[]time.Time:
				return true
			case string:
				return ark.Match(v, "datetime")
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case []time.Time:
				return v
			case *[]time.Time:
				return v
			default:
				return v
			}
		},
	}, false)

	ark.Register("timestamp", ark.Type{
		Name: "时间戳", Desc: "时间戳",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case time.Time:
				return true
			case string:
				return ark.Match(v, "datetime")
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case time.Time:
				return v.Unix()
			case string:
				dt, err := time.Parse("2006-01-02 15:04:05", v)
				if err == nil {
					return dt.Unix()
				} else {
					return v
				}
			}

			return value
		},
	}, false)

	ark.Register("[timestamp]", ark.Type{
		Name: "时间戳数组", Desc: "时间戳数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case time.Time:
				return true
			case []time.Time:
				return true
			case string:
				return ark.Match(v, "datetime")
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case time.Time:
				return []int64{v.Unix()}
			case []time.Time:
				{
					ts := []int64{}
					for _, dt := range v {
						ts = append(ts, dt.Unix())
					}
					return ts
				}
			case string:
				//应该JSON解析
				dt, err := time.Parse("2006-01-02 15:04:05", v)
				if err == nil {
					return dt.Unix()
				} else {
					return v
				}
			}

			return value
		},
	}, false)

	//---------- datetime end ----------------------------------

	//---------- enum begin ----------------------------------

	ark.Register("enum", ark.Type{
		Name: "枚举", Desc: "枚举",
		Valid: func(value Any, config Var) bool {

			sv := fmt.Sprintf("%v", value)

			if config.Option != nil {
				for k, _ := range config.Option {
					if k == sv {
						return true
					}
				}
			}
			return false
		},
		Value: func(value Any, config Var) Any {
			return fmt.Sprintf("%v", value)
		},
	}, false)

	ark.Register("[enum]", ark.Type{
		Name: "枚举数组", Desc: "枚举数组",
		Valid: func(value Any, config Var) bool {
			vals := []string{}
			switch v := value.(type) {
			case string:
				{
					//如果是 {},  []  包括的字串，就做split
					//postgres中的， {a,b,c} 格式
					if strings.HasPrefix(v, `{"`) && strings.HasSuffix(v, `"}`) {
						//cockroach字串数组返回格式 {"a","b","c"}
						s := v[2 : len(v)-2] //去掉{""}
						vals = strings.Split(s, `","`)
					} else if v[0:1] == "{" && v[len(v)-1:len(v)] == "}" {
						v = v[1 : len(v)-1]
						vals = strings.Split(v, ",")
					} else if v[0:1] == "[" && v[len(v)-1:len(v)] == "]" {
						//json数组格式
						ark.Unmarshal([]byte(v), &vals)
					} else {
						vals = append(vals, v)
					}
				}
			case []string:
				{
					vals = v
				}
			case []Any:
				{
					for _, va := range v {
						vals = append(vals, fmt.Sprintf("%v", va))
					}
				}
			default:
				vals = append(vals, fmt.Sprintf("%v", v))
			}

			//ark.Debug("enums", vals)

			oks := 0

			//遍历数组， 全部在enum里才行
			if config.Option != nil {
				for k, _ := range config.Option {

					for _, v := range vals {
						if k == v {
							oks++
						}
					}
				}
			}

			if oks >= len(vals) {
				return true
			} else {
				return false
			}

		},
		Value: func(value Any, config Var) Any {
			vals := []string{}

			switch v := value.(type) {
			case string:
				{

					//如果是 {},  []  包括的字串，就做split
					//postgres中的， {a,b,c} 格式
					//postgres中的， {a,b,c} 格式
					if strings.HasPrefix(v, `{"`) && strings.HasSuffix(v, `"}`) {
						//cockroach字串数组返回格式 {"a","b","c"}
						s := v[2 : len(v)-2] //去掉{""}
						vals = strings.Split(s, `","`)
					} else if v[0:1] == "{" && v[len(v)-1:len(v)] == "}" {
						v = v[1 : len(v)-1]
						vals = strings.Split(v, ",")
					} else if v[0:1] == "[" && v[len(v)-1:len(v)] == "]" {
						//json数组格式
						ark.Unmarshal([]byte(v), &vals)
					} else {
						vals = append(vals, v)
					}
				}
			case []string:
				{
					vals = v
				}
			case []Any:
				{
					for _, va := range v {
						vals = append(vals, fmt.Sprintf("%v", va))
					}
				}
			default:
				vals = append(vals, fmt.Sprintf("%v", v))
			}
			return vals
		},
	}, false)

	//---------- enum end ----------------------------------

	//---------- file begin ----------------------------------

	ark.Register("file", ark.Type{
		Name: "file", Desc: "file",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch vv := value.(type) {
			case Map:
				return vv
			}
			return value
		},
	}, false)

	ark.Register("[file]", ark.Type{
		Name: "文件数组", Desc: "文件数组",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			case []Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case Map:
				return []Map{v}
			case []Map:
				return v
			}
			return []Map{}
		},
	}, false)

	//---------- file end ----------------------------------

	//---------- float begin ----------------------------------

	ark.Register("float", ark.Type{
		Name: "浮点型", Desc: "布尔型",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case int, int8, int16, int32, int64:
				{
					return true
				}
			case float32, float64:
				{
					return true
				}
			default:
				sv := fmt.Sprintf("%v", v)
				sv = strings.TrimSpace(sv)
				if _, e := strconv.ParseFloat(sv, 64); e == nil {
					return true
				}
			}

			return false
		},
		Value: func(value Any, config Var) Any {
			switch v := value.(type) {
			case int:
				{
					return float64(v)
				}
			case int32:
				{
					return float64(v)
				}
			case int64:
				{
					return float64(v)
				}
			case int8:
				{
					return float64(v)
				}
			case float32:
				{
					return float64(v)
				}
			case float64:
				{
					return v
				}
			case string:
				{
					v = strings.TrimSpace(v)
					if v, e := strconv.ParseFloat(v, 64); e == nil {
						return v
					}
				}
			default:
				sv := fmt.Sprintf("%v", v)
				sv = strings.TrimSpace(sv)
				if v, e := strconv.ParseFloat(sv, 64); e == nil {
					return v
				}
			}

			return float64(0.0)
		},
	}, false)

	ark.Register("[float]", ark.Type{
		Name: "浮点数组", Desc: "浮点数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case int, int8, int16, int32, int64:
				{
					return true
				}
			case []int, []int8, []int16, []int32, []int64:
				{
					return true
				}
			case float32, float64:
				return true
			case []float32, []float64:
				return true
			case []string:
				{
					if len(v) > 0 {
						for _, sv := range v {
							sv = strings.TrimSpace(sv)
							if _, e := strconv.ParseFloat(sv, 64); e != nil {
								return false
							}
						}
						return true
					}
				}
			case []Any:
				{
					if len(v) > 0 {
						for _, av := range v {
							sv := strings.TrimSpace(fmt.Sprintf("%v", av))
							if _, e := strconv.ParseFloat(sv, 64); e != nil {
								return false
							}
						}
						return true
					}
				}
			case string:
				{

					//此为postgresql数组返回的数组格式
					//{1,2,3,4,5,6,7,8,9}
					if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
						s := v[1 : len(v)-1]
						if s == "" {
							return true
						}
						arr := strings.Split(s, ",")

						for _, sv := range arr {
							if _, e := strconv.ParseFloat(sv, 64); e != nil {
								return false
							}
						}
						return true
					} else if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
						jv := []interface{}{}
						e := ark.Unmarshal([]byte(v), &jv)
						if e == nil {
							return true
						} else {
							return false
						}
					} else {

						if _, e := strconv.ParseFloat(v, 64); e == nil {
							return true
						} else {
							return false
						}
					}

				}
			default:
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case int:
				{
					return []float64{float64(v)}
				}
			case int8:
				{
					return []float64{float64(v)}
				}
			case int16:
				{
					return []float64{float64(v)}
				}
			case int32:
				{
					return []float64{float64(v)}
				}
			case int64:
				{
					return []float64{float64(v)}
				}

			case []int:
				{
					arr := []float64{}
					for _, iv := range v {
						arr = append(arr, float64(iv))
					}
					return arr
				}
			case []int8:
				{
					arr := []float64{}
					for _, iv := range v {
						arr = append(arr, float64(iv))
					}
					return arr
				}
			case []int16:
				{
					arr := []float64{}
					for _, iv := range v {
						arr = append(arr, float64(iv))
					}
					return arr
				}
			case []int32:
				{
					arr := []float64{}
					for _, iv := range v {
						arr = append(arr, float64(iv))
					}
					return arr
				}
			case []int64:
				{
					arr := []float64{}
					for _, iv := range v {
						arr = append(arr, float64(iv))
					}
					return arr
				}

			case float32:
				return []float64{float64(v)}
			case float64:
				return []float64{float64(v)}

			case []float32:
				arr := []float64{}
				for _, iv := range v {
					arr = append(arr, float64(iv))
				}
				return arr
			case []float64:
				return v

			case []string:
				{
					arr := []float64{}
					for _, sv := range v {
						sv = strings.TrimSpace(sv)
						if iv, e := strconv.ParseFloat(sv, 64); e == nil {
							arr = append(arr, float64(iv))
						}
					}
					return arr
				}

			case []Any:
				{
					arr := []float64{}
					for _, av := range v {
						sv := strings.TrimSpace(fmt.Sprintf("%v", av))
						if iv, e := strconv.ParseFloat(sv, 64); e == nil {
							arr = append(arr, float64(iv))
						}
					}
					return arr
				}
			case string:
				{

					if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
						arr := []float64{}
						s := v[1 : len(v)-1]
						if s != "" {
							strArr := strings.Split(v[1:len(v)-1], ",")
							for _, sv := range strArr {
								sv = strings.TrimSpace(sv)
								if iv, e := strconv.ParseFloat(sv, 64); e == nil {
									arr = append(arr, iv)
								}
							}
						}
						return arr
					} else if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
						jv := []interface{}{}
						e := ark.Unmarshal([]byte(v), &jv)
						if e == nil {

							arr := []float64{}
							//所以符合的类型,才写入数组
							//json回转,所有的数都是float64
							for _, anyVal := range jv {
								if newVal, ok := anyVal.(float64); ok {
									arr = append(arr, newVal)
								}
							}
							return arr
						}
					} else {

						v = strings.TrimSpace(v)
						if vvv, e := strconv.ParseFloat(v, 64); e == nil {
							return []float64{vvv}
						}
					}

				}
			default:
			}

			return []float64{}
		},
	}, false)

	//---------- float end ----------------------------------

	//---------- image begin ----------------------------------

	ark.Register("image", ark.Type{
		Name: "image", Desc: "image",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch vv := value.(type) {
			case Map:
				return vv
			}
			return Map{}
		},
	}, false)

	ark.Register("[image]", ark.Type{
		Name: "image数组", Desc: "image数组",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			case []Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case Map:
				return []Map{v}
			case []Map:
				return v
			}
			return []Map{}
		},
	}, false)

	ark.Register("audio", ark.Type{
		Name: "audio", Desc: "audio",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch vv := value.(type) {
			case Map:
				return vv
			}
			return Map{}
		},
	}, false)

	ark.Register("[audio]", ark.Type{
		Name: "audio数组", Desc: "audio数组",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			case []Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case Map:
				return []Map{v}
			case []Map:
				return v
			}
			return []Map{}
		},
	}, false)

	ark.Register("video", ark.Type{
		Name: "video", Desc: "video",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch vv := value.(type) {
			case Map:
				return vv
			}
			return Map{}
		},
	}, false)

	ark.Register("[video]", ark.Type{
		Name: "video数组", Desc: "video数组",
		Valid: func(value Any, config Var) bool {

			switch value.(type) {
			case Map:
				return true
			case []Map:
				return true
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case Map:
				return []Map{v}
			case []Map:
				return v
			}
			return []Map{}
		},
	}, false)

	//---------- image end ----------------------------------

	//---------- json begin ----------------------------------

	ark.Register("json", ark.Type{
		Name: "JSON", Desc: "JSON",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case Map:
				return true
			case string:
				m := Map{}
				err := ark.Unmarshal([]byte(v), &m)
				if err == nil {
					return true
				}
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch vv := value.(type) {
			case Map:
				return vv
			case string:
				m := Map{}
				err := ark.Unmarshal([]byte(vv), &m)
				if err == nil {
					return m
				}
			}
			return value
		},
	}, false)

	ark.Register("[json]", ark.Type{
		Name: "JSON数组", Desc: "JSON数组",
		Valid: func(value Any, config Var) bool {

			switch v := value.(type) {
			case Map:
				return true
			case []Map:
				return true
			case []Any: //而是这个
				return true
			case string:
				m := []Map{}
				err := ark.Unmarshal([]byte(v), &m)
				if err == nil {
					return true
				}
			}

			return false
		},
		Value: func(value Any, config Var) Any {

			switch v := value.(type) {
			case Map:
				return []Map{v}
			case []Map:
				return v
			case []Any: //而是这个
				vvvv := []Map{}
				for _, m := range v {
					if mv, ok := m.(Map); ok {
						mm := Map{}
						for kkk, vvv := range mv {
							mm[kkk] = vvv
						}
						vvvv = append(vvvv, mm)
					}
				}
				return vvvv

			case string:
				m := []Map{}
				err := ark.Unmarshal([]byte(v), &m)
				if err == nil {
					return m
				}
			}
			return []Map{}
		},
	}, false)

}
