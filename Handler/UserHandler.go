package handler

import (
	repository "FirstAPI/Repository/User"
	types "FirstAPI/Types"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func (U *UserHandler) Register(ctx *fiber.Ctx) error {
	newUser := types.User{}
	if err := ctx.BodyParser(&newUser); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(newUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  err.Error(),
		})
	}
	if err := U.UserRepository.Register(newUser); err != nil {
		return err
	}
	ctx.Status(fiber.StatusCreated)
	return nil
}

func (U *UserHandler) Login(ctx *fiber.Ctx) error {
	loginCred := types.LoginCredentials{}
	if err := ctx.BodyParser(&loginCred); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(loginCred); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  err.Error(),
		})
	}
	user, err := U.UserRepository.Login(loginCred)
	if err != nil {
		if err.Error() == "no rows in result set" {
			ctx.Status(fiber.StatusUnauthorized)
			return nil
		}
		return err
	}

	claims := types.UserClaim{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
			ID:        fmt.Sprint(user.Id),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SecretKey")))
	if err != nil {
		return err
	}
	ctx.JSON(map[string]string{"Token": signedToken})
	return nil
}
