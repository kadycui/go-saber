package main

import (
	"fmt"
	"gostart/libs/common"
	"gostart/libs/system"
)

func main() {
	var evn = system.Args.Env
	var id = system.Args.ServiceId
	var root = system.Root

	fmt.Println("Env:", evn)
	fmt.Println("ServiceId", id)
	fmt.Println("root", root)

	var ip = common.GetLocalIp()
	fmt.Println(ip)
}
