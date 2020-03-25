package zz

import "testing"

func TestVerifyEmailFormat(t *testing.T) {
	a := ""
	aVerify := VerifyEmailFormat(a)
	if aVerify == true {
		t.Errorf("a check %+v", aVerify)
		t.Fail()
		return
	}

	b := "qingfeng.lian@test.com"
	bVerify := VerifyEmailFormat(b)
	if bVerify == false {
		t.Errorf("b check %+v", bVerify)
		t.Fail()
		return
	}

	c := "qingfeng.--lian@test.com"
	cVerify := VerifyEmailFormat(c)
	if cVerify == false {
		t.Errorf("c check %+v", cVerify)
		t.Fail()
		return
	}
}
