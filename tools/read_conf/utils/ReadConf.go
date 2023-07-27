package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type Mysql struct {
	Url  string
	Port int
}

type Redis struct {
	Host string
	Port int
}

func ReadConf() {
	// // 使用 gopkg.in/yaml.v3 读取配置文件
	dataBytes, err := os.ReadFile("read_conf/conf/conf.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	fmt.Printf("config → %+v\n", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}

	mp := make(map[string]any, 2)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	fmt.Printf("map → %+v", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}

}

func ViperConf() {
	// 使用 github.com/spf13/viper  读取配置文件
	// 设置配置文件的名字
	viper.SetConfigName("conf")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath("./read_conf/conf")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(viper.Get("mysql"))     // map[port:3306 url:127.0.0.1]
	fmt.Println(viper.Get("mysql.url")) // 127.0.0.1

}
