package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-echo-gorm-practice/service"
)

func GetStudents(c echo.Context) error {
	students, _ := service.Students()
	return c.JSON(http.StatusOK, students)
}