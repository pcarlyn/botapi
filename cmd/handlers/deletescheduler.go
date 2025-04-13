package handlers

import (
	"fmt"
	"net/http"
	"start/internal/models"
	"start/internal/utils/database"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSchedulerById(c echo.Context) error {
	idstr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.SchedulerAnswer{})
	}

	result := database.Remove(id)
	return c.JSON(http.StatusOK, result)
}

// DeleteScheduler отменяет отложенную задачу по ID
// @Summary Отмена задачи
// @Description Останавливает таймер и помечает задачу как "canceled" по переданному ID
// @Tags scheduler
// @Param id path string true "ID задачи (ObjectID)"
// @Success 200 {object} map[string]interface{} "Задача успешно отменена"
// @Failure 400 {object} map[string]interface{} "Неверный ID"
// @Failure 404 {object} map[string]interface{} "Задача не найдена"
// @Router /frontapi/v1/schedulers/{id} [delete]
func DeleteScheduler(c echo.Context) error {
	idStr := c.Param("id")
	insertedID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	mutex.Lock()
	timer, exists := timers[insertedID]
	mutex.Unlock()

	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "task not found"})
	}

	if timer.Stop() {
		fmt.Printf("📅 [Task %s] Timer stopped before execution\n", insertedID.Hex())
	}

	mutex.Lock()
	delete(timers, insertedID)
	mutex.Unlock()

	database.Patch(insertedID, models.UpdateMsg{Msg: "canceled"})

	return c.JSON(http.StatusOK, echo.Map{
		"status": "task canceled",
		"id":     insertedID.Hex(),
	})
}
