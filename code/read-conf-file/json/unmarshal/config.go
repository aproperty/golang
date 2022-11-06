package juconf

import (
	"encoding/json"

	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var Conf JsonConfig

type JsonConfig struct {
	SQLType       string `json:"mysqlType"`
	SQLAddress    string `json:"mysqlAddress"`
	SQLWriteQueue uint32 `json:"mysqlQueue"`

	BoolValue bool   `json:"boolValue"`
	MyName    string `json:"myName"`
}

// LoadFormFile ...
func LoadFormFile(filePath string) error {

	fileSuffix := path.Ext(filePath)[1:]
	if fileSuffix == "" {
		return errors.New("can not find file suffix")
	}
	fileSuffix = strings.ToLower(fileSuffix)

	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	switch fileSuffix {
	case "json":
		err = json.Unmarshal(content, &Conf)
		return err
	default:
		return errors.New("unknow configure file type")
	}
}
