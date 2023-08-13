package repository

import types "FirstAPI/Types"

type ITodoRepository interface {
	GetAll() ([]types.Todo, error)
	GetById(id int) (types.Todo, error)
	Add(NewTodo types.Todo) error
	Update(updatedTodo types.Todo) error
	Delete(id int) error
	Toggole(id int) error
}
