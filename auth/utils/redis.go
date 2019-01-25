package utils

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	redisdb *redis.Client
}

func (client *Client) Ping() error {
	pong, err := client.redisdb.Ping().Result()
	fmt.Println(pong, err)
	return err
}

func New() *Client {

	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	return &Client{
		redisdb: redisdb,
	}
}

func (client *Client) GenerateString() (string, error) {
	id, _ := uuid.NewV4()
	exp := time.Duration(600 * time.Second) // 10 minutes

	fmt.Printf("UUIDv4: %s\n", id)
	err := client.redisdb.Set(id.String(), "testuserid", exp).Err()
	return id.String(), err
}

func (client *Client) GetToken(key string) (string, error) {
	val, err := client.redisdb.Get(key).Result()
	return val, err
}
