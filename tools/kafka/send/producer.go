package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	type User struct {
		Id   int
		Name string
	}

	user := &User{
		1,
		"A",
	}

	str, err2 := json.Marshal(user)
	if err2 != nil {
		fmt.Printf("err: %v\n", err2)
	}

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("172.20.166.56:9082"),
		Topic:    "quickstart",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: str,
		},
		// kafka.Message{
		// 	Key:   []byte("Key-B"),
		// 	Value: []byte("One!"),
		// },
		// kafka.Message{
		// 	Key:   []byte("Key-C"),
		// 	Value: []byte("Two!"),
		// },
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}

