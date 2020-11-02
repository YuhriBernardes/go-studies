package studies

import (
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type consumer struct {
	instance *kafka.Consumer
	stopChan chan bool
}

func NewConsumer(bootstrapServers string, groupId string) *consumer {
	instance, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.WithError(err).Error("Failed do create consumer")
		panic(err)
	}

	return &consumer{
		instance: instance,
	}
}

func NewSubscribedConsumer(bootstrapServers string, groupId string, topics []string) *consumer {
	instance := NewConsumer(bootstrapServers, groupId)

	instance.SubscribeToTopics(topics)

	return instance
}
func (c *consumer) SubscribeToTopic(topic string) {
	if err := c.instance.Subscribe(topic, nil); err != nil {
		log.WithError(err).WithField("topic", topic).Error("Failed to subscribe to topic")
		panic(err)
	}
}

func (c *consumer) SubscribeToTopics(topics []string) {
	if err := c.instance.SubscribeTopics(topics, nil); err != nil {
		log.WithError(err).WithField("topic", strings.Join(topics, ",")).Error("Failed to subscribe to topics")
		panic(err)
	}
}

func (c *consumer) PollAsync(timeoutMs int, processor func(key string, value string)) {

	c.stopChan = make(chan bool, 1)

	go func() {
		for {
			select {
			case <-c.stopChan:
				c.instance.Close()
				return

			default:
				event := c.instance.Poll(timeoutMs)
				switch e := event.(type) {
				case *kafka.Message:
					key, value := string(e.Key), string(e.Value)
					processor(key, value)
				case *kafka.Error:
					log.WithError(e).Error("Failed to poll messages")
				}
			}
		}
	}()

}

func (c *consumer) Stop() {
	log.Warn("Stopping consumer")
	c.stopChan <- true
}
