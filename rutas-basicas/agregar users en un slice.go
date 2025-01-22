package main

import "github.com/gin-gonic/gin"

// creamos un slice
var mySlice = []string{"Pablo", "Martin"}

func main() {
	r := gin.Default()

	// creamos una ruta GET donde se mostrara la lista
	r.GET("/slice", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"List: ": mySlice,
		})
	})

	// creamos una ruta donde puedas agregar una persona a la lista
	r.GET("/addSlice/:name", func(c *gin.Context) {
		var mySlice = []string{"Pablo", "Martin"}
		// obtiene el parametro name
		name := c.Param("name")
		// agrega el parametro que obtuvimos a mySlice
		mySlice = append(mySlice, name)
		//devuelve un JSON
		// con los resultados
		c.JSON(200, gin.H{
			"new: ":      name,
			"List Now: ": mySlice,
		})
	})
	// Runeamos el sv
	r.Run(":8080")
}
