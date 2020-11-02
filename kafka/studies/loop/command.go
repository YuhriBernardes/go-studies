package loop

import (
	"os"
	"os/signal"
	"time"

	"github.com/YuhriBernardes/kafka/studies"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	name = "basics"
)

var (
	Cmd = &Command{}
)

type Command struct {
}

func (c *Command) run(cmd *cobra.Command, args []string) {
	bootstrapServers, _ := cmd.Parent().PersistentFlags().GetString("bs")

	timeout := time.Second * 2
	topic, _ := cmd.PersistentFlags().GetString("topic")

	log.WithFields(log.Fields{
		"bootstrapServers": bootstrapServers,
		"topic":            topic,
	}).Warn("Starting Basic example")

	adm := studies.NewAdmin(bootstrapServers, timeout)

	adm.Newtopic(topic)

	producer := studies.NewProducer("Producer", bootstrapServers, []string{topic})
	producer.Start(2 * time.Second)

	consumer := studies.NewSubscribedConsumer(bootstrapServers, "some.consumer", []string{topic})
	consumer.PollAsync(100, func(key string, value string) {
		log.WithFields(log.Fields{
			"key":   key,
			"value": value,
		}).Info("Message received")
	})

	mainChan := make(chan os.Signal, 1)

	signal.Notify(mainChan, os.Interrupt)

	<-mainChan

	consumer.Stop()
	producer.Stop()

	adm.DeleteTopic(topic)

	adm.Close()
}

func (c *Command) Register() {
	cmd := &cobra.Command{
		Use:   "loop",
		Short: "Basic consume and produce operations",
		Long:  "Create a topic, with a cosumer subscribe to it and a producer. Every 2 seconds the produces produce a message to this topic and the consumer logs the message",
		Run:   c.run,
	}

	cmd.PersistentFlags().String("topic", "basic.topic", "topic name")

	studies.Register(cmd)

}

func (c *Command) Name() string {
	return name
}
