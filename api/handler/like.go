package handler

import (
	"log"
	"matcher/api/input"
	"matcher/api/output"
	"matcher/services/like"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

// createLike creates a new like between 2 users. input is the input.CreateLike struct and out put is 204
func createLike(likeService like.Usecase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var json input.CreateLike

		if err := c.ShouldBindJSON(&json); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := likeService.CreateLike(json.FromID, json.ToID)

		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnt create like"})
			return
		}

		c.String(http.StatusCreated, "")
	})
}

// getLikes gets who likes the user paginated max page size 100.
func getLikes(likeService like.Usecase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		var pagination input.ListLikes

		if err := c.ShouldBindWith(&pagination, binding.Query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId, err := uuid.Parse(c.Param("userId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		results, err := likeService.ListLikes(userId, pagination.Page, pagination.PageSize)

		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnt get likes"})
			return
		}

		var likes []*output.ListLikes

		for _, v := range results {
			likes = append(likes, &output.ListLikes{
				Name:      v.From.FirstName,
				FromID:    v.FromID.String(),
				ToID:      v.ToID.String(),
				CreatedAt: v.CreatedAt,
			})
		}

		c.JSON(http.StatusOK, gin.H{"likes": likes})
	})
}

// LikesHandler creates all the routes + functions that the likes /likes is responsible for
// closure used to pass the service through to the handlers
func LikesHandler(r *gin.Engine, likeService like.Usecase) {
	v1 := r.Group("v1")
	{
		v1.GET("/likes/:userId", getLikes(likeService))
		v1.POST("/likes", createLike(likeService))
	}
}
