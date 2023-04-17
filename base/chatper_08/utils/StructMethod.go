package utils

import "fmt"

type User struct {
	Name  string
	Email string
}

// User结构体创建了ChangeName方法，在该方法内部将接收者的Name值修改为Tom，接收者的类型为指针类型

// 如果方法需要修改接收者，接收者必须是指针类型。

func (u *User) ChangeName() { // 指针类型接受者
	u.Name = "Tom"
}

func Demo6() {
	u := &User{"Peter", "kadycui@qq.com"} // 创建指针类型实例
	fmt.Println("Name:", u.Name, " Email:", u.Email)
	u.ChangeName()
	fmt.Println("Name:", u.Name, " Email:", u.Email)

}

func Demo7() {
	// 只要结构体方法的接收者为指针类型，即使实例不是指针类型，修改也能生效
	u := User{"Jam", "kk@qq.com"}
	fmt.Println("Name:", u.Name, " Email:", u.Email)
	u.ChangeName()
	fmt.Println("Name:", u.Name, " Email:", u.Email)
}
