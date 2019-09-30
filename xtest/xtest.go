package xtest

import "github.com/williamchanrico/xtest/xtest/redis"

// Service struct for xtest
type Service struct {
	redis *redis.Xtest
}

// New xtest service
func New(redis *redis.Xtest) *Service {
	return &Service{
		redis: redis,
	}
}

// TestRedis runs the test for redis backend
func (s *Service) TestRedis() (string, error) {
	return s.redis.TestPing()
}
