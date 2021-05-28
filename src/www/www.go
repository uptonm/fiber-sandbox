package www

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptonm/fiber-sandbox/src/config"
	"github.com/uptonm/fiber-sandbox/src/pkg/stubbify"
)

// WireHandlers handles initial setup of middlewares and routes
func WireHandlers(r *fiber.App, c *config.Configuration, p bool) {
	WireMiddlewares(r, c, p)
	v1 := r.Group("/api/v1")

	r.Get("/stubify/:stub", stubbify.HandleGetStubbedRoute)
	v1.Post("/stubify", stubbify.HandleStubifyRoute(c))
}
