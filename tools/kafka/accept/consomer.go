package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	type User struct {
		Id   int
		Name string
	}

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"172.20.166.56:9082"},
		GroupID:        "consumer-group-id",
		Topic:          "quickstart",
		Partition:      0,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second 要 加 GroupID 才能自动提交
	})
	// r.SetOffset(2)

	ctx := context.Background()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		var user User
		err2 := json.Unmarshal(m.Value, &user)
		if err2 != nil {
			fmt.Printf("json unmarshal err: %v \n", err2)
		}

		fmt.Printf("--user-: %v\n", user)
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

	}

	// if err := r.Close(); err != nil {
	// 	log.Fatal("failed to close reader:", err)
	// }
}
