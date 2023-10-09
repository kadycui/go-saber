package curd

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

var esClient *elastic.Client

func init() {
	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.20.166.56:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}
	esClient = client

}

func GetEsClient() *elastic.Client {
	return esClient
}

func CreateIndex(esClient *elastic.Client, indexName string, mapping string) {
	// 检查索引是否存在
	exists, err := esClient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Fatalf("检查索引是否存在失败：%s", err)
	}

	if exists {
		// 索引存在，跳过检查
		log.Printf("索引%s已存在\n", indexName)
	} else {
		// 索引不存在，创建索引
		createIndex, err := esClient.CreateIndex(indexName).BodyString(mapping).Do(context.Background())
		if err != nil {
			log.Fatalf("创建索引失败：%s", err)
		}
		if !createIndex.Acknowledged {
			log.Fatalf("创建索引未被确认")
		}
		log.Printf("索引%s已创建\n", indexName)
	}

}
