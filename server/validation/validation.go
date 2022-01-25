package validation

import "github.com/xeipuuv/gojsonschema"

func ValidateJSONFile(jsonObjectFile string, jsonSchemaFile string) (bool, error) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + jsonSchemaFile)
	objectLoader := gojsonschema.NewReferenceLoader("file://" + jsonObjectFile)
	return validateJSON(schemaLoader, objectLoader)
}

func ValidateJSONObject(jsonObjectString string, jsonSchemaFile string) (bool, error) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + jsonSchemaFile)
	objectLoader := gojsonschema.NewStringLoader(jsonObjectString)
	return validateJSON(schemaLoader, objectLoader)
}

func validateJSON(schemaLoader, objectLoader gojsonschema.JSONLoader) (bool, error) {
	result, err := gojsonschema.Validate(schemaLoader, objectLoader)
	if err != nil {
		return false, err
	}
	return result.Valid(), nil
}
