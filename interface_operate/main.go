package main

import (
	"fmt"
	. "gostart/interface_operate/service"
	"gostart/interface_operate/utils"
)

func main() {
	// var service IService = NewUserService()
	// service.Save()

	var is IService = NewProdService()
	is.Save()

	utils.Demo1()
	utils.Demo2()
	utils.Demo3()
	utils.Demo4()
	utils.Demo5()
	fmt.Println("-----------------------------")
	utils.Demo6()

}
