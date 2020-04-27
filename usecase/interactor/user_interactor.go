package interactor

import (
	"fmt"
	"github.com/linda/auth/domain/model"
	"github.com/linda/auth/usecase/presenter"
	"github.com/linda/auth/usecase/repository"

	ir "github.com/linda/auth/interface/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	GetForSignIn(u model.User) (err string,e error)
	GetForSignUp(u model.User)(rm model.User, e error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}


func (us *userInteractor) GetForSignIn(u model.User) (string, error) {
	u, err := us.UserRepository.FindUser(u)
	if err.Error() == ir.NOT_FOUND_USER {
		return "",err
	}
	if err != nil {
		return  "", err
	}

	return us.UserPresenter.ResponseUserSignIn(u)
}


func(us *userInteractor) GetForSignUp(u model.User)(rm model.User,e error){
	result := us.UserRepository.SaveUser(u)
	if result != nil{
		fmt.Print("43::user_interactor")
		return model.User{},result
	}
	return us.UserPresenter.ResponseUserSignUp(u)
}