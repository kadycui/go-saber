package utils

import "fmt"

type NoticeInterface interface {
	SeedEmail()
	SeedSMS()
}

type Studentt struct {
	Name   string
	Email  string
	Pthone string
}

// 如果方法集使用的是指针类型，那么我们必须传入指针类型的值，如果方法集使用的是值类型，那么我们既可以传入值类型也可以传入指针类型

func (stt *Studentt) SeedEmail() {
	fmt.Printf("发送邮件给%s\r\n", stt.Email)
}

func (stt *Studentt) SeedSMS() {
	fmt.Printf("发送短信给%s\r\n", stt.Pthone)
}

func Demo6() {
	stu := Studentt{"Tom", "tom@gmail.com", "10086"}
	SeedNotice(&stu)

}

func SeedNotice(notice NoticeInterface) {
	notice.SeedEmail()
	notice.SeedSMS()
}
