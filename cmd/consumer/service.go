package consumer

import (
	"github.com/gokafka/service"
	"github.com/gokafka/service/kafka"
)

func InitService(conf *service.Config) error {
	var err error
	kafka.Conn, err = kafka.InitKafkaConnect(&conf.Kafka)
	if err != nil {
		return err
	}
	return nil
}
