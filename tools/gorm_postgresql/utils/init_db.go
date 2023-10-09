package utils

import (
	"fmt"

	"github.com/kadycui/go-saber/tools/gorm_postgresql/model"

	"gorm.io/gorm/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// root:123456@tcp(10.16.168.30:3336)/gms_xunxian01

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

func init() {
	//配置postgres连接参数
	username := "kadycui" //账号
	password := "123456"  //密码
	host := "127.0.0.1"   //数据库地址，可以是Ip或者域名
	port := 5432          //数据库端口
	Dbname := "saber"     //数据库名
	// timeout := "10s"      //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?sslmode=disable TimeZone=Asia/Shanghai&timeout=%s", username, password, host, port, Dbname, timeout)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, username, password, Dbname, port)
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接postgres, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 禁用表明复数
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	// 数据迁移
	err = Migrate(_db)
	if err != nil {
		return
	}

}

func GetDB() *gorm.DB {
	return _db
}

func Migrate(db *gorm.DB) error {
	// err := db.Table("dynamic").AutoMigrate(&Dynamic{})
	err := db.AutoMigrate(&model.Book{})
	return err
}
