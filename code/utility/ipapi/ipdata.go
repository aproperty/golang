package ipapi

// IPData struct holds the data acquired from the ip-api.com for a given Ip Address ...
type IPData struct {
	Status string `json:"status"`

	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`

	Region     string `json:"region"`
	RegionName string `json:"regionName"`

	City string `json:"city"`

	Zip string `json:"zip"`

	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone"`

	Isp string `json:"isp"`
	Org string `json:"org"`
	As  string `json:"as"`

	Query string `json:"query"`
}

/*
{
	"status": "success",

	"country": "中国",
	"countryCode": "CN",

	"region": "GD",
	"regionName": "广东",

	"city": "深圳市",

	"zip": "",

	"lat": 22.5318,
	"lon": 114.1374,
	"timezone": "Asia/Shanghai",

	"isp": "Chinanet",
	"org": "Chinanet GD",
	"as": "AS4134 CHINANET-BACKBONE",

	"query": "61.141.73.132"
}
*/
