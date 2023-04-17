package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// 匿名结构体

/*
结构体实例 := struct{
	//匿名结构体定义
	 成员变量1 类型1
	 成员变量2 类型2
	 成员变量3 类型3
	 ...
	}{
	 //成员变量初始化（可选）
	 成员变量1:值1,
	 成员变量2:值2,
	 成员变量3:值3,
	 ...
}

*/

type Book3 struct {
	title  string
	author string
	num    int
	id     int
}

func Demo11() {
	book3 := struct {
		title  string
		author string
		num    int
		id     int
	}{
		title:  "python",
		author: "Jam",
		num:    10,
		id:     1004,
	}
	fmt.Println("title:", book3.title)
	fmt.Println("author:", book3.author)
	fmt.Println("num:", book3.num)
	fmt.Println("id:", book3.id)
}

//匿名结构体一般可用于组织全局变量、构建数据模板和解析JSON等。
//例如，我们可以通过自定义匿名结构体将同一类的全局变量组织在一起。

func Demo12() {
	data := &struct {
		Code int
		Msg  string
	}{}

	jsonData := `{"code":200,"msg":"success"}`
	if err := json.Unmarshal([]byte(jsonData), data); err != nil {
		fmt.Println(err)
	}

	fmt.Println("code:", data.Code)
	fmt.Println("msg:", data.Msg)
}

// Go语言中是利用反射机制将来自XML文件中的数据反射成对应的struct对象的，其中缺少的元
//素或空属性值将被解析为零值。

type Result struct {
	Person []Person
}

type Person struct {
	Name      string
	Age       string
	Interests Interests
}

type Interests struct {
	Interest []string
}

func Demo13() {
	var res Result
	content, err := ioutil.ReadFile("chatper_08/utils/test.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = xml.Unmarshal(content, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("XML文件解析后内容为：")
	fmt.Println(res)
}
