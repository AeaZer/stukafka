package service

import (
	"github.com/gokafka/service/kafka"
)

type Config struct {
	Kafka kafka.Config `yaml:"kafka"`
}
