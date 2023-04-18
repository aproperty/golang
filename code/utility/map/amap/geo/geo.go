package geo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

const AmapKey string = "27780a1f06c4a3817fef6b67b21ef123"
const BaseUrl string = "https://restapi.amap.com/v3/geocode/geo"

func GetGeoByAddress(s string) (address string, location string, err error) {

	url := fmt.Sprintf("%s?address=%s&key=%s",
		BaseUrl,
		s,
		AmapKey)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	amapGeo := AmapGeo{}
	err = jsoniter.Unmarshal(body, &amapGeo)
	if err != nil {
		logrus.Error(err)
		return
	}

	if amapGeo.Status == "1" {
		var count int
		count, err = strconv.Atoi(amapGeo.Count)
		if err != nil {
			logrus.Error(err)
			return
		}
		if count > 0 && amapGeo.Geocodes != nil && len(*amapGeo.Geocodes) > 0 {
			address = (*amapGeo.Geocodes)[0].FormattedAddress
			location = (*amapGeo.Geocodes)[0].Location
			return
		}
	}

	errmsg := "amap geo api error"
	if amapGeo.Status == "0" {
		errmsg = amapGeo.Info
	}

	err = errors.New(errmsg)
	return
}
