package handler

import (
	"log"
	"matcher/api/input"
	"matcher/api/output"
	"matcher/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getUser gets the user by userId. Returns a output.User
// take a userId as a query param NEEDS TO BE UUID
func getUser(userService user.Usecase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		userId, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userService.GetUser(userId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser := &output.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.LastName,
		}

		c.JSON(http.StatusOK, gin.H{"user": newUser})
	})
}

// createUser creates a new user
func createUser(userService user.Usecase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var json input.CreateUser

		if err := c.ShouldBindJSON(&json); err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := userService.CreateUser(json.FirstName, json.LastName, json.Email, json.Email)

		if err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnt create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"userId": id})
	})
}

func UsersHandler(r *gin.Engine, userService user.Usecase) {
	v1 := r.Group("v1")
	{
		v1.GET("/users/:id", getUser(userService))
		v1.POST("/users", createUser(userService))
	}
}
