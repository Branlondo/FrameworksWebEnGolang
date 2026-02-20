//se encarga de crear el router se encarga de iniciar el serve y importar route y usar las funciones necesarias para registrar rutas en el router

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hola, mundo!")

	})

	r.Run()
}
