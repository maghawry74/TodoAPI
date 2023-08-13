package filters

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	fmt.Printf("%v", ctx.Locals("IsAuthunticated"))
	if ctx.Locals("IsAuthunticated") == "true" {
		return ctx.Next()
	}
	return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "UnAuthunticated"}
}

func Authorize(role string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if ctx.Locals("Role") == role {
			return ctx.Next()
		}
		return &fiber.Error{Code: fiber.StatusForbidden, Message: "UnAuthorized"}
	}
}
