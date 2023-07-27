package utils

import "fmt"

type Pet interface {
	Eat()
}

type Dog struct {
	name string
}

// func (d Dog) Eat(name string) string {
// 	d.name = "小泉"
// 	fmt.Printf("name: %v\n", name)
// 	return "吃的狗粮"
// }

func (d *Dog) Eat(name string) string {
	d.name = "小泉"
	fmt.Printf("name: %v\n", name)
	return "吃的狗粮"
}

func Demo1() {
	// d := Dog{
	// 	name: "川普",
	// }
	// r := d.Eat("骨头")
	// fmt.Printf("r: %v\n", r)
	// fmt.Printf("d: %v\n", d)

	d := &Dog{
		name: "川普",
	}
	r := d.Eat("骨头")
	fmt.Printf("r: %v\n", r)
	fmt.Printf("d: %v\n", d)
}
