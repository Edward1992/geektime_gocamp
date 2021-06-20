package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/pioz/faker"
)

const KEY_AMOUNT = 100*100*5

func main() {
	redis_client := redis.NewClient(&redis.Options{
		DB: 7,
	})

	redis_client.FlushAll()
	memory_info, _ := redis_client.Info("memory").Result()
	fmt.Println(memory_info)
	for i := 0; i < KEY_AMOUNT; i++ {
		key := faker.StringWithSize(12)
		value := faker.StringWithSize(5120) // 10 20 50 100 200 1024 5120
		redis_client.Set(key, value, 0)
	}
	memory_info, _ = redis_client.Info("memory").Result()
	fmt.Println(memory_info)
}
