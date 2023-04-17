package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

// 数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "gostart"
)

// Db数据库连接池

var DB *sql.DB

type User struct {
	id    int64  `db:"id"`
	name  string `db:"name"`
	age   int8   `db:"age"`
	sex   int8   `db:"sex"`
	phone string `db:"phone"`
}

// 方法名大写就是 publish

func InitDb() *sql.DB {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Println("打开数据库失败")
		return nil
	}
	log.Println("连接成功")
	return DB
}

// 创建表

func CreateTable() {
	DB = InitDb()
	err := DB.Ping()
	if err != nil {
		return
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(DB)

	createTable := "CREATE TABLE `gostart`.`user` (" +
		"`id` INT(10) NOT NULL AUTO_INCREMENT," +
		"`name` VARCHAR(64) NULL," +
		"`sex` INT(20) NULL DEFAULT NULL," +
		"`age` INT(20) NULL DEFAULT NULL," +
		"`phone` VARCHAR(64) NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`)" +
		");"

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("创建表失败", err)

	}
	log.Println("数据表创建成功!")

}

// 插入数据

func InsertData() {
	DB = InitDb()
	err := DB.Ping()
	if err != nil {
		return
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(DB)

	_, err2 := DB.Query("INSERT INTO user (name, sex, age,phone) VALUES('崔护',1, 30,'182353623245')")
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("写入数据成功")

}

// 查询操作

func SelectData() {
	DB = InitDb()
	err := DB.Ping()
	if err != nil {
		return
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(DB)
	var user User
	// rows, err := DB.Query("select * from user")
	var s int = 2
	// 条件语句查询
	rows, err := DB.Query("select * from user where sex=?", s)
	if err != nil {
		fmt.Println("数据库执行更新出错")
	}
	for rows.Next() {
		err = rows.Scan(&user.id, &user.name, &user.age, &user.sex, &user.phone)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Id: %d, Name: %s, Age: %d, Sex: %d, Phone: %s\n", user.id, user.name, user.age, user.sex, user.phone)
		//log.Println(user)
	}

}

func DeleteData() {
	DB = InitDb()
	err := DB.Ping()
	if err != nil {
		return
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(DB)

	var s = "DELETE FROM user WHERE name = '崔护'"
	res, err2 := DB.Exec(s)
	if err2 != nil {
		panic(err2.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("受到影响 %d 行\n", affectedRows)
}

func UpdateData() {
	DB = InitDb()
	err := DB.Ping()
	if err != nil {
		return
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(DB)

	var s = "update user set name = ? WHERE id = ?"
	res, err2 := DB.Exec(s, "BaiJuyi", 10)
	if err2 != nil {
		panic(err2.Error())
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("更新成功, 受影响 %d 行 \n", affectedRows)
}
