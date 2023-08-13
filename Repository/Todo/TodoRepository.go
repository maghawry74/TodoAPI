package repository

import (
	types "FirstAPI/Types"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

type TodoRepository struct {
	Db *pgx.Conn
}

func (t *TodoRepository) GetAll() ([]types.Todo, error) {
	rows, err := t.Db.Query(context.Background(), "select id,name,completed,createdat from Todos;")
	if err != nil {
		fmt.Println("error Querying", err)
		return nil, err
	}
	todos := []types.Todo{}
	for rows.Next() {
		todo := &types.Todo{}
		if err := rows.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, *todo)
	}
	return todos, nil
}

func (t *TodoRepository) GetById(id int) (types.Todo, error) {
	todo := types.Todo{}
	err := t.Db.QueryRow(context.Background(), "Select id,name,completed,createdat from Todos where id = $1", id).Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.CreatedAt)
	return todo, err
}

func (t *TodoRepository) Add(NewTodo types.Todo) error {
	NewTodo.CreatedAt = time.Now()
	_, err := t.Db.Exec(context.Background(), "Insert into Todos(name,completed,createdat) values($1,$2,$3)", NewTodo.Name, NewTodo.Completed, NewTodo.CreatedAt)
	return err
}

func (t *TodoRepository) Update(updatedTodo types.Todo) error {
	ct, err := t.Db.Exec(context.Background(), "Update Todos set name=$1,completed=$2 where id=$3", updatedTodo.Name, updatedTodo.Completed, updatedTodo.Id)
	if ct.RowsAffected() == 0 {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: fmt.Sprintf("Todo With Id %d Not Found", updatedTodo.Id)}
	}
	return err
}

func (t *TodoRepository) Delete(id int) error {
	ct, err := t.Db.Exec(context.Background(), "Delete from Todos Where id=$1", id)
	if ct.RowsAffected() == 0 {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: fmt.Sprintf("Todo With Id %d Not Found", id)}
	}
	return err
}

func (t *TodoRepository) Toggole(id int) error {
	ct, err := t.Db.Exec(context.Background(), "Update Todos set completed= not completed where id=$1", id)
	if ct.RowsAffected() == 0 {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: fmt.Sprintf("Todo With Id %d Not Found", id)}
	}
	return err
}
