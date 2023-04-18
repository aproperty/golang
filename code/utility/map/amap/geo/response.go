package geo

type AmapGeo struct {
	Status   string          `json:"status"`   // 返回结果状态值:返回值为 0 或 1，0 表示请求失败；1 表示请求成功
	Info     string          `json:"info"`     // 返回状态说明: 当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”
	Count    string          `json:"count"`    // 返回结果的个数
	Geocodes *[]amapGeocodes `json:"geocodes"` // 地理编码信息列表
}

type amapGeocodes struct {
	FormattedAddress string `json:"formatted_address"` // 结构化地址信息
	Location         string `json:"location"`          // 坐标点：经度，纬度
}
