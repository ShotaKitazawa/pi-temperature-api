package rest

import (
	"bytes"
	"net/http"
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
