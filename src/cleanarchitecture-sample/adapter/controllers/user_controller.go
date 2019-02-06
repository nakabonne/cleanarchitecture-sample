package controllers

import (
	"cleanarchitecture-sample/adapter/gateway"
	"cleanarchitecture-sample/adapter/interfaces"
	"cleanarchitecture-sample/domain"
	"cleanarchitecture-sample/usecase"

	"github.com/jinzhu/gorm"
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
