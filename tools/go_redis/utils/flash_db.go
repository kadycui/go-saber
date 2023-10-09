package utils

import "fmt"

func FlashDb() {
	client := NewRedisClient()

	err := client.FlushDB()
	if err != nil {
		fmt.Println("FlushDB error:", err)
		return
	} else {
		fmt.Println("FlushDB Success")
	}
}

func FlashAll() {
	client := NewRedisClient()
	err := client.FlushAll()
	if err != nil {
		fmt.Println("FlushAll error:", err)
		return
	} else {
		fmt.Println("FlushAll Success")
	}

}
