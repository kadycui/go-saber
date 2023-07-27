package main

import (
	"fmt"
	. "gostart/feature/interface_operate/service"
	"gostart/feature/interface_operate/utils"
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
	utils.Demo6()
	fmt.Println("-----------------------------")
	utils.AreaCircle()
	utils.Demo7()

}
