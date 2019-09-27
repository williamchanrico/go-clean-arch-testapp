package xtest

import (
	"github.com/go-redis/redis"
	"github.com/williamchanrico/xtest/log"
	"github.com/williamchanrico/xtest/server"
)

// Flags for xtest
type Flags struct {
	HTTPAddress  string
	RedisAddress string
	LogLevel     string
}

// Run is the entry point of xtest
func Run(flags Flags) (int, error) {
	log.Debugf("xtest runs with flags: %+v", flags)

	redisClient := redis.NewClient(&redis.Options{
		Addr: flags.RedisAddress,
	})

	s := server.Server{
		HTTPAddress: flags.HTTPAddress,
		Redis:       redisClient,
	}
	log.Infof("xtest HTTP server listens on: %v\n", s.HTTPAddress)
	if err := s.Run(); err != nil {
		return 1, err
	}

	return 0, nil
}
