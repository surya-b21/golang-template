package service

import "github.com/redis/go-redis/v9"

var Cache *redis.Client

func InitCache() {
	if Cache != nil {
		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	Cache = rdb
}
