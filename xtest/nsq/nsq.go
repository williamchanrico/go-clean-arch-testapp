package nsq

import (
	"github.com/nsqio/go-nsq"
)

// Xtest for postgresql
type Xtest struct {
	producer *nsq.Producer
}

// New client for postgres xtest backend
func New(producer *nsq.Producer) *Xtest {
	return &Xtest{
		producer: producer,
	}
}

// TestPingDefaultAddr ping postgres backend
func (x *Xtest) TestPingDefaultAddr() (string, error) {
	err := x.producer.Ping()
	if err != nil {
		return "Ping failed!", err
	}
	return "Ping success!", nil
}

// TestPingNewAddr ping nsqd backend
// e.g. addr: "127.0.0.1:4150"
func (x *Xtest) TestPingNewAddr(addr string) (string, error) {
	nsqConfig := nsq.NewConfig()
	nsqProducerClient, err := nsq.NewProducer(addr, nsqConfig)
	if err != nil {
		return "NSQd new producer client failed!", err
	}

	err = nsqProducerClient.Ping()
	nsqProducerClient.Stop()
	if err != nil {
		return "Ping failed!", err
	}
	return "Ping success!", nil
}
