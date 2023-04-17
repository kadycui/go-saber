package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// root:123456@tcp(10.16.168.30:3336)/gms_xunxian01

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

func init() {
	//配置MySQL连接参数
	username := "root"        //账号
	password := "123456"      //密码
	host := "10.16.168.30"    //数据库地址，可以是Ip或者域名
	port := 3336              //数据库端口
	Dbname := "gms_xunxian01" //数据库名
	timeout := "10s"          //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

}

func GetDB() *gorm.DB {
	return _db

}

func main() {

	db := GetDB()
	fmt.Println(db)

}
