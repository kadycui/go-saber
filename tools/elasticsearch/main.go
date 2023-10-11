package main

import (
	"github.com/kadycui/go-saber/tools/elasticsearch/order"
)

func main() {
	// 生成es图书数据
	// book.MakeBookData()
	//order.GetOrderNo()
	//order.GetChannel()
	//order.GetPayPlatform()
	//order.GetTime()
	//order.GetCharge()
	//order.GetServerId()
	order.MakeOrderData()
	//client := curd.GetEsClient()

	// fmt.Println(client.Ping("http://172.20.166.56:9200").Do(context.Background()))

	// fmt.Println(client.ElasticsearchVersion("http://172.20.166.56:9200"))

	/*
			settings是修改分片和副本数的。
		    mappings是修改字段和类型的。
	*/

	//const mapping = `
	//		{
	//			"settings": {
	//				"number_of_shards": 1,
	//				"number_of_replicas": 0
	//			}
	//		}
	//	`

	// const body = `{
	// 	  "settings" : {
	// 		"index" : {
	// 		  "number_of_shards" : "5",
	// 		  "number_of_replicas" : "1",
	// 		  "version" : { }
	// 		}
	// 	  }
	// 	}`

	// curd.CreateIndex(esClient, "test2_index", body)

	// body := `{
	// 	"settings": {
	// 	  "index": {
	// 		"number_of_shards": "3",
	// 		"number_of_replicas": "0",
	// 		"version": {}
	// 	  }
	// 	},
	// 	"mappings": {
	// 	  "properties": {
	// 		"name": {
	// 		  "type": "keyword"
	// 		},
	// 		"age": {
	// 		  "type": "long"
	// 		}
	// 	  }
	// 	}
	//   }`

	//curd.CreateIndex(client, "struct_index", body)
	// curd.StructForm(client, "str_index")

	// curd.CreateIndex(client, "str_index", body)
	// curd.StrForm(client, "str_index")

	// curd.SingleUpdate(client, "str_index")
	// curd.MultipleUpdate(client, "book")

	// curd.TrueDelete(client, "book")
	// curd.IdQuery(client, "book", "539")

	// curd.MatchQuery(client)
	// curd.MatchQueryToJson(client)

	//curd.RangeQuery(client)
	// curd.BoolQuery(client)
	//curd.PrefixQuery2(client)

}
