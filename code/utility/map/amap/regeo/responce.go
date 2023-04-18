package regeo

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
