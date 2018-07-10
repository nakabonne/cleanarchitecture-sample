package controllers

import (
	"github.com/nakabonne/cleanarc-sample/src/adapter/gateway"
	"github.com/nakabonne/cleanarc-sample/src/adapter/interfaces"
	"github.com/nakabonne/cleanarc-sample/src/domain"
	"github.com/nakabonne/cleanarc-sample/src/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(handler interfaces.DBHandler) *UserController {
	return &UserController{
		Interactor: usercase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				DBHandler: handler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
}
