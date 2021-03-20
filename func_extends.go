package validator

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/2 16:09
 * @Desc: 自定义扩展func
 */

const (
	DefaultLocal   = "Asia/Shanghai"
	DefaultData    = "2006-01-02"
	DefaultErrData = "1000-01-01"
)

var (
	loc, _ = time.LoadLocation(DefaultLocal) //设置时区
)

//自定义函数
type ValidationFuncRule struct {
	Func func(val string) bool
	Msg  string
}

//正则
type ValidationRegexpRule struct {
	Regexp string
	Msg    string
}

// 检查id格式 val => "1,2,3" id必须int
func ValidationIdArrayData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			valList := strings.Split(val, ",")
			for _, v := range valList {
				id, err := strconv.Atoi(v)
				if err != nil || id <= 0 {
					return false
				}
			}
			return true
		},
		"%s 格式不正确",
	}
}

// 检查标识格式 如token
func ValidationTokenArrayData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			valList := strings.Split(val, ",")
			for _, v := range valList {
				if isMatch, _ := regexp.MatchString("^[0-9a-f]{32}$", v); !isMatch {
					return false
				}
			}
			return true
		},
		"%s 格式不正确",
	}
}

// 检查生日格式
func ValidationBirthdayData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			t, err := time.ParseInLocation(DefaultData, val, loc)
			if err != nil {
				t, _ = time.Parse(DefaultData, DefaultErrData)
			}
			if t.Year() > 1905 && t.Year() < time.Now().In(loc).Year() {
				return true
			}
			return false
		},
		"%s 必须介于 1905年 - 至今 之间",
	}
}

// 检查身份证号码
func ValidationIdCardCodeData() ValidationFuncRule {
	return ValidationFuncRule{
		ValifyIdCardCode,
		"%s 格式错误",
	}
}

//验证身份证号码
func ValifyIdCardCode(val string) bool {
	if len(val) != 18 {
		return false
	}

	// 加权因子
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	//校验码范围
	checkCode := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	sum := 0
	for i := 0; i < 17; i++ {
		//第一步：前17位的本体码乘以加权因子：
		v, _ := strconv.Atoi(string(val[i]))
		sum = sum + v*weight[i]
	}
	//第二步：校验码[求和后除以11(校验码的长度)的余数] == 身份证最后一位
	return string(val[17]) == checkCode[sum%11]
}

// 检查开始时间 不能在当前时间之后
func ValidationStartAtData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			intVal, err := strconv.Atoi(val)
			if err == nil && intVal < int(time.Now().Unix()) {
				return true
			}
			return false
		},
		"%s 错误",
	}
}

// 检查手机格式
func ValidationMobileData() ValidationRegexpRule {
	return ValidationRegexpRule{
		`^(1[3-9]\d{9})$`,
		"%s 格式不正确",
	}
}

// 检查邮箱格式
func ValidationEmailData() ValidationRegexpRule {
	return ValidationRegexpRule{
		`^([\dA-Za-z_\.-]+)@([\dA-Za-z\.-]+)\.([A-Za-z\.]+)$`,
		"%s 格式不正确",
	}
}

// 检查身高格式
func ValidationHeightData() []int {
	return []int{0, 300}
}

// 检查体重格式
func ValidationWeightData() []int {
	return []int{0, 500}
}

// 检查用户名格式
func ValidationUsernameData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			flag := true
			isIncludeLetter := false
			if len(val) >= 5 && len(val) <= 25 {
				for _, v := range val {
					if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57) || v == 95) {
						flag = false
						break
					}

					if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
						isIncludeLetter = true
					}
				}
				//不能以下划线开头和结尾
				if val[0] == 95 || val[len(val)-1] == 95 {
					flag = false
				}
			} else {
				flag = false
			}

			if flag && isIncludeLetter == false {
				flag = false
			}

			return flag
		},
		"%s 5~25位数字字母下划线组合，必须包含字母,不能以下划线开后和结尾",
	}
}

// 检查真实姓名格式
func ValidationRealnameData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			flag := true
			l := len([]rune(val))
			if l == 0 || l > 20 {
				flag = false
			} else {
				for _, code := range val {
					//code>128中文
					if !(code >= 128 || (code >= 48 && code <= 57) || (code >= 65 && code <= 90) || (code >= 97 && code <= 122) || code == 46) {
						flag = false
						break
					}
				}
			}
			return flag
		},
		"%s 1~20位中文，英文，字母,.的组合",
	}
}

// 检查编号是否符合mongoDB格式
func ValidationObjectId() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			if val == "" {
				return true
			}
			return CheckMongoIdFormat(val)
		},
		"%s 格式错误",
	}
}

// 验证密码格式
func ValidationPasswordData() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			m := 0
			if len(val) >= 8 && len(val) <= 32 {
				for _, v := range val {
					//[0-9A-Za-z!"#$%&'()*+,-./:;<=>?@ [\]^_`{|}~
					if v >= 48 && v <= 57 {
						m = m | 1
					} else if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
						m = m | 2
					} else if (v >= 33 && v <= 47) || (v >= 58 && v <= 64) || (v >= 91 && v <= 96) || (v >= 123 && v <= 126) {
						m = m | 4
					} else {
						m = 0
						break
					}
				}
				if m != 0 && m != 1 && m != 2 && m != 4 {
					return true
				}
			}
			return false
		},
		"%s 8~32位字母,数字,特殊符号的组合，且包含2种以上组合",
	}
}

// 检查编号是否符合mongoDB格式
func ValidationObjectIds() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			if val == "" {
				return true
			}
			valList := strings.Split(val, ",")
			for _, v := range valList {
				if !CheckMongoIdFormat(v) {
					return false
				}
			}
			return true
		},
		"%s  格式错误",
	}
}

// 检查编号是否符合mongoDB格式
func CheckMongoIdFormat(val string) bool {
	isMatch, _ := regexp.MatchString("^[0-9a-f]{24}$", val)
	return isMatch
}
