package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FileRead() {

	file, err := os.Open("./onego/File/r.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	content := make([]byte, 1000)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("读取文件失败", err)
		file.Close()
		return
	}
	fmt.Println(string(content[0:n]))
	// fmt.Println(content)

	file.Close()
}

func ReadLineFile() {
	dir := GetPath()
	fileName := "r.txt"
	fullPath := filepath.Join(dir, fileName)
	fmt.Println("读文件路径: ", fullPath)
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line)
				break
			} else {
				fmt.Println("读取文件发生错误", err)
				return
			}
		} else {
			fmt.Print(line)
		}
	}

}
