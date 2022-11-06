package geo

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"temporary/geocoding/common"
)

type GeoClient struct {
	common.CommonClient
}

func GetGeoClient(config common.ClientConfig) (*GeoClient, error) {
	commonClient, err := common.GetClient(config)
	if err != nil {
		return nil, err
	}
	client := GeoClient{
		CommonClient: *commonClient,
	}
	return &client, nil
}

// GeocodeAddress implements the geocoding functionality
func (g GeoClient) GeocodeAddress(ctx context.Context, request GeoRequest) (*GeocodeAddressResponse, error) {

	err := request.Validate(g.Validator)
	if err != nil {
		return nil, err
	}

	formattedAddress := ""
	if request.AddressLine1 != nil {
		formattedAddress = formattedAddress + *request.AddressLine1 + ", "
	}
	if request.AddressLine2 != nil {
		formattedAddress = formattedAddress + *request.AddressLine2 + ", "
	}

	if request.City != nil {
		formattedAddress = formattedAddress + *request.City + ", "
	}
	if request.State != nil {
		formattedAddress = formattedAddress + *request.State + ", "
	}
	if request.Zipcode != nil {
		formattedAddress = formattedAddress + *request.Zipcode + ", "
	}

	req, err := g.getGeoRequest(formattedAddress)
	if err != nil {
		return nil, err
	}
	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var geocodeResponse GeocodingResponse
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &geocodeResponse)
	if err != nil {
		return nil, err
	}

	return &GeocodeAddressResponse{
		Latitude:  geocodeResponse.Results[0].Geometry.Location.Lat,
		Longitude: geocodeResponse.Results[0].Geometry.Location.Lng,
	}, nil
}

// getGeoRequest gets the formatted  request
func (g GeoClient) getGeoRequest(formattedAddress string) (*http.Request, error) {
	req, err := http.NewRequest("GET", g.APIBaseURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", g.APIKey)
	q.Add("address", formattedAddress)

	req.URL.RawQuery = q.Encode()

	return req, nil
}
