package utils

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var WeekDayMap = map[string]string{
	"Monday":    "1",
	"Tuesday":   "2",
	"Wednesday": "3",
	"Thursday":  "4",
	"Friday":    "5",
	"Saturday":  "6",
	"Sunday":    "7",
}

type DataInfo struct {
	nc         int64
	pa         float64
	pn         int64
	new_user   int64
	new_amount float64
	old_user   int64
	old_amount float64
	lg         int64
}

func StartPush() {

	day := "2023-02-14 00:00:00"
	var ip = "cTEfZKKC23ehnIwX"
	// var ip ="172.17.64.7"
	var exchangeRate = 7.0635

	now := time.Now().Unix()
	fmt.Println("now是时间戳", now)
	StartTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	EndTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 59, time.Local).Unix()

	StartTimeStr := time.Unix(StartTime, 0).Format("2006-01-02 15:04:05")
	EndTimeStr := time.Unix(EndTime, 0).Format("2006-01-02 15:04:05")
	fmt.Println("StartTime是时间戳", StartTime, "| StartTimeStr是时间字符串", StartTimeStr)
	fmt.Println("EndTime是时间戳", EndTime, "| EndTimeStr是时间字符串", EndTimeStr)

	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", day, time.Local)
	fmt.Println("t1是时间类型", t1)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", StartTimeStr, time.Local)
	fmt.Println("t2是时间类型", t2)

	fmt.Println("================================================")
	for {
		startT1 := t1.Format("2006-01-02 15:04:05")
		fmt.Println("startT1=", t1.Format("2006-01-02 15:04:05"))
		fmt.Println("今天周", WeekDayMap[t1.Weekday().String()])
		week := WeekDayMap[t1.Weekday().String()]
		fmt.Println("日期是", t1.Format("2006-01-02"))
		date := t1.Format("2006-01-02")
		formatT1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 59, 59, 59, time.Local).Unix()
		endT1 := time.Unix(formatT1, 0).Format("2006-01-02 15:04:05")
		fmt.Println("endT1=", time.Unix(formatT1, 0).Format("2006-01-02 15:04:05"))
		cur := DbCur()
		PushDbData(cur, week, date, startT1, endT1, ip, exchangeRate)

		fmt.Println("================================================")
		time.Sleep(time.Duration(2) * time.Second)

		t1 = t1.AddDate(0, 0, 1)

		if t1.After(t2) || t1.Equal(t2) {
			break
		}

	}
}

func PushDbData(cur *sql.DB, week string, date string, startT1 string, endT1 string, ip string, exchangeRate float64) {
	var QueryStr = fmt.Sprintf("SELECT a.nc, a.pa,  a.pn,c.new_user, c.new_amount,c.old_user,c.old_amount,a.lg FROM (SELECT DATE(r.log_time) as dd, log_channel channel_id,CONVERT(SUM(if(log_type=23, log_previous , 0)),  SIGNED) pa,CONVERT(SUM(if(log_type=23, log_result, 0)),  SIGNED) pn,CONVERT(SUM(if(log_type=9, log_now , 0)),  SIGNED) lg,CONVERT(SUM(if(log_type=3, log_previous , 0)),  SIGNED) nc FROM log_statistic_result r WHERE log_time>='%s' AND log_time<='%s' GROUP BY dd) a LEFT JOIN (SELECT b.last_date, SUM(CASE WHEN x = 0 THEN 1 ELSE 0 END) new_user, SUM(CASE WHEN j = 0 THEN 1 ELSE 0 END) old_user, SUM(CASE WHEN x = 0 THEN b.post_amount ELSE 0 END) new_amount, SUM(CASE WHEN j = 0 THEN b.post_amount ELSE 0 END) old_amount FROM (SELECT date(post_time) last_date,sum(post_amount) post_amount,if((select id from pay_action where pay_status=4 AND pay_amount>0 AND pay_user=p.pay_user AND date(player_create_time) = date(p.post_time) limit 1), 0, 1) x, if((select id from pay_action where pay_status=4 AND pay_amount>0 AND pay_user=p.pay_user AND date(player_create_time) = date(p.post_time) limit 1), 1, 0) j FROM pay_action p WHERE p.pay_status = 4 AND p.pay_amount > 0 AND p.post_time BETWEEN '%s' AND '%s' GROUP BY pay_user, last_date ,channel_id) b GROUP BY b.last_date) c ON a.dd = c.last_date", startT1, endT1, startT1, endT1)
	rows, err := cur.Query(QueryStr)
	if err != nil {
		fmt.Println("数据库执行更新出错")
	}
	fmt.Println("推送DB数据!!")

	var datainfo DataInfo
	for rows.Next() {
		err = rows.Scan(&datainfo.nc, &datainfo.pa, &datainfo.pn, &datainfo.new_user, &datainfo.new_amount, &datainfo.old_user, &datainfo.old_amount, &datainfo.lg)
		if err != nil {
			log.Println(err)
		}
		amount, _ := strconv.Atoi(fmt.Sprintf("%1.0f", datainfo.pa*exchangeRate))
		amount1, _ := strconv.Atoi(fmt.Sprintf("%1.0f", datainfo.new_amount*exchangeRate))
		amount2, _ := strconv.Atoi(fmt.Sprintf("%1.0f", datainfo.old_amount*exchangeRate))

		fmt.Printf("新增注册: %d, 充值金额: %d, 充值人数: %d, 新增用户: %d, 新付费金额: %d, 旧付费人数: %d, 旧付费金额: %d, 登录用户数: %d \n", datainfo.nc, amount, datainfo.pn, datainfo.new_user, amount1, datainfo.old_user, amount2, datainfo.lg)

		rb := fmt.Sprintf(`{
							"registered": "%d",
							"amount": "%d",
							"pay_user": "%d",
							"new_pay_user": "%d",
							"new_amount": "%d",
							"old_pay_user": "%d",
							"old_amount": "%d",
							"land": "%d",
							"ip": "%s",
							"date":"%s",
							"week":"%s"
							}`, datainfo.nc, amount, datainfo.pn, datainfo.new_user, amount1, datainfo.old_user, amount2, datainfo.lg, ip, date, week)
		fmt.Println("请求的参数", rb)
		var jsonStr = []byte(rb)
		url := "http://mhtx-statistics.ios.shyouai.com:8001/game/statistics/addStatistics"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}

func DbCur() *sql.DB {
	// 寻仙港澳台
	dsn := "ef946573-9609-40aa-a306-ac25d24a948c:Q952BJ9EEG1XsJWb@tcp(login.yuanhui.work:30004)/mohuan"

	// 异兽国服
	// dsn := "23bf980d-7153-4e1f-9061-a2137861aac7:Swq0cJ7hwluBvReg@tcp(login.yuanhui.work:30006)/gms_ys_200103"

	// 寻仙国服
	// dsn := "gms:7FTFGbNPbY@tcp(sh-cdb-j14oh158.sql.tencentcdb.com:58409)/mohuan"


	Db, err := sql.Open("mysql", dsn) //Open函数第一个参数就是驱动的名字（不能随意写的）
	if err != nil {                   //有一个坑，如果连接数据库配置写错了，不会在此处报错
		log.Printf("打开数据库失败,err:%v\n", err)
		return nil
	}
	err = Db.Ping() //ping一下没报错证明连接数据库成功
	if err != nil {
		log.Printf("ping 失败,err :%v\n", err)
		return nil
	}
	log.Println("连接成功")
	return Db

}
