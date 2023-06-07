package utils

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(124.220.3.231:3309)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 指定动态表名称为 dynamic
func (Dynamic) TableName() string {
	return "dynamic"
}

// 创建动态表
func Migrate(db *gorm.DB) error {
	// err := db.Table("dynamic").AutoMigrate(&Dynamic{})
	err := db.AutoMigrate(&Dynamic{}, &User{})
	return err
}

// 定义动态表模型（表名为 dynamic）
type Dynamic struct {
	gorm.Model           // 增加created_at, updated_at, deleted_at字段
	ID         int       `gorm:"primaryKey;autoIncrement;column:id"` // ID 对应表中的 id 字段
	EffBgID    int       `gorm:"not null;column:effBgId"`            // EffBgID 对应表中的 effBgId 字段
	PName      string    `gorm:"not null;column:pName"`              // PName 对应表中的 pName 字段
	FaceName   string    `gorm:"column:faceName"`                    // FaceName 对应表中的 faceName 字段
	SID        int       `gorm:"column:sid"`                         // SID 对应表中的 sid 字段
	Width      int       `gorm:"column:width"`                       // Width 对应表中的 width 字段
	PropID     int       `gorm:"column:propId"`                      // PropID 对应表中的 propId 字段
	PID        int64     `gorm:"column:pid"`                         // PID 对应表中的 pid 字段
	Hight      int       `gorm:"column:hight"`                       // Hight 对应表中的 hight 字段
	TitleID    int       `gorm:"column:titleId"`                     // TitleID 对应表中的 titleId 字段
	TextName   int       `gorm:"column:textName"`                    // TextName 对应表中的 textName 字段
	LastTime   time.Time `gorm:"default:null;column:lastTime"`       // LastTime 对应表中的 lastTime 字段
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GormToDb() {
	db, err := InitDB()
	if err != nil {
		// 处理初始化数据库失败的情况
		log.Fatal(err)
	} else {
		log.Println("DB初始化完成!")
	}

	err = Migrate(db)
	if err != nil {
		// 处理创建表失败的情况
		log.Fatal(err)
	} else {
		log.Println("数据库建表完成!")
	}

}

func GormInsert() {
	db, err := InitDB()
	if err != nil {
		// 处理初始化数据库失败的情况
		log.Fatal(err)
	} else {
		log.Println("DB初始化完成!", db)
	}

	// 单个插入
	// user := User{Name: "吴承恩", Age: 49, Email: "wuchengen@ming.com"}
	// result := db.Create(&user)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	// 批量插入
	users := []User{
		{Name: "Tom", Age: 30, Email: "tom@example.com"},
		{Name: "Jerry", Age: 28, Email: "jerry@example.com"},
		{Name: "Mike", Age: 35, Email: "mike@example.com"},
	}

	result := db.Create(&users)
	if result.Error != nil {
		panic(result.Error)
	}

}

func GormSelect() {
	db, err := InitDB()
	if err != nil {
		// 处理初始化数据库失败的情况
		log.Fatal(err)
	} else {
		log.Println("DB初始化完成!", db)
	}

	/* 查询全部
	var users []User
	result := db.Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Age, user.Email, user.CreatedAt, user.UpdatedAt)
	}
	*/

	// 条件查询
	var users []User
	// result := db.Where("age > ?", 30).Find(&users)

	// result := db.Where("age BETWEEN ? AND ?", 20, 30).Find(&users)

	// result := db.Order("age desc").Find(&users)

	result := db.Offset(5).Limit(3).Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Age, user.Email, user.CreatedAt, user.UpdatedAt)
	}

}

func GormUpdate() {

	db, err := InitDB()
	if err != nil {
		// 处理初始化数据库失败的情况
		log.Fatal(err)
	} else {
		log.Println("DB初始化完成!", db)
	}

	var user User
	r1 := db.First(&user, "name = ?", "Tom")

	if r1.Error != nil {
		panic(r1.Error)
	}

	log.Println("修改前的数据", user.Name, user.Age)

	update_result := db.Model(&User{}).Where("name = ?", "Tom").Update("age", 35)
	if update_result.Error != nil {
		panic(update_result.Error)
	}

	r1 = db.First(&user, "name = ?", "Tom")

	if r1.Error != nil {
		panic(r1.Error)
	}

	log.Println("修改后的数据", user.Name, user.Age)

}

func GormDelete() {
	db, err := InitDB()
	if err != nil {
		// 处理初始化数据库失败的情况
		log.Fatal(err)
	} else {
		log.Println("DB初始化完成!", db)
	}

	result := db.Model(&User{}).Where("name = ?", "Tom").Delete(&User{})
	if result.Error != nil {
		panic(result.Error)
	}

}
