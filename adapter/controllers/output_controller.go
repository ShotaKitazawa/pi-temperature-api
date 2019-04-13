package controllers

import (
	"fmt"
	"net/http"

	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/interfaces"
	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/usecase"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type OutputController struct {
	Interactor usecase.OutputInteractor
}

func NewOutputController(dbConn *gorm.DB, httpConn *http.Request, logger interfaces.Logger) *OutputController {
	return &OutputController{
		Interactor: usecase.OutputInteractor{
			OutputRepository: &gateway.OutputRepository{
				DBConn:   dbConn,
				HttpConn: httpConn,
			},
			Logger: logger,
		},
	}
}

func (controller *OutputController) Post(c interfaces.Context) {

	type (
		Request struct {
			Value string `json:"value"`
		}
		Response struct {
		}
	)
	req := Request{}
	c.Bind(&req)
	status := &domain.Output{State: req.Value}

	sc, err := controller.Interactor.Post(status)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "output_controller: cannot post data"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	if sc >= 400 {
		msg := fmt.Sprintf("output_controller: POST StatusCode is %d", sc)
		controller.Interactor.Logger.Log(msg)
		c.JSON(500, NewError(500, msg))
		return
	}

	res := Response{}
	c.JSON(201, res)
}
