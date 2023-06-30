package kafka

import (
	gokafka "github.com/segmentio/kafka-go"
)

var (
	Conn *gokafka.Conn
)

type Config struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Network   string `yaml:"network"`
	Topic     string `yaml:"topic"`
	Partition int    `yaml:"partition"`
}
