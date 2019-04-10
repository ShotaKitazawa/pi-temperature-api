package usecase

import (
	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/usecase/interfaces"
)

type OutputInteractor struct {
	OutputRepository interfaces.OutputRepository
	Logger         interfaces.Logger
}

func (i *OutputInteractor) Post(input *domain.Output) (int, error) {
	i.Logger.Log("post state data!")
	return i.OutputRepository.Post(input)
}
