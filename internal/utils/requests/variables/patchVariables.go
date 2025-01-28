package variables

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"start/internal/models"
)

func PatchVariables(variables models.Variable) (*models.UpdateVariablesOK, int) {

	variablesJson, err := json.Marshal(variables)
	if err != nil {
		return &models.UpdateVariablesOK{}, http.StatusInternalServerError
	}

	req, err := http.NewRequest("PATCH", models.BaseUrl+"/variables/", bytes.NewBuffer(variablesJson))
	if err != nil {
		return &models.UpdateVariablesOK{}, http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

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
