package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		c.JSON(200, gin.H{
			"status ": "recibido",
			"name: ":  name,
			"email: ": email,
		})
	})
	r.Run(":8080")
}
