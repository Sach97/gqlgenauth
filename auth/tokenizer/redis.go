package tokenizer

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisClient struct {
	redisdb *redis.Client
}

func NewRedisClient() *RedisClient {

	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	return &RedisClient{
		redisdb: redisdb,
	}
}

func (client *RedisClient) Ping() error {
	pong, err := client.redisdb.Ping().Result()
	fmt.Println(pong, err)
	return err
}

//TODO: rename this to GenerateToken

func (client *RedisClient) GenerateString() (string, error) {
	id, _ := uuid.NewV4()
	exp := time.Duration(600 * time.Second) // 10 minutes

	fmt.Printf("UUIDv4: %s\n", id)
	err := client.redisdb.Set(id.String(), "testuserid", exp).Err()
	return id.String(), err
}

func (client *RedisClient) GetToken(key string) (string, error) {
	val, err := client.redisdb.Get(key).Result()
	return val, err
}
