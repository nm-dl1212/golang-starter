package handlers

import (
	"app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetTasks endpoint
func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTasks())
}
