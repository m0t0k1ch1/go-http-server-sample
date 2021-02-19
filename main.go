package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/app"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

func main() {
	var confPath = flag.String("conf", "", "the path to the config file")
	flag.Parse()

	conf, err := common.LoadConfig(*confPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := conf.Validate(); err != nil {
		log.Fatalf("invalid config: %v", err)
	}

	app, err := app.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				app.Logger.Info(err)
			} else {
				app.Logger.Fatal(err)
			}
		}
	}()

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM)
	for {
		sig := <-sigch
		if sig != syscall.SIGTERM {
			continue
		}
		app.Logger.Info("received SIGTERM")
		if err := app.Shutdown(context.Background()); err != nil {
			app.Logger.Fatal(err)
		}
		return
	}
}
