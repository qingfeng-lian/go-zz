package zz

import (
	"encoding/base64"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	orgData := "fWTJlun0V1UVo8ubE67EmQ==" //"这里请换上需要解密的字符串， 此字符串是经过base64.Encod"
	iv := "1234567890123456"              //"初始向量"
	aseKey := "keyqingfeng.lian"          //"秘钥"
	orgDataByte, err := base64.StdEncoding.DecodeString(orgData)
	if err != nil {
		t.Errorf("err %s", err.Error())
		t.Fail()
	}
	d, err := AESDecrypt(orgDataByte, []byte(aseKey), []byte(iv))
	if err != nil {
		t.Errorf("err %s", err.Error())
		t.Fail()
	}
	t.Logf("TestAESDecrypt %+v", d)
}

func TestAESEncrypt(t *testing.T) {
	byteCode := AESEncrypt([]byte("qingfeng.lian"), []byte("keyqingfeng.lian"), []byte("1234567890123456"))

	t.Logf("str %s", base64.StdEncoding.EncodeToString(byteCode))
}
