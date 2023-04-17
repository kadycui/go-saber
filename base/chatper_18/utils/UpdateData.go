package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func UpdateData() {
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

	rs, err := db.Exec("update `userinfo` set password=? where username=?", "88888", "john")
	checkErr(err)

	rowCount, err := rs.RowsAffected()
	checkErr(err)

	if rowCount > 0 {
		log.Println("数据更新成功")
	}

}
