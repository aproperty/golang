package regeo

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// RegeoRequest defines the response type for geocoding
type RegeoRequest struct {
	Latitude  float64 `validate:"required,max=90,min=-90" json:"latitude"`
	Longitude float64 `validate:"required,max=180,min=-180" json:"longitude"`
}

// Validate validates the response struct
func (g RegeoRequest) Validate(val validator.Validate) error {

	// perform validations over struct
	if err := val.Struct(g); err != nil {

		return fmt.Errorf("RegeoRequest.Validate.error: %v", err)
	}
	return nil
}
