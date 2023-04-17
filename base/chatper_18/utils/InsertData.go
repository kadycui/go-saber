package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/djcache")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// 连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败!", err)
	}

	log.Println("数据库连接成功!")

	// 创建插入语句
	// rs, err := db.Exec("INSERT INTO `userinfo`(username,gender,password,created) VALUES (?,?,?,?)", "john", 1, "123456", time.Now())

	// 使用预编译模式
	stmt, err := db.Prepare("INSERT INTO `userinfo`(username,gender,password,created) VALUES (?,?,?,?)")

	defer stmt.Close()
	rs, err := stmt.Exec("Ailsa", 0, "111111", time.Now())
	checkErr(err)

	rowCount, err := rs.RowsAffected()
	checkErr(err)
	log.Printf("插入了 %d 行", rowCount)

}
