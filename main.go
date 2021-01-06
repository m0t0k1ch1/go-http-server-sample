package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := NewApp()

	go func() {
		if err := app.Start(":1323"); err != nil {
			if err == http.ErrServerClosed {
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
