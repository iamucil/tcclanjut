package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{"module": "main"})

var (
	revisionID     = "unknown"
	buildTimestamp = "unknown"
)

func main() {
	fmt.Fprintf(os.Stdout, "goweb revision %s built at %s\n", revisionID, buildTimestamp)

	logLevelStr := os.Getenv("LOG_LEVEL")

	if logLevelStr != "" {
		logLevel, err := logrus.ParseLevel(logLevelStr)
		if err != nil {
			panic(err)
		}
		logrus.SetLevel(logLevel)
	}

	httpServicePort := os.Getenv("HTTP_PORT")
	if httpServicePort == "" {
		httpServicePort = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

  // serve static resources
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("resources/assets"))))

	var shuttingDown bool
	shutdownSignal := make(chan os.Signal)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	httpServer := &http.Server{Addr: ":" + httpServicePort}

	go func() {
		log.Infof("HTTP Service is ready and listen on port %s", httpServicePort)
		err := httpServer.ListenAndServe()

		if err != nil && (err != http.ErrServerClosed || !shuttingDown) {
			log.Errorf("HTTP Service error: %v", err)
		}
	}()

	// catch shutting down signal from os
	<-shutdownSignal
	shuttingDown = true
	log.Infof("Shutting down the server....")

	go func() {
		<-shutdownSignal
		os.Exit(0)
	}()

	shutdownCtx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	httpServer.Shutdown(shutdownCtx)
	log.Infof("Done.")
}
