package rest

import (
	"bytes"
	"net/http"

	"github.com/ShotaKitazawa/pi-temperature-api/utils"
)

//var db *gorm.DB
var req *http.Request

func Request(target string) *http.Request {
	var err error

	req, err := http.NewRequest(
		"POST",
		target,
		bytes.NewBuffer([]byte("{}")),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}
	return req
}

func GetEnv() string {
	return utils.GetEnvOrDefault("REST_URL", "http://localhost:8000/")
}
