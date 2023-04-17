package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 遍历目录

func FileDir() {
	dir, err := ioutil.ReadDir("C:\\Users")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range dir {
		fmt.Println(file.Name())
	}
}

// 遍历文件和目录

func ListDir(dirPth string) error {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range dir {
		if file.IsDir() { // 忽略目录
			fmt.Println("目录:" + file.Name())

		} else {
			fmt.Println("文件:" + file.Name())
		}

	}
	return nil

}

/*
Walk函数会遍历root指定的目录下的文件树，对每一个该文件树中的目录和文件都会调用
walkFn，包括root自身。所有访问文件/目录时遇到的错误都会传递给walkFn过滤。文件是按词法顺
序遍历的，这让输出显得更漂亮，但也导致处理非常大的目录时效率会降低。Walk函数不会遍历
文件树中的符号链接（快捷方式）文件包含的路径。
*/

//获取指定目录及所有子目录下的所有文件

func WalkDir(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

// 创建目录
// func Mkdir(name string, perm FileMode) error

func createDir(path string, dirName string) {
	dirPath := path + "\\" + dirName
	err := os.Mkdir(dirPath, 0777) //0777也可以os.ModePerm
	if err != nil {
		fmt.Println(err)
	}
	err = os.Chmod(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建文件夹=>" + path + dirName)
}

func Demo1() {
	createDir("C:\\Users", "text")
	createDir("C:\\Windows\\Temp", "test")

}

func createDirAll(path string, dirName string) {
	dirPath := path + "\\" + dirName
	fmt.Println("创建目录=> " + dirPath)
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("创建成功")
	}
	err2 := os.Chmod(dirPath, 0777)
	if err != nil {
		fmt.Println(err2)
	}

}

func Demo2() {
	createDirAll("D:\\vvv", "dir1\\dir2\\dir3")
}

// 删除空目录

func deleteEmptyDir(dirPath string) {
	fmt.Println("Delete Dir => " + dirPath)
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除成功!")
	}
}

func Demo3() {
	deleteEmptyDir("D:\\templ\\test")
}

// 非空文件夹删除

func deleteNotEmptyDir(dirPath string) {
	fmt.Println("Delete Dir => " + dirPath)
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除成功!")
	}
}

func Demo4() {
	deleteNotEmptyDir("D:\\templ\\test")
}
