package utilidades

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kitabisa/teler-waf"
)

func Waf(h fiber.Handler) fiber.Handler {
	waf := teler.New()
	return adaptor.HTTPHandler(waf.Handler(adaptor.FiberHandlerFunc(h)))
}
