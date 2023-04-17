package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type UserInfo struct {
	Uid      int
	Username string
	Gender   bool
	Password string
	Created  string
}

func SelectData() {
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

	rows, err := db.Query("select * from `userinfo` where username=?", "john")
	defer rows.Close()

	for rows.Next() {
		user := UserInfo{}
		err := rows.Scan(&user.Uid, &user.Username, &user.Gender, &user.Password, &user.Created)
		checkErr(err)
		log.Println(user)

	}

}
