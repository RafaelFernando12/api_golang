package main

import (
	"os"
	"os/signal"

	"github.com/RafaelFernando12/api-golang/controller"
	"github.com/RafaelFernando12/api-golang/internal/http"
	"github.com/RafaelFernando12/api-golang/pkg/env"
	"github.com/RafaelFernando12/api-golang/pkg/log"
)

const (
	applicationName = "api-golang"
)

func main() {
	logger := log.NewLogger(applicationName, env.GetEnv(env.EnvLoggerLevel, env.DefaultLoggerLevel))

	handler := controller.NewHandler(
		logger,
		env.GetEnv(env.EnvOrigin, env.DefaultOrigin))

	server := http.NewServer(env.GetEnv(env.EnvPort, env.DefaultPort), handler, logger)
	server.Start()

	<-interrupt()

	server.Shutdown()

}

func interrupt() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	return c
}
