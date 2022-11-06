package juconf

import (
	"path/filepath"
	"testing"
	"time"
)

func TestReadJsonConfig(t *testing.T) {

	// var settingFilePath string

	// if 1 == len(os.Args) {
	// 	settingFilePath = filepath.Join(filepath.Dir(os.Args[0]), "config.json")
	// } else {
	// 	settingFilePath = os.Args[1]
	// }
	// t.Log(filepath.Dir(os.Args[0]))
	settingFilePath := "./config.json"

	settingFilePath, _ = filepath.Abs(settingFilePath)
	t.Log("configFileName: ", settingFilePath)

	err := LoadFormFile(settingFilePath)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("Run successful: %+v \n", Conf)

	time.Sleep(time.Duration(1) * time.Second)
}
