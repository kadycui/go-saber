package utils

import "io/ioutil"

func WalkDir(filepath string) ([]string, error) {
	files, err := ioutil.ReadDir(filepath) // files为当前目录下的所有文件名称【包括文件夹】
	if err != nil {
		return nil, err
	}

	var allfile []string
	for _, v := range files {
		fullPath := filepath + "\\" + v.Name() // 全路径 + 文件名称
		if v.IsDir() {                         // 如果是目录
			a, _ := WalkDir(fullPath) // 遍历改路径下的所有文件
			allfile = append(allfile, a...)
		} else {
			allfile = append(allfile, fullPath) // 如果不是文件夹，就直接追加到路径下
		}
	}

	return allfile, nil
}

//层次显示目录

func WalkDir2(filepath string, level int) ([]string, error) {
	prefix := "|"
	for i := 0; i < level; i++ {
		prefix += "------"
	}
	files, err := ioutil.ReadDir(filepath) // files为当前目录下的所有文件名称【包括文件夹】
	if err != nil {
		return nil, err
	}

	var allfile []string
	for _, v := range files {
		fullPath := filepath + "\\" + v.Name() // 全路径 + 文件名称
		if v.IsDir() {                         // 如果是目录
			allfile = append(allfile, prefix+v.Name())
			a, _ := WalkDir2(fullPath, level+1) // 遍历改路径下的所有文件
			allfile = append(allfile, a...)
		} else {
			allfile = append(allfile, prefix+v.Name()) // 如果不是文件夹，就直接追加到路径下
		}
	}

	return allfile, nil
}
