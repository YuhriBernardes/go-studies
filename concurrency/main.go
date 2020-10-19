package main

import (
	"errors"
	"flag"
	"os"

	"github.com/YuhriBernardes/go-studies-concurrency/sample"

	log "github.com/sirupsen/logrus"
)

var (
	samples = map[string]sample.Sample{
		"countdown": sample.Countdown{},
	}
	sampleName = ""
)

func discoverSample() (sample.Sample, error) {

	if sample, found := samples[sampleName]; found {

		return sample, nil

	}
	return nil, errors.New("Sample not found")

}

func processArgs() (err error) {
	flag.StringVar(&sampleName, "sample", "", "Select the desired sample to run")
	flag.Parse()

	if sampleName == "" {
		return errors.New("You need to specify the desired sample (e.g. -sample=countdown)")
	}

	return nil
}

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	log.SetLevel(log.InfoLevel)

	if err := processArgs(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Warningf("Loading sample %s", sampleName)

	if s, err := discoverSample(); err == nil {
		log.Warningf("Starting sample %s", sampleName)
		s.Start()
	} else {
		log.Error(err.Error())
		os.Exit(1)
	}

}
