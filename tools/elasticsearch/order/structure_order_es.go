package order

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// GetOrderNo 随机生成订单号
func GetOrderNo() string {
	order_no := uuid.New()
	fmt.Println(order_no)
	return order_no.String()
}

// GetChannel 随机生成渠道
func GetChannel() string {
	channels := []string{
		"趣乐渠道",
		"vivo渠道",
		"oppo渠道",
		"华为渠道",
		"小米渠道",
		"鲸旗渠道",
		"元起渠道",
		"游爱渠道",
		"光卡渠道",
		"宜搜渠道",
		"爱奇艺渠道",
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomIndex := randomGenerator.Intn(len(channels))
	randomChannel := channels[randomIndex]

	fmt.Println(randomChannel)

	return randomChannel

}

// GetPayPlatform 随机生成充值平台
func GetPayPlatform() string {
	platforms := []string{
		"IOS",
		"Andriod",
		"Web",
		"Mac",
		"Windows",
		"Ubuntu",
		"Mint",
		"CentOS",
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomIndex := randomGenerator.Intn(len(platforms))
	randomPlatform := platforms[randomIndex]

	fmt.Println(randomPlatform)

	return randomPlatform
}

func GetTime() string {
	// 获取当前时间
	now := time.Now()

	// 生成随机时间
	minDate := now.AddDate(-1, 0, 0) // 截至到今天一年前
	maxDate := now

	// 获取时间范围
	minUnix := minDate.Unix()
	maxUnix := maxDate.Unix()

	// 生成随机时间戳
	randUnix := rand.Int63n(maxUnix-minUnix) + minUnix

	// 转换为Time对象
	randomTime := time.Unix(randUnix, 0)

	// 格式化为字符串
	// randomTimeString := randomTime.Format("2006-01-02 15:04:05")
	randomTimeString := randomTime.Format("2006-01-02T15:04:05Z07:00")

	fmt.Println(randomTimeString)

	return randomTimeString
}

func GetCharge() (string, int64) {
	charges := []map[string]int64{
		{"1001": 1},
		{"1002": 6},
		{"1003": 30},
		{"1004": 68},
		{"1005": 128},
		{"1006": 198},
		{"1007": 328},
		{"1008": 648},
		{"1010": 1000},
		{"1020": 2000},
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomIndex := randomGenerator.Intn(len(charges))
	randomCharge := charges[randomIndex]

	var k string
	var v int64
	for key, value := range randomCharge {
		k = key
		v = value
	}
	fmt.Println(k, v)
	return k, v
}

func GetServerId() int {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomIndex := randomGenerator.Intn(10) + 1

	fmt.Println(randomIndex)
	return randomIndex
}

func GetProvince() (string, string) {

	CodeProvince := map[string]string{
		"CN-AH": "安徽省",
		"CN-BJ": "北京市",
		"CN-CQ": "重庆市",
		"CN-FJ": "福建省",
		"CN-GD": "广东省",
		"CN-GS": "甘肃省",
		"CN-GX": "广西壮族自治区",
		"CN-GZ": "贵州省",
		"CN-HI": "海南省",
		"CN-HE": "河北省",
		"CN-HL": "黑龙江省",
		"CN-HA": "河南省",
		"CN-HK": "香港特别行政区",
		"CN-HB": "湖北省",
		"CN-HN": "湖南省",
		"CN-JS": "江苏省",
		"CN-JX": "江西省",
		"CN-JL": "吉林省",
		"CN-LN": "辽宁省",
		"CN-MO": "澳门特别行政区",
		"CN-NM": "内蒙古自治区",
		"CN-NX": "宁夏回族自治区",
		"CN-QH": "青海省",
		"CN-SN": "陕西省",
		"CN-SC": "四川省",
		"CN-SD": "山东省",
		"CN-SH": "上海市",
		"CN-SX": "山西省",
		"CN-TJ": "天津市",
		"CN-XJ": "新疆维吾尔自治区",
		"CN-XZ": "西藏自治区",
		"CN-YN": "云南省",
		"CN-ZJ": "浙江省",
		"CN-TW": "台湾省",
	}
	// var provinces = []string{
	// 	"北京市",
	// 	"天津市",
	// 	"河北省",
	// 	"山西省",
	// 	"辽宁省",
	// 	"吉林省",
	// 	"黑龙江省",
	// 	"上海市",
	// 	"江苏省",
	// 	"浙江省",
	// 	"安徽省",
	// 	"福建省",
	// 	"江西省",
	// 	"山东省",
	// 	"河南省",
	// 	"湖北省",
	// 	"湖南省",
	// 	"广东省",
	// 	"广西壮族自治区",
	// 	"海南省",
	// 	"重庆市",
	// 	"四川省",
	// 	"贵州省",
	// 	"云南省",
	// 	"西藏自治区",
	// 	"陕西省",
	// 	"甘肃省",
	// 	"青海省",
	// 	"宁夏回族自治区",
	// 	"新疆维吾尔自治区",
	// 	"香港特别行政区",
	// 	"澳门特别行政区",
	// }

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomIndex := randomGenerator.Intn(len(CodeProvince))
	i := 0
	var randomCode, randomProvince string
	for code, province := range CodeProvince {
		if i == randomIndex {
			randomCode = code
			randomProvince = province
			break
		}
		i++
	}

	fmt.Println(randomCode, randomProvince)

	return randomCode, randomProvince

}

func GetIP() string {
	// 生成四个随机数分别代表 IP 的四个部分
	part1 := rand.Intn(256)
	part2 := rand.Intn(256)
	part3 := rand.Intn(256)
	part4 := rand.Intn(256)

	// 将四个数拼接成 IP 地址的形式
	ip := strconv.Itoa(part1) + "." + strconv.Itoa(part2) + "." + strconv.Itoa(part3) + "." + strconv.Itoa(part4)
	fmt.Println(ip)
	return ip
}
