package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/kadycui/go-saber/tools/gorm_postgresql/model"
)

func Start() {
	var wg sync.WaitGroup

	// 设置等待的协程数量
	wg.Add(5)

	// 创建并启动5个协程
	for i := 0; i < 5; i++ {
		go CreateBookData(&wg)
	}

	// 等待所有协程执行完毕
	wg.Wait()
	fmt.Println("所有任务完成!!!")
}

func CreateBookData(wg *sync.WaitGroup) {
	// 在函数退出时通知WaitGroup，表示当前协程已完成
	defer wg.Done()
	db := GetDB()
	num := 100
	// 循环遍历输出书籍信息
	for i := 1; i <= num; i++ {
		book := model.Book{
			No:       time.Now().UnixNano(),
			Title:    BookName(),
			Author:   GetFullName(),
			Price:    Price(),
			Rating:   Rating(),
			Date:     RandomTimeStr(),
			Category: Category(),
			Publish:  Publish(),
		}

		result := db.Create(&book)
		fmt.Println(result)
	}
}
