package messenger

import (
	"github.com/nsqio/go-nsq"
)

type Messenger interface {
	CreateConsumer(topic, channel string) (*nsq.Consumer, error)
	CreateProducer(address string) (*nsq.Producer, error)
}

type messenger struct {
	config *nsq.Config
}

func (m messenger) CreateConsumer(topic, channel string) (*nsq.Consumer, error) {
	consumer, consumerError := nsq.NewConsumer(topic, channel, m.config)

	return consumer, consumerError
}

func (m messenger) CreateProducer(address string) (*nsq.Producer, error) {
	producer, producerError := nsq.NewProducer(address, m.config)

	return producer, producerError
}

func New() Messenger {
	config := nsq.NewConfig()

	return messenger{
		config,
	}
}
