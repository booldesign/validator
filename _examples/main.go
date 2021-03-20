package main

import (
	"fmt"
	"github.com/booldesign/validator"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/2 15:47
 * @Desc:
 */

const ListMaxCount = 100

func main() {
	var FuncName = func(key string) string {
		params := make(map[string]string)
		params["orgId"] = "1"
		params["status"] = "DELETED"
		params["ids"] = "1,3"
		params["keywords"] = "booldesign"
		params["pageNum"] = "1"
		params["pageSize"] = "10"
		params["username"] = "gegeg122"
		params["password"] = "$$$$$$$a"
		params["idCardCode"] = "110103200301013718"
		params["email"] = "booldesign@163.com"
		params["mobile"] = "13501691436"
		params["isSync"] = "true"
		params["userIds"] = "1,2,4,q,5"
		params["cids"] = "1,2,4,2,4,5"
		params["birthday"] = "2010-01-01"
		return params[key]
	}

	data, err := Validator(FuncName, []validator.ValidationItem{
		{
			Key: "orgId", Name: "组织机构id",
			Rules: []validator.ValidationRule{
				{Rule: "required"},
				{Rule: "integer"},
				{Rule: "min", Data: 1},
			},
		}, {
			Key: "status", Name: "状态",
			Rules: []validator.ValidationRule{
				{Rule: "in", Data: []string{"DELETED", "ENABLED", "DISABLED"}},
			},
		}, {
			Key: "ids", Name: "显示列表编号",
			Rules: []validator.ValidationRule{
				{Rule: "arrayInArray", Data: []interface{}{",", []int{1, 2, 3}}},
			},
		}, {
			Key: "keywords", Name: "关键词",
			Rules: []validator.ValidationRule{
				{Rule: "required"},
				{Rule: "filterChar", Data: []string{"%", "_"}},
			},
		}, {
			Key: "sort", Name: "排序",
			Rules: []validator.ValidationRule{
				{Rule: "arrayInArray", Data: []interface{}{",", []string{"id", "username", "realName"}}},
			},
		}, {
			Key: "pageNum", Name: "页号",
			Rules: []validator.ValidationRule{
				{Rule: "min", Data: 1},
				{Rule: "max", Data: 100},
			},
		}, {
			Key: "pageSize", Name: "每页记录条数",
			Rules: []validator.ValidationRule{
				{Rule: "between", Data: []int{-1, ListMaxCount}},
			},
		}, {
			Key: "username", Name: "用户名",
			Rules: []validator.ValidationRule{
				{Rule: "required"},
				{Rule: "func", Data: validator.ValidationUsernameData()},
			},
		}, {
			Key: "password", Name: "密码",
			Rules: []validator.ValidationRule{
				{Rule: "required"},
				{Rule: "func", Data: validator.ValidationPasswordData()},
			},
		}, {
			Key: "idCardCode", Name: "身份证号码",
			Rules: []validator.ValidationRule{
				{Rule: "func", Data: validator.ValidationIdCardCodeData()},
			},
		}, {
			Key: "email", Name: "邮箱",
			Rules: []validator.ValidationRule{
				{Rule: "regexp", Data: validator.ValidationEmailData()},
			},
		}, {
			Key: "mobile", Name: "手机号",
			Rules: []validator.ValidationRule{
				{Rule: "required"},
				{Rule: "regexp", Data: validator.ValidationMobileData()},
			},
		}, {
			Key: "isSync", Name: "是否同步",
			Rules: []validator.ValidationRule{
				{Rule: "required"},
				{Rule: "bool"},
			},
		}, {
			Key: "userIds", Name: "用户ids",
			Rules: []validator.ValidationRule{
				{Rule: "distinct", Data: ","},
			},
		}, {
			Key: "cids", Name: "ids",
			Rules: []validator.ValidationRule{
				{Rule: "func", Data: validator.ValidationIdArrayData()},
			},
		}, {
			Key: "birthday", Name: "生日",
			Rules: []validator.ValidationRule{
				{Rule: "func", Data: validator.ValidationBirthdayData()},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
