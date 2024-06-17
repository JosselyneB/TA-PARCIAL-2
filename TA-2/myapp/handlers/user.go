package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "myapp/models"
)

func CreateUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return err
    }
    if err := c.Validate(user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
    id := c.Param("id")
    // Lógica para obtener usuario
    return c.JSON(http.StatusOK, id)
}

func UpdateUser(c echo.Context) error {
    id := c.Param("id")
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return err
    }
    if err := c.Validate(user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    // Lógica para actualizar usuario
    return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
    id := c.Param("id")
    // Lógica para eliminar usuario
    return c.NoContent(http.StatusNoContent)
}
