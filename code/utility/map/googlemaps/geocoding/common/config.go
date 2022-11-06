package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ClientConfig ...
type ClientConfig struct {
	APIKey     string `validate:"required" env:"GEOCODING_SERVICE_API_KEY"`
	APIBaseURL string `validate:"required,url" env:"GEOCODING_SERVICE_BASE_URL"`
}

// Validate validates the configuration struct
func (g ClientConfig) Validate(val validator.Validate) error {

	// perform validations over struct
	if err := val.Struct(g); err != nil {

		return fmt.Errorf("ClientConfig.Validate.error: %v", err)
	}

	return nil
}
