package geo

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// GeocodingResponse defines the result of an API call to the Google geocoding service
type GeocodingResponse struct {
	Results []GeocodingResult `json:"results"`
	Status  string            `json:"status"`
}

// GeocodingResult ...
type GeocodingResult struct {
	AddressComponents []GeocodingAddressComponents `json:"address_components"`
	FormattedAddress  string                       `json:"formatted_address"`
	Geometry          GeocodingGeometry            `json:"geometry"`
	PlaceID           string                       `json:"place_id"`
	Types             []string                     `json:"types"`
}

// GeocodingAddressComponents ...
type GeocodingAddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// GeocodingGeometry ...
type GeocodingGeometry struct {
	Bounds       Bounds   `json:"bounds"`
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

// Location ...
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Geometry ...
// type Geometry struct {
// 	Location     Location `json:"location"`
// 	LocationType string   `json:"location_type"`
// 	Viewport     Viewport `json:"viewport"`
// }

// Viewport ...
type Viewport struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}

// Bounds ...
type Bounds struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}

// Northeast ...
type Northeast struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Southwest ...
type Southwest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// GeocodeAddressResponse defines the response type for geocoding
type GeocodeAddressResponse struct {
	Latitude  float64 `validate:"required,max=90,min=-90" json:"latitude"`
	Longitude float64 `validate:"required,max=180,min=-180" json:"longitude"`
}

// Validate validates the response struct
func (g GeocodeAddressResponse) Validate(val validator.Validate) error {
	if err := val.Struct(g); err != nil {
		return fmt.Errorf("GeocodeAddressResponse.Validate.error: %v", err)
	}
	return nil
}
