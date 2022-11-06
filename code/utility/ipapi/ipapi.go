package ipapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"net/http"
)

const ipapiURL string = "http://ip-api.com/json/"
const lang string = "?lang=zh-CN"

// const fields string = "?fields=66846719"

// FromString creates a new Ipdata object from String.
func FromString(ip string) (*IPData, error) {

	resp, err := http.Get(ipapiURL + ip + lang)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	ipData := new(IPData)
	err = json.Unmarshal(buf.Bytes(), &ipData)
	if ipData.Status != "success" || err != nil {
		return nil, errors.New("IPAPI: Ip Address is Invalid")
	}

	return ipData, nil
}

// FromIP creates a new Ipdata object from IP Address.
func FromIP(ip *net.IP) (*IPData, error) {
	ipString := ip.String()
	ipdata, err := FromString(ipString)
	return ipdata, err
}
