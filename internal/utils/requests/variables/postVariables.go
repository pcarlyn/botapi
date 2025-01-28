package variables

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"start/internal/models"
)

func PostVariables(variables models.Variable) (*models.UpdateVariablesOK, int) {

	variablesJson, err := json.Marshal(variables)
	if err != nil {
		return &models.UpdateVariablesOK{}, http.StatusInternalServerError
	}

	resp, err := http.Post(models.BaseUrl+"/variables/", "application/json", bytes.NewBuffer(variablesJson))

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		return &models.UpdateVariablesOK{}, http.StatusInternalServerError
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		return &models.UpdateVariablesOK{}, http.StatusBadRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &models.UpdateVariablesOK{}, http.StatusInternalServerError
	}

	var response models.UpdateVariablesOK
	err = json.Unmarshal(body, &response)
	if err != nil {
		return &models.UpdateVariablesOK{}, http.StatusInternalServerError
	}

	return &response, resp.StatusCode
}
