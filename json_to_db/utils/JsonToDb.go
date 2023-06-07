package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func JsonToDb() {
	file, err := ioutil.ReadFile("json_to_db\\utils\\data.json")
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
		effBgId := int(item.(map[string]interface{})["effBgId"].(float64))
		pName := item.(map[string]interface{})["pName"]
		faceName := item.(map[string]interface{})["faceName"]
		sid := int(item.(map[string]interface{})["sid"].(float64))
		with := int(item.(map[string]interface{})["with"].(float64))
		propId := int(item.(map[string]interface{})["propId"].(float64))
		pid := int(item.(map[string]interface{})["pid"].(float64))
		hight := int(item.(map[string]interface{})["hight"].(float64))
		titleId := int(item.(map[string]interface{})["titleId"].(float64))
		textName := int(item.(map[string]interface{})["textName"].(float64))
		lastTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("res: %d,%s,%s,%d,%d,%d,%d,%d,%d,%d\n", effBgId, pName, faceName, sid, with, propId, pid, hight, titleId, textName)

		result, err := db.Exec("INSERT INTO dynamic (effBgId, pName, faceName, sid, width, propId, pid, hight, titleId, textName, lastTime) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", effBgId, pName, faceName, sid, with, propId, pid, hight, titleId, textName, lastTime)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.LastInsertId())
		elapsed := time.Since(t1)
		fmt.Println(elapsed)

	}

}
