package handler

import (
	"github.com/gin-gonic/gin"
	"sycret-test-task/internal/service"
)

type Handlers struct {
	services *service.Services
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		docsGenerator := api.Group("/docs")
		{
			docsGenerator.GET("/generate", h.generateDoc)
		}
	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
