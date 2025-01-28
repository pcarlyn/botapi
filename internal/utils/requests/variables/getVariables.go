package variables

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"start/internal/models"
)

func GetVariables(id int) ([]models.VariablesResponse, int) {
	url := models.BaseUrl + "/variables/?id=" + strconv.Itoa(id)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in GetVariables: %v", err)
		return []models.VariablesResponse{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in GetVariables: %v", err)
		return []models.VariablesResponse{}, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in GetVariables: %v", err)
		return []models.VariablesResponse{}, http.StatusInternalServerError
	}

	var variables []models.VariablesResponse
	err = json.Unmarshal(body, &variables)
	if err != nil {
		fmt.Printf("Error in GetVariables: %v", err)
		return []models.VariablesResponse{}, http.StatusInternalServerError
	}

	return variables, resp.StatusCode
}
