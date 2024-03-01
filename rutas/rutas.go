package rutas

import (
	"github.com/gofiber/fiber/v2"
	"nd-back/controladores"
	"nd-back/middlewares"
)

func Configuracion(app *fiber.App) {
	app.Post("/api/registrar", controladores.Registrar)
	app.Post("/api/ingresar", controladores.Ingresar)
	app.Get("/api/entrada/:id", controladores.LeerEntrada)
	app.Get("/api/entradas", controladores.TodasEntradas)   // Muestra las entradas paginadas
	app.Get("/api/todas", controladores.EntradasSinPaginar) // Muestra las entradas paginadas
	app.Get("/api/especial", controladores.TodasEntradas)   // Muestra las entradas ver parametros del query
	app.Get("/api/comentarios", controladores.TodosComentarios)
	app.Post("/api/KzQ987h29KkYem", controladores.Escribeme)
	app.Use(middlewares.Autenticado)
	app.Get("/api/usuario", controladores.Usuario)
	app.Post("/api/salir", controladores.Salir)
	app.Post("/api/usuario", controladores.CrearUsuario)
	app.Post("/api/entrada", controladores.CrearEntrada)
	app.Post("/api/comentario", controladores.CrearComentario)
	app.Get("/api/usuario/:id", controladores.LeerUsuario)
	app.Get("/api/comentario/:id", controladores.LeerComentario)
	app.Put("/api/usuario/:id", controladores.ActualizarUsuario)
	app.Put("/api/entrada/:id", controladores.ActualizarEntrada)
	app.Put("/api/comentario/:id", controladores.ActualizarComentario)
	app.Delete("/api/usuario/:id", controladores.BorrarUsuario)
	app.Delete("/api/entrada/:id", controladores.BorrarEntrada)
	app.Delete("/api/comentario/:id", controladores.BorrarComentario)
	app.Get("/api/usuarios", controladores.TodosUsuarios)

}
