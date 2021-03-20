package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/2 15:23
 * @Desc: validation核心逻辑
 */

const (
	ValidateMethodNotAllowSth   = "验证方法 %s 不允许 %s"
	ValidateValCanNotEmpty      = "%s 不能为空"
	ValidateValNotExists        = "%s 不存在"
	ValidateValMustBool         = "%s 必须为 true 或 false"
	ValidateValMustInteger      = "%s 必须是整数"
	ValidateValNotBetweenInt    = "%s 必须是 %d - %d 之间的整数"
	ValidateValNotBetweenFloat  = "%s 必须是 %f - %f 之间数字"
	ValidateValNotBetweenStr    = "%s 长度必须在 %d - %d 之间"
	ValidateValNotMinInt        = "%s 必须是大等于 %d 的整数"
	ValidateValNotMinFloat      = "%s 必须是大等于 %f 的数字"
	ValidateValNotMinStr        = "%s 长度不能小于 %d"
	ValidateValNotMaxInt        = "%s 必须是小等于 %d 的整数"
	ValidateValNotMaxFloat      = "%s 必须是小等于 %f 的数字"
	ValidateValNotMaxStr        = "%s 长度不能大于 %d"
	ValidateValArrayNotInArray  = "%s 不能含有 %v 以外的值"
	ValidateValExistsFilterChar = "%s 不允许包含 %v"
	ValidateValMustDistinct     = "%s 含有重复的值 [%s]"
)

// 验证规则，多个验证规则组合成一个验证项
type ValidationRule struct {
	Rule string      // 规则名称
	Data interface{} // 扩展数据
}

// 验证项，用于验证某个参数
type ValidationItem struct {
	Key   string           // 参数键
	Name  string           // 参数名称
	Rules []ValidationRule // 规则
}

// 参数验证
func Validation(params func(string) string, rules []ValidationItem) (map[string]string, string, error) {
	var err error
	data := map[string]string{}

	for _, v := range rules {
		val := params(v.Key)
		for vIk, vI := range v.Rules {
			switch vI.Rule {
			case "required":
				err = ValidationRequired(&v, vIk, val)
			case "in":
				err = ValidationIn(&v, vIk, val)
			case "bool":
				err = ValidationBool(&v, vIk, val)
			case "integer":
				err = ValidationInteger(&v, vIk, val)
			case "between":
				err = ValidationBetween(&v, vIk, val)
			case "min":
				err = ValidationMin(&v, vIk, val)
			case "max":
				err = ValidationMax(&v, vIk, val)
			case "arrayInArray":
				err = ValidationArrayInArray(&v, vIk, val)
			case "filterChar":
				err = ValidationFilterChar(&v, vIk, val)
			case "regexp":
				err = ValidationRegexp(&v, vIk, val)
			case "func":
				err = ValidationFunc(&v, vIk, val)
			case "distinct":
				err = ValidationDistinct(&v, vIk, val)
			}
			if err != nil {
				return nil, v.Key, err
			}
		}
		data[v.Key] = val
	}

	return data, "", nil
}

// 是否为空或未提交
func ValidationRequired(rule *ValidationItem, _ int, val string) error {
	if val == "" {
		return fmt.Errorf(ValidateValCanNotEmpty, rule.Name)
	}
	return nil
}

// 提交的数据是否包含在允许的数组内
func ValidationIn(rule *ValidationItem, index int, val string) error {
	if val != "" {
		for _, v := range rule.Rules[index].Data.([]string) {
			if v == val {
				return nil
			}
		}
		return fmt.Errorf(ValidateValNotExists, rule.Name)
	}
	return nil
}

// 是否为布尔类型
func ValidationBool(rule *ValidationItem, _ int, val string) error {
	if val != "" {
		_, err := strconv.ParseBool(val)
		if err != nil {
			return fmt.Errorf(ValidateValMustBool, rule.Name)
		}
	}
	return nil
}

// 是否为整数类型
func ValidationInteger(rule *ValidationItem, _ int, val string) error {
	if val != "" {
		_, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf(ValidateValMustInteger, rule.Name)
		}
	}
	return nil
}

// 字符串长度或数值是否在范围内
func ValidationBetween(rule *ValidationItem, index int, val string) error {
	if val != "" {
		switch reflect.TypeOf(rule.Rules[index].Data).String() {
		case "[]int":
			size := rule.Rules[index].Data.([]int)
			valInt, err := strconv.Atoi(val)
			if err == nil && size[0] <= valInt && valInt <= size[1] {
				return nil
			}
			return fmt.Errorf(ValidateValNotBetweenInt, rule.Name, size[0], size[1])
		case "[]float64":
			size := rule.Rules[index].Data.([]float64)
			valFloat, err := strconv.ParseFloat(val, 64)
			if err == nil && size[0] <= valFloat && valFloat <= size[1] {
				return nil
			}
			return fmt.Errorf(ValidateValNotBetweenFloat, rule.Name, size[0], size[1])
		case "[]string":
			sizeStr := rule.Rules[index].Data.([]string)
			size := [2]int{}
			size[0], _ = strconv.Atoi(sizeStr[0])
			size[1], _ = strconv.Atoi(sizeStr[1])
			if size[0] <= utf8.RuneCount([]byte(val)) && utf8.RuneCount([]byte(val)) <= size[1] {
				return nil
			}
			return fmt.Errorf(ValidateValNotBetweenStr, rule.Name, size[0], size[1])
		}
		return fmt.Errorf(ValidateMethodNotAllowSth, "ValidationBetween",
			reflect.TypeOf(rule.Rules[index].Data).String())
	}
	return nil
}

