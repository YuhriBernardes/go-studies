package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/YuhriBernardes/kafka/studies/basics"
	log "github.com/sirupsen/logrus"
)

func configLogs() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

}

func main() {
	configLogs()

	timeout := time.Second * 3
	bootstrapServers := "localhost:9092"
	topic := "my.topic"

	a := basics.NewAdmin(bootstrapServers, timeout)

	a.Newtopic(topic)

	p := basics.NewProducer("Producer 1", bootstrapServers, []string{topic})
	p.Start(2 * time.Second)

	c := basics.NewSubscribedConsumer(bootstrapServers, "some.consumer", []string{topic})
	c.PollAsync(100, func(key string, value string) {
		log.WithFields(log.Fields{
			"key":   key,
			"value": value,
		}).Info("Message received")
	})

	mainChan := make(chan os.Signal, 1)

	signal.Notify(mainChan, os.Interrupt)

	<-mainChan

	c.Stop()
	p.Stop()

	a.DeleteTopic(topic)

	a.Close()
}
