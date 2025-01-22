package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// se debe de importar el archivo html
	r.LoadHTMLFiles("index.html")
	// se maneja la peticion /home
	r.GET("/home", func(c *gin.Context) {
		// La respuesta es un HTML
		c.HTML(200, "index.html", nil)
	})
	// Runea el server
	r.Run(":8080")
}
