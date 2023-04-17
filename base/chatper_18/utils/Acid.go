package utils

import (
	"database/sql"
	"log"
)

func checkErrWithTx(err error, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
	}
	log.Fatal(err)

}

func Acid() {
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

	var password string

	tx, err := db.Begin()
	checkErr(err)

	// 查找john的密码，如果密码为123123就将密码改为111111，否则不执行任何操作
	err = tx.QueryRow("select password from `userinfo` where username=?", "john").Scan(&password)
	checkErrWithTx(err, tx)

	if password == "123123" {
		rs, err := tx.Exec("update `userinfo` set password=? where username=?", "111111", "john")
		checkErrWithTx(err, tx)
		rowCount, err := rs.RowsAffected()
		checkErrWithTx(err, tx)
		if rowCount > 0 {
			log.Println("密码更新完成！")
		}

	}
	tx.Commit()
	log.Println("事务处理完成！")
}
