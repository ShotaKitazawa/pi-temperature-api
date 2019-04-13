package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpPost(req *http.Request, bodyStr string) (*http.Response, error) {
	body := bytes.NewBuffer([]byte(bodyStr))
	rc, ok := io.Reader(body).(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}

	req.Body = rc
	req.ContentLength = int64(len(bodyStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
