package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func ParseConfiguration(file string) (*Configuration, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	var configuration Configuration
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}
