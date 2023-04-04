package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hungtg7/backend-app/app/authen/config"
)

type consumerBody struct {
	username string `json:"username"`
}

func grant_athorize_service_to_user(user_id string)

func create_consumer_for_new_comer(ctx context.Context, user_id string) error {
	body := consumerBody{
		username: user_id,
	}
	// Encode the struct to a JSON byte slice
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return err
	}

	// Wrap the byte slice in a bytes.Reader
	ioBody := bytes.NewReader(buf.Bytes())
	// Construct request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, config.AppServerConfig().KongAddr, ioBody)

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client and send the request
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Handle the response
	if response.StatusCode != http.StatusOK {
		// Handle the error
		fmt.Printf("Error: %v", response)
	}

	// Handle the response body
	responseBody := make([]byte, response.ContentLength)
	_, err = response.Body.Read(responseBody)
	if err != nil {
		return err
	}

	return nil
}

func get_authorize_code() {}

func get_access_token() {}

func refresh_access_token() {}
