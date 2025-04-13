package handlers

import (
	"net/http"
	"start/internal/models"
	"start/internal/utils/database"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateScheduler(c echo.Context) error {
	idstr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.SchedulerAnswer{})
	}
	var msg models.UpdateMsg

	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	result := database.Patch(id, msg)
	return c.JSON(http.StatusOK, result)
}
