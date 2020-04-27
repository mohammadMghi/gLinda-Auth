package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/linda/auth/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindUser(u model.User) (model.User, error)
	SaveUser(u model.User)(e error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
//store,find data in database !
func (ur *userRepository)SaveUser(u model.User)( e error){
	err := ur.db.Save(&u)
	if err != nil{
		return err.Error
	}
	return nil
}

func (ur *userRepository) FindUser(u model.User) (model.User, error) {
	println(&u)
	result := ur.db.Where("email = ?", &u.Email).First(&u)

	if result.RowsAffected <= 0{
		 return model.User{},errors.New(NOT_FOUND_USER)
	}

	if result != nil {
		return model.User{},result.Error
	}

	return u, nil
}

const (
	NOT_FOUND_USER = "NotFoundUser"
)
