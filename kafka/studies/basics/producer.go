package basics

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type producer struct {
	instance *kafka.Producer
	topics   []kafka.TopicPartition
	name     string
	stopChan chan bool
}

func NewProducer(name string, bootstrapServers string, topics []string) *producer {
	prod, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	partitions := make([]kafka.TopicPartition, 0)

	for _, topic := range topics {
		partitions = append(partitions, kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny})
	}

	if err != nil {
		log.WithError(err).Error("Failed to create producer")
		panic(err)
	}

	return &producer{
		instance: prod,
		topics:   partitions,
		name:     name,
	}
}

func (p *producer) createMessage(topic kafka.TopicPartition, key string, value string) *kafka.Message {
	return &kafka.Message{
		TopicPartition: topic,
		Value:          []byte(value),
		Key:            []byte(key),
	}
}

func (p *producer) Start(interval time.Duration) {
	p.stopChan = make(chan bool, 1)

	go func() {
		var count int
		for {
			select {
			case <-p.stopChan:
				p.instance.Close()
				return

			default:
				for _, topic := range p.topics {
					p.instance.ProduceChannel() <- p.createMessage(topic, p.name, fmt.Sprintf("Message number %d", count))
				}
				count++
				time.Sleep(interval)
			}
		}
	}()

}

func (p *producer) Stop() {
	log.Warnf("Stopping producer %s", p.name)
	p.stopChan <- true
}
