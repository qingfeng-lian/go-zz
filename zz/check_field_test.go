package zz

import (
	"testing"
)

func TestCheckStructField(t *testing.T) {
	type param struct {
		Name    string `json:"name" check:"required"`
		Address string `json:"address" check:"required"`
	}
	a := param{
		Name:    "dafdf",
		Address: "",
	}
	err := CheckStructField(a)
	if err == nil {
		t.Errorf("\n最终检测结果应该有错误的\n")
		return
	}

	//指针
	b := &param{
		Name:    "aaaa",
		Address: "",
	}
	err = CheckStructField(b)
	if err != nil {
		t.Errorf("\n最终检测结果应该有错误的\n")
		return
	}

	t.Logf("success")
}
