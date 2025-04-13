package handlers

import (
	"fmt"
	"net/http"
	"start/internal/models"
	"start/internal/utils/database"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	timers = make(map[primitive.ObjectID]*time.Timer)
	// messages = make(map[primitive.ObjectID]models.SchedulerMsg)
	mutex sync.Mutex
)

// PostScheduler создает новую задачу с отложенной отправкой сообщения
// @Summary Создать задачу
// @Description Создает задачу, которая будет выполнена в будущем времени
// @Tags scheduler
// @Accept json
// @Produce json
// @Param payload body models.SchedulerMsg true "Данные новой задачи"
// @Success 200 {object} map[string]string "Задача успешно запланирована"
// @Failure 400 {object} map[string]string "Неверные входные данные или время отправки в прошлом"
// @Router /frontapi/v1/schedulers [post]
func PostScheduler(c echo.Context) error {
	var msg models.SchedulerMsg

	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	if time.Until(msg.SendAt) <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "send_at must be in the future"})
	}

	insertedID := database.Save(msg)

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
