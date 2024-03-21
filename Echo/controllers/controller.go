package controllers

import (
	"echo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllUsers(c echo.Context) error {
	result, err := models.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func InsertUser(c echo.Context) error {
	username := c.FormValue("username")
	age := c.FormValue("age")
	email := c.FormValue("email")
	ageint, err := strconv.Atoi(age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.InsertUser(username, ageint, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	username := c.FormValue("username")
	age := c.FormValue("age")
	email := c.FormValue("email")
	ageint, err := strconv.Atoi(age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(conv_id, username, ageint, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
