package main

import (
	"os"
	"os/signal"

	"github.com/RafaelFernando12/api-golang/controller"
	"github.com/RafaelFernando12/api-golang/domain/user"
	"github.com/RafaelFernando12/api-golang/internal/http"
	"github.com/RafaelFernando12/api-golang/internal/repository"
	"github.com/RafaelFernando12/api-golang/pkg/env"
	"github.com/RafaelFernando12/api-golang/pkg/log"
)

const (
	applicationName = "api-golang"
)

func main() {
	/*db*/
	db, err := repository.NewDb(env.GetEnv(env.DbUrl))
	if err != nil {
		panic("Connection database failed!")
	}

	logger := log.NewLogger(applicationName, env.GetEnv(env.EnvLoggerLevel, env.DefaultLoggerLevel))

	/*Repository*/
	userRepository := repository.NewUserRepository(db)

	/*Services*/
	userService := user.NewUserService(userRepository, logger)

	handler := controller.NewHandler(
		userService,
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
