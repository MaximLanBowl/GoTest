package main

import (
	"my-microservice/cmd/internal/handlers"
	"my-microservice/cmd/internal/repository"
	"my-microservice/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	rep := repository.NewPostgresRepository()
	kafkaSrv := service.NewKafkaSrv()

	handlers.RegisterHandlers(rep, kafkaSrv)

	r.Run()
	
}