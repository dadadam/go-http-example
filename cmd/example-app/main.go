package main

import (
	"os"

	"github.com/dadadam/go-http-example/api/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)

	log.WithFields(log.Fields{
		"module": "API",
	}).Info("API server started")

	s := &server.Server{}
	s.Start(":8080")
}
