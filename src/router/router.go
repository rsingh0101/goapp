package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rsingh0101/src/handlers"
	"github.com/rsingh0101/src/mariadb"
)

func SetupRouter(db *mariadb.DB) *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// r.GET("/produce", handlers.ProduceMessageHandler(producer, config.Kafka.Topic))
	// r.GET("/create_kafka")
	r.GET("/users", handlers.GetUsersHandler(db))
	r.POST("/insert", handlers.InsertUsersHandler(db))
	r.POST("/delete", handlers.DeleteUsersHandler(db))
	return r

}
