package xtest

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/nsqio/go-nsq"

	nsqBackend "github.com/williamchanrico/xtest/xtest/nsq"
	postgresBackend "github.com/williamchanrico/xtest/xtest/postgres"
	redisBackend "github.com/williamchanrico/xtest/xtest/redis"
)

// Service struct for xtest
type Service struct {
	redis    *redisBackend.Xtest
	postgres *postgresBackend.Xtest
	nsq      *nsqBackend.Xtest
}

// New xtest service
func New(redisClient *redis.Client, postgresClient *sql.DB, nsqProducerClient *nsq.Producer) *Service {
	redisBackend := redisBackend.New(redisClient)
	postgresBackend := postgresBackend.New(postgresClient)
	nsqBackend := nsqBackend.New(nsqProducerClient)

	return &Service{
		redis:    redisBackend,
		postgres: postgresBackend,
		nsq:      nsqBackend,
	}
}

// TestRedis runs the test for redis backend
func (s *Service) TestRedisDefaultAddr() (string, error) {
	return s.redis.PingDefaultAddr()
}

func (s *Service) TestRedisNewAddr(addr string) (string, error) {
	return s.redis.PingNewAddr(addr)
}

// TestPostgresDefaultDSN runs the test for redis backend
func (s *Service) TestPostgresDefaultDSN() (string, error) {
	return s.postgres.TestPingDefaultDSN()
}

// TestPostgresNewDSN runs the test for redis backend
func (s *Service) TestPostgresNewDSN(dsn string) (string, error) {
	return s.postgres.TestPingNewDSN(dsn)
}

// TestNSQDefaultAddr runs the test for redis backend
func (s *Service) TestNSQDefaultAddr() (string, error) {
	return s.nsq.TestPingDefaultAddr()
}

// TestNSQNewAddr runs the test for redis backend
func (s *Service) TestNSQNewAddr(addr string) (string, error) {
	return s.nsq.TestPingNewAddr(addr)
}
