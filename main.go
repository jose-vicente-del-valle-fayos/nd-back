package main

import (
	"fmt"
	"github.com/corazawaf/coraza/v3"
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
		waf, err := coraza.NewWAF(coraza.NewWAFConfig().WithDirectivesFromFile("./coraza.conf").WithDirectivesFromFile("./crs-setup.conf").WithDirectivesFromFile("./rules/*.conf"))
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		tx := waf.NewTransaction()
		defer func() {
			tx.ProcessLogging()
			err = tx.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()
		tx.ProcessConnection(c.IP(), 443, "216.24.57.4", 10000)
		if it1 := tx.ProcessRequestHeaders(); it1.Status != 0 {
			fmt.Printf("transacción durante el procesamiento de headers interrumpida con estado %d\n", it1.Status)
			fmt.Printf(it1.Data)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		it2, err := tx.ProcessRequestBody()
		if it2.Status != 0 || err != nil {
			fmt.Printf("transacción durante el procesamiento de body interrumpida con estado %d\n", it2.Status)
			fmt.Printf(it2.Data)
			fmt.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Next()
	})
	app.Use(func(c *fiber.Ctx) error {
		if c.Hostname() != os.Getenv("HOSTNAME_PERMITIDO") {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin, Authorization",
		AllowOrigins:     os.Getenv("CORS_DOMINIO_PERMITIDO"), // http://localhost:3000
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
