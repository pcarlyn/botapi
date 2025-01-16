package main

import (
	"fmt"
	"start/internal/models"
	"start/internal/utils/requests/variables"
)

func main() {

	profile, status := variables.PatchVariables(models.Variable{ProfileId: 1, Name: "test", Value: "test"})
	fmt.Println("Status: ", status)
	fmt.Printf("Response: %+v\n", profile)
}
