package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nd-back/bbdd"
	"nd-back/rutas"
	"os"
)

func main() {
	bbdd.Conectar()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin, Authorization",
		AllowOrigins:     "https://nd-front-945b3a8a4358.herokuapp.com, https://www.nuestrodiario.es, http://localhost:3000",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
		MaxAge:           86400,
	}))
	rutas.Configuracion(app)
	app.Listen(":" + os.Getenv("PORT"))
}
