package stubbify

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/uptonm/fiber-sandbox/src/config"
)

type StubifyBody struct {
	Destination string `json:"destination"`
}

// HandleStubifyRoute handles the shortening and storing of urls
func HandleStubifyRoute(c *config.Configuration) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		body := new(StubifyBody)

		err := ctx.BodyParser(body)
		if err != nil || len(body.Destination) == 0 {
			return fiber.ErrBadRequest
		}

		stub := Encode(body.Destination)

		err = ctx.JSON(fiber.Map{
			"shortened_url": fmt.Sprintf("%s:%s/stubify/%s", c.Ingress.Host, c.Ingress.Port, stub),
		})
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return nil
	}
}

// HandleGetStubbedRoute handles the retrieval of stored stubbed routes
func HandleGetStubbedRoute(ctx *fiber.Ctx) error {
	stub := ctx.Params("stub")
	if len(stub) == 0 {
		return fiber.ErrBadRequest
	}

	dest, err := Decode(stub)
	if err != nil {
		fmt.Printf("error decoding url: %s\n", err.Error())
		return fiber.ErrNotFound
	}

	err = ctx.Redirect(dest)
	if err != nil {
		fmt.Printf("error redirecting user: %s\n", err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
