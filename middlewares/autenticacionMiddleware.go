package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"nd-back/utilidades"
)

func Autenticado(c *fiber.Ctx) error {
	cookie := c.Cookies("nd-jwt")
	if _, err := utilidades.ParsearJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"Mensaje": "Sin autenticar.",
		})
	}
	return c.Next()
}
