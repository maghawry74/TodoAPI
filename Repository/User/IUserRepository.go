package repository

import types "FirstAPI/Types"

type IUserRepository interface {
	Register(NewUser types.User) error
	Login(Logincred types.LoginCredentials) (types.User,error)
}
