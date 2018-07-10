package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/nakabonne/cleanarc-sample/src/adapter/gateway"
	"github.com/nakabonne/cleanarc-sample/src/adapter/interfaces"
	"github.com/nakabonne/cleanarc-sample/src/domain"
	"github.com/nakabonne/cleanarc-sample/src/usecase"
	"github.com/pkg/errors"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(conn *gorm.DB, logger interfaces.Logger) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				Conn: conn,
			},
			Logger: logger,
		},
	}
}

func (controller *UserController) Create(c interfaces.Context) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot add user"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	c.JSON(201, user)
}
