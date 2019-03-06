package tokenizer

import (
	"fmt"
	"time"

	"github.com/Sach97/ninshoo/context"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

//RedisClient holds our redis connexion
type RedisClient struct {
	redisdb *redis.Client
}

//NewRedisClient creates a new redis connexion
func NewRedisClient(config *context.Config) *RedisClient {
	if config.RedisPassword == "" {
		panic("You must fill REDISPASSWORD env variable with your redis password")
	}

	redisdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,      // use default Addr
		Password: config.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
	return &RedisClient{
		redisdb: redisdb,
	}
}

//Ping pings redis to see if we are connected
func (client *RedisClient) Ping() error {
	pong, err := client.redisdb.Ping().Result()
	fmt.Println(pong, err)
	return err
}

// GenerateToken generate a random string
func (client *RedisClient) GenerateToken(userID string) (string, error) {
	id := uuid.NewV4()
	exp := time.Duration(600 * time.Second) // 10 minutes

	err := client.redisdb.Set(id.String(), userID, exp).Err()
	return id.String(), err
}

// GetUserID token retrieves the value of the token from our storage
func (client *RedisClient) GetUserID(token string) (string, error) {
	val, err := client.redisdb.Get(token).Result()
	return val, err
}
