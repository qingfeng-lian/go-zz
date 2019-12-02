package zz

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//检测结构体字段
func CheckStructField(data interface{}) (err error) {
	rv := reflect.ValueOf(data)
	numField := rv.NumField()
	for i := 0; i < numField; i++ {
		field := rv.Type().Field(i)
		tagName := getTagName(field.Tag.Get("check"))
		tagNameSlice := strings.Split(tagName, ";")
		for _, v := range tagNameSlice {
			switch v {
			case "":
				//这是一个特殊的判断， strings.Split 返回空字符串
			case "required":
				if rv.Field(i).IsZero() {
					err = errors.New(fmt.Sprintf("%s is required", field.Name))
					return
				}
			case "string":

			default:
				err = errors.New(fmt.Sprintf("field check type not found, current type %+v", v))
				return
			}
		}
	}
	return err
}

func getTagName(tagNameStr string) string {
	tagName := strings.Split(tagNameStr, ",")
	return tagName[0]
}
