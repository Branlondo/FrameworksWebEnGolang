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
	//se define una ruta GET para la raíz ("/") que responde con un mensaje de saludo
	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "¡Hola, mundo!")
	})
	//se define una ruta GET para "/saludo/:nombre" que responde con un mensaje de saludo personalizado utilizando el valor del parámetro "nombre" en la URL
	r.GET("/saludo/:nombre", func(c *gin.Context) {
		//se obtiene el valor del parámetro "nombre" de la URL utilizando c.Param("nombre") y se almacena en la variable nombre
		nombre := c.Param("nombre")
		//el codigo de estado http.StatusOK es 200, se devuelve un mensaje de saludo personalizado
		c.String(http.StatusOK, "¡Hola, %s!", nombre)
	})
	r.POST("/usuarios", func(c *gin.Context) {
		var nuevoUsuario Usuario

		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el JSON"})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y correo electronico son datos requerido"})
		}

		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario Registrado", "datos": usuarios})
	})
}
