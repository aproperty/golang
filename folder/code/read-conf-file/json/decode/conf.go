package jdconf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var Conf Config

type Config struct {
	Name    string `json:"name"`
	Threads int    `json:"threads"`
	Version string `json:"version"`
}

func ReadConfig() error {

	cfg := &Conf

	configFileName := "./config.json"
	// if len(os.Args) > 1 {
	// 	configFileName = os.Args[1]
	// }
	// fmt.Println(len(os.Args))
	// fmt.Printf("%+v   \n\n", os.Args)
	configFileName, _ = filepath.Abs(configFileName)
	fmt.Println("configFileName: ", configFileName)
	// fmt.Println()
	// fmt.Println("----")
	// fmt.Println()

	configFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(cfg); err != nil {
		return err
	}
	return nil
}
