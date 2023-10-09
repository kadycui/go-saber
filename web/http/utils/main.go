package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/kadycui/go-saber/web/http/utils/logger"

	"gopkg.in/ini.v1"
)

func init() {
	p := GetPath()
	logPath := p + "/log/"
	// 初始化日志文件
	logger.InitLogger(logPath)

}

type RequestData struct {
	Sign string `json:"sign"`
	Time int64  `json:"time"`
}

type GmsInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
}

func GetPath() string {
	// 获取当前执行文件绝对路径（go run）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}
	dir := path.Dir(filename)
	return dir
}

func ReadIni(gmsId string) (string, string) {
	p := GetPath()
	fmt.Println(p)
	// 加载INI文件
	iniPath := p + "/conf.ini"
	cfg, err := ini.Load(iniPath)
	if err != nil {
		panic(err)
	}
	// 获取指定section的值
	section := cfg.Section(gmsId)
	host := section.Key("host").String()
	key := section.Key("key").String()
	name := section.Key("name").String()

	gms := GmsInfo{
		Id:   gmsId,
		Name: name,
		Host: host,
	}

	jsonData, err := json.Marshal(gms)
	if err != nil {
		fmt.Println("转换为 JSON 出错:", err)
		panic(err)
	}

	logger.Info(string(jsonData))

	return host, key
}

func GetMd5(ts int64, key string) string {
	// 将密钥和时间戳连接在一起
	data := fmt.Sprintf("%s%d", key, ts)

	// 计算 MD5 值
	hash := md5.Sum([]byte(data))
	md5Value := hex.EncodeToString(hash[:])
	return md5Value
}

func WriteCsv(data string) {
	// 将body内容解析为二维字符串切片
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 创建CSV文件并写入数据
	p := GetPath()
	csvPath := p + "/csv/server.csv"
	file, err := os.Create(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func GetCsv(host string, key string) {
	timestamp := time.Now().Unix()
	sign := GetMd5(timestamp, key)
	logger.Info(fmt.Sprintf("%s--%s--%d", host, sign, timestamp))
	url := fmt.Sprintf("%s?sign=%s&time=%d", host, sign, timestamp)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET request failed:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭失败!")
		}
	}(response.Body)

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}

	// 输出响应内容
	fmt.Println(string(body))
	fmt.Println(reflect.TypeOf(string(body)))
	WriteCsv(string(body))

}

func main() {
	// host := "http://xx-center-tw.lsqy.tw/client_api/server_csv"
	//host := "http://xx-center-tw.lsqy.tw/client_api/channel_csv"
	//key := "LBrEjqRCnqsf5mWb"
	//GetCsv(host, key)
	gmsId := "170103"
	logger.Info(gmsId)
	host, key := ReadIni(gmsId)
	host = host + "/client_api/server_csv"
	GetCsv(host, key)

}
