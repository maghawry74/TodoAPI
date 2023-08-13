package main

import (
	filters "FirstAPI/Filters"
	handler "FirstAPI/Handler"
	middlewares "FirstAPI/Middlewares"
	todoRepository "FirstAPI/Repository/Todo"
	userRepository "FirstAPI/Repository/User"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DBURL"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())
	todoHandler := handler.TodoHandler{
		Repo: &todoRepository.TodoRepository{
			Db: conn,
		},
	}
	userHandler := handler.UserHandler{
		UserRepository: userRepository.UserRepository{
			Db: conn,
		},
	}
	server := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(map[string]string{"Error": err.Error()})
		},
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))
	server.Use(logger.New())
	server.Use(middlewares.Authunticate)

	app := server.Group("/Api/v1")
	app.Get("/Todo", todoHandler.GetAll)
	app.Get("/Todo/:id", todoHandler.GetTodobyId)
	app.Post("/Todo",filters.Authenticate ,todoHandler.AddTodo)
	app.Put("/Todo",filters.Authenticate ,todoHandler.UpdateTodo)
	app.Patch("/Todo/:id",filters.Authenticate, todoHandler.ToggoleTodo)
	app.Delete("/Todo/:id",filters.Authorize("Admin"), todoHandler.DeleteTodo)
	app.Post("/User", userHandler.Register)
	app.Post("/User/Login", userHandler.Login)
	log.Fatal(server.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
