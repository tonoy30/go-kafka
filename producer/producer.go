package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PushCommitToTopic(topic string, message []byte) (string, error) {
	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return "", err
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(producer)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return "", err
	}
	successMsg := fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)
	return successMsg, nil
}
