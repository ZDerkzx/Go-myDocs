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
	r.GET("/math/:a/:b/:option", func(c *gin.Context) {
		//ya pedidos los valores tenemos que guardarlos en la memoria del servidor
		// temporalmente mientras los multiplicamos
		// bueno entonces los guardamos
		a := c.Param("a")
		b := c.Param("b")
		option := c.Param("option")
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
		// creamos un switch para validar cosas con option
		// y dependiendo la opcion elegida ara una cosa u otra
		switch option {
		case "multiplicar":
			result := numberA * numberB
			c.JSON(200, gin.H{
				"Resultado: ": result,
			})
		case "suma":
			result := numberA + numberB
			c.JSON(200, gin.H{
				"Resultado: ": result,
			})
		case "resta":
			result := numberA - numberB
			c.JSON(200, gin.H{
				"Resultado: ": result,
			})
		}

	})
	// Runeamos el server
	// en el puerto 8080
	r.Run(":8080")
}
