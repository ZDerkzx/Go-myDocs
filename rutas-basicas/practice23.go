package main

import "github.com/gin-gonic/gin"

func main() {
	// crea una ruta
	r := gin.Default()
	// con la misma ruta osea "r" le dice que estara aceptando peticiones Tipo GET en /ping
	// la funcion que tiene ahi creo "CREO"(ando aprendiendo apenas)
	// que es lo que contendra
	r.GET("/ping", func(c *gin.Context) {
		// devuelve un 200(Osea que si se pudo encontrar la pagina)
		// y devuelve un JSON
		c.JSON(200, gin.H{
			//Clave: Valor
			"message": "pong",
		})
	})
	// asemos que maneje otra peticion GET con el enrutador que creamos
	// y si vez le puse :name en "r.GET("/apply/:name", func(c *gin.Context) {"
	// esto lo que ase es capturar un valor ingresado en la URL tipo: localhost:8080/apply/Pedro
	r.GET("/apply/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			// y devolvera el JSON Clave: Valor
			// mas el nombre que capturamos en el GET con :name
			"test": "Hola " + name,
		})
	})
	// Runea el Router principal que creamos al empezar
	r.Run(":8080")
}
