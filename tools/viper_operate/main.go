package main

import (
	"fmt"

	"github.com/kadycui/go-saber/tools/viper_operate/configs"
)

func main() {
	fmt.Println("hello hello")

	fmt.Println(configs.Conf.AppName)
}
