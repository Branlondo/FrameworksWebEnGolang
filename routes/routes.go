//define las rutas de la app, importa gin y usa las funciones para definir las rutas define las funciones para procesar las solicitudes

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

// SetupRoutes configura las rutas para la aplicación Gin
func SetupRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	//se define una ruta GET para la raíz ("/") que responde con un mensaje de saludo
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Mi primera app con Gin",
			"Heading": "¡Hola, Mundo!",
			"Message": "Bienvenido a mi primera aplicación web con Gin y plantillas HTML.",
		})
	})

	r.Static("/static", "./static")
}
