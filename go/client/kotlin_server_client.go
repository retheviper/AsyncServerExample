package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type request struct {
	ID int `json:"id"`
}

type Response struct {
	ID     int    `json:"id"`
	Result string `json:"result"`
}

func Post(ID int) (*Response, error) {
	requestContent := request{ID: ID}
	responseContent := &Response{}

	request, err := json.Marshal(requestContent)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(
		"http://localhost:8900/api/v1/some-process",
		"application/json",
		bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &responseContent)
	if err != nil {
		return nil, err
	}

	return responseContent, nil
}
