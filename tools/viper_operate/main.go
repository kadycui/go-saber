package main

import (
	"fmt"
	"gostart/tools/viper_operate/configs"
)

func main() {
	fmt.Println("hello hello")

	fmt.Println(configs.Conf.AppName)
}
