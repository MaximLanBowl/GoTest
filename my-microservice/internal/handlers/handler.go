package handlers

import (
	"my-microservice/internal/repository"
	"my-microservice/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, repo repository.Repository, kafkaSvc service.KafkaService) {
	r.POST("/messages", handleCreateMessage(repo, kafkaSvc))
	r.GET("/messages", handleGetMessages(repo))
}

func handleCreateMessage(repo repository.Repository, kafkaSvc service.KafkaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg model.Message
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := repo.SaveMessage(&msg); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := kafkaSvc.SendMessage(&msg); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, msg)
	}
}

func handleGetMessages(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		messages, err := repo.GetMessages()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, messages)
	}
}
