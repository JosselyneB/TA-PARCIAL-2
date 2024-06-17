package handlers

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    if username == "user" && password == "password" {
        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["name"] = "John Doe"
        claims["admin"] = true
        claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

        t, err := token.SignedString([]byte("secret"))
        if err != nil {
            return err
        }

        return c.JSON(http.StatusOK, echo.Map{
            "token": t,
        })
    }

    return echo.ErrUnauthorized
}

func Restricted(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    return c.String(http.StatusOK, "Welcome "+name+"!")
}
