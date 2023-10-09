package order

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	_ "github.com/olivere/elastic/v7"
	"log"
	"strconv"
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
}

func MakeOrderData() {
	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.20.166.56:9200"),
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

	num := 200000
	// 循环写入订单信息
	for i := 1; i <= num; i++ {
		chargeId, amount := GetCharge()
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
		}

		indexResp, err := client.Index().
			Index(indexName).
			// Id(strconv.Itoa(id)).
			BodyJson(order).
			Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Indexed book: %s\n", indexResp.Id)

	}
}
