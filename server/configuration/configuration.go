package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"stormlab.fr/stormbudget/v1/server/errors"
	"stormlab.fr/stormbudget/v1/server/validation"
)

type Configuration struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func ParseConfiguration(file string) (*Configuration, error) {
	// Validate the JSON file with schema
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// The second argument of ValidateJSONFile permits to use the validator on Linux AND Windows
	valid, err := validation.ValidateJSONFile(file, strings.Replace(path+"/schemas/configuration.json", "\\", "/", -1))
	if err != nil {
		return nil, err
	}
	if !valid {
		validationError := errors.ValidationError{
			Message: "The JSON file is not valid",
		}
		return nil, &validationError
	}
	// Parse the configuration file
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
