package zz

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

//map to key=val&key=val 按key正序排列
func MapSortKeyToString(data map[string]interface{}) string {
	keyMap := make([]string, 0)
	dataStr := ""
	if len(data) <= 0 {
		return ""
	}
	for k, _ := range data {
		keyMap = append(keyMap, k)
	}
	sort.Strings(keyMap)
	for _, k := range keyMap {
		v := data[k]
		tmpS := ""
		fmt.Println(reflect.TypeOf(v).String())
		switch reflect.TypeOf(v).String() {
		case "string":
			tmpS = v.(string)
		case "int64":
			tmpS = strconv.FormatInt(v.(int64), 10)
		case "uint64":
			tmpS = strconv.FormatUint(v.(uint64), 10)
		case "float64":
			tmpS = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case "uint":
			tmpS = strconv.FormatUint(uint64(v.(uint)), 10)
		case "int":
			tmpS = strconv.Itoa(v.(int))
		default:
			tmpS = ""
		}
		dataStr += k + "=" + tmpS + "&"
	}
	dataStr = dataStr[:len(dataStr)-1]
	fmt.Println("MapSortKeyToString", dataStr)
	return dataStr
}

//map生成md5签名
func MapToMd5String(data map[string]interface{}) string {
	dataStr := MapSortKeyToString(data)
	has := md5.Sum([]byte(dataStr))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
