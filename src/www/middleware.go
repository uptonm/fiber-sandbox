package www

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/uptonm/fiber-sandbox/src/config"
	"strings"
)

const logFormatProd = "${ip} ${header:x-forwarded-for} ${header:x-real-ip} " +
	"[${time}] ${pid} ${locals:requestid} \"${method} ${path} ${protocol}\" " +
	"${status} ${latency} \"${referrer}\" \"${ua}\"\n"

const logFormatDev = "${ip} [${time}] \"${method} ${path} ${protocol}\" " +
	"${status} ${latency}\n"

// WireMiddlewares is a function which handles the initialization and chaining of fiber middlewares
func WireMiddlewares (r *fiber.App, c *config.Configuration, prod bool) {
	r.Use(requestid.New())

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: parseOrigins(c.Ingress.CorsOrigins),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// STDOUT request logger
	r.Use(logger.New())
}

// ParseOrigins accepts an []string of allow origins and formats them in a way accepted by the cors middleware
func parseOrigins(allowOrigins []string) string {
	return strings.Join(allowOrigins, ", ")
}

func logFormat(prod bool) string {
	if prod {
		return logFormatProd
	}
	return logFormatDev
}