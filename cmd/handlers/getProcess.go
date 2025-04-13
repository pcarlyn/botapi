package handlers

import (
	"net/http"
	"start/internal/models"
	"start/internal/utils/database"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProcess возвращает процесс по ID
// @Summary Получение процесса по ID
// @Description Получаем процесс по его ID
// @Tags processes
// @Accept json
// @Produce json
// @Param id path string true "ID процесса"
// @Success 200 {object} models.ProcessAnswer "Процесс"
// @Failure 400 {object} map[string]string "Ошибка: неверный ID"
// @Failure 404 {object} map[string]string "Ошибка: процесс не найден"
// @Router /frontapi/v1/processes/{id} [get]
func GetProcess(c echo.Context) error {
	idstr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ProcessAnswer{})
	}
	result := database.GetProcessById(id)
	loc, _ := time.LoadLocation("Europe/Moscow")
	if t, ok := result["runat"].(primitive.DateTime); ok {
		result["runat"] = t.Time().In(loc)
		// .Format("02.01.2006 15:04")
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, models.ProcessAnswer{})
	}
	return c.JSON(http.StatusOK, result)
}

// GetProcesses возвращает все процессы
// @Summary Получение всех процессов
// @Description Получаем список всех процессов
// @Tags processes
// @Accept json
// @Produce json
// @Success 200 {array} models.SchedulerAnswer "Список процессов"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /frontapi/v1/processes [get]
func GetProcesses(c echo.Context) error {
	result, err := database.GetAllProcesses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, []models.SchedulerAnswer{})
	}
	loc, _ := time.LoadLocation("Europe/Moscow")

	for _, doc := range result {
		// Перевод RunAt
		if sendAt, ok := doc["runat"].(primitive.DateTime); ok {
			doc["runat"] = sendAt.Time().In(loc)
		}
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, models.ProcessAnswer{})
	}
	return c.JSON(http.StatusOK, result)
}
