package utils

/*
在Go语言中，接口（interface）是一个自定义类型。接口类型是一种抽象的类型，它不会暴露
出它所代表的内部属性的结构，而只会展示出它自己的方法，因此，不能将接口类型实例化。
根据Go语言规范，单个函数的接口命名为函数名加上“er”作为后缀，例如Reader、Writer、
Formatter等。接口命名规范如下：
◇ 单个函数的接口名以“er”作为后缀，接口的实现则去掉“er”。
◇ 两个函数的接口名综合两个函数名，以“er”作为后缀，接口的实现则去掉“er”。
◇ 三个以上函数的接口，抽象这个接口的功能，类似于结构体命名。
*/

// 定义了一个叫Animal的接口，结构体Cat实现了Animal接口的两个方法，因此我们就可以认为，Cat实现了Animal接口

// go语言中的接口只定义方法,没有数据字段

type Animal interface {
	Name() string
	Speak() string
}

type Cat struct {
}

func (cat Cat) Name() string {
	return "Cat"
}

func (cat Cat) Speak() string {
	return "喵喵喵!"
}

/*
type InterfaceName interface{
 Method()
}
*/

// 创建一个数据库操作接口

//type IDdtabaser interface {
//	Connect() error
//	Disconnect() error
//}
//
//// MySQL数据库操作
//
//type Mysql struct {
//	DBName    string
//	isConnect bool
//}
//
//func (mysql *Mysql) Connect() error {
//	fmt.Println("Mysql Connect DB => " + mysql.DBName)
//
//	// 数据库连接
//	mysql.isConnect = true
//
//	if mysql.isConnect {
//		fmt.Println("Mysql Connect Success!")
//		return nil
//	} else {
//		return errors.New("Connect failure! ")
//	}
//
//}
//
//func (mysql *Mysql) Disconnect() error {
//	//关闭连接
//	fmt.Println("Mysql Disconnect Success!")
//	return nil
//}
//
//// Redis数据库操作
//
//type Redis struct {
//	DBName string
//}
//
//func (redis *Redis) Connect() error {
//	fmt.Println("Redis Connect DB => " + redis.DBName)
//	//do Connect
//	fmt.Println("Redis Connect Success!")
//	return nil
//}
//
//func (redis *Redis) Disconnect() error {
//	//do Disconnect
//	fmt.Println("Redis Disconnect Success!")
//	return nil
//}
//
//func Demo() {
//
//	var mysql = Mysql{DBName: "student"}
//	fmt.Println("开始连接")
//	mysql.Connect()
//	//do something
//	fmt.Println("断开连接")
//	mysql.Disconnect()
//
//	var redis = Redis{DBName: "teacher"}
//	fmt.Println("开始连接")
//	redis.Connect()
//	//do something
//	fmt.Println("断开连接")
//	redis.Disconnect()
//
//}
//
//// HandleDB()函数只有一个，却能实现处理多个不同类型的数据，这也称为Go的多态
//
//func HandleDB(db IDdtabaser) {
//	fmt.Println("开始连接")
//	db.Connect()
//	// do something
//	fmt.Println("断开连接")
//	db.Disconnect()
//}
//
//func Demo1() {
//	var mysql = Mysql{DBName: "student"}
//	HandleDB(&mysql)
//	var redis = Redis{DBName: "teacher"}
//	HandleDB(&redis)
//}
