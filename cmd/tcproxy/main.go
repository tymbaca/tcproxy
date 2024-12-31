package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/tymbaca/tcproxy/internal/config"
	"github.com/tymbaca/tcproxy/internal/tcproxy"
)

var configPathFlag = flag.String("config", "tcproxy.yaml", "path to config file")

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.Parse(*configPathFlag)
	if err != nil {
		return fmt.Errorf("can't parse config from %s: %w", *configPathFlag, err)
	}

	tcproxy, err := tcproxy.New(cfg)
	if err != nil {
		return fmt.Errorf("can't init tcproxy: %w", err)
	}

	err = tcproxy.Run(ctx)
	if err != nil {
		return fmt.Errorf("can't run tcproxy: %w", err)
	}

	return nil
}
