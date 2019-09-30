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

// TestPing ping redis backend
func (x *Xtest) TestPing() (string, error) {
	redisOpts := x.client.Options()
	result, err := x.client.Ping().Result()

	ret := fmt.Sprintf("%v [%v]\n", result, redisOpts.Addr)
	return ret, err
}
