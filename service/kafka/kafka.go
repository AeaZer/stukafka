package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func InitKafkaConnect(conf *Config) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(),
		conf.Network, fmt.Sprintf("%s:%d", conf.Host, conf.Port), conf.Topic, conf.Partition)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type Options struct {
	Topic  string
	Escape bool
}

func WriteMessages(contents []string) error {
	messages := make([]kafka.Message, 0, len(contents))
	for index := range contents {
		messages = append(messages, kafka.Message{Value: []byte(contents[index])})
	}
	_, err := Conn.WriteMessages(messages...)
	if err != nil {
		return err
	}
	return nil
}

func WriteMessageWithOptions(content []string, opts ...*Options) error {
	return nil
}
