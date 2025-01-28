package profiles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
)

func PatchProfile(profile *models.LVUpdate) (*models.UpdateResponse, int) {

	url := models.BaseUrl + "/profiles/lvupdate/"

	jsonData, err := json.Marshal(profile)
	if err != nil {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusBadRequest
	}

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}

	var lvupdate models.UpdateResponse
	err = json.Unmarshal(body, &lvupdate)
	if err != nil {
		fmt.Printf("Error in PatchProfile: %v", err)
		return &models.UpdateResponse{}, http.StatusInternalServerError
	}

	return &lvupdate, resp.StatusCode
}
