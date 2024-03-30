package rutas

import (
	"github.com/gofiber/fiber/v2"
	"nd-back/controladores"
	"nd-back/middlewares"
	"nd-back/utilidades"
)

func Configuracion(app *fiber.App) {
	// app.Post("/registrar", utilidades.Waf(controladores.Registrar))
	app.Post("/ingresar", utilidades.Waf(controladores.Ingresar))
	app.Get("/entrada/:id", utilidades.Waf(controladores.LeerEntrada))
	app.Get("/entradas", utilidades.Waf(controladores.TodasEntradas))   // Muestra las entradas paginadas
	app.Get("/todas", utilidades.Waf(controladores.EntradasSinPaginar)) // Muestra las entradas sin paginar
	app.Get("/especial", utilidades.Waf(controladores.TodasEntradas))   // Muestra las entradas ver parametros del query
	app.Get("/comentarios", utilidades.Waf(controladores.TodosComentarios))
	app.Post("/escribeme", utilidades.Waf(controladores.Escribeme))
	app.Use(middlewares.Autenticado)
	app.Get("/usuario", utilidades.Waf(controladores.Usuario))
	app.Post("/salir", utilidades.Waf(controladores.Salir))
	app.Post("/usuario", utilidades.Waf(controladores.CrearUsuario))
	app.Post("/entrada", utilidades.Waf(controladores.CrearEntrada))
	app.Post("/comentario", utilidades.Waf(controladores.CrearComentario))
	app.Get("/:id", utilidades.Waf(controladores.LeerUsuario))
	app.Get("/comentario/:id", utilidades.Waf(controladores.LeerComentario))
	app.Put("/usuario/:id", utilidades.Waf(controladores.ActualizarUsuario))
	app.Put("/entrada/:id", utilidades.Waf(controladores.ActualizarEntrada))
	app.Put("/comentario/:id", utilidades.Waf(controladores.ActualizarComentario))
	// app.Delete("/usuario/:id", utilidades.Waf(controladores.BorrarUsuario))
	app.Delete("/entrada/:id", utilidades.Waf(controladores.BorrarEntrada))
	app.Delete("/comentario/:id", utilidades.Waf(controladores.BorrarComentario))
	app.Get("/usuarios", utilidades.Waf(controladores.TodosUsuarios))

}
