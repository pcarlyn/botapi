package handlers

import (
	"net/http"
	"os/exec"
	"start/internal/models"
	"start/internal/utils/database"

	"github.com/labstack/echo/v4"
)

// StartProcess запускает новый процесс с заданным путем и аргументами
// @Summary Запуск процесса
// @Description Запускает новый процесс на сервере с переданными данными
// @Tags processes
// @Accept json
// @Produce json
// @Param payload body models.Process true "Данные для запуска процесса"
// @Success 200 {object} map[string]interface{} "Процесс успешно запущен"
// @Failure 400 {object} map[string]string "Неверные входные данные"
// @Failure 500 {object} map[string]string "Ошибка запуска процесса"
// @Router /frontapi/v1/processes [post]
func StartProcess(c echo.Context) error {

	var proc models.Process

	if err := c.Bind(&proc); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	cmd := exec.Command(proc.Path, proc.Args...)

	if err := cmd.Start(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	pid := cmd.Process.Pid

	insertedID := database.PostProcess(proc, uint32(pid))

	go func() {
		cmd.Wait()
	}()

	return c.JSON(http.StatusOK, echo.Map{
		"status": "started",
		"pid":    pid,
		"id":     insertedID,
	})
}
