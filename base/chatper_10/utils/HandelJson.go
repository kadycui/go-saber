package utils

import (
	"encoding/json"
	"fmt"
)

/*
{
 "name":"小王",
 "age":24,
 "sex":true,
 "birthday":"1995-01-01",
 "company":"百度",
 "language":[
 "Go",
 "PHP",
 "Python"
 ]
}

*/
// 标准库提供了encoding/json库来处理JSON。编码JSON，即从其他的数据类型编码成JSON字符串，这个过程我们会使用如下的接口
//func Marshal(v interface{}) ([]byte, error)

//Marshal函数返回interface{}类型的JSON编码，通常interface{}类型会使用map或者结构体。为了
//让输出的JSON字符串更加直观，可以使用另一个JSON编码接口，对输出的JSON进行格式化操作。
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

func Demo10() {
	// 创建map
	m := make(map[string]interface{}, 6)
	m["name"] = "小王"
	m["age"] = 24
	m["sex"] = true
	m["birthday"] = "1995-01-01"
	m["company"] = "百度"
	m["language"] = []string{"Go", "PHP", "Python"}

	// 编码成JSON
	result, _ := json.Marshal(m)
	resultFormat, _ := json.MarshalIndent(m, "", "    ")
	fmt.Println("result = ", string(result))
	fmt.Println("resultFofmat = ", string(resultFormat))

}

/*
在定义struct字段的时候，可以在字段后面添加标签来控制编码/解码的过程：是否要编码或解
码某个字段，JSON中的字段名称是什么。可以选择的控制字段有三种：

1.-：不要解析这个字段。
2.omitempty：当字段为空（默认值）时，不要解析这个字段。比如false、0、nil、长度为0的array、map、slice、string。
3.FieldName：当解析JSON的时候，使用这个名字。
在动手写11.3.2中，“`json:"name"`”就是定义的第三类标签，表示将Name属性的key值解析为name。
*/

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Sex      bool     `json:"sex"`
	Birthday string   `json:"birthday"`
	Company  string   `json:"company,omitempty"`
	Language []string `json:"language"`
}

func Demo11() {
	// 定义一个结构体变量
	person := Person{"小王", 24, true, "1995-01-01", "零壹快学", []string{"Go", "PHP", "Python"}}

	// result, err := json.Marshal(person)
	result, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result = ", string(result))

}

// json解析
// func Unmarshal(data []byte, v interface{}) error
/*
Bool 						对应JSON布尔类型
float64 					对应JSON数字类型
string 						对应JSON字符串类型
[]interface{} 				对应JSON数组
map[string]interface{} 		对应JSON对象
nil 						对应JSON的null
*/

func ParseJson() {
	jsonstr := `
{
	 "name":"小王",
	 "age":24,
	 "sex":true,
	 "birthday":"1995-01-01",
	 "company":"百度",
	 "language":[
		 "Go",
		 "PHP",
		 "Python"
	 ]
}
`

	// 创建一个map
	m := make(map[string]interface{}, 6)
	err := json.Unmarshal([]byte(jsonstr), &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("m = ", m)

	// 类型断言
	for key, value := range m {
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]的值类型为string，value = %s\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为int，value = %f\n", key,
				data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool，value = %t\n", key,
				data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string，value = %v\n",
				key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface{}，value =%v\n", key, data)

		}
	}
}

// struct来解码JSON，JSON库会自动对结构体的类型进行解析，无须类型判断。

type Person2 struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Sex      bool     `json:"sex"`
	Birthday string   `json:"birthday"`
	Company  string   `json:"company"`
	Language []string `json:"language"`
}

func ParseJson2() {
	jsonstr := `
{
	 "name":"小王",
	 "age":24,
	 "sex":true,
	 "birthday":"1995-01-01",
	 "company":"百度",
	 "language":[
		 "Go",
		 "PHP",
		 "Python"
	 ]
}
`

	var person Person2
	err := json.Unmarshal([]byte(jsonstr), &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("person = %+v", person)
}
