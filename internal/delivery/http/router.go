package http

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/clients", func(c *gin.Context) {
			c.JSON(501, gin.H{"message": "not implemented"})
		})
	}

	return r
}