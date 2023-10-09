package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func JsonToDb() {

	path := GetPath()
	p1, _ := filepath.Abs(path)
	p2 := filepath.Join(p1, "../conf/data.json")
	file, err := os.ReadFile(p2)
	if err != nil {
		log.Fatal(err)
	}

	var data []interface{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, item := range data {
		t1 := time.Now()
		effBgId, err := strconv.Atoi(item.(map[string]interface{})["effBgId"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		pName := item.(map[string]interface{})["pName"]
		faceName := item.(map[string]interface{})["faceName"]
		sid, err := strconv.Atoi(item.(map[string]interface{})["sid"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		width, err := strconv.Atoi(item.(map[string]interface{})["width"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		propId, err := strconv.Atoi(item.(map[string]interface{})["propId"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		pid, err := strconv.Atoi(item.(map[string]interface{})["pid"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		hight, err := strconv.Atoi(item.(map[string]interface{})["hight"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		titleId, err := strconv.Atoi(item.(map[string]interface{})["titleId"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		textName, err := strconv.Atoi(item.(map[string]interface{})["textName"].(string))
		if err != nil {
			log.Println("转换 int 失败：", err)
			return
		}
		lastTime := item.(map[string]interface{})["lastTime"].(string)

		// fmt.Printf("res: %T,%T,%T,%T,%T,%T,%T,%T,%T,%T\n", effBgId, pName, faceName, sid, width, propId, pid, hight, titleId, textName)
		// fmt.Printf("res: %v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n", effBgId, pName, faceName, sid, width, propId, pid, hight, titleId, textName)
		// fmt.Printf("res: %d,%s,%s,%d,%d,%d,%d,%d,%d,%d\n", effBgId, pName, faceName, sid, width, propId, pid, hight, titleId, textName)
		result, err := db.Exec("INSERT INTO dynamic (effBgId, pName, faceName, sid, width, propId, pid, hight, titleId, textName, lastTime) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", effBgId, pName, faceName, sid, width, propId, pid, hight, titleId, textName, lastTime)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.LastInsertId())
		elapsed := time.Since(t1)
		fmt.Println(elapsed)

	}

}
