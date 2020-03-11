package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client facilitates making HTTP requests to the GitHub API
type Client struct {
	url  string
	http *http.Client
}

// NewClient initializes a Client
func NewClient(url string) *Client {
	tr := http.DefaultTransport
	http := &http.Client{Transport: tr}
	client := &Client{http: http, url: url}
	return client
}

func handleResponse(resp *http.Response, data interface{}) error {
	success := resp.StatusCode >= 200 && resp.StatusCode < 300

	if !success {
		return handleHTTPError(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	gr := &GraphQLResponse{Data: data}
	err = json.Unmarshal(body, &gr)

	if err != nil {
		return err
	}

	if len(gr.Errors) > 0 {
		return &GraphQLErrorResponse{Errors: gr.Errors}
	}
	return nil
}

func handleHTTPError(resp *http.Response) error {
	var message string
	var parsedBody struct {
		Message string `json:"message"`
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		message = string(body)
	} else {
		message = parsedBody.Message
	}
	fmt.Println(string(body))
	return fmt.Errorf("http error, '%s' failed (%d): '%s'", resp.Request.URL, resp.StatusCode, message)
}
