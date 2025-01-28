package profiles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
	"strconv"
)

func GetProfile(id int) (models.GetResponseProfile, int) {
	var url string
	switch {
	case id == 0:
		url = models.BaseUrl + "/profiles/"
	case id > 0:
		url = models.BaseUrl + "/profiles/?id=" + strconv.Itoa(id)
	default:
		return models.GetResponseProfile{}, http.StatusInternalServerError
	}

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in GetProfile: %v", err)
		return models.GetResponseProfile{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in GetProfile: %v", err)
		return models.GetResponseProfile{}, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in GetProfile: %v", err)
		return models.GetResponseProfile{}, http.StatusInternalServerError
	}

	var profiles models.GetResponseProfile
	err = json.Unmarshal(body, &profiles)
	if err != nil {
		fmt.Printf("Error in GetProfile: %v", err)
		return models.GetResponseProfile{}, http.StatusInternalServerError
	}

	return profiles, resp.StatusCode
}
