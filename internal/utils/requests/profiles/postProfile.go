package profiles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
)

func PostProfile(profile *models.PostRequestProfile) (*models.ResponseProfile, int) {
	url := models.BaseUrl + "/profiles/"

	jsonData, err := json.Marshal(profile)
	if err != nil {
		fmt.Printf("Error in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusInternalServerError
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Error in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusBadRequest
	}

	if resp.StatusCode == http.StatusConflict {
		fmt.Printf("Error in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusConflict
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusInternalServerError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusInternalServerError
	}

	var profiles models.ResponseProfile
	err = json.Unmarshal(body, &profiles)
	if err != nil {
		fmt.Printf("Error in PostProfile: %v", err)
		return &models.ResponseProfile{}, http.StatusInternalServerError
	}

	return &profiles, resp.StatusCode
}
