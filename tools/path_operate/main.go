package main

import (
	"errors"
	"fmt"
	"gostart/path_operate/utils"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	files := "D:\\GoCode\\gostart\\path_operate\\utils\\test.txt"
	paths, fileName := filepath.Split(files)
	fmt.Println(paths, fileName)      //获取路径中的目录及文件名 E:\data\  test.txt
	fmt.Println(filepath.Base(files)) //获取路径中的文件名test.txt
	fmt.Println(path.Ext(files))      //获取路径中的文件的后缀 .txt

	str, _ := os.Getwd() // D:\GoCode\gostart
	fmt.Println(str)
	str2, _ := filepath.Abs(str)
	fmt.Println(str2)

	str3 := filepath.Join(str2, "../") // 获取父级路径
	fmt.Println(str3)

	str4, _ := filepath.Abs(str3)
	fmt.Println(str4)

	cpath := getCurrentPath()
	filePath := CurrentFile()
	fmt.Println(cpath)
	fmt.Println(filePath)

	// 遍历目录
	fs, _ := utils.WalkDir(`D:\GoCode\gostart`)
	for _, v := range fs {
		fmt.Println(v)
	}

	// 层次显示目录
	fs2, _ := utils.WalkDir2(`D:\GoCode\gostart`, 2)
	for _, v := range fs2 {
		fmt.Println(v)
	}

}

// CurrentFile 获取当前文件的详细路径   D:/GoCode/gostart/path_operate/main.go
func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}

//获取当前的执行路径 C:\Users\company\AppData\Local\Temp\GoLand\
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
