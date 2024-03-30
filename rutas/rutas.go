package rutas

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nd-back/controladores"
	"nd-back/middlewares"
	"os"
	"strconv"
)

func Configuracion(app *fiber.App) {
	reg, err := strconv.ParseBool(os.Getenv("REGISTRAR_ENABLED"))
	if err != nil {
		fmt.Println(err)
	}
	if reg {
		app.Post("/registrar", controladores.Registrar)
	}
	app.Post("/ingresar", controladores.Ingresar)
	app.Get("/entrada/:id", controladores.LeerEntrada)
	app.Get("/entradas", controladores.TodasEntradas)   // Muestra las entradas paginadas
	app.Get("/todas", controladores.EntradasSinPaginar) // Muestra las entradas sin paginar
	app.Get("/especial", controladores.TodasEntradas)   // Muestra las entradas ver parametros del query
	app.Get("/comentarios", controladores.TodosComentarios)
	app.Post("/escribeme", controladores.Escribeme)
	app.Use(middlewares.Autenticado)
	app.Get("/usuario", controladores.Usuario)
	app.Post("/salir", controladores.Salir)
	app.Post("/usuario", controladores.CrearUsuario)
	app.Post("/entrada", controladores.CrearEntrada)
	app.Post("/comentario", controladores.CrearComentario)
	app.Get("/:id", controladores.LeerUsuario)
	app.Get("/comentario/:id", controladores.LeerComentario)
	app.Put("/usuario/:id", controladores.ActualizarUsuario)
	app.Put("/entrada/:id", controladores.ActualizarEntrada)
	app.Put("/comentario/:id", controladores.ActualizarComentario)
	// app.Delete("/usuario/:id", controladores.BorrarUsuario)
	app.Delete("/entrada/:id", controladores.BorrarEntrada)
	app.Delete("/comentario/:id", controladores.BorrarComentario)
	// app.Get("/usuarios", controladores.TodosUsuarios)
}
