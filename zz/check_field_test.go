package zz

import (
	"testing"
)

func TestCheckStructField(t *testing.T) {
	type param struct {
		Name    string `json:"name" check:"required"`
		Address string `json:"address"`
	}
	a := param{
		Name:    "dafdf",
		Address: "daffdsfasdfsdf",
	}
	err := CheckStructField(a)
	if err != nil {
		t.Errorf("\n最终检测结果: %+v\n", err.Error())
		return
	}
	t.Logf("success")
}
