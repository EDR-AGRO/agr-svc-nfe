package handlers

import (
	"br/com/agr/nfe/api/v1/schemas"
	"br/com/agr/nfe/domain/services"
	"br/com/agr/nfe/infrastructure/apm"
	"br/com/agr/nfe/infrastructure/httpclient"
	"br/com/agr/nfe/infrastructure/jsonschema"
	"br/com/agr/nfe/infrastructure/logger"
	"io"
	"net/http"
)

type NfeRouterHandler struct {
	nfeService services.NfeServiceIF
}

func NewNfeRouterHandler(nfeService services.NfeServiceIF) *NfeRouterHandler {
	return &NfeRouterHandler{
		nfeService: nfeService,
	}
}

// version godoc
// @Summary Emite nfp-e
// @Description  Emite nota fiscal eletronica de produtor
// @Accept  json
// @Produce  json
// @Param        nfe body dto.InfNFe true "Corpo da mensagem"
// @Success 200 {object} helpers.JsonResponse "Response"
// @failure 400 {object} helpers.JsonResponse "Response"
// @failure 500 {object} helpers.JsonResponse "Response"
// @Router /agr-svc-nfe/v1/emit [post]
func (h *NfeRouterHandler) EmitNfe(w http.ResponseWriter, r *http.Request) {

	apmt := apm.StartHttpTransaction(r, "EmitNfe")

	var status int = http.StatusOK
	var response httpclient.AppResponse = httpclient.AppResponse{
		Success: true,
		Result:  "Mensagem recebida com sucesso",
		Errors:  nil,
	}

	defer func() {
		httpclient.WriteJSON(w, status, response, r)
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf(apmt.Ctx, "erro ao ler o corpo da requisição: %v", err)
		status = http.StatusInternalServerError
		response = httpclient.AppErr(status, "Erro ao ler o corpo da requisição")
		return
	}

	logger.Infof(apmt.Ctx, "nfe payload: %s", string(body))

	schemaSpan := apmt.StartSpan("NfeSchema", "SCHEMA_VALIDATOR")
	jsonValidation := jsonschema.Validate(schemas.NfeSchema, string(body))
	schemaSpan.EndSpan()

	if jsonValidation != nil {
		logger.Errorf(apmt.Ctx, "erro no body da requisição, schema invalido: %v", jsonValidation.Errors)
		status = http.StatusBadRequest
		response = *jsonValidation
		return
	}

	/**TODO
	create a service file
	create xml payload with tags
	implement build xml
	implement load cert
	sign nfe
	create map with with states and url sefaz gov
	send nfe
	*/

}
