# go-zz  go常用方法
这里封装了一些go常用的方法，包括常用的方法和微信小程序接口等

## 安装使用
#### go mod
```go
require (
	github.com/qingfeng-lian/go-zz master
)
```
#### go path
```go
go get github.com/qingfeng-lian/go-zz
```


## go 常用方法封装
## 常用方法集合
| 说明 | 对应文件 | 是否有测试案例|
| :--- | :--- | :---:|
| 结构体字段规则校验 | check_field.go |&radic;|
| aes加密解密 | aes.go |&radic;|
| http请求 | http.go |&times;|
|微信小程序部分接口实现|wx.go|&radic;|
|常用验证规则（邮箱验证，手机号验证等等）|verify_format.go|&radic;|
|目录文件操作(创建、写入文件等)|file.go|&times;|
|map常用操作|utility_map.go|&times;|

### map处理
map按key从小到大排序返回字符串 以"&"链接
map生成按key从小到大排序返回字符串以"&"链接的字符串的md5

