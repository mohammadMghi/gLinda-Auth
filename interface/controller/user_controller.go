package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/linda/auth/domain/model"

	"github.com/linda/auth/usecase/interactor"

	ir "github.com/linda/auth/interface/repository"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) SignIn(c echo.Context) (err error) {

	userData := new(model.User)
	if err = c.Bind(userData); err != nil {
		return
	}

	//invalidation data before pass
 	if !(userData.Email != "" && userData.Password != ""){
 		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Username or Password is empty"}
	}
	if !(len(userData.Email) <= 100 && len(userData.Password) <= 200){
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Username or Password len's is so long !"}
	}


	token, err := uc.userInteractor.GetForSignIn(*userData)

	if err != nil {
		if err.Error() == ir.NOT_FOUND_USER{
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Email or password is wrong please try with other or sign up!"}
		}
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func (uc *userController)SignUp(c echo.Context) (err error){
	userData := new(model.User)
	//invalidation data before pass
	if err = c.Bind(userData); err != nil {
		return
	}
	if !(userData.Email != "" && userData.Password != ""){
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Username or Password is empty"}
	}
	if !(len(userData.Email) <= 100 && len(userData.Password) <= 200){
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Username or Password len's is so long !"}
	}
	u,err := uc.userInteractor.GetForSignUp(*userData)
	if err != nil{
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	return c.JSON(http.StatusOK ,u)
}