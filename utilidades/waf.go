package utilidades

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kitabisa/teler-waf"
	"os"
	"strings"
)

func Waf(h fiber.Handler) fiber.Handler {
	// c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	waf := teler.New(teler.Options{
		Whitelists: []string{
			`request.Headers matches "(curl|Go-http-client|okhttp)/*" && threat == BadCrawler`,
			`request.URI startsWith "/escribeme"`,
			fmt.Sprintf(`request.Headers contains "authorization" && (request.Method == "POST" || request.Method == "PUT" || request.Method == "DELETE") && request.Host == "%s"`, strings.TrimPrefix(os.Getenv("CORS_DOMINIO_PERMITIDO"), "https://")),
			`request.Method == "POST" && request.URI == "/escribeme"`,
		},
	})
	return adaptor.HTTPHandler(waf.Handler(adaptor.FiberHandlerFunc(h)))
}
