package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/ise-car-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 3 * time.Second,
		ProducerCount:  28,
		WorkerCount:    2,
	}

	r := retranslator.NewRetranslator(cfg)
	r.Start(context.Background())

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	r.Close()
}
