package repository

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/linda/auth/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindUser(u model.User) (model.User, error)
	SaveUser(u model.User)(err error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
//store,find data in database !
func (ur *userRepository)SaveUser(u model.User)(err error){
    hashPass := sha256.Sum256([]byte(u.Password + "this is salt for this enc"))
	err = ur.db.Create(&model.User{Email: u.Email,Password:base64.URLEncoding.EncodeToString(hashPass[:])}).Error
	if err != nil{
		fmt.Print("32::user_repository")
		return err
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
	ERROR_SIGN_UP = "SignUpError"
)

