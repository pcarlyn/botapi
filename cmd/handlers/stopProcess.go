package handlers

import (
	"net/http"
	"os"
	"start/internal/utils/database"
	"strconv"

	"github.com/labstack/echo/v4"
)

// StopProcess завершает процесс по его ID
// @Summary Завершение процесса по ID
// @Description Завершаем процесс по переданному ID процесса
// @Tags processes
// @Accept json
// @Produce json
// @Param id path string true "ID процесса для завершения"
// @Success 200 {object} map[string]interface{} "Процесс успешно завершен"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 500 {object} map[string]string "Не удалось завершить процесс"
// @Router /frontapi/v1/processes/{id} [delete]
func StopProcess(c echo.Context) error {
	idStr := c.Param("id")
	pid, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "cannot find process"})
	}

	if err := proc.Kill(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to kill process"})
	}

	answer := database.RemoveProcess(uint32(pid))

	return c.JSON(http.StatusOK, echo.Map{
		"status": "killed",
		"pid":    pid,
		"info":   answer,
	})
}
