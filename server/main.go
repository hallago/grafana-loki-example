package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/coolishbee/grafana-loki-example/server/pkg/logging"
)

func init() {
	logging.Setup()
}

func main() {
	TestLog()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func TestLog() {
	for range time.Tick(time.Second * time.Duration(10)) {
		random := rand.Int()
		logging.Info("API Server", fmt.Sprintf("%d", random))
	}
}
