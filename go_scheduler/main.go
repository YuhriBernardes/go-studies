package main

import (
	"github.com/YuhriBernardes/gs-scheduler/sample"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	log.SetLevel(log.InfoLevel)

	rootCmd := cobra.Command{
		Use:   "gs",
		Short: "Run go schedule samples",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	sample.LoadSamples(&rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Error("Failed to execute command")
	}

}
