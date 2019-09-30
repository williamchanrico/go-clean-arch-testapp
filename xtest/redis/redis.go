package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Xtest for redis
type Xtest struct {
	client *redis.Client
}

// New client for redis xtest backend
func New(client *redis.Client) *Xtest {
	return &Xtest{
		client: client,
	}
}

// PingDefaultAddr pings the default redis
func (x *Xtest) PingDefaultAddr() (string, error) {
	redisOpts := x.client.Options()
	result, err := x.client.Ping().Result()

	ret := fmt.Sprintf("%v [%v]\n", result, redisOpts.Addr)
	return ret, err
}

// PingNewAddr pings redis located on addr
func (x *Xtest) PingNewAddr(addr string) (string, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: addr})
	result, err := redisClient.Ping().Result()
	redisClient.Close()
	return result, err
}
