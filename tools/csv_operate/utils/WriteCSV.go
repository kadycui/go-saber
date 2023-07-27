package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteData(filePath string, data [][]string) {

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	var header []string = []string{"name", "age", "address"}

	// 写入文件头
	err = writer.Write(header)

	if err != nil {
		fmt.Println(err)
	}

	for _, value := range data {
		// err = writer.WriteAll(data)    // 全部写入
		err := writer.Write(value)
		if err != nil {
			fmt.Println(err)
		}

	}

}
