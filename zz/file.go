package zz

import (
	"fmt"
	"io"
	"os"
	"path"
)

//判断文件是否存在 绝对路径
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

//递归创建文件夹
func Mkdir(path string, modePerm os.FileMode) error {
	//递归创建文件夹
	err := os.MkdirAll(path, modePerm)
	if err != nil {
		return err
	}
	return nil
}

//创建文件，如果目录不存在递归创建目录
func Mkfile(filePath string, modePerm os.FileMode) error {
	//递归创建文件夹
	dirPath := path.Dir(filePath)
	err := os.MkdirAll(dirPath, modePerm)
	if err != nil {
		return err
	}
	_, err = os.Create(filePath)
	if err != nil {
		return err
	}
	return nil
}

//简单的文件写入内容
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	//打开文件
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	//文件的写入
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	//关闭文件
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
