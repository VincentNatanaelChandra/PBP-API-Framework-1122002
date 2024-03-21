package routes

import (
	"echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/user", controllers.FetchAllUsers)
	e.POST("/user", controllers.InsertUser)
	e.PUT("/user", controllers.UpdateUser)
	e.DELETE("/user", controllers.DeleteUser)

	return e
}
