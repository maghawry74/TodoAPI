package handler

import (
	repository "FirstAPI/Repository/Todo"
	types "FirstAPI/Types"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

type TodoHandler struct {
	Repo repository.ITodoRepository
}

func (t *TodoHandler) GetAll(ctx *fiber.Ctx) error {
	todos, err := t.Repo.GetAll()
	if err != nil {
		return err
	}
	ctx.JSON(todos)
	return nil
}
func (t *TodoHandler) GetTodobyId(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	todo, err := t.Repo.GetById(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &fiber.Error{Code: 404, Message: fmt.Sprintf("Todo With Id %d Not Found", id)}
		}
		return err
	}
	ctx.JSON(todo)
	return nil
}
func (t *TodoHandler) AddTodo(ctx *fiber.Ctx) error {
	todo := types.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return err
	}
	fmt.Printf("%v", todo)
	validate := validator.New()
	err := validate.Struct(todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  err.Error(),
		})
	}
	err = t.Repo.Add(todo)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusCreated)
	return nil
}

func (t *TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	todo := types.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  err.Error(),
		})
	}
	err = t.Repo.Update(todo)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusNoContent)
	return nil
}

func (t *TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	err = t.Repo.Delete(id)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusNoContent)
	return nil
}
func (t *TodoHandler) ToggoleTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	err = t.Repo.Toggole(id)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusNoContent)
	return nil
}
