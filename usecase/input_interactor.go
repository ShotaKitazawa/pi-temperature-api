package usecase

import (
	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/usecase/interfaces"
)

type InputInteractor struct {
	InputRepository interfaces.InputRepository
	Logger          interfaces.Logger
}

func (i *InputInteractor) Get1() (*domain.Input, error) {
	i.Logger.Log("get data!")
	return i.InputRepository.Get1()
}

/*
func (i *InputInteractor) Get10(input *domain.Input) (int, error) {
	i.Logger.Log("get data!")
	return i.InputRepository.Get10(input)
}
*/
