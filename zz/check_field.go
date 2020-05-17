package zz

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

//检测结构体字段 ,检查通过err=nil
//required :必填 |  email:邮箱  |  mobile:手机号  |  maxLen:字符串最大长度(包含)  |  minLen:字符串最小长度(包含)
func CheckStructField(data interface{}) (err error) {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	numField := rv.NumField()
	for i := 0; i < numField; i++ {
		field := rv.Type().Field(i)
		tagName := getTagName(field.Tag.Get("check"))
		tagFieldName := field.Tag.Get("json")
		if tagFieldName == "" {
			tagFieldName = field.Name
		}
		tagNameSlice := strings.Split(tagName, ";")
		for _, v := range tagNameSlice {
			tmpSlice := strings.Split(v, ":")
			if len(tmpSlice) == 2 {
				switch tmpSlice[0] {
				case "maxLen":
					lenNum, lenErr := strconv.ParseInt(tmpSlice[1], 10, 0)
					if lenErr != nil {
						err = errors.New(fmt.Sprintf("%s is type err, err msg:%+v", tmpSlice[0], lenErr.Error()))
						return
					}
					if utf8.RuneCountInString(rv.Field(i).String()) > int(lenNum) {
						err = errors.New(fmt.Sprintf("%s is maxLen  greater than %+v", tmpSlice[0], tmpSlice[1]))
						return
					}
				case "minLen":
					lenNum, lenErr := strconv.ParseInt(tmpSlice[1], 10, 0)
					if lenErr != nil {
						err = errors.New(fmt.Sprintf("%s is type err err msg:%+v", tmpSlice[0], lenErr.Error()))
						return
					}
					//fmt.Println("bbbbbbbbb", utf8.RuneCountInString(rv.Field(i).String()), int(lenNum))
					if utf8.RuneCountInString(rv.Field(i).String()) < int(lenNum) {
						err = errors.New(fmt.Sprintf("%s is maxLen  less than %+v", tmpSlice[0], tmpSlice[1]))
						return
					}
				default:
					err = errors.New(fmt.Sprintf("%s not found tag", tmpSlice[0]))
					return
				}

			} else if len(tmpSlice) == 1 {
				switch v {
				case "":
					//这是一个特殊的判断， strings.Split 返回空字符串
				case "required": //必填
					if rv.Field(i).IsZero() {
						err = errors.New(fmt.Sprintf("%s is required", tagFieldName))
						return
					}
				case "email":
					if !VerifyEmailFormat(rv.Field(i).String()) {
						err = errors.New(fmt.Sprintf("%s is invalid eamil", tagName))
						return
					}
				case "mobile":
					if !VerifyMobileFormat(rv.Field(i).String()) {
						err = errors.New(fmt.Sprintf("%s is invalid mobile", tagName))
						return
					}
				default:
					err = errors.New(fmt.Sprintf("field check type not found, current type %+v", v))
					return
				}
			} else {
				err = errors.New(fmt.Sprintf("%s is format err", tagName))
				return
			}
		}
	}
	return
}

func getTagName(tagNameStr string) string {
	tagName := strings.Split(tagNameStr, ",")
	return tagName[0]
}
