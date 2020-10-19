package sample

import (
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Countdown struct{}

func (c Countdown) counter(name string, start int) {

	log.WithFields(log.Fields{
		"Thread": name,
	}).Warn("Starting Counter")

	for i := start; i > 0; i-- {

		log.WithFields(log.Fields{
			"Thread": name,
			"count":  strconv.Itoa(i),
		}).Info("Counting")

		time.Sleep(1 * time.Second)

	}

	log.WithFields(log.Fields{
		"Thread": name,
	}).Warn("Counter finished")

}

func (c Countdown) Start() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		c.counter("1", 10)
		wg.Done()
	}()
	go func() {
		c.counter("2", 40)
		wg.Done()
	}()

	wg.Wait()
	log.Warn("Finished")

}
