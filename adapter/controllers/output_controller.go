package controllers

import (
	"github.com/ShotaKitazawa/pi-temperature-api/usecase"
)

type OutputController struct {
	Interactor usecase.OutputInteractor
}
