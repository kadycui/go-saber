package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileWrite() {
	// os.O_APPEND 追加
	// os.O_TRUNC 覆盖
	dir := GetPath()
	fileName := "w.txt"
	fullPath := filepath.Join(dir, fileName)
	fmt.Println("写入文件路径: ", fullPath)
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	content := "今天阳光明媚,\n宜赏花!!"
	n, err := file.Write([]byte(content))
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	} else {
		fmt.Printf("成功写入文件%d个字节\n", n)
	}

}
