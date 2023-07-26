package main

import (
	"fmt"
	"gostart/viper_operate/configs"
)

func main() {
	fmt.Println("hello hello")

	fmt.Println(configs.Conf.AppName)
}
