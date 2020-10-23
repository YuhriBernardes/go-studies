package basics

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

func NewAdmin(config *kafka.ConfigMap, timeout time.Duration) *admin {
	adm, err := kafka.NewAdminClient(config)

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

	ctx := context.Background()
	clusterId, _ := a.adm.ClusterID(ctx)

	config := defaultTopicConfig(name)

	log.WithFields(log.Fields{
		"cluster ID":     clusterId,
		"name":           config.Topic,
		"partitionCount": config.NumPartitions,
	}).Warn("Creating topic")

	_, err := a.adm.CreateTopics(ctx, []kafka.TopicSpecification{config}, kafka.SetAdminRequestTimeout(a.timeout), kafka.SetAdminOperationTimeout(a.timeout))

	if err != nil {
		log.WithError(err).Errorf("Failed to create topic %s", config.Topic)
		panic(err)
	}
}

func (a admin) DeleteTopic(name string) {

	ctx := context.Background()
	clusterId, _ := a.adm.ClusterID(ctx)

	log.WithFields(log.Fields{
		"cluster ID": clusterId,
		"name":       name,
	}).Warn("Deleting topic")

	if _, err := a.adm.DeleteTopics(ctx, []string{name}, kafka.SetAdminRequestTimeout(a.timeout), kafka.SetAdminOperationTimeout(a.timeout)); err != nil {
		log.WithError(err).Errorf("Failed to delete topic %s", name)
		panic(err)
	}

}
