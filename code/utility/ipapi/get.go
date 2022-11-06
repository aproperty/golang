package ipapi

// GetCountryCode returns the two-letter country code ISO 3166-1 alpha-2e for a given Ip.
func (ipdata *IPData) GetCountryCode() string {
	return ipdata.CountryCode
}

// GetRegion returns the region/state short code (FIPS or ISO).
func (ipdata *IPData) GetRegion() string {
	return ipdata.Region
}

// GetRegionName returns the Region/state for a given Ip
func (ipdata *IPData) GetRegionName() string {
	return ipdata.RegionName
}

// GetCity returns the CITY
func (ipdata *IPData) GetCity() string {
	return ipdata.City
}

// GetTimezone returns the Timezone
func (ipdata *IPData) GetTimezone() string {
	return ipdata.Timezone
}

// GetIsp gives the ISP.
func (ipdata *IPData) GetIsp() string {
	return ipdata.Isp
}

// GetOrg returns the organisation
func (ipdata *IPData) GetOrg() string {
	return ipdata.Org
}

// GetAs returns the AS number and organization, separated by space (RIR). Empty for IP blocks not being announced in BGP tables.
func (ipdata *IPData) GetAs() string {
	return ipdata.As
}

// GetLatitude gives Latitude
func (ipdata *IPData) GetLatitude() float64 {
	return ipdata.Lat
}

//GetLongitude gives Longitude
func (ipdata *IPData) GetLongitude() float64 {
	return ipdata.Lon
}

// GetZip gives Zip
func (ipdata *IPData) GetZip() string {
	return ipdata.Zip
}

// GetCountry is
func (ipdata *IPData) GetCountry() string {
	return ipdata.Country
}
