package sample

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Init()
// run(*cobra.Command, []string)
// Command() *cobra.Command

type Servers struct {
	command *cobra.Command
}

func createHandler(handlerFunc http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", handlerFunc)
	return mux
}

func createDummyHandlerFunc(serverNumber int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Infof("Request received on server %d", serverNumber)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"server": serverNumber,
		})

	}
}

func createServer(serverNumber int, serverPort int) http.Server {
	h := createHandler(createDummyHandlerFunc(serverNumber))

	return http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", serverPort),
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

}

func (s *Servers) run(cmd *cobra.Command, args []string) {
	port1, _ := cmd.PersistentFlags().GetInt("port1")
	port2, _ := cmd.PersistentFlags().GetInt("port2")

	log.WithFields(log.Fields{
		"number": 1,
	}).Warn("Creating server")

	server1 := createServer(1, port1)

	log.WithFields(log.Fields{
		"number": 2,
	}).Warn("Creating server")

	server2 := createServer(2, port2)

	go func() {
		log.WithFields(log.Fields{
			"port": port1,
		}).Warn("Starting server")
		server1.ListenAndServe()
	}()

	go func() {
		log.WithFields(log.Fields{
			"port": port2,
		}).Warn("Starting server")
		server2.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx := context.Background()

	log.Warn("Shutting down servers")
	server1.Shutdown(ctx)
	server2.Shutdown(ctx)

}

func (s *Servers) Init() {
	cmd := &cobra.Command{
		Use:   "servers",
		Short: "Starts two servers at same time on different ports",
		Run:   s.run,
	}

	cmd.PersistentFlags().Int("port1", 3000, "Server 1 port (defaults to 3000)")
	cmd.PersistentFlags().Int("port2", 3001, "Server 2 port (defaults to 3001)")

	s.command = cmd
}

func (s *Servers) Command() *cobra.Command {
	return s.command
}
