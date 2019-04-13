package gateway

import (
	"net/http"
	"time"

	"github.com/ShotaKitazawa/pi-temperature-api/domain"
	"github.com/ShotaKitazawa/pi-temperature-api/utils"

	"github.com/jinzhu/gorm"
)

type (
	OutputRepository struct {
		Conn *http.Request
	}

	Output struct {
		gorm.Model
		ID          int
		Temperature float64
		Humidity    float64
		CreatedAt   *time.Time
	}
)

func (r *OutputRepository) Post(o *domain.Output) (int, error) {
	resp, err := utils.HttpPost(r.Conn, o.State)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
