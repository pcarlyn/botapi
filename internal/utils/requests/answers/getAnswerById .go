package answers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
	"strconv"
)

func GetAnswerById(id int) (models.GetResponseAnswer, int) {

	url := models.BaseUrl + "/answers/?id=" + strconv.Itoa(id)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in GetAnswerById: %v", err)
		return models.GetResponseAnswer{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in GetAnswerById: %v", err)
		return models.GetResponseAnswer{}, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in GetAnswerById: %v", err)
		return models.GetResponseAnswer{}, http.StatusInternalServerError
	}

	var answer models.GetResponseAnswer
	err = json.Unmarshal(body, &answer)
	if err != nil {
		fmt.Printf("Error in GetAnswerById: %v", err)
		return models.GetResponseAnswer{}, http.StatusInternalServerError
	}

	return answer, resp.StatusCode
}
