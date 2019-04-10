package gateway

import (
	"time"

	"github.com/ShotaKitazawa/pi-temperature-api/domain"

	"github.com/jinzhu/gorm"
)

type (
	OutputRepository struct {
		Conn *gorm.DB
	}

	Output struct {
		gorm.Model
		ID          int
		Temperature float64
		Humidity    float64
		CreatedAt   *time.Time
	}
)

func (r *OutputRepository) Post(*domain.Output) (int, error) {
	// TODO
	input := new(Output)

	if err := r.Conn.Create(input); err != nil {
		return 0, err.Error
	}
	return input.ID, nil
}

