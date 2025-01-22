package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"chat-platform-api.com/chat-platform-api/src/tool/logging_tool"
	"chat-platform-api.com/chat-platform-api/src/variable/common_variable"
	"github.com/IBM/sarama"
)

var (
	brokers = []string{os.Getenv("KAFKA_HOST")}
)

var (
	config   *sarama.Config
	producer = getProducer()
)

func getConfig() *sarama.Config {
	if config == nil {
		config = sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.Timeout = time.Second * 5
	}
	return config
}

func getProducer() sarama.SyncProducer {
	getProducer, err := sarama.NewSyncProducer(brokers, getConfig())
	if err == nil {
		log.Println(err)
		return nil
	}
	return getProducer
}

func KafkaSendMessage[T any](topic string, data T) {
	var err error

	defer func() {
		if err != nil {
			logging_tool.PrintErrorLog("kafka:SendMessage", err.Error())
		}
	}()

	if producer == nil {
		err = errors.New("kafka not work")
		return
	}

	valueBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(valueBytes)),
	})

	if err != nil {
		return
	}

	if common_variable.ENVIRONMENT == "development" {
		fmt.Printf("Message is stored in partition %d, offset %d\n", partition, offset)
	}
}
