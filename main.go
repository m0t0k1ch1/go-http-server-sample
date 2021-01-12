package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/app"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/common"
)

func main() {
	conf, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(conf)

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
