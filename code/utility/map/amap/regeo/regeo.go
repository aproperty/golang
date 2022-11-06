package regeo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"temporary/constant"

	jsoniter "github.com/json-iterator/go"
)

type AmapReGeo struct {
	Status    string         `json:"status"` // 返回结果状态值: 返回值为 0 或 1，0 表示请求失败；1 表示请求成功
	Info      string         `json:"info"`   // 返回状态说明: 当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”
	InfoCode  string         `json:"infocode"`
	ReGeoCode *AmapReGeoCode `json:"regeocode"` // 逆地理编码列表
}

type AmapReGeoCode struct {
	FormattedAddress string                `json:"formatted_address"` // 结构化地址信息
	AddressComponent *AmapAddressComponent `json:"addressComponent"`  // 地址元素列表
}

type AmapAddressComponent struct {
	Province string      `json:"province"` // 坐标点所在省名称
	City     interface{} `json:"city"`     // 坐标点所在城市名称
	// District string      `json:"district"` // 坐标点所在区
}

func GetReGeoByAddress(location string) (address string, province string, city string, err error) {

	var baseUrl string = "https://restapi.amap.com/v3/geocode/regeo"
	url := fmt.Sprintf("%s?location=%s&key=%s", baseUrl, location, constant.AmapKey)

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
