package handlers

import (
	"net/http"
	"start/internal/models"
	"start/internal/utils/database"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSchedulers возвращает список всех отложенных задач
// @Summary Получить список всех задач
// @Description Возвращает массив всех задач с датами в часовом поясе Europe/Moscow
// @Tags scheduler
// @Success 200 {array} map[string]interface{} "Список задач"
// @Failure 500 {array} models.SchedulerAnswer "Ошибка сервера при получении задач"
// @Router /frontapi/v1/schedulers [get]
func GetSchedulers(c echo.Context) error {
	result, err := database.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, []models.SchedulerAnswer{})
	}
	loc, _ := time.LoadLocation("Europe/Moscow")

	for _, doc := range result {
		// Перевод SendAt
		if sendAt, ok := doc["sendat"].(primitive.DateTime); ok {
			doc["sendat"] = sendAt.Time().In(loc)
		}

		// Перевод CreatedAt
		if createdAt, ok := doc["createdat"].(primitive.DateTime); ok {
			doc["createdat"] = createdAt.Time().In(loc)
		}

		// Перевод UpdatedAt
		if updatedAt, ok := doc["updatedat"].(primitive.DateTime); ok {
			doc["updatedat"] = updatedAt.Time().In(loc)
		}
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, models.ProcessAnswer{})
	}

	return c.JSON(http.StatusOK, result)
}

// GetSchedulerById возвращает данные отложенной задачи по ID
// @Summary Получить задачу по ID
// @Description Возвращает полную информацию о задаче, включая время создания, обновления и отправки
// @Tags scheduler
// @Param id path string true "ID задачи (ObjectID)"
// @Success 200 {object} map[string]interface{} "Информация о задаче"
// @Failure 400 {object} models.SchedulerAnswer "Неверный ID"
// @Router /frontapi/v1/schedulers/{id} [get]
func GetSchedulerById(c echo.Context) error {
	idstr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.SchedulerAnswer{})
	}
	result := database.Get(id)

	loc, _ := time.LoadLocation("Europe/Moscow")

	if t, ok := result["sendat"].(primitive.DateTime); ok {
		result["sendat"] = t.Time().In(loc)
		// .Format("02.01.2006 15:04")
	}
	if t, ok := result["createdat"].(primitive.DateTime); ok {
		result["createdat"] = t.Time().In(loc)
		// .Format("02.01.2006 15:04")
	}
	if t, ok := result["updatedat"].(primitive.DateTime); ok {
		result["updatedat"] = t.Time().In(loc)
		// .Format("02.01.2006 15:04")
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, models.ProcessAnswer{})
	}
	return c.JSON(http.StatusOK, result)
}
