package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadDataLine(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	var data [][]string

	for {
		csvdata, err := reader.Read()
		if err == io.EOF {
			break
		}
		data = append(data, csvdata)
		for _, line := range data {
			fmt.Println(line)
		}
	}

}

func ReadDataAll(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	csvdata, err := reader.ReadAll()

	for _, line := range csvdata {
		fmt.Println(line)
	}
}
