package presenter

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/linda/auth/domain/model"
	"time"
)

type userPresenter struct {
}



type UserPresenter interface {
	ResponseUserSignIn(us model.User) (string,error)
	ResponseUserSignUp(us model.User) error
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUserSignIn(u model.User) (s string,err error) {
	//handles user ,convert data before pass to view
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(model.Key))
	if err != nil {
		return "",err
	}
	return u.Token,nil
}
func (up *userPresenter) ResponseUserSignUp(us model.User) error {
	//handles user ,convert data before pass to view

	return nil
}