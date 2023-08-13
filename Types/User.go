package types

type User struct {
	Id        int
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Password  string `validate:"required"`
	Role      string `validate:"required"`
}

type LoginCredentials struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
