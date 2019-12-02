package zz

import (
	"encoding/base64"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	t.Skip() //这个方法由于解密原始值换成了 文字描述，所以该测试方法跳过
	d := WxEncryptedData{}
	orgData := "这里请换上需要解密的字符串， 此字符串是经过base64.Encod"
	iv := "初始向量"
	aseKey := "秘钥"
	orgDataByte, err := base64.StdEncoding.DecodeString(orgData)
	if err != nil {
		t.Errorf("err %s", err.Error())
		t.Fail()
	}
	aseKeyByte, err := base64.StdEncoding.DecodeString(aseKey)
	if err != nil {
		t.Errorf("err %s,%s", err.Error(), string(aseKeyByte))
		t.Fail()
	}

	ivByte, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		t.Errorf("err %s", err.Error())
		t.Fail()
	}

	err = AESDecrypt(orgDataByte, aseKeyByte, ivByte, &d)
	if err != nil {
		t.Errorf("err %s", err.Error())
		t.Fail()
	}
	t.Logf("TestAESDecrypt %+v", d)
}
