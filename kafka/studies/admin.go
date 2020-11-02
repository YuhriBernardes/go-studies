package studies

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type admin struct {
	adm     *kafka.AdminClient
	timeout time.Duration
}

func NewAdmin(bootstrapServers string, timeout time.Duration) *admin {
	log.Warn("Creating new admin")
	adm, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	if err != nil {
		log.WithError(err).Error("Failed to start client")
		panic(err)
	}

	return &admin{
		adm: adm,
	}

}

func defaultTopicConfig(name string) kafka.TopicSpecification {
	return kafka.TopicSpecification{
		Topic:             name,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
}

func (a admin) Newtopic(name string, opts ...kafka.AdminOption) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topicConfig := defaultTopicConfig(name)

	log.WithFields(log.Fields{
		"topic": name,
	}).Warn("Creating topic")

	results, err := a.adm.CreateTopics(ctx, []kafka.TopicSpecification{
		topicConfig,
	}, kafka.SetAdminOperationTimeout(a.timeout))

	if err != nil {
		log.WithError(err).Error("Failed to create topic")
	}

	for _, result := range results {
		log.WithFields(log.Fields{
			"Topic": result,
		}).Info("Topic created")
	}
}

func (a admin) DeleteTopic(name string) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.WithFields(log.Fields{
		"name": name,
	}).Warn("Deleting topic")
	results, err := a.adm.DeleteTopics(ctx, []string{name}, kafka.SetAdminOperationTimeout(a.timeout))

	if err != nil {
		log.WithError(err).Errorf("Failed to delete topic %s", name)
		panic(err)
	}

	for _, result := range results {
		log.WithFields(log.Fields{
			"Topic": result,
		}).Info("Topic deleted")
	}

}

func (a admin) Close() {
	a.adm.Close()
}
