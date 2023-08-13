package types

import (
	"time"
)

type Todo struct {
	Id        int
	Name      string `validate:"required"`
	Completed bool
	CreatedAt time.Time
}
