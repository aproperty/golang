package geocoding

import (
	"context"

	"temporary/geocoding/geo"
)

// CommonClientInterface defines the necessary contract for a geocoding client
type CommonClientInterface interface {
	GeocodeAddress(ctx context.Context, request geo.GeoRequest) (*geo.GeocodeAddressResponse, error)
}
