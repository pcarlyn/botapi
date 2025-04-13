package handlers

import (
	"fmt"
	"net/http"
	"start/internal/models"
	"start/internal/utils/database"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChangeScheduler обновляет существующую задачу
// @Summary Обновить задачу
// @Description Обновляет параметры ранее запланированной задачи по ID
// @Tags scheduler
// @Accept json
// @Produce json
// @Param id path string true "ID задачи"
// @Param payload body models.SchedulerMsg true "Обновленные данные задачи"
// @Success 200 {object} map[string]string "Задача успешно обновлена и перепланирована"
// @Failure 400 {object} map[string]string "Неверные входные данные или время отправки в прошлом"
// @Router /frontapi/v1/schedulers/{id} [put]
func ChangeScheduler(c echo.Context) error {
	idstr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.SchedulerAnswer{})
	}
	var msg models.SchedulerMsg

	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	if time.Until(msg.SendAt) <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "send_at must be in the future"})
	}

	database.Put(id, msg)
	insertedID := id

	delay := time.Until(msg.SendAt)
	timer := time.AfterFunc(delay, func() {
		fmt.Printf("📨 [Task %s] Sending message: %s to %s via %s\n", insertedID.Hex(), msg.Content, msg.Recipient, msg.Channel)

		database.Patch(insertedID, models.UpdateMsg{Msg: "sent"})
		mutex.Lock()
		delete(timers, insertedID)
		mutex.Unlock()
	})

	mutex.Lock()
	timers[insertedID] = timer
	mutex.Unlock()

	return c.JSON(http.StatusOK, echo.Map{
		"status": "scheduled",
		"id":     insertedID.Hex(),
	})
}
