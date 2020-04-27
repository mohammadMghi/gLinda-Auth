package registry

import (
	"github.com/linda/auth/interface/controller"
	ip "github.com/linda/auth/interface/presenter"
	ir "github.com/linda/auth/interface/repository"
	"github.com/linda/auth/usecase/interactor"
	up "github.com/linda/auth/usecase/presenter"
	ur "github.com/linda/auth/usecase/repository"
)



func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
