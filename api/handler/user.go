package handler

import (
	"matcher/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(userService user.Usecase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"user": "get user"})
	})
}

func createUser(userService user.Usecase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"user": "create user"})
	})
}
func UsersHandler(r *gin.Engine, userService user.Usecase) {
	v1 := r.Group("v1")
	{
		v1.GET("/users/:id", getUser(userService))
		v1.POST("/users", createUser(userService))
	}
}
