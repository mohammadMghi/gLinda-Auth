package repository

import "github.com/linda/auth/domain/model"

type UserRepository interface {
	FindUser(u model.User) (model.User, error)
	SaveUser(u model.User) (e error)
}
