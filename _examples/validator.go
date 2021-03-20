package main

import "github.com/booldesign/validator"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/2 15:23
 * @Desc: 对核心包的调用封装，不同框架不同封装
 */

const (
	CodeRequestParamsInvalid = "request.params.invalid"
)

type Error struct {
	Code    string            `json:"code"`    // 错误代码
	Message string            `json:"message"` // 错误信息
	Fields  map[string]string `json:"fields"`  // 错误字段信息
}

// 参数验证&错误响应
func Validator(params func(string) string, rules []validator.ValidationItem) (map[string]string, *Error) {
	data, field, err := validator.Validation(params, rules)
	if err != nil {
		return nil, &Error{
			Code:    CodeRequestParamsInvalid,
			Message: err.Error(),
			Fields:  map[string]string{field: err.Error()},
		}
	}

	return data, nil
}
