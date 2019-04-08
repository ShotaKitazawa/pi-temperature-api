package usecase

import (
	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	Logger         interfaces.Logger
}

func (i *UserInteractor) Add(u domain.User) (int, error) {
	i.Logger.Log("store user!")
	return i.UserRepository.Store(u)
}
