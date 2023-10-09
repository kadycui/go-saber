package order

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
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
