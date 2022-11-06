package common

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// CommonClient defines the explicit client to be used for geocoding
type CommonClient struct {
	APIKey     string
	APIBaseURL string

	HTTPClient http.Client
	Validator  validator.Validate
}

func GetClient(config ClientConfig) (*CommonClient, error) {

	val := validator.New()
	err := config.Validate(*val)
	if err != nil {
		return nil, err
	}

	client := CommonClient{
		APIBaseURL: config.APIBaseURL,
		APIKey:     config.APIKey,
		HTTPClient: *http.DefaultClient,
		Validator:  *val,
	}
	return &client, nil
}
