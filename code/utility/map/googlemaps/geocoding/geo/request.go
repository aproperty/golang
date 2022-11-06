package geo

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// GeoRequest defines the request message for geocoding
type GeoRequest struct {
	AddressLine1 *string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`

	City    *string `json:"city"`
	State   *string `json:"state"`
	Zipcode *string `json:"zipcode"`
}

// Validate validates the request struct
func (g GeoRequest) Validate(val validator.Validate) error {

	if g.AddressLine1 == nil &&
		g.AddressLine2 == nil &&
		g.City == nil &&
		g.State == nil &&
		g.Zipcode == nil {

		return fmt.Errorf("GeoRequest.Validate.error: %v", "request empty")
	}

	return nil
}
