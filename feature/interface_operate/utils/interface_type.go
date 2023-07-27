package utils

import "fmt"

// 一个类型实现多个接口

type Music interface {
	PlayMusic()
}

type Video interface {
	PlayVideo()
}

type Mobile struct {
}

func (m Mobile) PlayMusic() {
	fmt.Println("播放音乐")
}

func (m Mobile) PlayVideo() {
	fmt.Println("播放视频")
}

func Demo2() {
	m := Mobile{}
	m.PlayVideo()
	m.PlayMusic()
}

// 多个类型实现同一个接口

type Pet2 interface {
	Eat()
}

type Dog2 struct {
}

type Cat2 struct {
}

func (d Dog2) Eat() {
	fmt.Println("dog eat")
}

func (c Cat2) Eat() {
	fmt.Println("cat eat")
}

func Demo3() {
	dog := Dog2{}
	cat := Cat2{}

	dog.Eat()
	cat.Eat()

	// 实现多态
	var p Pet2
	p = Dog2{}
	p.Eat()
	p = Cat2{}
	p.Eat()

}

// 定义一个接口
type Phone interface {
	speak()
	read()
}

// 以下结构体可以分别设置自己的属性
type IPhone struct {
	name string
}
type Oppo struct {
	id int
}
type Mi struct {
	f bool
}

// 手机讲话区域（对相关的接口进行实现）
func (a IPhone) speak() {
	fmt.Println("我叫sir,您好！")
}
func (a Oppo) speak() {
	fmt.Println("我是oppo小精灵!")
}
func (a Mi) speak() {
	fmt.Println("大家好,我是小爱童鞋!")
}

func Demo4() {
	// 将新建对象传入展示大舞台,大舞台代码不变,展示不同的效果
	i := IPhone{
		name: "苹果",
	}
	i.speak()

	o := Oppo{
		id: 20,
	}
	o.speak()

	m := Mi{
		f: true,
	}
	m.speak()
}
