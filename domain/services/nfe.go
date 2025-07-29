package services

import (
	"br/com/agr/nfe/domain/dto"
	"br/com/agr/nfe/infrastructure/apm"
	"br/com/agr/nfe/infrastructure/httpclient"
	"encoding/json"
	"fmt"
	"net/http"
)

type NfeServiceIF interface {
	PostEmitNfe(apmt *apm.ApmTransaction, message []byte) *httpclient.Error
}

type NfeService struct {
}

func NewNfeService() NfeServiceIF {
	return &NfeService{}
}

func (s *NfeService) PostEmitNfe(apmt *apm.ApmTransaction, request []byte) *httpclient.Error {

	nfeRequest := dto.NewInfNFe()
	if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request)), nfeRequest); err != nil {
		appError := httpclient.AppError{
			Message: "Error to Unmarshal nfe to struct",
			Code:    http.StatusInternalServerError,
		}
		return httpclient.NewError(http.StatusInternalServerError, err, appError)
	}

	return nil
}
