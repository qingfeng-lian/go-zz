package zz

import (
	"testing"
)

func TestCheckStructField(t *testing.T) {
	type param struct {
		Name    string `json:"name" check:"required;maxLen:2000;minLen:7"`
		Address string `json:"address" check:"required;email"`
		Mobile  string `json:"mobile" check:"required;mobile"`
	}

	a := param{
		Name:    "我已超a出度",
		Address: "qingfeng.lian@163.com",
		Mobile:  "18000000000",
	}
	err := CheckStructField(a)
	if err == nil {
		t.Errorf("\n最终检测结果应该有错误的\n")
		return
	}
	t.Logf("这里检查结果应该是err ,err:%+v", err)

	t.Logf("success")
}
