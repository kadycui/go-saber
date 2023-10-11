package order

import (
	"context"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/olivere/elastic/v7"
	"gopkg.in/ini.v1"
)

type Order struct {
	OrderNo     string `json:"order_no"`
	Amount      int64  `json:"amount"`
	GameCoin    int64  `json:"game_coin"`
	ChargeId    string `json:"charge_id"`
	ServerId    int    `json:"server_id"`
	PlayerId    int64  `json:"player_id"`
	PlayerName  string `json:"player_name"`
	PayPlatform string `json:"pay_platform"`
	Channel     string `json:"channel"`
	CreateTime  string `json:"date"`
	IP          string `json:"ip"`
	Province    string `json:"province"`
}

func GetPath() string {
	// 获取当前执行文件绝对路径（go run）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}
	p := path.Dir(filename)
	// 获取父级目录
	parentPath := filepath.Join(p, "../")
	return parentPath
}

type EsConf struct {
	Name string
	Host string
}

func MakeOrderData() {

	p := GetPath()
	// 加载INI文件
	iniPath := p + "/config/conf.ini"
	cfg, err := ini.Load(iniPath)
	if err != nil {
		panic(err)
	}
	// 获取指定section的值
	section := cfg.Section("es0")
	name := section.Key("name").String()
	host := section.Key("host").String()

	fmt.Println(name, host)

	esconf := EsConf{
		Name: name,
		Host: host,
	}

	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL(esconf.Host),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个book索引（如果还不存在）
	indexName := "go_pay_order"
	// 检查索引是否存在
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Fatalf("检查索引是否存在失败：%s", err)
	}

	if exists {
		// 索引存在，跳过检查
		fmt.Printf("索引%s存在\n", indexName)
	} else {
		// 索引不存在，创建索引
		createIndex, err := client.CreateIndex(indexName).BodyString(`
			{
				"settings": {
					"number_of_shards": 1,
					"number_of_replicas": 0
				}
			}
		`).Do(context.Background())
		if err != nil {
			log.Fatalf("创建索引失败：%s", err)
		}
		if !createIndex.Acknowledged {
			log.Fatalf("创建索引未被确认")
		}
		fmt.Printf("索引%s已创建\n", indexName)
	}

	num := 10
	// 循环写入订单信息
	for i := 1; i <= num; i++ {
		chargeId, amount := GetCharge()
		code, _ := GetProvince()
		order := Order{
			OrderNo:     GetOrderNo(),
			Amount:      amount,
			GameCoin:    amount * 10,
			ChargeId:    chargeId,
			ServerId:    GetServerId(),
			PlayerId:    int64(GetServerId()<<20) + int64(i),
			PlayerName:  "玩家-" + strconv.Itoa(i),
			PayPlatform: GetPayPlatform(),
			Channel:     GetChannel(),
			CreateTime:  GetTime(),
			IP:          GetIP(),
			Province:    code,
		}

		indexResp, err := client.Index().
			Index(indexName).
			// Id(strconv.Itoa(id)).
			BodyJson(order).
			Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Indexed Order Id: %s\n", indexResp.Id)

	}
}
