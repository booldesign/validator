package validator

import (
	"regexp"
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/3/19 19:10
 * @Desc:
 */

func TestValidationIdArrayData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"1,3,4,5", true},
		{"1,a,b,3", false},
		{"1,-,4,4,5", false},
	}

	for _, test := range tests {
		vFunc := ValidationIdArrayData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Log("ValidationIdArrayData() failed.")
		}
	}
}

func TestValidationBirthdayData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"2011-12-31", true},
		{"2030-01-01", false},
		{"1900-01-01", false},
	}

	for _, test := range tests {
		vFunc := ValidationBirthdayData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Log("ValidationBirthdayData() failed.")
		}
	}
}

func TestValidationIdCardCodeData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"110103200301013718", true},
		{"110103200301013719", false},
	}

	for _, test := range tests {
		vFunc := ValidationIdCardCodeData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Log("ValidationIdCardCodeData() failed.")
		}
	}
}

func TestValidationStartAtData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"1616152846", true},
		{"1893427200", false},
	}

	for _, test := range tests {
		vFunc := ValidationStartAtData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Log("ValidationStartAtData() failed.")
		}
	}
}

func TestValidationMobileData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"13501691436", true},
		{"12909090909", false},
	}

	for _, test := range tests {
		vFunc := ValidationMobileData()
		ok, _ := regexp.MatchString(vFunc.Regexp, test.in)
		if ok != test.expect {
			t.Log("ValidationMobileData() failed.")
		}
	}
}

func TestValidationEmailData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"booldesign@163.com", true},
		{"booldesign163.com", false},
	}

	for _, test := range tests {
		vFunc := ValidationEmailData()
		ok, _ := regexp.MatchString(vFunc.Regexp, test.in)
		if ok != test.expect {
			t.Log("ValidationEmailData() failed.")
		}
	}
}

func TestValidationUsernameData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"feg12_4", true},
		{"244jjijiji", true},
		{"244jjijttetq153535tny4yn4y4y4ijia", false},
		{"1", false},
		{"114155", false},
		{"_gegg124", false},
		{"我hi hi1515", false},
		{"123_", false},
		{"123w1_", false},
		{"123w1_", false},
	}

	for _, test := range tests {
		vFunc := ValidationUsernameData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Errorf("ValidationUsernameData() failed."+vFunc.Msg, test.in)
		}
	}
}

func TestValidationRealnameData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"wei_jian", false},
		{"weijianwen2", true},
		{"wei.jianwen1", true},
		{"卫建文", true},
		{"卫建文_", false},
		{"_卫建文", false},
		{"wei jianwen", false},
	}

	for _, test := range tests {
		vFunc := ValidationRealnameData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Errorf("ValidationRealnameData() failed. "+vFunc.Msg, test.in)
		}
	}
}

func TestValidationPasswordData(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"1", false},
		{"weijianwen2", true},
		{"wei.jianwen1", true},
		{"卫建文", true},
		{"卫建文_", false},
		{"_卫建文", false},
		{"wei jianwen", false},
	}

	for _, test := range tests {
		vFunc := ValidationRealnameData()
		if ok := vFunc.Func(test.in); ok != test.expect {
			t.Errorf("ValidationRealnameData() failed. "+vFunc.Msg, test.in)
		}
	}
}
