package gateway

import (
	"net/http"

	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/utils"

	"github.com/jinzhu/gorm"
)

type (
	OutputRepository struct {
		DBConn   *gorm.DB
		HttpConn *http.Request
	}

	Output struct {
		gorm.Model
		State string
	}
)

func (r *OutputRepository) Post(o *domain.Output) (int, error) {
	resp, err := utils.HttpPost(r.HttpConn, o.State)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func (r *OutputRepository) Store(o *domain.Output) (int, error) {
	output := &Output{
		State: o.State,
	}

	if err := r.DBConn.Create(output).Error; err != nil {
		return 0, err
	}

	return int(output.ID), nil
}
