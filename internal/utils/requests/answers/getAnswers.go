package answers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
)

func GetAnswers(param, exp string) ([]models.ResponseAnswer, int) {

	var url string
	switch param {
	case "clb":
		url = models.BaseUrl + "/answers/clb/?callback=" + exp
	case "cmd":
		url = models.BaseUrl + "/answers/cmd/?cmd=" + exp
	case "txt":
		url = models.BaseUrl + "/answers/txt/?text=" + exp
	default:
		return []models.ResponseAnswer{}, http.StatusInternalServerError
	}

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in GetAnswers: %v", err)
		return []models.ResponseAnswer{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in GetAnswers: %v", err)
		return []models.ResponseAnswer{}, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in GetAnswers: %v", err)
		return []models.ResponseAnswer{}, http.StatusInternalServerError
	}

	var answers []models.ResponseAnswer
	err = json.Unmarshal(body, &answers)
	if err != nil {
		fmt.Printf("Error in GetAnswers: %v", err)
		return []models.ResponseAnswer{}, http.StatusInternalServerError
	}

	return answers, resp.StatusCode
}
