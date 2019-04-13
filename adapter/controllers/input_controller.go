package controllers

import (
	"fmt"
	"time"

	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/interfaces"
	"github.com/ShotaKitazawa/pi-temperature-api/usecase"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
)

type InputController struct {
	Interactor usecase.InputInteractor
}

func NewInputController(dbConn *gorm.DB, logger interfaces.Logger) *InputController {
	return &InputController{
		Interactor: usecase.InputInteractor{
			InputRepository: &gateway.InputRepository{
				DBConn: dbConn,
			},
			Logger: logger,
		},
	}
}

func (controller *InputController) Get1(c interfaces.Context) {
	type (
		Request struct {
		}
		Response struct {
			ID          uint       `json:"id"`
			Temperature float64    `json:"temperature"`
			Humidity    float64    `json:"humidity"`
			CreatedAt   *time.Time `json:"created_at"`
		}
	)
	req := Request{}
	c.Bind(&req)

	data, err := controller.Interactor.Get1()
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "input_controller: cannot get data"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{
		ID:          data.ID,
		Temperature: data.Temperature,
		Humidity:    data.Humidity,
		CreatedAt:   data.CreatedAt,
	}
	c.JSON(201, res)
}

/*
func (controller *InputController) Get10(c interfaces.Context) {
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
	input := new(domain.Input)

	//id, err := controller.Interactor.Add(user)
	id, err := controller.Interactor.Get1(input)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot add user"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{UserID: id}
	c.JSON(201, res)
}
*/
