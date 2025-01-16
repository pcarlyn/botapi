package states

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"start/internal/models"
	"strconv"
)

func GetStatesById(id int) (*models.RState, int) {
	var state *models.RState

	url := models.BaseUrl + "/states/?id=" + strconv.Itoa(id)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Error in GetStatesById: %v", err)
		return nil, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Error in GetStatesById: %v", err)
		return nil, http.StatusNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in GetStatesById: %v", err)
		return nil, http.StatusInternalServerError
	}

	err = json.Unmarshal(body, &state)
	if err != nil {
		fmt.Printf("Error in GetStatesById: %v", err)
		return nil, http.StatusInternalServerError
	}

	return state, resp.StatusCode
}
