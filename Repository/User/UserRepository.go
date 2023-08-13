package repository

import (
	types "FirstAPI/Types"
	"context"

	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	Db *pgx.Conn
}

func (u *UserRepository) Register(NewUser types.User) error {
	_, err := u.Db.Exec(context.Background(), "Insert into Users(firstname, lastname, role, password, email) Values($1,$2,$3,$4,$5)", NewUser.FirstName, NewUser.LastName, NewUser.Role, NewUser.Password, NewUser.Email)
	return err
}
func (u *UserRepository) Login(cred types.LoginCredentials) (types.User, error) {
	user := types.User{}
	err := u.Db.QueryRow(context.Background(), "Select id,firstname,lastname,email,role from Users where email=$1 and password=$2", cred.Email, cred.Password).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Role)
	return user, err
}
