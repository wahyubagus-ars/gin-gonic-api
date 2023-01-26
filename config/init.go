package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
	RoleRepo repository.RoleRepository
}

func NewInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,
	}
}