// 字符串长度或数值是否小于最小值
func ValidationMin(rule *ValidationItem, index int, val string) error {
	if val != "" {
		switch reflect.TypeOf(rule.Rules[index].Data).String() {
		case "int":
			size := rule.Rules[index].Data.(int)
			valInt, err := strconv.Atoi(val)
			if err == nil && size <= valInt {
				return nil
			}
			return fmt.Errorf(ValidateValNotMinInt, rule.Name, size)
		case "float64":
			size := rule.Rules[index].Data.(float64)
			valFloat, err := strconv.ParseFloat(val, 64)
			if err == nil && size <= valFloat {
				return nil
			}
			return fmt.Errorf(ValidateValNotMinFloat, rule.Name, size)
		case "string":
			sizeStr := rule.Rules[index].Data.(string)
			size, _ := strconv.Atoi(sizeStr)
			if size <= utf8.RuneCount([]byte(val)) {
				return nil
			}
			return fmt.Errorf(ValidateValNotMinStr, rule.Name, size)
		}
		return fmt.Errorf(ValidateMethodNotAllowSth, "ValidationBetween",
			reflect.TypeOf(rule.Rules[index].Data).String())
	}
	return nil
}

// 字符串长度或数值是否超过最大值
func ValidationMax(rule *ValidationItem, index int, val string) error {
	if val != "" {
		switch reflect.TypeOf(rule.Rules[index].Data).String() {
		case "int":
			size := rule.Rules[index].Data.(int)
			valInt, err := strconv.Atoi(val)
			if err == nil && size >= valInt {
				return nil
			}
			return fmt.Errorf(ValidateValNotMaxInt, rule.Name, size)
		case "float64":
			size := rule.Rules[index].Data.(float64)
			valFloat, err := strconv.ParseFloat(val, 64)
			if err == nil && size >= valFloat {
				return nil
			}
			return fmt.Errorf(ValidateValNotMaxFloat, rule.Name, size)
		case "string":
			sizeStr := rule.Rules[index].Data.(string)
			size, _ := strconv.Atoi(sizeStr)
			if size >= utf8.RuneCount([]byte(val)) {
				return nil
			}
			return fmt.Errorf(ValidateValNotMaxStr, rule.Name, size)
		}
		return fmt.Errorf(ValidateMethodNotAllowSth, "ValidationBetween",
			reflect.TypeOf(rule.Rules[index].Data).String())
	}
	return nil
}

// 提交的数组是否包含在允许的数组内
func ValidationArrayInArray(rule *ValidationItem, index int, val string) error {
	if val != "" {
		valList := strings.Split(val, rule.Rules[index].Data.([]interface{})[0].(string))
		list := rule.Rules[index].Data.([]interface{})[1]
		switch reflect.TypeOf(list).String() {
		case "[]string":
			listStr := list.([]string)
			for _, v := range valList {
				exists := false
				for _, lv := range listStr {
					if v == lv {
						exists = true
					}
				}
				if !exists {
					return fmt.Errorf(ValidateValArrayNotInArray, rule.Name, listStr)
				}
			}
			return nil
		case "[]int":
			listInt := list.([]int)
			for _, v := range valList {
				exists := false
				vInt, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf(ValidateValArrayNotInArray, rule.Name, listInt)
				}
				for _, lv := range listInt {
					if vInt == lv {
						exists = true
					}
				}
				if !exists {
					return fmt.Errorf(ValidateValArrayNotInArray, rule.Name, listInt)
				}
			}
			return nil
		}
		return fmt.Errorf(ValidateMethodNotAllowSth, "ValidationArrayInArray",
			reflect.TypeOf(list).String())
	}
	return nil
}

// 是否包含不允许出现的字符
func ValidationFilterChar(rule *ValidationItem, index int, val string) error {
	if val != "" {
		chars := rule.Rules[index].Data.([]string)
		for _, v := range chars {
			if strings.Index(val, v) >= 0 {
				return fmt.Errorf(ValidateValExistsFilterChar, rule.Name, chars)
			}
		}
	}
	return nil
}

// 是否匹配正则表达式
func ValidationRegexp(rule *ValidationItem, index int, val string) error {
	if val != "" {
		data := rule.Rules[index].Data.(ValidationRegexpRule)
		m, _ := regexp.MatchString(data.Regexp, val)
		if !m {
			return fmt.Errorf(data.Msg, rule.Name)
		}
	}
	return nil
}

// 自定义验证方法
func ValidationFunc(rule *ValidationItem, index int, val string) error {
	if val != "" {
		data := rule.Rules[index].Data.(ValidationFuncRule)
		if !data.Func(val) {
			return fmt.Errorf(data.Msg, rule.Name)
		}
	}
	return nil
}

// 是否有重复值
func ValidationDistinct(rule *ValidationItem, index int, val string) error {
	if val != "" {
		valList := strings.Split(val, rule.Rules[index].Data.(string))
		for i := 0; i < len(valList); i++ {
			for j := i + 1; j < len(valList); j++ {
				if valList[i] == valList[j] {
					return fmt.Errorf(ValidateValMustDistinct, rule.Name, valList[i])
				}
			}
		}
	}
	return nil
}
