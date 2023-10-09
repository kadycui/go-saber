package curd

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func TrueDelete(client *elastic.Client, index string) {
	_, err := client.Delete().Index("book").Id("2151").Do(context.Background())
	if err != nil {
		panic(err)
	}
}
