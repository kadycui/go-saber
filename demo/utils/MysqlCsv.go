package utils

// 从Mysql中导出数据到CSV文件。

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// 设置配置文件的名字
	viper.SetConfigName("db")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath("./")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(viper.Get("mysql")) // map[port:3306 url:127.0.0.1]

}

var (
	tables = []string{"pay_action"}
	count  = len(tables)
	ch     = make(chan bool, count)
)

func RechargeToCsv() {
	sDate := flag.StringP("sdate", "s", "2028-01-01", "开始日期")
	eDate := flag.StringP("edate", "e", "2028-01-02", "结束日期")
	dBConf := flag.StringP("conf", "c", "./", "数据库配置")
	flag.Parse()
	fmt.Println(*sDate, *eDate, *dBConf)
	sdate := *sDate + " 00:00:00"
	edate := *eDate + " 23:59:59"
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", sdate, time.Local)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", edate, time.Local)

	day := SubDays(t1, t2)
	if day > 31 {
		fmt.Println("时间区间超过一个月, 重新设置时间区间!")
		return

	}
	if day == 0 {
		fmt.Println("请重新输入时间区间 格式: 2023-01-02 2023-01-31 ")
		return

	}
	s := time.Now()
	fmt.Println("导出数据的日期范围: ", sdate, "--", edate)
	// 寻仙-内网
	p := viper.Get("mysql.port")
	port := fmt.Sprintf("%d", p)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		viper.Get("mysql.user"), viper.Get("mysql.password"), viper.Get("mysql.url"), port, viper.Get("mysql.name"))
	db, err := sql.Open("mysql", connStr)

	// 寻仙-港澳台
	// db, err := sql.Open("mysql", "ef946573-9609-40aa-a306-ac25d24a948c:Q952BJ9EEG1XsJWb@tcp(login.yuanhui.work:30004)/mohuan?charset=utf8")
	// defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	for _, table := range tables {
		go querySQL(db, table, sdate, edate, ch)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	elapsed := time.Since(s)
	fmt.Println("导出完成!!, 总耗时:", elapsed)
}

func querySQL(db *sql.DB, table string, sdate string, edate string, ch chan bool) {
	fmt.Println("数据库表：", table)
	rows, _ := db.Query(fmt.Sprintf("SELECT * from %s where last_time between  '%s'  and  '%s' and pay_status = 4;", table, sdate, edate))

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	//values：一行的所有值，长度==列数
	values := make([]sql.RawBytes, len(columns))
	// print(len(values))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	totalValues := [][]string{}
	for rows.Next() {
		var s []string
		err = rows.Scan(scanArgs...) //把每行的内容添加到scanArgs，也添加到了values
		if err != nil {
			panic(err.Error())
		}

		for _, v := range values {
			s = append(s, string(v))
			// print(len(s))
		}
		totalValues = append(totalValues, s)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	writeToCSV(table+".csv", columns, totalValues)
	ch <- true
}

func writeToCSV(file string, columns []string, totalValues [][]string) {
	// fmt.Println(columns)
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	w := csv.NewWriter(f)
	for a, i := range totalValues {
		if a == 0 {
			w.Write(columns)
			w.Write(i)
		} else {
			// fmt.Println(i)
			w.Write(i)
		}
	}
	w.Flush()
	fmt.Println("生成文件：", file)
}

func SubDays(t1, t2 time.Time) (day int) {
	timestamp1 := t1.Unix()
	timestamp2 := t2.Unix()

	day = int((timestamp2 - timestamp1) / 86400)

	fmt.Println("时间间隔是: ", day+1)

	return day + 1

}
