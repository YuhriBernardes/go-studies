package main

import (
	"github.com/YuhriBernardes/kafka/studies"
	"github.com/YuhriBernardes/kafka/studies/loop"
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

	loop.Cmd.Register()

	if err := studies.Execute(); err != nil {
		log.WithError(err).Error("Failed to start cli")
	}

}
