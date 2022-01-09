package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func ConsumePartition(brokersUrl []string, topic string) {
	worker, err := connectConsumer(brokersUrl)
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	count := 0
	done := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				count++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n",
					count, msg.Topic, string(msg.Value))
			}
		}
	}()
	<-done
}
