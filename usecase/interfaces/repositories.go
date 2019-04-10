package interfaces

import "github.com/ShotaKitazawa/pi-temperature-api/domain"

type InputRepository interface {
	Get1() (*domain.Input, error)
	//Get10(* domain.Input) (int, error)
}

type OutputRepository interface {
	Post(*domain.Output) (int, error)
}
