package regeo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const AmapKey string = "27780a1f06c4a3817fef6b67b21ef123"
const BaseUrl string = "https://restapi.amap.com/v3/geocode/regeo"

func GetReGeoByAddress(location string) (address string, province string, city string, err error) {

	url := fmt.Sprintf("%s?location=%s&key=%s",
		BaseUrl,
		location,
		AmapKey)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	amapReGeo := AmapReGeo{}
	err = jsoniter.Unmarshal(body, &amapReGeo)
	if err != nil {
		return
	}

	if amapReGeo.Status == "1" {
		if amapReGeo.ReGeoCode != nil {
			address = amapReGeo.ReGeoCode.FormattedAddress
			province = amapReGeo.ReGeoCode.AddressComponent.Province

			cityUnknow := amapReGeo.ReGeoCode.AddressComponent.City
			switch cityUnknow := cityUnknow.(type) {
			case string:
				city = cityUnknow
			default:
				city = ""
			}

			return
		}
	}

	errmsg := "amap regeo api error"
	if amapReGeo.Status == "0" {
		errmsg = amapReGeo.Info
	}
	err = errors.New(errmsg)
	return
}
