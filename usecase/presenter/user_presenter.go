package presenter

import "github.com/linda/auth/domain/model"


type UserPresenter interface {
	ResponseUserSignIn(us model.User) (string,error)
	ResponseUserSignUp(us model.User) (model.User,error)
}