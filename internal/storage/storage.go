package storage

import (
	"os"

	"github.com/redis/go-redis/v9"
)

type WhiteList struct {
	client *redis.Client
}

func NewWhiteList() *WhiteList {
	return &WhiteList{
		client: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_ADDR"),
		}),
	}
}
