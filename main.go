//se encarga de crear el router se encarga de iniciar el serve y importar route y usar las funciones necesarias para registrar rutas en el router

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, Mundo!",
		})

	})

	r.Run(":8080")
}
