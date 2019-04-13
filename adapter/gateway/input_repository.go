package gateway

import (
	"github.com/ShotaKitazawa/pi-temperature-api/domain"

	"github.com/jinzhu/gorm"
)

type (
	InputRepository struct {
		Conn *gorm.DB
	}

	Input struct {
		gorm.Model
		// Id          int
		Temperature float64
		Humidity    float64
		//CreatedAt   *time.Time
	}
)

func (r *InputRepository) Get1() (*domain.Input, error) {
	input := Input{}

	if err := r.Conn.First(&input); err.Error != nil {
		return nil, err.Error
	}
	return &domain.Input{
		//ID:          input.ID,
		Temperature: input.Temperature,
		Humidity:    input.Humidity,
		//CreatedAt:   input.CreatedAt,
	}, nil
}

/*
func (r *InputRepository) Get10(u *domain.Input) (id int, err error) {
	input := new(Input)

	if err = r.Conn.Create(input).Error; err != nil {
		return
	}

	return int(input.ID), nil
}

/*
func (r *InputRepository) FindByName(name string) (d []domain.Input, err error) {
	users := []Input{}
	if err = r.Conn.Where("name = ?", name).Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]domain.Input, n)
	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].Name = users[i].Name
		d[i].Email = users[i].Email
	}
	return
}

func (r *InputRepository) FindAll() (d []domain.Input, err error) {
	users := []Input{}
	if err = r.Conn.Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]domain.Input, n)
	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].Name = users[i].Name
		d[i].Email = users[i].Email
	}
	return
}
*/
