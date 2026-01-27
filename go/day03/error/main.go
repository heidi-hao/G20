package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

// 全局兜底“异常钩子”
func panicHandler() {
	fmt.Println("panic handler")
	if r := recover(); r != nil { //recover只能在defer里面
		fmt.Println("Recoverd from panic:", r)
		os.Exit(1)
	}
}

func main() {
	defer panicHandler() //只要main退出一定会执行panicHandler 因为defer

	//变量目录 打印了文件名称
	err := walkDir("./xxxx", func(filepath string) {
		fmt.Println(filepath)
	})
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("not exist", err)
		} else {
			fmt.Println("%s", err)
		}
	}

}

// 最常见的就是 遍历目录里面的文件
func walkDir(path string, fn func(string)) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	//读取path目录下的文件和子目录
	for _, file := range files {
		//如果是目录，继续往下找
		if file.IsDir() {
			walkDir(path+"/"+file.Name(), fn)
		} else {
			//文件 调用fn函数处理
			fn(path + "/" + file.Name())
		}
	}
	return nil
}
