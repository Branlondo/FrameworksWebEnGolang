//se encarga de crear el router se encarga de iniciar el serve y importar route y usar las funciones necesarias para registrar rutas en el router

package main

import (
	//Manda a importar las rutas
	"Gin/routes"
	// se trabaja el servidor en gin
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.SetupRoutes(r)
	//el servidor 8080 el cual es por defecto tiene un problema asi que se cambia a 8181 para evitar conflictos con otros servidores que puedan estar corriendo en el puerto 8080
	r.Run(":8181")
}
