package interfaces

import "github.com/ShotaKitazawa/pi-temterature-api/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindByName(string) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}
