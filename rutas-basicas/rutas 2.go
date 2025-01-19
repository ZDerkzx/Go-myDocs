package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// creamos como siempre lo necesario el router
	r := gin.Default()

	// creamos la peticion que manejara el servidor
	// le pedimos 2 valores para multiplicarlos
	r.GET("/multiply/:a/:b", func(c *gin.Context) {
		//ya pedidos los valores tenemos que guardarlos en la memoria del servidor
		// temporalmente mientras los multiplicamos
		// bueno entonces los guardamos
		a := c.Param("a")
		b := c.Param("b")
		// y los convertimos con strconv a int
		// y agregamos un manejo de errores
		numberA, errorA := strconv.Atoi(a)
		if errorA != nil {
			fmt.Println("ERROR IN A")
		}
		// igual aqui(es lo mismo que numberA)
		numberB, errorB := strconv.Atoi(b)
		if errorB != nil {
			fmt.Println("ERROR IN B")
		}
		// creamos una variable multiplicando los 2 valores
		result := numberA * numberB
		// por ultimo damos la respuesta al cliente con un 200(EXITO)
		// y le mostramos el Resultado en JSON
		c.JSON(200, gin.H{
			"El ": result,
			"OMG": "ALV",
		})
	})
	// Runeamos el server
	// en el puerto 8080
	r.Run(":8080")
}
