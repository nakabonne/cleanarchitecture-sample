package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/nakabonne/cleanarchitecture-sample/pkg/adapter/gateway"
	"github.com/nakabonne/cleanarchitecture-sample/pkg/adapter/interfaces"
	"github.com/nakabonne/cleanarchitecture-sample/pkg/domain"
	"github.com/nakabonne/cleanarchitecture-sample/pkg/usecase"
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
	type (
		Request struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		Response struct {
			UserID int `json:"user_id"`
		}
	)
	req := Request{}
	c.Bind(&req)
	user := domain.User{Name: req.Name, Email: req.Email}

	id, err := controller.Interactor.Add(user)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot add user"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{UserID: id}
	c.JSON(201, res)
}
