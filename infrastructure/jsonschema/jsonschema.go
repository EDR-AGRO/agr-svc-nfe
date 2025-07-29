package jsonschema

import (
	"br/com/agr/nfe/infrastructure/httpclient"
	"net/http"

	"github.com/xeipuuv/gojsonschema"
)

type SchemaError struct {
	Errors []string `json:"errors"`
}

func (se *SchemaError) ToAppResponse() *httpclient.AppResponse {
	ar := &httpclient.AppResponse{
		Success: false,
		Result:  nil,
		Errors:  []httpclient.AppError{},
	}

	for _, err := range se.Errors {
		ar.Errors = append(ar.Errors, httpclient.AppError{
			Message: err,
			Code:    http.StatusBadRequest,
		})
	}

	return ar
}

func Validate(schema string, json string) (appError *httpclient.AppResponse) {
	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewStringLoader(json)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return &httpclient.AppResponse{
			Success: false,
			Result:  nil,
			Errors: []httpclient.AppError{
				{
					Message: "Error validating schema",
					Code:    http.StatusInternalServerError,
				},
			},
		}
	}

	return validateSchemaResult(result)
}

func getJsonSchemaError(errors []gojsonschema.ResultError) *SchemaError {
	schemaError := &SchemaError{
		Errors: []string{},
	}

	for _, desc := range errors {
		newError := desc.Description()
		schemaError.Errors = append(schemaError.Errors, newError)
	}

	return schemaError
}

func validateSchemaResult(result *gojsonschema.Result) (appError *httpclient.AppResponse) {
	if result.Valid() {
		return nil
	}

	return getJsonSchemaError(result.Errors()).ToAppResponse()
}
