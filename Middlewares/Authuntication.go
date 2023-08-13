package middlewares

import (
	types "FirstAPI/Types"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authunticate(ctx *fiber.Ctx) error {
	tokenBytes := ctx.Request().Header.Peek("Authorization")
	tokenArray := strings.Split(string(tokenBytes), "Bearer")
	StringToken := strings.TrimSpace(strings.Join(tokenArray, ""))
	if StringToken == "" {
		return ctx.Next()
	}
	token, err := jwt.ParseWithClaims(StringToken, &types.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SecretKey")), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fiber.ErrUnauthorized
	}
	claims, ok := token.Claims.(*types.UserClaim)
	if !ok {
		return &fiber.Error{Code: 400, Message: "Invalid Claims Type"}
	}
	ctx.Locals("IsAuthunticated", "true")
	ctx.Locals("Role", claims.Role)
	ctx.Locals("Email", claims.Email)
	ctx.Locals("Id", claims.ID)
	return ctx.Next()
}
