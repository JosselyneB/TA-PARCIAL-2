package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "myapp/handlers"
    "myapp/utils"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.JWT([]byte("secret")))

    // Validators
    e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}

    // Rutas públicas
    e.POST("/login", handlers.Login)
    e.POST("/users", handlers.CreateUser)

    // Grupo de rutas protegidas
    r := e.Group("/restricted")
    r.Use(middleware.JWT([]byte("secret")))
    r.GET("", handlers.Restricted)

    // GraphQL
    e.POST("/graphql", handlers.GraphQL)

    // WebSocket
    e.GET("/ws", handlers.ServeWs)

    e.Logger.Fatal(e.Start(":1323"))
}
