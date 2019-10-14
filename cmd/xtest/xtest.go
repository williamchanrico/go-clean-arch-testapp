package xtest

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/nsqio/go-nsq"
	"github.com/williamchanrico/xtest/log"
	"github.com/williamchanrico/xtest/server"
	"github.com/williamchanrico/xtest/xtest"
)

// Flags for xtest
type Flags struct {
	HTTPAddress  string
	GRPCAddress  string
	RedisAddress string
	PostgresDSN  string
	NSQDAddress  string
	LogLevel     string
}

// Run is the entry point of xtest
func Run(flags Flags) (int, error) {
	log.Debugf("xtest runs with flags: %+v", flags)

	redisClient := redis.NewClient(&redis.Options{Addr: flags.RedisAddress})

	postgresClient, err := sql.Open("postgres", flags.PostgresDSN)
	if err != nil {
		return 1, err
	}

	nsqConfig := nsq.NewConfig()
	nsqProducerClient, err := nsq.NewProducer(flags.NSQDAddress, nsqConfig)
	if err != nil {
		return 1, err
	}

	xtestSvc := xtest.New(redisClient, postgresClient, nsqProducerClient)

	s := server.Server{
		HTTPAddress: flags.HTTPAddress,
		GRPCAddress: flags.GRPCAddress,
		Xtest:       xtestSvc,
	}
	log.Infof("Starging xtest servers")
	if err := s.Run(context.Background()); err != nil {
		return 1, err
	}

	return 0, nil
}
