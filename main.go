package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rsingh0101/src/config"
	"github.com/rsingh0101/src/mariadb"
	"github.com/rsingh0101/src/router"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Create a new database connection
	db, err := mariadb.NewDB(config.Database.User, config.Database.Password, config.Database.Host, config.Database.DBName)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	defer db.Conn.Close()

	r := router.SetupRouter(db)
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()
	// Signal handling to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	// Perform any necessary cleanup here
	if err := db.Conn.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	log.Println("Server exited")
}
