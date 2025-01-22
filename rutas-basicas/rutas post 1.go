package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// inicializamos el formulario
	r.GET("/", func(c *gin.Context) {
		c.File("myForm.html")
	})

	// creamos una ruta de metodo POST
	// esta ruta obtiene lo que estaba en el formulario
	r.POST("/submit", func(c *gin.Context) {
		// estas 2 variables guardan las variables que estan en los input de myForm
		// cuales variables? pues las: name="texto de ejemplo"
		// esas variables se encuentran en myForm
		name := c.PostForm("name")
		password := c.PostForm("password")
		// devolvemos un JSON con lo que ingreso el usuario
		c.JSON(201, gin.H{
			"Name: ":   name,
			"password": password,
		})
	})
	r.Run()
}
