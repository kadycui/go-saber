package redis

import (
	"gostart/libs/dict"
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	redisClient *redis.Client
	prefix      string
}

func NewClient(redisConfig map[string]interface{}) (*Client, error) {
	prefix := dict.GetString(redisConfig, "prefix")
	host := dict.GetString(redisConfig, "host")
	port := dict.GetString(redisConfig, "port")
	pass := dict.GetString(redisConfig, "auth_pass")
	db := dict.GetInt(redisConfig, "db")

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Client{
		redisClient: client,
		prefix:      prefix,
	}, nil
}

func (client *Client) GetKey(key string) string {
	return client.prefix + "." + key
}

func (client *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	key = client.GetKey(key)
	return client.redisClient.Set(key, value, expiration)
}

func (client *Client) Get(key string) *redis.StringCmd {
	key = client.GetKey(key)
	return client.redisClient.Get(key)
}

func (client *Client) HSet(key, field string, value interface{}) *redis.BoolCmd {
	key = client.GetKey(key)
	return client.redisClient.HSet(key, field, value)
}

func (client *Client) HGet(key, field string) *redis.StringCmd {
	key = client.GetKey(key)
	return client.redisClient.HGet(key, field)
}

func (client *Client) HGetAll(key string) *redis.StringStringMapCmd {
	key = client.GetKey(key)
	return client.redisClient.HGetAll(key)
}
