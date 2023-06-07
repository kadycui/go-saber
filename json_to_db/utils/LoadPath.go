package utils

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/ini.v1"
)

func GetPath() string {
	// 获取当前执行文件绝对路径（go run）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}
	dir := path.Dir(filename)
	return dir
}

func LoadPath() {
	path := GetPath()
	p1, _ := filepath.Abs(path)
	p2 := filepath.Join(p1, "../")

	cfg, err := ini.Load(p2 + "\\conf\\my.ini")
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
	}

	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())

	//获取一个类型为字符串（string）的值
	fmt.Println("MySQL IP:", cfg.Section("mysql").Key("ip").String())
	//获取一个类型为整形（int）的值
	mysqlPort, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Port:", mysqlPort)
	fmt.Println("MySQL User:", cfg.Section("mysql").Key("user").String())
	fmt.Println("MySQL Password:", cfg.Section("mysql").Key("password").String())
	fmt.Println("MySQL Database:", cfg.Section("mysql").Key("database").String())

	//go-ini也提供对应的MustType（Type 为Init/Uint/Float64等）方法，这个方法只返回一个值。
	//fmt.Println("redis Port:", cfg.Section("redis").Key("port").MustInt(6381))

	//获取一个类型为字符串（string）的值
	fmt.Println("Redis IP:", cfg.Section("redis").Key("ip").String())
	redisPort, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis Port:", redisPort)

}
