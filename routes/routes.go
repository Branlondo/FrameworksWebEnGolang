//define las rutas de la app, importa gin y usa las funciones para definir las rutas define las funciones para procesar las solicitudes

package routes

import (
	//se necesita para trabajar con el protocolo HTTP y manejar las respuestas
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// ejemplo de uuso aun no tiene funcion
type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

// SetupRoutes configura las rutas para la aplicación Gin
func SetupRoutes(r *gin.Engine) {
	//esto funciona para servir archivos estáticos desde el directorio "./static" cuando se accede a la ruta "/static"
	r.Static("/static", "./static")
	//esto indica a Gin que cargue las plantillas HTML desde el directorio "templates" con la extensión ".html"
	r.LoadHTMLGlob("templates/*.html")

	//se define una ruta para la raíz ("/") que renderiza la plantilla "index.html" cuando se accede a ella
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//se define una ruta para cualquier página ("/:page") que verifica si la plantilla correspondiente existe y la renderiza, o muestra una página de error 404 si no existe
	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")

		if !strings.HasSuffix(page, ".html") {
			page += ".html"
		}

		if _, err := os.Stat("templates/" + page); err == nil {
			c.HTML(http.StatusOK, page, nil)
		} else {
			c.HTML(http.StatusNotFound, "404.html", nil)
		}
	})

}
