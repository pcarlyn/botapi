package states

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
)

func PostStates(state *models.RState) (*models.RState, int) {
	url := models.BaseUrl + "/states/"

	jsonData, err := json.Marshal(state)
	if err != nil {
		fmt.Printf("Error in PostStates: %v", err)
		return &models.RState{}, http.StatusInternalServerError
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in PostStates: %v", err)
		return &models.RState{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Error in PostStates: %v", err)
		return &models.RState{}, http.StatusBadRequest
	}
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error in PostStates: %v", err)
		return &models.RState{}, http.StatusInternalServerError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in PostStates: %v", err)
		return &models.RState{}, http.StatusInternalServerError
	}

	err = json.Unmarshal(body, &state)
	if err != nil {
		fmt.Printf("Error in PostStates: %v", err)
		return &models.RState{}, http.StatusInternalServerError
	}

	return state, resp.StatusCode
}
