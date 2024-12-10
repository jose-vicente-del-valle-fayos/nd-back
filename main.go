package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nd-back/bbdd"
	"nd-back/rutas"
	"os"
)

func main() {
	bbdd.Conectar()
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		if c.Hostname() != os.Getenv("HOSTNAME_PERMITIDO") {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin, Authorization",
		AllowOrigins:     os.Getenv("CORS_DOMINIO_PERMITIDO_1") + ", " + os.Getenv("CORS_DOMINIO_PERMITIDO_2"), // http://localhost:3000,
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE",
		MaxAge:           86400,
	}))
	rutas.Configuracion(app)
	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
	}
}
