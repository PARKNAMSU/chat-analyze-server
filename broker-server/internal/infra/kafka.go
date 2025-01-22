package infra

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"chat-analyze.com/chat-analyze-server/internal/tools"
	"github.com/IBM/sarama"
)

var (
	brokers = []string{os.Getenv("KAFKA_HOST")}
)

var (
	consumer = getConsumer()
)

func getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.Timeout = time.Second * 5
	return config
}

func getConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokers, getConfig())
	if err == nil {
		log.Println(err)
		return nil
	}
	return consumer
}

func CloseConsummer() {
	if consumer == nil {
		return
	}
	consumer.Close()
}

func KafkaSubscribeTopic(topic string) sarama.PartitionConsumer {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		tools.PrintErrorLog("KafkaSubscribeTopic", err.Error())
		return nil
	}
	return partitionConsumer
}

func KafkaPolling[T any](topicConsumber sarama.PartitionConsumer, msgChan chan T) {
	for msg := range topicConsumber.Messages() {
		var data T
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			tools.PrintErrorLog("KafkaPolling", err.Error())
			continue
		}
		msgChan <- data
	}
}
