package sample

import (
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CountDown struct {
	command *cobra.Command
}

func (cnt *CountDown) counter(name string, start int) {

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

func (cnt *CountDown) run(cmd *cobra.Command, args []string) {

	var wg sync.WaitGroup
	wg.Add(2)

	c1, _ := cmd.PersistentFlags().GetInt("c1")
	c2, _ := cmd.PersistentFlags().GetInt("c2")

	go func() {
		cnt.counter("1", c1)
		wg.Done()
	}()
	go func() {
		cnt.counter("2", c2)
		wg.Done()
	}()

	wg.Wait()
	log.Warn("Finished")

}

func (cnt *CountDown) Init() {

	cmd := &cobra.Command{
		Use:   "countdown",
		Short: "Multithread countdown",
		Run:   cnt.run,
	}

	cmd.PersistentFlags().Int("c1", 10, "Counter 1 start value")
	cmd.PersistentFlags().Int("c2", 10, "Counter 2 start value")

	cnt.command = cmd

}

func (cnt *CountDown) Command() *cobra.Command {
	return cnt.command
}
