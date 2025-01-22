package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un servidor Gin
	r := gin.Default()

	r.LoadHTMLFiles("test.html")

	// Sirve el archivo HTML
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "test.html", nil) // Envía el archivo prueba.html al navegador
	})

	// Maneja los datos enviados desde el formulario
	r.POST("/submit", func(c *gin.Context) {
		// Obtiene los valores enviados desde el formulario
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Responde con los datos recibidos
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// Ejecuta el servidor en localhost:8080
	r.Run(":8080")
}
