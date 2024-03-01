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
		AllowOrigins:     "https://www.nuestrodiario.es, http://localhost:3000, http://localhost:8000,",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	rutas.Configuracion(app)
	app.Listen(":" + os.Getenv("PORT"))
}
