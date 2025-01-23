package states

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
)

func PatchStates(state models.RState) (*models.UpdateResponse, int) {

	url := models.BaseUrl + "/states/"

	jsonData, err := json.Marshal(state)
	if err != nil {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusBadRequest
	}
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}

	var response *models.UpdateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error in PatchStates: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}

	return response, resp.StatusCode
}
