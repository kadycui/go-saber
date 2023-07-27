package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	LogLevel string
	MySQL    MySQLConfig
	Redis    RedisConfig
}

type MySQLConfig struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

type RedisConfig struct {
	IP   string
	Port int
}

var Conf Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./viper_operate/configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件错误：%w", err))
	}

	
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("配置文件格式错误：%w", err))
	}

	fmt.Println("AppName:", Conf.AppName)
	fmt.Println("LogLevel:", Conf.LogLevel)
	fmt.Println("MySQL Host:", Conf.MySQL.Host)
	fmt.Println("MySQL Port:", Conf.MySQL.Port)
	fmt.Println("MySQL User:", Conf.MySQL.User)
	fmt.Println("MySQL Pwd:", Conf.MySQL.Pwd)
	fmt.Println("MySQL DbName:", Conf.MySQL.DbName)
	fmt.Println("Redis IP:", Conf.Redis.IP)
	fmt.Println("Redis Port:", Conf.Redis.Port)
}
