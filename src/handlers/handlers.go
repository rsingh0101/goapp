package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsingh0101/src/mariadb"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsersHandler(db *mariadb.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := db.QueryUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error querying users: %v", err)})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func InsertUsersHandler(db *mariadb.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User

		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can't read request body"})
			return
		}

		err = json.Unmarshal(bodyBytes, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		err = db.InsertUser(user.Name, user.Age)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User inserted successfully"})
	}
}

func DeleteUsersHandler(db *mariadb.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can't read request body"})
			return
		}

		err = json.Unmarshal(bodyBytes, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}
		err = db.DeleteUser(user.Name, user.Age)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		err = db.DeleteUser(user.Name, user.Age)
		if err != nil {
			log.Fatalf("Error deleting user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
