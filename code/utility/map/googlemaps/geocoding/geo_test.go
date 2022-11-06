package geocoding

import (
	"context"
	"testing"
	"time"

	"temporary/geocoding/common"
	"temporary/geocoding/geo"
)

func TestClient(t *testing.T) {

	g := common.ClientConfig{
		APIKey:     "",
		APIBaseURL: "https://maps.googleapis.com/maps/api/geocode/json",
	}

	var commonClient CommonClientInterface
	commonClient, err := geo.GetGeoClient(g)
	if err != nil {
		t.Errorf("ClientConfig.Validate() error = %v", err)
	}

	addressRequest := new(geo.GeoRequest)
	city := "深圳"
	addressRequest.City = &city

	ctx, cancel := context.WithCancel(context.Background())
	res, err := commonClient.GeocodeAddress(ctx, *addressRequest)
	if err != nil {
		t.Errorf("GeocodeAddress error = %v", err)
	}

	time.Sleep(time.Second * 3)
	cancel()

	t.Logf("%+v", res)
}
